# yopen

A clone of the concept in github.com/brentp/xopen with a few other
changes

*  No bufio wrapper (let the caller wrap)

*  No http,piped-cmd,stdin check, or auto detection of gzip contents

*  Automatic re-attempt to open a file with .gz extension (specific to
   my usecase) if read on other file failed

*  Create files like rsync (where a dot-prefixed file is created first
   and then renamed to correct name at the end). This way you can be
sure the process wrote everything and the data is complete if the file
exists.
