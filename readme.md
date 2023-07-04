# Temporal in Golang Tutorial - Example Code

This repository contains an example code for using Temporal Workflow SDK in Golang to build resilient, scalable, and distributed applications with ease. Temporal is an open-source workflow and coordination platform that allows you to write your business logic as workflows and handle all the underlying complexity of running and scaling them in the cloud or on-premises.

# Dependencies

To run this example, you’ll need the following dependencies:

- Go 1.18.x or later (brew install go)
- Temporal Server (docker run --rm -p 7233:7233 temporalio/temporal:latest)
- Temporal CLI (brew install temporalio/tap/tctl)

Setup

1. Clone this repository or download the source code.

```
```

2. Install the dependencies using `go mod`.

```
go mod tidy
```

3. Start the Temporal Server using Docker (or manually on your machine).

```
docker-compose up
```

4. Start the go http server

```
go run main.go
```

5. Run the example code:

```
curl -X GET http://localhost:5000/weather\?city\=Cairo
```

# Example Description

The example code defines a simple workflow that orchestrates an activity to get Weather for a given City name and returns WeatherData. The workflow receives the name as an input parameter and passes it to the activity.

[Tutorial](https://younisjad.medium.com/using-temporal-to-build-scalable-and-fault-tolerant-applications-in-golang-99ed2f47bf68)