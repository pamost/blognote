+++
title = "Functions"
description = "Notes about Functions"
date = 2024-11-12

[author]
name = "pamost"
email = "pamost@yandex.ru"
+++

A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do, see the article on Go's declaration syntax.)

``` golang
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```