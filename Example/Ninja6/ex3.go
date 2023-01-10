/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex3.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/05 16:50:38 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/05 17:56:00 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja6 ex3")
	defer bar()
	foo()
}

func foo() {
	defer func() {
		fmt.Println("Anonymous")
	}()
	fmt.Println("test defer")
}

func bar() {
	fmt.Println("Defer")
}
