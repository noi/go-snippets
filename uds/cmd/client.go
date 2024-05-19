package cmd

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use: "client",
	RunE: func(cmd *cobra.Command, args []string) error {
		unixSocket := cmd.Flags().Lookup("unix-socket").Value.String()
		url := args[0]

		transport := http.DefaultTransport.(*http.Transport).Clone()
		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		transport.DialContext = func(ctx context.Context, _, _ string) (net.Conn, error) {
			return dialer.DialContext(ctx, "unix", unixSocket)
		}
		client := &http.Client{
			Transport: transport,
			Timeout:   10 * time.Second,
		}

		res, err := client.Get(url)
		if err != nil {
			return err
		}
		defer func() {
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
		}()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Printf("%s", body)

		return nil
	},
}

func init() {
	clientCmd.Flags().String("unix-socket", "", "")
	rootCmd.AddCommand(clientCmd)
}
