# Testcontainers Cloud Go Example

The current repository helps you to verify that you configured your [Testcontainers Cloud][tcc] agent correctly in your
local environment.

## Clone the repository and run the first Testcontainer test suite

```
git clone https://github.com/AtomicJar/testcontainers-cloud-go-example
cd testcontainers-cloud-go-example
go mod download
go test -v -count=1
```

## Verify the agent is running

✅ __Passive State__: Agent awaiting a Testcontainers test to be executed.

![agent-running](./docs/passive-connection.png)

✅ __Running State__: Agent connected to Testcontainers Cloud.

![agent-running](./docs/active-connection.png)

⚠️ __Stopped State__: Agent is stopped and will not accept connections.

Please, Start the agent to continue.

![agent-stopped](./docs/stopped.png)

To download the agent for local usage, check the [download page here][tcc-download].

## Run the test suite

`go test -v -count=1`

### Your environment is correctly configured if

Test output:

```shell
=== RUN   TestWithRedis/Connected_to_Testcontainers_Cloud
    main_test.go:91:                                          /
                                               /////////
                                            ///////////////
                                           /////////////////
                                              /////////////
                                             %%   ////   %
                                             %%    //   %%
                                           %%      //      %
                                         %%        ////      %
                                         %     /////////     %
                                          % /////////////// %%
                                            %%%%%%%%%%%%%%%

            /%%%%%%    /%%                             /%%              /%%%%%
           /%%__  %%  | %%                            |__/             |__  %%
          | %%  \ %% /%%%%%%    /%%%%%%  /%%%%%%/%%%%  /%%  /%%%%%%%      | %%  /%%%%%%   /%%%%%%
          | %%%%%%%%|_  %%_/   /%%__  %%| %%_  %%_  %%| %% /%%_____/      | %% |____  %% /%%__  %%
          | %%__  %%  | %%    | %%  \ %%| %% \ %% \ %%| %%| %%       /%%  | %%  /%%%%%%%| %%  \__/
          | %%  | %%  | %% /%%| %%  | %%| %% | %% | %%| %%| %%      | %%  | %% /%%__  %%| %%
          | %%  | %%  |  %%%%/|  %%%%%%/| %% | %% | %%| %%|  %%%%%%%|  %%%%%%/|  %%%%%%%| %%
          |__/  |__/   \___/   \______/ |__/ |__/ |__/|__/ \_______/ \______/  \_______/|__/


          Congratulations on running your first test on Testcontainers Cloud! 🎉
          You can now return to the website to complete your onboarding

--- PASS: TestWithRedis (4.38s)
    --- PASS: TestWithRedis/Container_can_start (0.06s)
    --- PASS: TestWithRedis/Connected_to_Testcontainers_Cloud (0.00s)
PASS
ok  	github.com/AtomicJar/testcontainers-cloud-go-example	4.750s
```

Agent status:

![agent-running](./docs/active-connection.png)

[tcc]: https://testcontainers.cloud/

[tcc-download]: https://app.testcontainers.cloud/start/download?mode=update
