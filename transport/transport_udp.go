package transport

/*
import (
	"github.com/stefankopieczek/gossip/base"
	"github.com/stefankopieczek/gossip/log"
	"github.com/stefankopieczek/gossip/parser"
)

import (
	"net"
)

type Udp struct {
	laddr  *net.UDPAddr
	in     *net.UDPConn
	output chan base.SipMessage
}

func NewUdp(output chan base.SipMessage) (udp *Udp, err error) {
	var laddr *net.UDPAddr
	laddr, err = net.ResolveUDPAddr("udp", address)
	if err == nil {
		udp = &Udp{laddr: laddr, output: output}
	}

	return
}

func (udp *Udp) Listen(parser parser.Parser) error {
	var err error = nil
	udp.in, err = net.ListenUDP("udp", udp.laddr)

	if err == nil {
		go udp.listen(parser)
	}

	return err
}

func (udp *Udp) IsStreamed() bool {
	return false
}

func (udp *Udp) Send(addr string, msg base.SipMessage) error {
	log.Debug("Sending message %s to %s", msg.Short(), addr)
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	log.Debug("About to dial")
	var conn *net.UDPConn
	conn, err = net.DialUDP("udp", udp.laddr, raddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Debug("About to write")
	_, err = conn.Write([]byte(msg.String()))
	log.Debug("Wrote")

	return err
}

func (udp *Udp) listen(parser parser.Parser) {
	log.Info("Begin listening for UDP on address " + udp.laddr.String())

	buffer := make([]byte, c_BUFSIZE)
	for {
		num, _, err := udp.in.ReadFromUDP(buffer)
		if err != nil {
			log.Severe("Failed to read from UDP buffer: " + err.Error())
			continue
		}

		pkt := append([]byte(nil), buffer[:num]...)
		parser.Write(pkt)
	}
}

func (udp *Udp) Stop() {
	udp.in.Close()
}
*/
