PACKAGE = study-goroutine
BUILDPATH ?= $(CURDIR)
BASE	= $(BUILDPATH)
BIN		= $(BASE)/bin

# 이 부분은 사실상 도커 때문에 해주는 거임
UNAME := $(shell uname)                 # uname 은 쉘 명령어고, 이걸 치면 OS(?) 정보가 나옴. 맥의 경우 Darwin 이라고 뜸.
ifeq ($(UNAME), Linux)                  # 만약 OS 정보가 Linux 인 경우, GOENV 에 이하의 값들을 넣어주는 듯. 터미널에서 GOENV="GOOS=linux" 이런 식으로 치는 거랑 같은 거임.
	GOENV   ?= CGO_ENABLED=0 GOOS=linux
endif
GOBUILD = ${GOENV} go
GO      = go

BUILDTAG=-tags 'studyGoroutine'         # go build 명령어에서 -tags 옵션 줄 수 있음.
export GO111MODULE=on

# V랑 Q는 특정 명령어들의 경우, 출력 여부를 결정하게 하려고 만든 변수들
# Q만 찍어보면 @ 가 대입되어서 나옴.
# 명령어 앞에 @ 붙이면 출력 안 되니까, V 값에 따라서 출력 여부 결정하려고 만들어둠.
# 참고로, Q 만드는 구문은 filter로 나온 조건에 충족했으면 빈 값을 Q에 대입하고, 충족 못 했으면 문자열 @ 를 Q에 대입 하라는 뜻임.
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: build ; $(info $(M) building all steps… ) @ ## Build all steps


.PHONY: build
build: ; $(info $(M) building executable… ) @ ## Build program binary
	$Q cd $(BASE)/api && $(GOBUILD) build -i \
		$(BUILDTAG) \
		-o $(BIN)/$(PACKAGE)


# Test for V, Q
# run $ make echoTest
.PHONY: echoTest
echoTest: echoTestTwo ; $(info $(M) this is first test…) @ ## Example 1.. info 뒤에 주석을 붙였는데, 안 보이게 하고 싶으면 골뱅이를 앞에 붙여주면 됨!
	@echo "I'm echo test"
	$Q echo $Q "-------------------------"
# 문자열이 대입된 변수는 그냥 출력해도 나오지만, 직접 문자열을 입력할 때는 쌍따옴표나 `` 이거 안에 문자열을 써줘야 됨.


# echo 라는 명령어 자체는 출력이 안 되지만, 문자열은 출력이 됨.
# 왜냐면 echo 는 문자열을 출력하는 명령어니까, 걔한테 출력하라고 시킨 문자열은 출력이 되는 것
# 만약 $Q 빼면, "echo "---"" 라는 명령어도 출력되고, echo 에 의해서 "---" 부분도 출력이 됨.
.PHONY: echoTestTwo
echoTestTwo: ; $(info $(M) this is second test…) @ ## Example 2..
	echo "I'm second echo test"
	$Q @echo "-------------------------"
