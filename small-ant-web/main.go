package main

import "small-ant-parent/small-ant-web/routers"

func main() {
	routers.Run("")
	select {}
}
