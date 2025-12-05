package cmd

import (
	"fmt"
	"testing"

	"github.com/ayuxdev/cachex/pkg/config"
)

func TestBuildHelpMessage(t *testing.T) {
	cfg := config.DefaultConfig()
	helpMessage := buildHelpMessage(cfg)
	fmt.Print(helpMessage)
}
