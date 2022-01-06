# HANDLING STREAMS


## THE BYTES READER

The **bytes** package contains a useful structure that treats a slice of bytes as an **io.Reader** interface, and it implements many more I/O interfaces:

```CS
1. io.Reader: This can act as a regular reader
2. io.ReaderAt: This makes it possible to read from a certain position onward
3. io.WriterTo: This makes it possible to write the contents with an offset
4. io.Seeker: This can move the readers cursor freely
5. io.ByteScanner: This can execute a read operation for each byte separately
6. io.RuneScanner: This can do the same with characters that are made of more bytes
```

There is also **bytes.Buffer**, which adds writing capabilities on top of bytes.Reader and makes it possible to access the underlying slice or get the content as a string.

The **Buffer.String** method converts bytes to string, and this type of casting in Go is done by making a copy of the bytes, because strings are immutable. This means that eventual changes to the buffer are made after the copy will
not propagate to the string.


## STRING READER

The strings package contains another structure that is very similar to the io.Reader interface, called strings.Reader. This works exactly like the first but the underlying value is a string instead of a slice of bytes.

One of the main advantages of using a string instead of the byte reader, when
dealing with strings that need to be read, is the avoidance of copying the data
when initializing it. This subtle difference helps with both performance and
memory usage because it does fewer allocations and requires the Garbage
Collector (GC) to clean up the copy.

