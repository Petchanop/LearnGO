/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex6.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/06 16:08:04 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/06 16:20:08 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja 6 ex 6")
	ano := func() string {
		return "Hello I'm Anonymous func"
	}

	fmt.Println(ano())
}
