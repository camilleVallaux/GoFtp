package main

import (
	"bufio"
	"container/list"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/*********************************************/
/*                Server core                */
/*********************************************/

type piServer struct {
	root       string
	port       string
	maxClients int
	clients    list.List
	logFile    string
}

/*********************************************/
/*                  Server                   */
/*********************************************/

func (p *piServer) CheckUser(user string) bool {
	// TODO
	if user == "foo" {
		return true
	}
	return false
}

func (p *piServer) removeClient(c *client) {
	var next *list.Element
	for e := p.clients.Front(); e != nil; e = next {
		if e.Value == c {
			c.conn.Close()
			p.clients.Remove(e)
			logMsg("client disconnected")
			break
		}
	}
}

func (p *piServer) CheckUserPassword(c *client) bool {
	//TODO
	if c.user == "foo" && c.pass == "42" {
		return true
	}
	return false
}

func (p *piServer) ClientAuth(c *client) bool {

	if !c.send("220 Service ready") || !c.read() {
		goto closeClient
	}

	if c.getCmd() != "USER" {
		logError("bad sequence of Authentification")
		c.sendFatalError("503 Bad sequence of commands")
		goto closeClient
	}
	c.user = c.getValue()
	if !p.CheckUser(c.user) {
		logError("Bad user Authentification : \"" + c.user + "\"")
		c.sendFatalError("530 Not logged in")
		goto closeClient
	}
	if !c.send("331 User name ok, need password") || !c.read() {
		goto closeClient
	}

	if c.getCmd() != "PASS" {
		logError("bad sequence of Authentification")
		c.sendFatalError("503 Bad sequence of commands")
		goto closeClient
	}
	c.pass = c.getValue()
	if !p.CheckUserPassword(c) {
		logError("Bad user Authentification : \"" + c.user + "\"")
		c.sendFatalError("530 Not logged in")
		goto closeClient
	}

	if !c.send("230 User logged in") {
		goto closeClient
	}
	p.clients.PushFront(c)
	logMsg("New client accepted")
	return true

closeClient:
	c.conn.Close()
	return true
}

/*
 * Main client while
 */

func (p *piServer) HandleClient(c *client) {

	if !p.ClientAuth(c) {
		return
	}
	for {
		if c.read() != true {
			goto closeClient
		}
		cmd, ok := cmdMap[c.getCmd()]
		if !ok {
			cmdWrongCmd(p, c, c.getCmd()+" "+c.getValue())
			continue
		}
		if !cmd(p, c, c.getValue()) {
			break
		}
	}
closeClient:
	p.removeClient(c)
}

func (p *piServer) Start() {
	listener, error := net.Listen("tcp", p.port)
	if error != nil {
		panic(error)
	}
	logMsg("Starting Server...")
	for {
		conn, error := listener.Accept()
		if error != nil {
			logError("failed to accept new client")
			continue
		}

		c := &client{conn, "", "", false, [2]string{}, p.root, nil, ""}
		if p.clients.Len() == p.maxClients {
			logError("connection refuse, to many clients")
			c.sendFatalError("530 Not logged in")
			conn.Close()
			continue
		}
		c.conn.SetReadDeadline(time.Now().Add(10 * time.Minute))
		go p.HandleClient(c)
	}
}

func (p *piServer) stop() {

}

func newPiServer(confFile *string) *piServer {
	conf, error := os.Open(*confFile)
	if error != nil {
		panic(error)
	}

	p := piServer{}
	scanner := bufio.NewScanner(conf)
	for {
		eof := scanner.Scan()
		if error := scanner.Err(); error != nil {
			panic(error)
		}
		if !eof {
			break
		}
		kvalue := strings.Split(scanner.Text(), "=")
		switch kvalue[0] {
		case "port":
			if p.port != "" {
				panic("Invalid configuration file")
			}
			p.port = ":" + kvalue[1]
		case "root":
			if p.root != "" {
				panic("Invalid configuration file")
			}
			p.root = kvalue[1]
		case "maxClients":
			if p.maxClients != 0 {
				panic("Invalid configuration file")
			}
			p.maxClients, error = strconv.Atoi(kvalue[1])
			if error != nil {
				panic("Configuration file: invalid field")
			}
			p.maxClients = 1
		case "logFile":
			if p.logFile != "" {
				panic("Invalid configuration file")
			}
			p.logFile = kvalue[1]
		default:
			panic("Configuration file: invalid file")
		}
	}
	if p.port == "" {
		p.port = ":21"
	}
	if p.root == "" || p.logFile == "" {
		panic("Configuration file: Missing mandatory field")
	}
	if p.maxClients == 0 {
		p.maxClients = 1
	}
	logInit(p.logFile)
	return &p
}
