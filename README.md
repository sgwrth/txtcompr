This project aims to build a simple text compression CLI tool in Go. First, the text is tokenized.
Then tokens of at least two runes in length are replaced by bytes. The information on which byte
stands for which token is stored in a map. This map will then be used once more when decompressing
the text file.
