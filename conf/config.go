package conf

import (
	"github.com/spf13/viper"
)

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper // 임베딩 시켜서 Viper 의 메소드를 사용하려고 함.
}

// StudyGoroutine ...
var StudyGoroutine *ViperConfig

func init() {
	StudyGoroutine = readConfig(map[string]interface{}{
		"debug_route": false,
		"port":        1323,
	})
}

func readConfig(defaults map[string]interface{}) *ViperConfig {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	return &ViperConfig{
		Viper: v,
	}
}
