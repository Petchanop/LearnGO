/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   pointer.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/10 13:02:23 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/10 13:28:55 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is pointer ex")
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 43
	fmt.Println(b)
	fmt.Println(*b)

	x := 0
	foo(&x)
	fmt.Println(x)
}

func foo(y *int) {
	fmt.Println(y)
	*y = 44
	fmt.Println(y)
}
