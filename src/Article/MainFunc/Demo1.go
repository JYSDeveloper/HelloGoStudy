package main

import "flag"

var name, pwd string

func init() {
	flag.StringVar(&name, "name", "default", "log in user")
	flag.StringVar(&pwd, "password", "", "log in user's password")

}
func main() {
	flag.Parse() //暂停获取参数
	println(name)
}
