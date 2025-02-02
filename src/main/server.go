/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"hello"
	"io"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hello.Hoge())
}

func main() {
	portNumber := "5000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
