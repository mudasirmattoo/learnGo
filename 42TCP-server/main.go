package main

import (
	"fmt"
	"log"
	"net"
)

// main starts a TCP server listening on port 9000.
func main() {
	// net.Listen sets up a TCP listener on the specified address (here, port 9000).
	listener, err := net.Listen("tcp", ":9000") //func Listen(network, address string) (Listener, error)
	//The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
	//

	/*
			type Listener interface {
			// Accept waits for and returns the next connection to the listener.
			Accept() (Conn, error)

			// Close closes the listener.
			// Any blocked Accept operations will be unblocked and return errors.
			Close() error

			// Addr returns the listener's network address.
			Addr() Addr
		}
	*/
	if err != nil {
		log.Fatalf("error starting TCP server: %v ", err)
	}

	// Ensure the listener is closed when main exits.
	defer listener.Close()
	log.Println("TCP Server listening on port 9000")

	// Continuously accept new TCP connections.
	for {

		conn, err := listener.Accept() //Waits for and returns the next connection to the listener.
		if err != nil {
			log.Printf("error accepting connection : %v", err)
			continue
		}

		go handleConnection(conn)

	}
}

// handleConnection manages the lifecycle of an individual TCP connection.

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store incoming data (up to 1024 bytes).
	buffer := make([]byte, 1024)

	// Read data from the connection.
	n, err := conn.Read(buffer) //n: The number of bytes actually read into the buffer.

	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}

	log.Printf("Received: %s", string(buffer[:n]))

	response := "Hello from TCP server!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
	}
	fmt.Println("Response sent to client.")

}

/*
GO DOCUMENTATION

type Conn interface {
	// Read reads data from the connection.
	// Read can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetReadDeadline.
	Read(b []byte) (n int, err error)

	// Write writes data to the connection.
	// Write can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetWriteDeadline.
	Write(b []byte) (n int, err error)

	// Close closes the connection.
	// Any blocked Read or Write operations will be unblocked and return errors.
	Close() error

	// LocalAddr returns the local network address, if known.
	LocalAddr() Addr

	// RemoteAddr returns the remote network address, if known.
	RemoteAddr() Addr

	// SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	//
	// A deadline is an absolute time after which I/O operations
	// fail instead of blocking. The deadline applies to all future
	// and pending I/O, not just the immediately following call to
	// Read or Write. After a deadline has been exceeded, the
	// connection can be refreshed by setting a deadline in the future.
	//
	// If the deadline is exceeded a call to Read or Write or to other
	// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// The error's Timeout method will return true, but note that there
	// are other possible errors for which the Timeout method will
	// return true even if the deadline has not been exceeded.
	//
	// An idle timeout can be implemented by repeatedly extending
	// the deadline after successful Read or Write calls.
	//
	// A zero value for t means I/O operations will not time out.
	SetDeadline(t time.Time) error

	// SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	// A zero value for t means Read will not time out.
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means Write will not time out.
	SetWriteDeadline(t time.Time) error
}
Conn is a generic stream-oriented network connection.

Multiple goroutines may invoke methods on a Conn simultaneously.

func Dial ¶
func Dial(network, address string) (Conn, error)
Dial connects to the address on the named network.
*/
