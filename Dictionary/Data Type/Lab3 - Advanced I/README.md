# DICTIONARY OF IO FUNCTION
## filepath.Function()
  List Function():
  ```CS
  - Abs(path string) return(string, error) : Returns the absolute version of the path thats pased by joining it to the current working directory (if its not already absolute), and then cleans it.

  - Base(path string) return(string) : Gives the last element of the path (base). For example "path/to/some/file" returns the file. Note that if the path is empty, this functions returns a .(dot) path.

  - Clean(path string) return(string) : Returns the shortest version fo the path by applying a series of defines rules. It does operations like replacing .(dot) and .., or removing trailing separators.

  - Dir(path string) return(string) : Gets the path without its last element. This usuallya returns the parent directory of the element.

  - EvalSymlinks(path string) return(string, error): Returns the path after evaluating symbolic links. The path is relative if the provided path is also relative and doesnt contain symbolic links with absolute paths.

  - Ext(path string) return(string): Gets the file extension of the path, the suffix that starts with the final dot of the last element of the path, and its an empty string if theres no dot (for example, docs/file.txt returns .txt).

  - FromSlash(path string) return(string): Replaces all / (slashes) found in the path with the operative system path separator. This function does nothing if the OS is Windows, and it executes a replacement under Unix or macOS.

  - Glob(pattern string) return(matches []string, err error): Finds all files matching the specified pattern. If there are no matching files, the result is nil. It doesnt report eventual errors that occur during path exploration. It shares syntax with Match.

  - HasPrefix(p, prefix string) return(bool) : This function is deprecated.

  - IsAbs(path string) return(bool): Shows if the path is absolute or not.

  - Join(elem ...string) return(string): Concatenates multiple path elements by joining them with the filepath separator. Note that this also calls Clean on the result.

  - Match(pattern, name string) return(matched bool, err error): Verifies that the given name matches the pattern, allowing the use of the wild char characters * and ?, and groups or sequences of characters using square brackets.

  - Rel(basepath, targpath string) return(string, error): Returns the relative path from the base to the target path, or an error if this is not possible. This function calls Clean on the result.

  - Split(path string) return(dir, file string): Divides the path into two partsusing the final trailing slash. The result is usually the parent path and the file name of the input path. If there is no separator, dir will be empty and the file will be the path itself.

  - SplitList(path string) return([]string): Returns a list of paths, separating them with the list separator character, which is : in Unix and macOS and ; in Windows.

  - ToSlash(path string) return(string): Operates the opposite substitution that the FromSlash function executes, changing each path separator to a /, doing nothing on Unix and macOS, and executing the replacement in Windows.

  - VolumeName(path string) return(string): This does nothing in platforms that arent Windows. It returns the path component which refers to the volume. This is done for both local paths and network resources.

  - Walk(root string, walkFn WalkFunc) return(error): Starting from the root directory, this function travels recursively through the file tree,
  executing the walk function for each entry of the tree. If the walk function returns an error, the walk stops and that error is returned. The
  function is defined as follows:
      - type WalkFunc func(path string, info os.FileInfo, err error) return(error) Before moving on to the next example, lets introduce an important variable: os.Args. This variable contains at least one value, which is the path that invoked the current process. This can be followed by eventual arguments that are specified in the same call.

  ```