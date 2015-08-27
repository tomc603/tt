/*
   Copyright [yyyy] [name of copyright owner]

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
	"flag"
	"fmt"
	"os"
	// "time"
)

func main() {
	var (
		operation string
		project   string
		notes     string
	)

	// Parse the CLI flags. Only project and operation are required. If notes
	// is not specified, tt will prompt the user.
	flag.StringVar(&operation, "operation", "", "Specify the operation. [ pause | start | status | stop ]")
	flag.StringVar(&operation, "o", "", "Specify the operation. [ pause | start | status | stop ] (short)")
	flag.StringVar(&project, "project", "", "Specify the working project.")
	flag.StringVar(&project, "p", "", "Specify the working project. (short)")
	flag.StringVar(&notes, "notes", "", "Specify operation notes.")
	flag.StringVar(&notes, "n", "", "Specify operation notes. (short)")
	flag.Parse()

	if project == "" {
		fmt.Print("Error: You must specify a project\n")
		flag.Usage()
		os.Exit(2)
	}
	if operation == "" {
		fmt.Print("Error: You must specify an operation\n")
		flag.Usage()
		os.Exit(2)
	}

	fmt.Printf("Flags:\n  Operation: %s\n  Project: %s\n  Notes: %s\n", operation, project, notes)
}
