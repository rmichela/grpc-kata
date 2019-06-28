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
            Console.WriteLine("Hello C#!");
        }
    }
}
