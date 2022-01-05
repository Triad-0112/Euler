# DICTIONARY OF IO FUNCTION
## filepath.Function()
  List Function():
  ```CS
  - Abs(path string) return(string, error) : Returns the absolute version of the path thats pased by joining it to the current working directory (if it's not already absolute), and then cleans it.
  - Base(path string) return string : Gives the last element of the path (base). For example "path/to/some/file" returns the file. Note that if the path is empty, this functions returns a .(dot) path.
  - Clean(path string) return string :
  ```