package player

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Player struct {
	conn net.Conn
	reader *bufio.Scanner
	writer *bufio.Writer
	Char Character
}

type Character struct {
	Name string
	Class string
}

type Command struct {
	Name string
	Short string
	Emit string
}

type Invocation string

type Commands map[string]Command

var abilities Commands = Commands{
	"exit": Command{
		Name: "Exit",
		Short: "exit",
		Emit: "Someone exited.",
	},
	"carp": Command{
		Name: "Carp",
		Short: "carp",
		Emit: "Hey a command!",
	},
	"error": Command{
		Name: "Unknown",
		Short: "unknown",
		Emit: "That won't work!",
	},
	"echo": Command{
		Name: "Echo",
		Short: "echo",
		Emit: "%s",
	},
}

func NewPlayer(conn net.Conn) Player {
	return Player{
		conn: conn,
		reader: bufio.NewScanner(conn),
		writer: bufio.NewWriter(conn),
		Char: Character{
			Name: "Some newb",
			Class: "Noob",
		},
	}
}

func (p Player) Play() {
	var err error
	cont := true
	defer p.conn.Close()

	for cont {
		if p.reader.Scan() {
			log.Printf("read %s\n", p.reader.Text())
			c, args := parseCommand(p.reader.Text())
			if c.Name == "Exit" {
				cont = false
			}
			err = p.Send(c.Emit, args...)
			if err != nil {
				log.Panic(err)
			}
			log.Printf("%s invoked %s", p.Char.Name, c.Name)
		}
	}
}

func parseCommand(l string) (c Command, args []string) {
	var found bool
	invocation := strings.Fields(l)
	switch {
	case len(invocation) == 0:
		c = abilities["error"]
		return
	case len(invocation) > 1:
		args = invocation[1:]
	}

	if c, found = abilities[invocation[0]]; !found {
		c = abilities["error"]
	}
	return
}

func (p Player) Send(s string, args ...string) (err error) {
	bytesWritten, err := fmt.Fprintf(p.writer, s+"\n")
	p.writer.Flush()
	log.Printf("%v bytes sent to %s", bytesWritten, p.Char.Name)
	return
}

func (c Command) GoString() (s string) {

	return
}


