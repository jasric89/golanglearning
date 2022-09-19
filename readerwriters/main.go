package main

import "io"

//func processData(reader io.Reader, writer io.Writer) { // This is a read write function
//	b := make([]byte, 2) // This Slices the Kayak input by 2 bytes each time b is called.
//	for {
//		count, err := reader.Read(b) // Read()b reads data into the specified byte in this case 2 bytes, The method will return the bytes that were read in this case KA and ya
//		if count > 0 {
//			writer.Write(b[0:count])
//			Printfln("Read %v bytes: %v", count, string(b[0:count]))
//		}
//		if err == io.EOF {
//			break
//		}
//	}
//}

//func processData(reader io.Reader, writer io.Writer) {
//	count, err := io.Copy(writer, reader)
//	if err == nil {
//		Printfln("Read %v bytes", count)
//	} else {
//		Printfln("Error: %v", err.Error())
//	}
//}

func main() {
	//Printfln("Product: %v, Price: %v", Kayak.Name, Kayak.Price) //%v is string formatting to print the value of a structure
	//r := strings.NewReader("Kayak")
	//var builder strings.Builder
	//processData(r, &builder)
	//Printfln("String builder contents: %s", builder.String())
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
		pipeWriter.Close()
	}()
	ConsumeData(pipeReader)
}
