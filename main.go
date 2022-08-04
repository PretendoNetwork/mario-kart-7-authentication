package main

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	nexauth "github.com/PretendoNetwork/nex-protocols-common-go/authentication"
)

var nexServer *nex.Server

func main() {
	nexServer = nex.NewServer()
	nexServer.SetPrudpVersion(0)
	nexServer.SetNexVersion(2)
	nexServer.SetKerberosKeySize(32)
	nexServer.SetAccessKey("6181dff1")
	nexServer.SetPingTimeout(20)
	nexServer.SetKerberosPassword("test")

	nexServer.On("Data", func(packet *nex.PacketV0) {
		request := packet.RMCRequest()

		fmt.Println("==MK8 - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	authenticationServer := nexauth.NewCommonAuthenticationProtocol(nexServer)
	authenticationServer.SetSecureStationURL(nex.NewStationURL("prudps:/address=154.51.186.148;port=61001;CID=1;PID=2;sid=1;stream=10;type=2"))
	authenticationServer.SetBuildName("Pretendo MK7")
	authenticationServer.PasswordFromPID(getNEXAccountByPID)
	_ = authenticationServer

	nexServer.Listen(":61000")
}
