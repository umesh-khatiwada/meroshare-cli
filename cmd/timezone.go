package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// timezoneCmd represents the timezone command
var timezoneCmd = &cobra.Command{
	Use:   "timezone [timezone]",
	Short: "Display the current date in a specified timezone",
	Long: `This command allows you to display the current date in a specified timezone.
You can specify the timezone as an argument, and optionally provide a date format using the --date flag.`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run:  runTimezoneCmd,
}

func init() {
	rootCmd.AddCommand(timezoneCmd)

	// Adding a flag to specify the date format
	timezoneCmd.Flags().String("date", "", "Specify the date format (default is RFC3339)")
}

func runTimezoneCmd(cmd *cobra.Command, args []string) {
	// runTimezoneCmd executes the timezone command
	// It retrieves the current date in the specified timezone and formats it according to the provided date flag.
	// If no date flag is provided, it defaults to RFC3339 format.

	timezone := args[0]
	location, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Printf("Error loading timezone '%s': %v\n", timezone, err)
		return
	}

	dateFlag, _ := cmd.Flags().GetString("date")
	var date string

	if dateFlag != "" {
		date = time.Now().In(location).Format(dateFlag)
	} else {
		date = time.Now().In(location).Format(time.RFC3339)[:10]
	}
	fmt.Printf("Current date in %v: %v\n", timezone, date)
}
