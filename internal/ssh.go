package internal

import (
	"context"
	"log"
	"time"

	"github.com/melbahja/goph"
)

// LogError checkout th err
func LogError(err error, errType string) {
	if err != nil {
		if err == context.DeadlineExceeded && errType == "ssh" {
			// exec command timeout
			log.Fatalf("Runtime error: %v, detail: execute command timeout.", err)
		} else if err == context.DeadlineExceeded && errType == "scp" {
			// file transfer timeout
			log.Fatalf("Runtime error: %v, detail: Exceeded the maximum time to wait for the file transfer"+
				" to complete when using SCP. Please check if the remote file exists or check if timeout when file transfer?", err)
		} else {
			log.Fatalf("Runtime error: %v.", err)
		}
	}
}

// SshConn create a new ssh connection
func SshConn(config *goph.Config) *goph.Client {
	client, err := goph.NewConn(config)
	LogError(err, "")
	return client
}

// ExecRemoteCommand exec the command on the remote server
func ExecRemoteCommand(command string, timeout time.Duration, client *goph.Client) {
	log.Printf("Start executing command [%s]", command)
	bT := time.Now()
	// close client net connection when command is already done.
	defer func(client *goph.Client) {
		err := client.Close()
		LogError(err, "")
	}(client)
	if command != "" {
		ctx := context.Background()
		// set timeout param
		if timeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()
		}
		out, err := client.RunContext(ctx, command)
		log.Printf("-----------------------stdout && stderr---------------------->\n%s\n-------------------------------------------end----------------------------------->", string(out))
		LogError(err, "ssh")
		log.Printf("Execute command [%s] complete.", command)
		log.Printf("Total use time: %f s", time.Since(bT).Seconds())
	} else {
		log.Fatalf("Runtime error: command can't be empty.")
	}
}
