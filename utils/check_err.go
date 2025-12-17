package utils


func Check_err(err error){
	if err != nil{
		panic(err)
	}
}