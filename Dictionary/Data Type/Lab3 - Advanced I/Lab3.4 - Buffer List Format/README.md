# File Modes OS

Notes :
```
Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
```
```CS
1. O_RDONLY int = syscall.O_RDONLY // open the file read-only.
2. O_WRONLY int = syscall.O_WRONLY // open the file write-only.
3. O_RDWR int = syscall.O_RDWR // open the file read-write.

// The remaining values may be or'ed in to control behavior.

1. O_APPEND int = syscall.O_APPEND // append data to the file when writing.
2. O_CREATE int = syscall.O_CREAT // create a new file if none exists.
3. O_EXCL int = syscall.O_EXCL // used with O_CREATE, file must not exist.
4. O_SYNC int = syscall.O_SYNC // open for synchronous I/O.
5. O_TRUNC int = syscall.O_TRUNC // if possible, truncate file when opened.
```

Explanation :
```CS
O_APPEND "Before each write, the file offset is positioned at the end  of the file."
O_CREATE "Makes it possible to create the file if it doesn't exist."
O_EXCL "If this is used with create, it fails if the file already exists (exclusive creation)"
O_SYNC "Executes a read/write operation and verifies its competition."
O_TRUNC "If the file exists, its sizes is truncated to 0"
```
