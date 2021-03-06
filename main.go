package main // import "github.com/kidoda/quikalarm"

import (
	"github.com/spf13/pflag"
)

const Usage = `Usage:	quikalarm [options...] [arguments]
Simple alarm clock with few options.

Mandatory arguments to long options are mandatory for short options too.

Options:
	-z, --snooze-time 	set snooze duration; The snooze prompt will appear once the alert buzzer starts; it defaults to 'YES', so just pressing 'ENTER' will snooze for this amount of time.

	-s, --set		set the alarm wake-up time; argument in 12h or 24h format, in the form [hour:minute AM/PM]
(e.g., 9:00AM, 22:00, or 08:30PM)

	-t, --timer		timer mode; in this mode the arguments given to -s, --set will be treated as a duration instead of an absolute time.`

func init() {
	// p *string, name string, shorthand string, value string, usage string
	pflag.StringVarP(&COLOR, "color", "c", "red", "set color of the digital display")
	pflag.Parse()
}

func main() {

}
