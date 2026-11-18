package policy

import "time"

// dateLayout is the only form an acceptance expiry may take.
const dateLayout = "2006-01-02"

type Acceptance struct {
	Rule    string `toml:"rule"`
	Account string `toml:"account"`
	Reason  string `toml:"reason"`
	Expires *Date  `toml:"expires"`
}

// Date is a calendar day written as YYYY-MM-DD, the form an acceptance expiry
// uses. The type exists so the two questions a caller asks about an expiry, is it
// a date and has that day passed, are answered in one place instead of by every
// caller parsing the string itself.
type Date string

// Valid reports whether the date is a real calendar day in YYYY-MM-DD form.
func (d Date) Valid() bool {
	_, err := time.Parse(dateLayout, string(d))
	return err == nil
}

// Expired reports whether now is past the named day. The day itself is inclusive
// and the comparison is by UTC calendar day, so a clock ahead of or behind UTC can
// neither widen nor narrow the window. A value that is not a date has no last day
// to hold on, so it reads as expired: an expiry the loader would refuse must never
// keep an acceptance alive.
func (d Date) Expired(now time.Time) bool {
	if !d.Valid() {
		return true
	}
	return now.UTC().Format(dateLayout) > string(d)
}
