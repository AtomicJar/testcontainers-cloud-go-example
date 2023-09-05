package main

import (
	"context"
	"database/sql"
	"fmt"
	"gotest.tools/v3/assert"
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
	t.Cleanup(func() {
		if err := c.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	cs, _ := c.ConnectionString(ctx, "sslmode=disable")
	db, err := sql.Open("postgres", cs)
	defer db.Close()

	var numberOfGuides int
	result, err := db.Query("SELECT COUNT(*) FROM guides")
	defer result.Close()

	result.Next()
	result.Scan(&numberOfGuides)
	require.NoError(t, err)
	assert.Equal(t, numberOfGuides, 6)

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
	if !containsCloud {
		expectedRuntime = info.OperatingSystem
	}
	if containsDesktop {
		expectedRuntime += " via Testcontainers Desktop app"
	}

	fmt.Printf(logo, expectedRuntime)
}
