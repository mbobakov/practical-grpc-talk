# mbobakov.timemachine.api.v1 Package protocol

## Services
### TimeMachine

Info: TimeMachine is a service for communicating with the time engine server

| Method Name | Request Type | Response Type | Comments |
| ----------- | ------------ | ------------- | ------- |
| Jump | .mbobakov.timemachine.api.v1.JumpRequest | .google.protobuf.Duration | Jump to specific moment in time |

## Messages

### Passenger

Info:  Passenger for the journey

| Name | Type | Comments|
| ----------- | ------------ | ---------- |
| name |string| name of Passenger (min.length = 3) |

### JumpRequest

Info:  JumpRequest is a request for the jump throuth time

| Name | Type | Comments|
| ----------- | ------------ | ---------- |
| to |Timestamp||
|passenger |Passenger| Passengers for the trip (min=2 max=5)|

