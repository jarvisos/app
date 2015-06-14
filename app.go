/*
   Copyright 2015 W. Max Lees

   This file is part of jarvisos.

   Jarvisos is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Jarvisos is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with jarvisos.  If not, see <http://www.gnu.org/licenses/>.

   File: app.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.14.2015
*/

package app

import (
	"net"
	"net/rpc"
)

type App interface {
	Call(string, *[]byte) error

	Who(bool, *[]byte) error
}

type Info struct {
	Port      string
	Functions []string
}

func Run(application App, address string) error {
	// Register the application with rpc
	rpc.Register(application)

	// Listen on the specified port
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// Listen for calls
	rpc.Accept(l)

	return nil
}
