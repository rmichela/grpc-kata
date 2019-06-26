package com.example;

import io.grpc.*;
import io.grpc.stub.*;

public class Main {
    public static class GreeterImpl extends GreeterGrpc.GreeterImplBase {
        @Override
        public void sayHello(GreeterOuterClass.HelloRequest request, StreamObserver<GreeterOuterClass.HelloResponse> responseObserver) {
            System.out.println("Java service request: " + request.getName());

            responseObserver.onNext(GreeterOuterClass.HelloResponse.newBuilder()
                .addMessage("Hello " + request.getName())
                .addMessage("Aloha " + request.getName())
                .addMessage("Howdy " + request.getName())
                .build());
            responseObserver.onCompleted();
        }
    }

    public static void main(String[] args) throws Exception {
        // Build the server
        System.out.println("Starting Java server on port 9000");
        Server server = ServerBuilder.forPort(9000).addService(new GreeterImpl()).build().start();

        // Block for server termination
        server.awaitTermination();
    }
}
