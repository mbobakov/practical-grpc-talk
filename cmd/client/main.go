package main

import (
	"context"
	"net/http"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jessevdk/go-flags"
	"github.com/mbobakov/practical-grpc-talk/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

func main() {
	var opts = struct {
		HTTPListen        string `long:"http.listen" env:"HTTP_LISTEN" default:":8080" description:"HTTP server interface"`
		DebugListen       string `long:"debug.listen" env:"DEBUG_LISTEN" default:":6060" description:"Interface for serve debug information(metrics/health/pprof)"`
		ConsulServiceName string `long:"consul.sname" env:"CONSUL_SERVICE_NAME" default:"time" description:"Consul service name"`
		SD                string `long:"sd.provider" env:"SD_PROVIDER" default:"direct" description:"Service discovery config provider" choice:"consul" choice:"direct"`
		Target            string `long:"target" env:"TARGET" default:"server:50051" description:"Connect to this target. If Service-discovery is consul - consul address"`
		Verbose           bool   `short:"v" env:"VERBOSE" description:"Enable Verbose log  output"`
	}{}

	_, err := flags.Parse(&opts)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}
	logger := logrus.New()

	if opts.Verbose {
		logger.SetLevel(logrus.DebugLevel)
	}

	grpclog.SetLoggerV2(&grpcLog{logger})

	conn, err := connect(opts.SD, opts.Target, opts.ConsulServiceName)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	client := api.NewTimeClient(conn)

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		_, err := client.CurrentDayLength(context.Background(), &empty.Empty{})
		st, ok := status.FromError(err)
		if !ok {
			logger.Errorf("Unknow grpc status: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if st.Code() != codes.OK {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + st.Message()))
			return
		}
	})

	err = http.ListenAndServe(opts.HTTPListen, nil)
	if err != nil {
		logger.Fatal(err)
	}
}

type grpcLog struct {
	*logrus.Logger
}

func (l *grpcLog) V(lvl int) bool {
	return true
}
