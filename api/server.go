/*
Copyright 2020 Elkhan Ibrahimov

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

package api

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartAPICmd bla
type StartAPICmd struct{}

// Run bla
func (s StartAPICmd) Run() error {
	gs := grpc.NewServer()
	RegisterTagServer(gs, &TagService{})
	RegisterStatServer(gs, &StatService{})
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":45555")
	if err != nil {
		//fmt.Println("Unable to listen 45555 port", err.Error())
		return err
	}
	if err := gs.Serve(l); err != nil {
		//fmt.Println("Couldn't serve accept incoming gRPC connections", err.Error())
		return err
	}

	return nil
}
