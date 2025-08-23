package main

func main() {
	var p *int32 = new(int32)
	*p = 42
	println(p, *p)
	var a int32 = 27
	p = &a
	println(p, *p)
	*p = 42
	println(p, *p)

}
