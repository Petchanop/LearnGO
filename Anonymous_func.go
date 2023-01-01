/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Anonymous_func.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/31 16:44:23 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/01 17:21:09 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymous Func")
	}()
	//Anonymous func with variables
	func(num int) {
		fmt.Println("input number", num)
	}(42)
	//function expression
	/*assign function to variables*/
	f := func() {
		fmt.Println("this is function expression")
	}
	f()
	/*Returning function from function use case for when we
	want to return result from the function (http package)*/
	str := return_func2()
	fmt.Printf("%T\n", str)
	fmt.Println(str())
}

func return_func() string {
	s := "Return function"
	return s
}

func return_func2() func() string {
	return func() string {
		return return_func()
	}
}
