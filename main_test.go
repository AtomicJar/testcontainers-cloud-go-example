package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestTestcontainersCloud(t *testing.T) {
	ctx := context.Background()

	c, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("postgres:14-alpine"),
		postgres.WithInitScripts(filepath.Join("testdata", "init.sql")),
		postgres.WithDatabase("testcontainers-go"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(10*time.Second)),
	)
	require.NoError(t, err)
	defer c.Terminate(ctx)

	dockerClient, err := testcontainers.NewDockerClient()
	require.NoError(t, err)

	info, err := dockerClient.Info(ctx)
	require.NoError(t, err)

	serverVersion := info.ServerVersion

	containsCloud := strings.Contains(serverVersion, "testcontainerscloud")
	containsDesktop := strings.Contains(serverVersion, "Testcontainers Desktop")
	if !(containsCloud || containsDesktop) {
		fmt.Printf(ohNo)
		t.FailNow()
	}

	expectedRuntime := "Testcontainers Cloud"
	if containsCloud {
		expectedRuntime = info.OperatingSystem
	} else if containsDesktop {
		expectedRuntime = "via Testcontainers Desktop app"
	}

	fmt.Printf(logo, expectedRuntime)
}
