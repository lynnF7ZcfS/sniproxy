// Copyright (C) 2019 Antoine Tenart <antoine.tenart@ack.tf>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"
)

func (conn *Conn) logf(format string, v ...interface{}) {
	log.Printf("%s %s", conn.RemoteAddr(), fmt.Sprintf(format, v...))
}

func (conn *Conn) log(v ...interface{}) {
	log.Printf("%s %s", conn.RemoteAddr(), fmt.Sprint(v...))
}
