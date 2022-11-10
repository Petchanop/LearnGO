/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex3.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/07 18:49:48 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/07 19:20:33 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"io"
	"os"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%d %s %t\n", x, y, z)
	io.WriteString(os.Stdout, s)
	s = fmt.Sprintf("%d\n", x)
	io.WriteString(os.Stdout, s)
	s = fmt.Sprintf("%s\n", y)
	io.WriteString(os.Stdout, s)
	s = fmt.Sprintf("%t\n", z)
	io.WriteString(os.Stdout, s)
}
