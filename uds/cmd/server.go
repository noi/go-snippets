package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		unixSocket := cmd.Flags().Lookup("unix-socket").Value.String()

		ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt)
		defer stop()

		mux := http.NewServeMux()
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("world"))
		})
		s := &http.Server{
			Handler: mux,
		}
		l, err := net.Listen("unix", unixSocket)
		if err != nil {
			return err
		}

		done := make(chan error, 1)
		go func() {
			done <- s.Serve(l)
		}()
		fmt.Println("listening:", unixSocket)
		select {
		case err := <-done:
			if !errors.Is(err, http.ErrServerClosed) {
				return err
			}
		case <-ctx.Done():
			if err := s.Shutdown(context.Background()); err != nil {
				return err
			}
			if err := <-done; !errors.Is(err, http.ErrServerClosed) {
				return err
			}
		}
		fmt.Println("closed")

		return nil
	},
}

func init() {
	serverCmd.Flags().String("unix-socket", "", "")
	rootCmd.AddCommand(serverCmd)
}
