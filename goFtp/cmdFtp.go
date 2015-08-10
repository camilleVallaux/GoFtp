package main

import (
	"fmt"
	"path/filepath"
)

const (
	CTLF = "\r\n"
)

/*********************************************/
/*                 cmd tools                 */
/*********************************************/

func readDirectory(root string, path string, dir string) string {
	listing := filepath.Join(filepath.Dir(path), dir) + "/*"
	listing = filepath.Clean(listing)
	// FIXE ME
	files, _ := filepath.Glob(listing)
	var list string
	for i := range files {
		var file string
		file = filepath.Base(files[i])
		list += file + "\n"
		fmt.Println(file)
	}
	list += CTLF
	return list
}

/*********************************************/
/*                  ftp cmd                  */
/*********************************************/

type cmd func(*piServer, *client, string) bool

//ABOR, ACCT, ALLO, APPE, CWD, DELE, HELP, LIST, MODE, NLST, NOOP,
//PASS, PASV, PORT, QUIT, REIN, REST, RETR, RNFR, RNTO, SITE, STAT,
//STOR, STRU, TYPE, USER

var cmdMap = map[string]cmd{
	"ABOR": cmdAbor, "ACCT": cmdAcct, "ALLO": cmdAllo,
	"APPE": cmdAppe, "CWD": cmdCwd, "DELE": cmdDele,
	"HELP": cmdHelp, "LIST": cmdList, "MODE": cmdMode,
	"NLST": cmdNlst, "NOOP": cmdNoop, "PASS": cmdPass,
	"PASV": cmdPasv, "PORT": cmdPort, "QUIT": cmdQuit,
	"REIN": cmdRein, "REST": cmdRest, "RETR": cmdRetr,
	"RNFR": cmdRnfr, "RNTO": cmdRnto, "SITE": cmdSite,
	"STAT": cmdStat, "STOR": cmdStor, "STRU": cmdStru,
	"TYPE": cmdType, "USER": cmdUser,
}

func cmdAbor(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdAcct(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdAllo(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdAppe(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdCwd(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdDele(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdHelp(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdList(p *piServer, c *client, s string) bool {
	list := readDirectory(p.root, c.currentDir, s)
	c.send(list)
	return true
}

func cmdMode(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdNlst(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdNoop(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdPass(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdPasv(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdPort(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdQuit(p *piServer, c *client, s string) bool {
	return false
}

func cmdRein(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdRest(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdRetr(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdRnfr(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdRnto(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdSite(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdStat(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdStor(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdStru(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdType(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

func cmdUser(p *piServer, c *client, s string) bool {
	logError("command not implemented")
	return true
}

/********************************************************/

func cmdWrongCmd(p *piServer, c *client, s string) bool {
	logError("Wrong cmd:" + s)
	return true
}
