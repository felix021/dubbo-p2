namespace go echo2

include "java.thrift"

struct EchoRequest {
    1: required string message,
    2: optional java.Object obj,
}(JavaClassName="kitex.echo2.EchoRequest")

struct EchoResponse {
    1: required string message,
}(JavaClassName="kitex.echo2.EchoResponse")

service TestService2 {
    EchoResponse Echo(1: EchoRequest req)
} (JavaClassName="kitex.echo2.TestService2")
