![](https://image.cdn.wenqi.us/Apps/ObjectPacker/objectpacker-logo.png)  

[ObjectPacker](http://objectpacker.vinkdong.com) is a Data Packer for program use in pack bytes data to struct .

For more details about it please visit the official documentation:

http://vinkdong.com/objectpacker

## Quick Start 

```bash
go get github.com/ObjectPacker/objectpacker-go/pack
```

```go
type X struct {
	A float64
	B float64
	C interface{}
	D float64
}

func main() {

	p := X{A:20}

	b := []byte(`
{"A":99}
	`)
	pack.Packer(&b, &p)
	fmt.Println(p.A)
}

```


## License

This program is under the terms of the [Apache License v2.0](http://www.apache.org/licenses/LICENSE-2.0).

## Authors

[ObjectPacker](http://vinkdong.com/objectpacker) is made by [VinkDong](http://vinkdong.com) among other [contributors](https://github.com/ObjectPacker/objectpacker-go/graphs/contributors).
