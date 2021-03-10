package main()

import "fmt"
"time"

// Structure definition 
type Student struct {    // type Structurename struct 
id int
name string
collage string
created time.Time
}

// main function 

func main(){

    // declare vairable
    var stu Student // var variablename Structurename
    newstu := new(Strudent)

    // set values
    stu.id = "038"
      stu.name = "Gokul" 
      stu.collage =  "IIITN"
      stu.created =  time.Now()  // it shows the time

      newstu.id = "038"
      newstu.name = "Gokul" 
      newstu.collage =  "IIITN"
      newstu.created =  time.Now()  // it shows the time


      // Display

      fmt.Println("++++++++++++++++++++++")
      fmt.Println(stu.id)
      fmt.Println(stu.name)
      fmt.Println(stu.collage)
      fmt.Println(stu.craeted)

      fmt.Println("++++++++++++++++++++++")
      fmt.Println(newstu.id)
      fmt.Println(newstu.name)
      fmt.Println(newstu.collage)
      fmt.Println(newstu.craeted)

}