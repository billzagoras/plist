package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"plist/server"
	"plist/utils"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const DEFAULT_PORT = "65535"

func initCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:                   "init [OPTIONS]",
		Short:                 "Initiate ApplicationServer",
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

			log.Println("server starts")
			server := server.ApplicationServer{}

			go func() {
				server.Setup()
				if err := server.Run(sanitizePort(os.Args)); utils.CheckErr(err) {
					log.Fatal("error in server running")
				}
			}()

			// Interrupt signal to shutdown the server gracefully.
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
			sig := <-quit

			var intSig syscall.Signal
			switch s := sig.(type) {
			case syscall.Signal:
				intSig = s
			default:
			}

			log.Printf("interrupt received: %d shutting down...\n", intSig)
			if err := server.Close(); utils.CheckErr(err) {
				log.Fatal("failed to close server")
			}
			log.Println("successfully closed server")
		},
	}

	return initCmd
}

// Sanitizing user's port input so that it is one of the private/dynamic ports based on RFC6335.
// Dynamic Ports, also known as the Private or Ephemeral Ports, from 49152-65535 are never assigned from transport protocols.
// That way, we ensure that our localhost application is will not used a potentially used port in the future.
func sanitizePort(args []string) string {
	err := godotenv.Load("../env/.env")
	if err != nil {
		log.Printf("could not load env file: %s\n", err)
	}

	port := os.Getenv("APP_PORT")
	if port != "" {
		return port
	}

	if len(args) >= 2 {
		port = args[1]
	}
	if port < "49152" || port > "65535" {
		port = DEFAULT_PORT
	}
	return port
}

func main() {
	if err := initCmd().Execute(); err != nil {
		log.Println("error during initCmd execution")
	}
}
