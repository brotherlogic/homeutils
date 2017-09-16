package homeutils

import (
	"context"
	"strconv"

	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
)

var (
	hostCache = ""
	hostPort  = -1
)

func getIP(servername string) (string, int32) {
	conn, _ := grpc.Dial("192.168.86.42:50055", grpc.WithInsecure())
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	entry := pbdi.RegistryEntry{Name: servername}
	r, err := registry.Discover(context.Background(), &entry)
	if err != nil {
		return "", -1
	}
	return r.Ip, r.Port
}

//GetConnection gets the client connection
func GetConnection(server string) (*grpc.ClientConn, error) {
	ip, port := getIP(server)
	return grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
}
