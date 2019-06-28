using System;
using System.Threading.Tasks;
using Com.Example;
using Grpc.Core;

namespace example_csharp
{
    class GreeterImpl : Greeter.GreeterBase
    {
        public override Task<HelloResponse> SayHello(HelloRequest request, ServerCallContext context) {
            Console.WriteLine("Java service request: " + request.Name);
            return Task.FromResult(new HelloResponse {
                Message = {
                    "Hello " + request.Name,
                    "Aloha " + request.Name,
                    "Howdy " + request.Name 
                }
            });
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            // Build the server
            Console.WriteLine("Starting Java server on port 9002");
            Server server = new Server {
                Services = { Greeter.BindService(new GreeterImpl()) },
                Ports = { new ServerPort("localhost", 9002, ServerCredentials.Insecure)}
            };
            server.Start();

            // Set up gRPC client
            Channel channel = new Channel("localhost:9002", ChannelCredentials.Insecure);
            var client = new Greeter.GreeterClient(channel);

            // Call the service
            var req = new HelloRequest { Name = "C#" };
            var resp = client.SayHello(req);
            foreach (string msg in resp.Message) {
                Console.WriteLine(msg);
            }

            // Block for server termination
            Console.ReadKey();
            channel.ShutdownAsync().Wait();
            server.ShutdownAsync().Wait();
        }
    }
}
