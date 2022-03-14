package main

import (
	"os"

	"github.com/kubeshop/testkube/pkg/executor/agent"
	"github.com/lreimer/testkube-executor-maven/pkg/runner"
)

func main() {
	agent.Run(runner.NewRunner(), os.Args)
}
