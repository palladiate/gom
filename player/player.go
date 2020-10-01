package player

import (
	"bufio"
	"log"
	"net"
)

type Player struct {
	conn net.Conn
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
}

func NewPlayer(conn net.Conn) Player {
	return Player{
		conn: conn,
		Char: Character{
			Name: "Some newb",
			Class: "Noob",
		},
	}
}

func (p Player) Play() {
	cont := true
	defer p.conn.Close()

	read := bufio.NewScanner(p.conn)
	write := bufio.NewWriter(p.conn)

	for cont {
		if read.Scan() {
			log.Printf("read %s\n", read.Text())
			c := parseCommand(read.Text())
			if c.Name == "Exit" {
				cont = false
			}
			_, err := write.WriteString(c.Emit)
			_, err = write.WriteString("\n")
			write.Flush()
			log.Printf("%s invoked %s", p.Char.Name, c.Name)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}

func parseCommand(l string) (c Command) {
	var found bool
	if c, found = abilities[l]; !found {
		c = abilities["error"]
	}
	return
}

