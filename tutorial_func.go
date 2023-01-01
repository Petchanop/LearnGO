/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tutorial_func.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/29 13:57:36 by npiya-is          #+#    #+#             */
/*   Updated: 2022/12/29 14:12:07 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	Test()
	Print_Hello("NiiNoi")
}

//everything in Go is pass by value
func Test() {
	fmt.Println("Hello this is my first function")
}

func Print_Hello(name string) {
	fmt.Printf("Hello, %s\n", name)
}
