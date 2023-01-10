/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex5.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/05 18:10:45 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/05 18:44:14 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"math"
)

type square struct {
	width  float64
	length float64
}

type circle struct {
	radius float64
}

func (r square) area() float64 {
	area := r.width * r.length
	return area
}

func (r circle) area() float64 {
	area := (r.radius * r.radius) * math.Pi
	return area
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("square area :", s.area())
	case circle:
		fmt.Println("circle area :", s.area())
	}
}

func main() {
	fmt.Println("Ninja6 ex 5")
	sq := square{
		width:  5,
		length: 5,
	}

	ci := circle{
		radius: 5.0,
	}
	info(sq)
	info(ci)

}
