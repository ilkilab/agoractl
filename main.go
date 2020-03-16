/*
Copyright © 2020 ILKI - ILKILABS agorakube@ilki.fr

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"./cmd"
	"os"
	"fmt"
)


func main() {
	user, present := os.LookupEnv("USER")
	sudo_user, is_sudoer := os.LookupEnv("SUDO_USER")
	_ = present
	_ = sudo_user
	if user == "root" || is_sudoer {
		cmd.Execute()
	} else {
		fmt.Println("Please, run this command as root or sudo")
	}
}
