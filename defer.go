/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   defer.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/29 15:20:25 by npiya-is          #+#    #+#             */
/*   Updated: 2022/12/29 15:24:33 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	defer Ex_defer()
	Ex_defer2()
}

func Ex_defer() {
	fmt.Println("test")
}

func Ex_defer2() {
	fmt.Println("defer")
}
