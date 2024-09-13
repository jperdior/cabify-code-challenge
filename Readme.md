# Car Pooling Service Challenge

![test workflow](https://github.com/jperdior/user-service/actions/workflows/test.yml/badge.svg)

## Candidate commentary

After being learning Go for the last months and knowing Cabify uses it intensively, I decided to put in practice what I've learned in this challenge and implemented the service in Go. I used gin framework for the server and the request handling.

Update: I removed the carPool from the context and passed it via dependency injection in the bootstrapping of the application.

### Improvements

CQRS it's an overkill in this exercise, and it's not needed, but I did it just to show my practical knowledge of the pattern.

Command and event buses are implemented in memory using go routines so there may be performance problems with a high number of requests. In a real world(TM) project I would use a message broker like RabbitMQ or Kafka.

### Requirements and how to run

- Make
- Docker
- Docker Compose
- Port 9091 available

To run the service, just execute:

```bash
make start
```

To run the tests, just execute:

```bash
make tests
```

To stop the service, just execute:

```bash
make stop
```

You can also restart the service with:

```bash
make restart
```

I included a Postman collection in the root of the project to test the service.

### Architecture

I implemented the solution using a DDD approach, with CQRS to communicate the server handlers with the application services. I also used Events to retry journeys when new cars are added or a journey is dropped. The structure may not be 100% "standard" for golang but I tried to follow the best practices I've learned so far.

- In cmd/api there's the main.go file, and the bootstrapping of the server.
- In kit I created basic interfaces for the command, query and event bus which are common to all the services.
- In internal/platform there's the infrastructure implementations, the request handlers and the bus in memory implementations using go routines.
- In internal/carpool there's the domain and application services, and the carpool aggregate.

### Decisions

- To avoid the need of a message broker, all buses work in memory with go routines, except the query bus.
- To avoid the need of a database, I create an instance of carpool in the bootstrapping of the server and I pass it via dependency injection to the application services/use cases.
- For the retry journey use case I am aware that there's space for improvement in performance, but as I wanted to respect the arrival order of the waiting groups it's just a loop.
- On the other side, for the locate and assign journey I created mappings so cars by seats available and groups can be retrieved with a complexity of O(1). This should allow the service to work reasonably well with at least 10^4 / 10^5 cars / waiting groups.

## The challenge

Design/implement a system to manage car pooling.

At Cabify we provide the service of taking people from point A to point B.
So far we have done it without sharing cars with multiple groups of people.
This is an opportunity to optimize the use of resources by introducing car
pooling.

You have been assigned to build the car availability service that will be used
to track the available seats in cars.

Cars have a different amount of seats available, they can accommodate groups of
up to 4, 5 or 6 people.

People requests cars in groups of 1 to 6. People in the same group want to ride
on the same car. You can take any group at any car that has enough empty seats
for them. If it's not possible to accommodate them, they're willing to wait until
there's a car available for them. Once a car is available for a group
that is waiting, they should ride.

Once they get a car assigned, they will journey until the drop off, you cannot
ask them to take another car (i.e. you cannot swap them to another car to
make space for another group).

In terms of fairness of trip order: groups should be served as fast as possible,
but the arrival order should be kept when possible.
If group B arrives later than group A, it can only be served before group A
if no car can serve group A.

For example: a group of 6 is waiting for a car and there are 4 empty seats at
a car for 6; if a group of 2 requests a car you may take them in the car.
This may mean that the group of 6 waits a long time,
possibly until they become frustrated and leave.

## Evaluation rules

This challenge has a partially automated scoring system. This means that before
it is seen by the evaluators, it needs to pass a series of automated checks
and scoring.

### Checks

All checks need to pass in order for the challenge to be reviewed.

- The `acceptance` test step in the `.gitlab-ci.yml` must pass in master before you
  submit your solution. We will not accept any solutions that do not pass or omit
  this step. This is a public check that can be used to assert that other tests
  will run successfully on your solution. **This step needs to run without
  modification**
- _"further tests"_ will be used to prove that the solution works correctly.
  These are not visible to you as a candidate and will be run once you submit
  the solution

### Scoring

There is a number of scoring systems being run on your solution after it is
submitted. It is ok if these do not pass, but they add information for the
reviewers.

## API

To simplify the challenge and remove language restrictions, this service must
provide a REST API which will be used to interact with it.

This API must comply with the following contract:

### GET /status

Indicate the service has started up correctly and is ready to accept requests.

Responses:

* **200 OK** When the service is ready to receive requests.

### PUT /cars

Load the list of available cars in the service and remove all previous data
(reset the application state). This method may be called more than once during
the life cycle of the service.

**Body** _required_ The list of cars to load.

**Content Type** `application/json`

Sample:

```json
[
  {
    "id": 1,
    "seats": 4
  },
  {
    "id": 2,
    "seats": 6
  }
]
```

Responses:

* **200 OK** When the list is registered correctly.
* **400 Bad Request** When there is a failure in the request format, expected
  headers, or the payload can't be unmarshalled.

### POST /journey

A group of people requests to perform a journey.

**Body** _required_ The group of people that wants to perform the journey

**Content Type** `application/json`

Sample:

```json
{
  "id": 1,
  "people": 4
}
```

Responses:

* **200 OK** or **202 Accepted** When the group is registered correctly
* **400 Bad Request** When there is a failure in the request format or the
  payload can't be unmarshalled.

### POST /dropoff

A group of people requests to be dropped off. Whether they traveled or not.

**Body** _required_ A form with the group ID, such that `ID=X`

**Content Type** `application/x-www-form-urlencoded`

Responses:

* **200 OK** or **204 No Content** When the group is unregistered correctly.
* **404 Not Found** When the group is not to be found.
* **400 Bad Request** When there is a failure in the request format or the
  payload can't be unmarshalled.

### POST /locate

Given a group ID such that `ID=X`, return the car the group is traveling
with, or no car if they are still waiting to be served.

**Body** _required_ A url encoded form with the group ID such that `ID=X`

**Content Type** `application/x-www-form-urlencoded`

**Accept** `application/json`

Responses:

* **200 OK** With the car as the payload when the group is assigned to a car. See below for the expected car representation
```json
  {
  "id": 1,
  "seats": 4
}
```

* **204 No Content** When the group is waiting to be assigned to a car.
* **404 Not Found** When the group is not to be found.
* **400 Bad Request** When there is a failure in the request format or the
  payload can't be unmarshalled.

## Tooling

At Cabify, we use Gitlab and Gitlab CI for our backend development work.
In this repo you may find a [.gitlab-ci.yml](./.gitlab-ci.yml) file which
contains some tooling that would simplify the setup and testing of the
deliverable. This testing can be enabled by simply uncommenting the final
acceptance stage. Note that the image build should be reproducible within
the CI environment.

Additionally, you will find a basic Dockerfile which you could use a
baseline, be sure to modify it as much as needed, but keep the exposed port
as is to simplify the testing.

:warning: Avoid dependencies and tools that would require changes to the
`acceptance` step of [.gitlab-ci.yml](./.gitlab-ci.yml), such as
`docker-compose`

:warning: The challenge needs to be self-contained so we can evaluate it.
If the language you are using has limitations that block you from solving this
challenge without using a database, please document your reasoning in the
readme and use an embedded one such as sqlite.

You are free to use whatever programming language you deem is best to solve the
problem but please bear in mind we want to see your best!

You can ignore the Gitlab warning "Cabify Challenge has exceeded its pipeline
minutes quota," it will not affect your test or the ability to run pipelines on
Gitlab.

## Requirements

- The service should be as efficient as possible.
  It should be able to work reasonably well with at least $`10^4`$ / $`10^5`$ cars / waiting groups.
  Explain how you did achieve this requirement.
- You are free to modify the repository as much as necessary to include or remove
  dependencies, subject to tooling limitations above.
- Document your decisions using MRs or in this very README adding sections to it,
  the same way you would be generating documentation for any other deliverable.
  We want to see how you operate in a quasi real work environment.

## Feedback

In Cabify, we really appreciate your interest and your time. We are highly
interested on improving our Challenge and the way we evaluate our candidates.
Hence, we would like to beg five more minutes of your time to fill the
following survey:

- https://forms.gle/EzPeURspTCLG1q9T7

Your participation is really important. Thanks for your contribution!
