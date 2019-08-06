package repo_test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.SetEnvPrefix("BLOG_TEST")
	viper.AutomaticEnv()

	os.Exit(m.Run())
}
