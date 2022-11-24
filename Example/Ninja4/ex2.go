/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex2.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 12:20:32 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 12:25:50 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sl := make([]int, 10)
	for i := 0; i < 10; i++ {
		sl[i] = rand.Intn(100)
	}
	for _, val := range sl {
		fmt.Println(val)
	}
	fmt.Printf("sl Type is %T\n", sl)
}
