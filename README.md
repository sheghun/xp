There are two files.

`main_1.go` reads the whole file into memory sorts the strings and breaks the file into smaller parts of 500kb

`main.go` reads the file by maxfilesize parts sorts the string and  breaks it into smaller files of 500kb it repeats this whole process till it reads the whole file. 


How to run

`go run main.go --file ./test.txt`

or 

`go run main_1.go --file ./test.txt`

You can specify the max file size in mb using `--maxfilesize 2` it has to be an integer but it already has a default of 1mb
