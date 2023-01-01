package cmd

import (
	"astools/internal"
	"flag"
	"fmt"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	host           string
	user           string
	pass           string
	port           uint
	cmd            string
	localFilePath  string
	remoteFilePath string
	scpType        string
	timeout        uint
	execType       string
)

func init() {

	flag.StringVar(&host, "ip", "", "machine ip address.")
	flag.StringVar(&user, "user", "", "ssh username or scp username.")
	flag.StringVar(&pass, "pass", "", "ssh password or scp password.")
	flag.UintVar(&port, "port", 22, "ssh port number or scp port number.")
	flag.StringVar(&cmd, "cmd", "", "command to run.")
	flag.StringVar(&localFilePath, "local-path", "", "local file path when use SCP.")
	flag.StringVar(&remoteFilePath, "remote-path", "", "remote file path when use SCP.")
	flag.StringVar(&scpType, "scp-type", "upload", "upload: upload local file to remote server; download: download remote file to local.")
	flag.UintVar(&timeout, "timeout", 60, "execute command's timeout(s) when use ssh or "+
		"file transfer's timeout(s) when use SCP. (0 means no timeout(s), not recommended!).")
	flag.StringVar(&execType, "type", "", "ssh: execute a command with ssh; scp: file transfer with SCP.")
}

func showTips() {
	log.Printf("[error] Parse parmas error, Please check your input params!")
	fmt.Printf("Welcome to astools, you can type ./astools_android_arm64 -h to show help " +
		"message.\nUsages:\n1. To execute a command on remote server with ssh:\n\t./astools_android_arm64" +
		" -type ssh -ip 192.168.124.16 -user ubuntu -pass 123456 -port 22 -timeout 5 -cmd \"pwd\"\n" +
		"2. To upload local file to remote server:\n\t./astools_android_arm64 -type scp -ip 192.168.124.16 -user ubuntu" +
		" -pass 123456 -port 22 -timeout 5 -scp-type upload -local-path ./demo.txt -remote-path /home/ubuntu/demo1.txt\n" +
		"3. To download remote file to local:\n\t./astools_android_arm64 -type scp -ip 192.168.124.16 -user ubuntu " +
		"-pass 123456 -port 22 -timeout 5 -scp-type download -local-path ./demo2.txt -remote-path /home/ubuntu/demo1.txt\n")
	os.Exit(4)
}

func checkScpParams(host string, user string, pass string, localPath string, remotePath string) bool {
	if strings.Trim(host, " ") != "" && strings.Trim(user, " ") != "" && strings.Trim(pass, " ") != "" && strings.Trim(localPath, " ") != "" && strings.Trim(remotePath, " ") != "" {
		return true
	}
	return false
}

func Execute() {
	flag.Parse()
	if execType == "ssh" {
		if strings.Trim(host, " ") != "" && strings.Trim(user, " ") != "" && strings.Trim(pass, " ") != "" && strings.Trim(cmd, " ") != "" {
			sshConfig := &goph.Config{
				User:     user,
				Addr:     host,
				Auth:     goph.Password(pass),
				Port:     port,
				Callback: ssh.InsecureIgnoreHostKey(),
			}
			client := internal.SshConn(sshConfig)
			internal.ExecRemoteCommand(cmd, time.Second*time.Duration(timeout), client)
		} else {
			showTips()
		}
	} else if execType == "scp" {
		if scpType == "upload" {
			if checkScpParams(host, user, pass, localFilePath, remoteFilePath) {
				scpClient := internal.ScpConn(user, pass, host, strconv.Itoa(int(port)), time.Second*60)
				internal.CopyLocalFile2Remote(scpClient, localFilePath, remoteFilePath, time.Second*time.Duration(timeout))
			} else {
				showTips()
			}
		} else if scpType == "download" {
			if checkScpParams(host, user, pass, localFilePath, remoteFilePath) {
				scpClient := internal.ScpConn(user, pass, host, strconv.Itoa(int(port)), time.Second*60)
				internal.CopyRemoteFile2Local(scpClient, localFilePath, remoteFilePath, time.Second*time.Duration(timeout))
			} else {
				showTips()
			}
		} else {
			showTips()
		}
	} else {
		log.Printf("[error] -type param must be `ssh` or `scp`, not null.")
		showTips()
	}
}
