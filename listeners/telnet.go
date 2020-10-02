package listeners

import (
	"github.com/palladiate/gom/player"
	"log"
	"net"
	"strconv"
)

type Listener interface {
	Name() string
	Active() bool
	Start() error
	Stop() error
	Players() []player.Player
	PlayerChannel() chan *player.Player
}



type Telnet struct {
	Host string
	Port int
	listener net.Listener
	players chan *player.Player
}

func NewTelnet(host string, port int) (t Telnet) {
	return Telnet{
		Host: host,
		Port: port,
		players: make(chan *player.Player, 10),
	}
}

func (t *Telnet) Start() (err error) {
	t.listener, err = net.Listen("tcp", t.Host+":"+strconv.Itoa(t.Port))
	if err == nil {
		go func() {
			for {
				toon, err := t.listener.Accept()
				if err == nil {
					log.Print("New player!")
					p := player.NewPlayer(toon)
					t.players <- &p
				}
			}
		}()
	}
	return
}

func (t Telnet) Stop() error {
	return t.listener.Close()
}

func (t Telnet) Active() bool {
	return t.listener != nil
}

func (t Telnet) Name() string {
	return "Telnet"
}

func (t Telnet) Players() (p []player.Player) {
	return make([]player.Player, 0)
}

func (t Telnet) PlayerChannel() chan *player.Player {
	return t.players
}