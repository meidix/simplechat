# Simple Chat :satisfied:

A smaple chatroom project using using the Go programming language with its standard library and features like channels, goroutines, etc. he server and the client apps are included independently which allows for an easy setup.

## Get Started

To run the application you need to first navigate to the `server/` directory and run the following command:

```bash
cd server/
go run .
```

This command will run the chat server which accepts connection on `localhost:8080`. Afterwards, client apps can easily connect to it through the command line. to connect a client simply navigate to the `client/` directory and run the following command:

```bash
cd client/
go run .
```

This will run a client app and connects to the server app that you ran in the previous step. To run multiple clients, you need to start another terminal and simply repeat the process of running a client. When at least two clients are connected to the server the y can start chatting.

## Server Structure

When a server process is executed, a TCP server is started on `localhost:8080`. The main goroutine of the server handles the incomming connection requests and setups the chat logic. The chatting logic is implemented in the `chat` module. In this module, a goroutine is started for every connection which is responsible for receiving the incomming messeges from that connection, and passes them to a global messege dispatching goroutine which then sends the messege to every other user in the chatroom.

**Note:** Neither the users, nor the chat messeges are persisted to hard drive. The messeges are not saved in the RAM.

## Client Structure

The client process establishes a connection with the server, and spawns two goroutines. One resonsible for receiving input from the console and sending it over the established connection, while the other one listens for incomming messeges from the server and printing them to the terminal.
