package main

import (
	"github.com/20zinnm/blinky"
	"github.com/spf13/cobra"
	"time"
)

var rate time.Duration

var arduinoBlinkyCmd = &cobra.Command{
	Use:   "blinky",
	Short: "Flashes an LED light.",
	Long:  "Flashes an LED light on IO pin 13 at the given rate.",
	Run: func(cmd *cobra.Command, args []string) {
		blinky.Run(verbose, serial, rate)
	},
}

func init() {
	arduinoCmd.AddCommand(arduinoBlinkyCmd)

	arduinoBlinkyCmd.Flags().DurationVar(&rate, "rate", 500*time.Millisecond, `How often to toggle the LED.`)
}
