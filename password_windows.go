// +build windows,!linux,!darwin,!js

package dlgs

// Password displays a dialog, returning the entered value and a bool for success.
func Password(title, text string, optsA ...interface{}) (string, bool, error) {
	return editBox(title, text, "", "ClassPassword", true, optsA...)
}
