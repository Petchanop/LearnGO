/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   func_methods.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/29 15:26:06 by npiya-is          #+#    #+#             */
/*   Updated: 2022/12/29 15:31:35 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretAgents struct {
	person
	ltk bool
}

func (agents secretAgents) speak() {
	fmt.Println("I am", agents.firstname, agents.lastname)
}

func main() {
	Agent_1 := secretAgents{
		person: person{
			"Petch",
			"Noppadon",
		},
		ltk: true,
	}
	fmt.Println(Agent_1)
	Agent_1.speak()
}
