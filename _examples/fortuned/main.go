package main

import (
	"bufio"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/CyCoreSystems/radius"
	"github.com/pborman/uuid"
)

// fortuned is a TCP service for service up fortunes

func main() {

	ln, err := net.Listen("tcp", "0.0.0.0:9191")
	if err != nil {
		panic(err)
	}

	rc := radius.NewClient(&radius.Options{
		Host:         "localhost:1813",
		SharedSecret: "testing123",
	})

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		id := uuid.New()
		h, _ := os.Hostname()
		session := rc.NewSession(id,
			radius.StringAttribute(radius.AccessUserName, "admin"),
			radius.StringAttribute(32, h), // NAS-Identifier
		)

		go handle(conn, session)
	}
}

func handle(conn io.ReadWriteCloser, session *radius.Session) {

	a := &acct{
		session: session,
		w:       conn,
		r:       conn,
		closer:  conn,
	}

	conn = a

	for {
		conn.Write([]byte(" > "))

		r := bufio.NewReader(a)
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)
		switch line {
		case "quit":
			session.Stop(radius.UserRequest)
			conn.Close()
			return
		case "fortune":
			cmd := exec.Command("fortune")
			out, err := cmd.Output()
			if err != nil {
				conn.Write([]byte("error running fortune: "))
				conn.Write([]byte(err.Error()))
				conn.Write([]byte("\n"))
			} else {
				conn.Write(out)
			}
		case "cowsay":
			cowsay := exec.Command("cowsay")
			cowsayInput, err := cowsay.StdinPipe()
			if err != nil {
				conn.Write([]byte("error running cowsay: "))
				conn.Write([]byte(err.Error()))
				conn.Write([]byte("\n"))
				return
			}
			cowsayOutput, err := cowsay.StdoutPipe()
			if err != nil {
				conn.Write([]byte("error running cowsay: "))
				conn.Write([]byte(err.Error()))
				conn.Write([]byte("\n"))
				return
			}

			cowsay.Start()

			fortune := exec.Command("fortune")
			body, err := fortune.Output()
			if err != nil {
				conn.Write([]byte("error running fortune: "))
				conn.Write([]byte(err.Error()))
				conn.Write([]byte("\n"))
				return
			}
			cowsayInput.Write(body)
			cowsayInput.Close()

			io.Copy(conn, cowsayOutput)
		default:
			conn.Write([]byte("Unknown command: "))
			conn.Write([]byte(line))
			conn.Write([]byte("\n"))
		}

		session.InterimUpdate()
	}
}
