package main

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestTestcontainersCloud(t *testing.T) {
	ctx := context.Background()

	c, err := postgres.Run(
		ctx,
		"postgres:14-alpine",
		postgres.WithInitScripts(filepath.Join("testdata", "init.sql")),
		postgres.WithDatabase("testcontainers-go"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(10*time.Second)),
	)
	testcontainers.CleanupContainer(t, c)
	require.NoError(t, err)

	cs, err := c.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	db, err := sql.Open("postgres", cs)
	require.NoError(t, err)
	defer db.Close()

	var numberOfGuides int
	result, err := db.Query("SELECT COUNT(*) FROM guides")
	require.NoError(t, err)
	defer result.Close()
	require.NoError(t, err)

	result.Next()
	result.Scan(&numberOfGuides)
	require.Equal(t, numberOfGuides, 6)

	dockerClient, err := testcontainers.NewDockerClientWithOpts(context.Background())
	require.NoError(t, err)

	info, err := dockerClient.Info(ctx)
	require.NoError(t, err)

	serverVersion := info.ServerVersion

	containsCloud := strings.Contains(serverVersion, "testcontainerscloud")
	containsDesktop := strings.Contains(serverVersion, "Testcontainers Desktop")
	if !(containsCloud || containsDesktop) {
		fmt.Print(ohNo)
		t.FailNow()
	}

	expectedRuntime := "Testcontainers Cloud"
	if !containsCloud {
		expectedRuntime = info.OperatingSystem
	}
	if containsDesktop {
		expectedRuntime += " via Testcontainers Desktop app"
	}

	fmt.Printf(logo, expectedRuntime)
}
