package main

import (
    "tutorial/protobuf"
    "fmt"
    "net"
    "os"
    "encoding/csv"
    "io"
    "strconv"
    "github.com/golang/protobuf/proto"
)

const CLIENT_ID = 1
const CLIENT_NAME = "ClientOne"

func main() {

    fmt.Println("This is", CLIENT_NAME)
    filename := "input_one.csv"
    dest := "127.0.0.1:2000"

    data, err := readDataFromFile(&filename)
    processError(err)
    sendData(data, &dest)

}

func processError(err error) {

    if err != nil {
        fmt.Println("Fatal error : %s", err.Error() )
    }

}


func readDataFromFile(filename *string)([]byte, error) {

    file, err := os.Open(*filename)
    processError(err)
    defer file.Close()

    Message := new(protobuf.Message)
    Message.ClientId = *proto.Int32(CLIENT_ID)
    Message.ClientName = *proto.String(CLIENT_NAME)

    AddrBook := new(protobuf.AddressBook)
    reader := csv.NewReader(file)
    for {
        record, err := reader.Read()
        if err != io.EOF {
            processError(err)
        }else {
            break
        }

        Person := new(protobuf.Person)
        Phone := new(protobuf.Person_PhoneNumber)

        Person.Name = record[0]
        Person.Email = record[2]
        pId, err := strconv.Atoi(record[1])
        processError(err)
        Person.Id = *proto.Int32(int32(pId)) //Go's default int is 64bit but the int in .proto is 32bit. So need to convert to 32bit
        Phone.Number = record[3]
        phoneType, err := strconv.Atoi(record[4])
        processError(err)
        fmt.Println("phoneType:", phoneType)
        pType := protobuf.Person_PhoneType(phoneType)
        fmt.Println("type:", pType)
        Phone.Type =  pType

        Person.Phones = append(Person.Phones, Phone)
        AddrBook.People = append(AddrBook.People, Person)
    }

    Message.Book = AddrBook
    fmt.Println(Message.Book)
    return proto.Marshal(Message)

}

func sendData(data []byte, dest *string){
    conn, err := net.Dial("tcp", *dest)
    if err != nil {
        fmt.Println("Fatal error :", err.Error())
        return
    }
    cnt, err := conn.Write(data)
    processError(err)
    fmt.Println("Sent " + strconv.Itoa(cnt) + " bytes")
}

