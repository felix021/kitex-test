namespace go kitex.felix021.test

struct EchoRequest {
    1: required string message,
}

struct EchoResponse {
    1: required string message,
}

service TestService {
    EchoResponse Echo (1: EchoRequest req),
}
