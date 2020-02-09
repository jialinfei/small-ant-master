package main

import "small-ant-parent/small-ant-web/routers"

func main() {
	//http:localhost:8080 访问login.html
	routers.Run("")
	select {}
}
