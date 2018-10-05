package main

import (
	"github.com/spf13/cobra"
)

var serial string

var arduinoCmd = &cobra.Command{
	Use:   "arduino",
	Short: "Collection of programs which use an Arduino Uno",
}

func init() {
	rootCmd.AddCommand(arduinoCmd)

	arduinoCmd.Flags().StringVarP(&serial, "serial", "s", "/dev/tty.usbmodem1421", `Path to serial port. By default, it is set to the correct serial path for macOS. Use 'gort scan serial' to find correct path for other systems.`)
}