package main

import (
    "fmt"
    "os"
    "net"
    "strconv"
    "encoding/csv"
    "tutorial/protobuf"
    "github.com/golang/protobuf/proto"
)


func main() {

   fmt.Println("Started server")
   src := "127.0.0.1:2000"
   c := make(chan *protobuf.Message)
   go func(){
       for{
        message := <-c
        writeDataToFile(message)
       }
   }()
   listener, err := net.Listen("tcp", src)
   if err != nil {
        fmt.Println("Fatal error :", err.Error())
        return
   }
   for{
       fmt.Println("Listening..")
       conn, err := listener.Accept()  //conn is stream-oriented connection
       if err == nil{ //err != nil means there is data to read on this tcp port
           go handleConnection(conn, c)
       }else {
            continue
       }

   }
}

func writeDataToFile(data *protobuf.Message) {
    ClientId := data.GetClientId()
    ClientName := data.GetClientName()

    file, err := os.OpenFile("output.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
    defer file.Close()
    if err != nil {
        fmt.Println("Error:", err.Error())
        return
    }

    AddrBook := data.GetBook()
    Persons := AddrBook.GetPeople()
    writer := csv.NewWriter(file)
    for _,person := range Persons {
        phones := person.GetPhones()
        for _,phone := range phones {

            record := []string{strconv.Itoa(int(ClientId)), ClientName, person.GetName(), strconv.Itoa(int(person.GetId())), person.GetEmail(), phone.GetNumber(), strconv.Itoa(int(phone.GetType()))}
            writer.Write(record)
            fmt.Println(record)
        }
    }

    writer.Flush()

}

func handleConnection(conn net.Conn, c chan *protobuf.Message) {

    fmt.Println("Reading data on tcp port")
    defer conn.Close()
    data := make([]byte, 4096)
    n, err := conn.Read(data)
    if err != nil {
        fmt.Println("Error:", err.Error())
        return
    }
    Message := new(protobuf.Message)
    fmt.Println("unmarshaled data:", data[0:n])
    err = proto.Unmarshal(data[0:n], Message)
    if err != nil {
        fmt.Println("Error:", err.Error())
        return
    }

    c <- Message

}
