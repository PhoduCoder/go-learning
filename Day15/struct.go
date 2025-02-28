type SliceHeader struct {
    Length int
    ZerothElement *byte
}

slice := SliceHeader{
    Length: len(slice)-2,
    ZerothElement: &buffer[101],
}