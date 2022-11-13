/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   operators_ex.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/12 16:05:59 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/12 16:10:03 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	x := 42 / 2
	fmt.Println("/ operators")
	fmt.Println("42 divide by 2 equal", x)

	x = x * 2
	fmt.Println("* operators")
	fmt.Println("21 multiply by 2 equal", x)
}
