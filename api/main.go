package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	ctrl "study_goroutine/api/controller"
	conf "study_goroutine/conf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)

const (
	banner = "\n" +
		"   _______.___________. __    __   _______  ____    ____      _______   ______   .______        ______    __    __  .___________. __  .__   __.  _______ \n" +
		"  /       |           ||  |  |  | |       \\ \\   \\  /   /     /  _____| /  __  \\  |   _  \\      /  __  \\  |  |  |  | |           ||  | |  \\ |  | |   ____| \n" +
		" |   (----`---|  |----`|  |  |  | |  .--.  | \\   \\/   /     |  |  __  |  |  |  | |  |_)  |    |  |  |  | |  |  |  | `---|  |----`|  | |   \\|  | |  |__    \n" +
		"  \\   \\       |  |     |  |  |  | |  |  |  |  \\_    _/      |  | |_ | |  |  |  | |      /     |  |  |  | |  |  |  |     |  |     |  | |  . `  | |   __|   \n" +
		" .----)   |   |  |     |  `--'  | |  '--'  |    |  |        |  |__| | |  `--'  | |  |\\  \\----.|  `--'  | |  `--'  |     |  |     |  | |  |\\   | |  |____  \n" +
		" |_______/    |__|      \\______/  |_______/     |__|         \\______|  \\______/  | _| `._____| \\______/   \\______/      |__|     |__| |__| \\__| |_______| \n" +
		" => Starting listen %s\n"
)

// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU()) // 고루틴 멀티코어 활용 관련 CPU 설정 코드. 이대로 사용하면 모든 CPU 사용
// }

func main() {
	StudyGoroutine := conf.StudyGoroutine
	e := echoInit(StudyGoroutine)
	signal := sigInit(e)
	mqConn, mqCh := mqInit(StudyGoroutine)
	defer mqConn.Close()
	defer mqCh.Close()

	if err := ctrl.InitHandler(StudyGoroutine, e, mqCh, signal); err != nil {
		e.Logger.Error("InitHandler Error")
		os.Exit(1) // 프로그램을 지정된 status로 즉시 종료
	}

	startServer(StudyGoroutine, e)
}

func echoInit(studyGoroutine *conf.ViperConfig) (e *echo.Echo) {
	e = echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
	}))

	e.HideBanner = true

	return e
}

func sigInit(e *echo.Echo) chan os.Signal { // graceful shutdown을 위한 메소드
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	go func() {
		sig := <-quit // 버퍼없는 채널. os.Signal 에 신호가 들어왔을 때 수신 작업이 일어날 것. 수신이 일어나면, shutdown 로직들이 작동할 것
		e.Logger.Error("Got signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
		signal.Stop(quit)
		close(quit)
	}()
	return quit
}

func mqInit(studyGoroutine *conf.ViperConfig) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(studyGoroutine.GetString("mq_host"))
	if err != nil {
		log.Println("Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel", err)
	}

	q, err := ch.QueueDeclare(
		studyGoroutine.GetString("mq_name"), // name
		false,                               // durable
		false,                               // delete when unused
		false,                               // exclusive
		false,                               // no-wait
		nil,                                 // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue", err)
	}

	log.Println("MQ Init Success...", "MQ Name", q.Name)

	return conn, ch
}

func startServer(studyGoroutine *conf.ViperConfig, e *echo.Echo) {
	// Start Server
	apiServer := fmt.Sprintf("0.0.0.0:%d", studyGoroutine.GetInt("port"))
	e.Logger.Debugf("Starting server, Listen[%s]", apiServer)

	fmt.Printf(banner, apiServer)
	if err := e.Start(apiServer); err != nil {
		e.Logger.Fatal(err)
	}
}
