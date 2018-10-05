package main

import (
	"github.com/spf13/cobra"
)

var serial string

var arduinoCmd = &cobra.Command{
	Use:   "arduino",
	Short: "Programs which use an Arduino Uno",
	Long: `Programs are configured by default to run on macOS.
To change this, find the appropriate path to the serial port for your platform and specify it when invoking the command.`,
}

func init() {
	rootCmd.AddCommand(arduinoCmd)

	arduinoCmd.Flags().StringVar(&serial, "serial", "/dev/tty.usbmodem1421", `Path to serial port. By default, it is set to the correct serial path for macOS. Use 'gort scan serial' to find correct path for other systems.`)
}