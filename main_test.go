package main

import (
	"context"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const success = "                                         /                                      \n" +
	"                                       /////////                                  \n" +
	"                                    ///////////////                               \n" +
	"                                   /////////////////                              \n" +
	"                                      /////////////                               \n" +
	"                                     %%   ////   %                                \n" +
	"                                     %%    //   %%                                \n" +
	"                                   %%      //      %                              \n" +
	"                                 %%        ////      %                            \n" +
	"                                 %     /////////     %                            \n" +
	"                                  % /////////////// %%                            \n" +
	"                                    %%%%%%%%%%%%%%%       \n" +
	"  \n" +
	"    /%%%%%%    /%%                             /%%              /%%%%%                    \n" +
	"   /%%__  %%  | %%                            |__/             |__  %%                    \n" +
	"  | %%  \\ %% /%%%%%%    /%%%%%%  /%%%%%%/%%%%  /%%  /%%%%%%%      | %%  /%%%%%%   /%%%%%% \n" +
	"  | %%%%%%%%|_  %%_/   /%%__  %%| %%_  %%_  %%| %% /%%_____/      | %% |____  %% /%%__  %%\n" +
	"  | %%__  %%  | %%    | %%  \\ %%| %% \\ %% \\ %%| %%| %%       /%%  | %%  /%%%%%%%| %%  \\__/\n" +
	"  | %%  | %%  | %% /%%| %%  | %%| %% | %% | %%| %%| %%      | %%  | %% /%%__  %%| %%      \n" +
	"  | %%  | %%  |  %%%%/|  %%%%%%/| %% | %% | %%| %%|  %%%%%%%|  %%%%%%/|  %%%%%%%| %%      \n" +
	"  |__/  |__/   \\___/   \\______/ |__/ |__/ |__/|__/ \\_______/ \\______/  \\_______/|__/    \n" +
	"  \n" +
	"  \n" +
	"  You configured correctly your Testcontainers Cloud environment! 🎉\n" +
	"  Continue your journey at https://app.testcontainers.cloud\n"

func TestWithRedis(t *testing.T) {
	w := &strings.Builder{}
	l := log.New(w, "[TESTCONTAINERS] ", log.LstdFlags)

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:6.2.6-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Logger:           l,
	})
	if err == nil {
		t.Cleanup(func() {
			_ = redisC.Terminate(ctx)
		})
	}

	t.Run("Container can start", func(t *testing.T) {
		if err != nil {
			t.Fatalf("[Container creation] %s: %v", w, err)
		}

		logs, err := redisC.Logs(ctx)
		if err != nil {
			t.Fatalf("Logs: %s: %v", w, err)
		}

		bytes, err := io.ReadAll(logs)
		if err != nil {
			t.Fatalf("%s: %v", w, err)
		}
		t.Logf("Container logs: \n%s", string(bytes))

		if redisC.GetContainerID() == "" {
			t.Error("Container ID is empty, something went wrong starting the redis container")
		}
	})

	t.Run("Connected to Testcontainers Cloud", func(t *testing.T) {
		if err != nil {
			t.Skip("Container can't be started, seems there is an issue connecting to Testcontainers Cloud")
		}
		if !strings.Contains(w.String(), "testcontainerscloud") {
			t.Fatal("Can't find <testcontainerscloud> in logs, which means that most probably, you're not connected to Testcontainers Cloud:\n", w.String())
		}

		t.Log(success)
	})

}
