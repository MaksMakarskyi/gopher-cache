package tests

import (
	"os"
	"testing"
	"time"

	"github.com/MaksMakarskyi/gopher-cache/internal/app"
)

func TestMain(m *testing.M) {
	app := app.NewApp(100)

	go app.Run("server", "localhost", "6379")

	time.Sleep(200 * time.Millisecond)

	code := m.Run()

	os.Exit(code)
}
