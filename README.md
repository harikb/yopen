# yopen

A clone of the concept in github.com/brentp/xopen with a few other
changes

* No bufio wrapper (let the caller wrap)

* No http/url stuff, stdin check, or auto detection of gzip contents

* Automatic re-attempt to open a file with .gz extension (specific to
my usecase) if read on other file failed

* Implement piped-cmd support, but only for .lzo extension (no Go std lib support)

* Create files like rsync (where a dot-prefixed file is created first
 and then renamed to correct name at the end). This way you can be
 sure the process wrote everything and the data is complete if the file
 exists.
