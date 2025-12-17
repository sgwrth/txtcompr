package constants

// Add new constants to AllOperarionArgs() below!
const compress = "-c"
const decompress = "-d"

func CompressArg() string {
	return compress
}

func DecompressArg() string {
	return decompress
}

func AllOperationArgs() []string {
	return []string{
		compress,
		decompress,
	}
}
