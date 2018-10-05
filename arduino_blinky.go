package main

import (
	"github.com/20zinnm/blinky"
	"github.com/spf13/cobra"
	"time"
)

var rate time.Duration

var arduinoBlinkyCmd = &cobra.Command{
	Use:   "blinky",
	Short: "Toggles an LED light at a fixed rate.",
	Long:  `To run blinky, first ensure an LED is appropriately attached to an Arduino Uno on IO pin 13.`,
	Run: func(cmd *cobra.Command, args []string) {
		blinky.Run(verbose, serial, rate)
	},
}

func init() {
	arduinoCmd.AddCommand(arduinoBlinkyCmd)

	arduinoBlinkyCmd.Flags().DurationVar(&rate, "rate", 500*time.Millisecond, `How often to toggle the LED.`)
}
