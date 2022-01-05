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


## CREATE

In order to create an Empty file, we can call a helper funtion called **Create**, which opens a new file with a 0666 permissions and truncates it if it doesnt exist. Alternatively, we can use **OpenFile** with the "O_CREATE|O_TRUNCATE" mode to specift custom permissions. 
For example:

```Go Language
package main
import "os"
func main() {
    f, err := os.Create("file.txt")
    if err != nil {
        fmt.Println("Error:",err)
        return
    }
    f.Close()
}
``` 


## TRUNCATE

To truncate the content of a file under a certain dimension, and leave the file untouched if it's smaller, there is the **os.Truncate** method. Its usage is pretty simple. 
For Example:

```Go Language
package main
import (
    "fmt"
    "os"
)
func main() {
    //Keep thing under 4kB
    if err := os.Truncate("file.txt",4096); err != nil {
        fmt.Println("Error:",err)
    }
}
```


## DELETE

In order to delete a file, there is another simple function, called **os.Remove**.
For Example:

```Go Language
package main
import "os"
func main() {
    if err := os.Remove("file.txt"); err != nil {
        fmt.Println("Error:",err)
    }
}
```

## MOVE

The **os.Rename** function makes it possible to change a file name and/or its directory. Note that this operation replaces the destination file if it already exists.
For Example:

```Go Language
package main
import "os"
func main() {
    if err := os.Rename("file.txt", "../file.txt"); err != nil { //file, dir/filename.ext
    fmt.Println("Error:", err)
    }
}
```


## COPY

There's no unique function that makes it possible to copy a file, but this can easily be done with a reader and a writer with the **io.Copy** function.
For Example:

```GO Language
func CopyFile(from, to string) (int64, error) {
    src, err := os.Open(from)
    if err != nil {
        return 0, err
    }
    defer src.Close()
    dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return 0, err
    }
    defer dst.Close()
    return io.Copy(dst, src)
}
```


## STATS

The **os** package offers the *Fileinfo* interface, which returns the metadata of a file.
The interface looks like:

```GO Language
type FileInfo interface {
    Name() string // base name of the file
    Size() int64 // length in bytes for regular files; system-dependent for others
    Mode() FileMode // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool // abbreviation for Mode().IsDir()
    Sys() interface{} // underlying data source (can return nil)
}
```
The **os.Stat** function returns information about the file with the specified path.


## CHANGING PROPERTIES

In order to interact with the filesystem and change these properties, three functions are available:
```diff
+1. func Chmod(name string, mode FileMode) return error  // Changes the permissions of a file.
+2. func Chown(name string, uid, gid int) return error   // Changes the owner and group of a file
+3. func Chtimes(name string, atime time.Time, mtime time.Time) return error // Changes the access and modification time of a file
```


## VIRTUAL FILESYSTEM

#### 1. vfs : github.com/blang/vfs
#### 2. afero : github.com/spf13/afero
