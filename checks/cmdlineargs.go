package checks

func IsArgPresent(args []string) bool {
	if len(args) < 2 {
		return false
	} else {
		return true
	}
}
