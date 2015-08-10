package main

import (
	"net"
	"strings"
)

type client struct {
	conn       net.Conn
	user       string
	pass       string
	isAdmin    bool
	lastCmd    [2]string
	currentDir string
	dataCon    *net.Conn
	dataPort   string
}

func (c *client) read() bool {
	buffer := make([]byte, 1024)
	bytesRead, error := c.conn.Read(buffer)
	if bytesRead == 0 {
		logMsg("client disconnecterd")
		return false
	}
	if error != nil {
		logError("socket error")
		return false
	}

	cmd := string(buffer[0:bytesRead])
	cmd = strings.TrimSuffix(cmd, CTLF)
	split_cmd := strings.SplitN(cmd, " ", 2)

	c.lastCmd[0] = split_cmd[0]
	if len(split_cmd) > 1 {
		c.lastCmd[1] = split_cmd[1]
	} else {
		c.lastCmd[1] = ""
	}
	return true
}

func (c *client) getCmd() string {
	return c.lastCmd[0]
}

func (c *client) getValue() string {
	return c.lastCmd[1]
}

func (c *client) send(msg string) bool {
	bytesSend, err := c.conn.Write([]byte(msg + CTLF))
	if bytesSend == 0 {
		logMsg("client disconnecterd")
		return false
	}
	if err != nil {
		logError("socket error")
		return false
	}
	return true
}

func (c *client) sendFatalError(msg string) {
	_ = c.send(msg)
}
