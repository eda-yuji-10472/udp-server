// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"log"
	"net"
	"os"
)

func main() {
    udpAddr := &net.UDPAddr{
        IP:   net.ParseIP("0.0.0.0"),
        Port: 8080,
    }
    updLn, err := net.ListenUDP("udp", udpAddr)

    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }

    buf := make([]byte, 1024)
    log.Println("Starting udp server...")

    for {
        n, addr, err := updLn.ReadFromUDP(buf)
        if err != nil {
            log.Fatalln(err)
            os.Exit(1)
        }

        go func() {
            log.Printf("Reciving data: %s from %s", string(buf[:n]), addr.String())

            log.Printf("Sending data..")
            updLn.WriteTo([]byte("Pong"), addr)
            log.Printf("Complete Sending data..")
        }()
    }
}

