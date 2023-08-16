# testcontainers-cloud-go-example

This is an example repository with a simple test confirming proper connection from Testcontainers Desktop (or the CI agent) to your [Testcontainers Cloud](https://app.testcontainers.cloud) account.

For details on how to bootstrap Testcontainers in an actual project, please refer to the [Getting started with Testcontainers Go guide](https://github.com/testcontainers/tc-guide-getting-started-with-testcontainers-for-go).

## Clone the repository and run the first Testcontainers test suite

```
git clone https://github.com/AtomicJar/testcontainers-cloud-go-example
cd testcontainers-cloud-go-example
```

## Run the test suite

Note that it's important to add the `-v` flag to the `go test` command, otherwise the output of the tests will be suppressed.

```shell
go test -v
```

## Confirm your environment is configured correctly

The test output should show the Testcontainers logo and which container runtime was used:  

```shell

████████╗███████╗███████╗████████╗ ██████╗ ██████╗ ███╗   ██╗████████╗ █████╗ ██╗███╗   ██╗███████╗██████╗ ███████╗ 
╚══██╔══╝██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██║████╗  ██║██╔════╝██╔══██╗██╔════╝ 
   ██║   █████╗  ███████╗   ██║   ██║     ██║   ██║██╔██╗ ██║   ██║   ███████║██║██╔██╗ ██║█████╗  ██████╔╝███████╗ 
   ██║   ██╔══╝  ╚════██║   ██║   ██║     ██║   ██║██║╚██╗██║   ██║   ██╔══██║██║██║╚██╗██║██╔══╝  ██╔══██╗╚════██║ 
   ██║   ███████╗███████║   ██║   ╚██████╗╚██████╔╝██║ ╚████║   ██║   ██║  ██║██║██║ ╚████║███████╗██║  ██║███████║ 
   ╚═╝   ╚══════╝╚══════╝   ╚═╝    ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝╚══════╝ 
  
  
  Congratulations on running your first test! 🎉
  Runtime used: 
      Testcontainers Cloud
 
  You can now return to the website to complete your onboarding.

2023/08/16 17:33:13 🐳 Terminating container: c2dbde862d74
2023/08/16 17:33:13 🚫 Container terminated: c2dbde862d74
--- PASS: TestTestcontainersCloud (2.63s)
PASS
ok      github.com/AtomicJar/testcontainers-cloud-go-example    2.946s
```

## (optional) Use Testcontainers Desktop to easily debug the database

[Testcontainers Desktop](https://testcontainers.com/desktop/) helps developers with common tasks such as debugging your Testcontainers-powered dependencies. Let's practice!

The tests in this project create a PostgreSQL database and populate it with sample data. You can [set a fixed port](https://newsletter.testcontainers.com/announcements/set-fixed-ports-to-easily-debug-development-services) for the `postgres` service, then [freeze containers shutdown](https://newsletter.testcontainers.com/announcements/freeze-containers-to-prevent-their-shutdown-while-you-debug) to easily connect to the database from your IDE after your tests run.

See if you can inspect the database. Username: `postgres`. Password: `postgres`.
 