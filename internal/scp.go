package internal

import (
	"context"
	"log"
	"os"
	"time"

	goScp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

func ScpConn(username string, password string, host string, port string, timeout time.Duration) *goScp.Client {
	scpConfig, _ := auth.PasswordKey(username, password, ssh.InsecureIgnoreHostKey())
	scpConfig.Timeout = timeout
	scpClient := goScp.NewClient(host+":"+port, &scpConfig)
	err := scpClient.Connect()
	LogError(err, "")
	return &scpClient
}

func CopyLocalFile2Remote(scpClient *goScp.Client, localFilePath string, remoteFilePath string, timeout time.Duration) {
	bT := time.Now()
	log.Printf("Start copy local file[%s] to remote server[%s - %s], please wait...", localFilePath, scpClient.Host, remoteFilePath)
	defer scpClient.Close()
	f, err := os.Open(localFilePath)
	LogError(err, "")
	defer func(f *os.File) {
		err := f.Close()
		LogError(err, "")
	}(f)
	ctx := context.Background()
	// set timeout param
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	err = scpClient.CopyFromFile(ctx, *f, remoteFilePath, "0777")
	LogError(err, "scp")
	log.Printf("Copy local file[%s] to remote server[%s - %s] done.", localFilePath, scpClient.Host, remoteFilePath)
	log.Printf("Total use time: %f s", time.Since(bT).Seconds())
}

func CopyRemoteFile2Local(scpClient *goScp.Client, localFilePath string, remoteFilePath string, timeout time.Duration) {
	bT := time.Now()
	log.Printf("Start copy remote file[%s - %s] to local[%s], please wait...", scpClient.Host, remoteFilePath, localFilePath)
	defer scpClient.Close()
	f, err := os.OpenFile(localFilePath, os.O_RDWR|os.O_CREATE, 0777)
	LogError(err, "")
	defer func(f *os.File) {
		err := f.Close()
		LogError(err, "")
	}(f)
	ctx := context.Background()
	// set timeout param
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	err = scpClient.CopyFromRemote(ctx, f, remoteFilePath)
	LogError(err, "scp")
	log.Printf("Copy remote file[%s - %s] to local[%s] done.", scpClient.Host, remoteFilePath, localFilePath)
	log.Printf("Total use time: %f s", time.Since(bT).Seconds())
}
