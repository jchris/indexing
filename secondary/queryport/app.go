package queryport

import "github.com/couchbase/indexing/secondary/logging"
import c "github.com/couchbase/indexing/secondary/common"
import protobuf "github.com/couchbase/indexing/secondary/protobuf/query"
import "io"

// Application is example application logic that uses query-port server
func Application(config c.Config) {
	killch := make(chan bool)
	s, err := NewServer(
		"localhost:9990",
		func(req interface{},
			w io.Writer, quitch <-chan interface{}) {
			requestHandler(req, w, quitch, killch)
		},
		config)

	if err != nil {
		logging.Fatalf("Listen failed - %v", err)
	}
	<-killch
	s.Close()
}

// will be spawned as a go-routine by server's connection handler.
func requestHandler(
	req interface{},
	w io.Writer, // Write handle to the tcp socket
	quitch <-chan interface{}, // client / connection might have quit (done)
	killch chan bool, // application is shutting down the server.
) {

	var responses []*protobuf.ResponseStream

	switch req.(type) {
	case *protobuf.StatisticsRequest:
		// responses = getStatistics()
	case *protobuf.ScanRequest:
		// responses = scanIndex()
	case *protobuf.ScanAllRequest:
		// responses = fullTableScan()
	}

	buf := make([]byte, 1024, 1024)
loop:
	for _, resp := range responses {
		// query storage backend for request
		protobuf.EncodeAndWrite(w, buf, resp)
		select {
		case <-quitch:
			close(killch)
			break loop
		}
	}
	protobuf.EncodeAndWrite(w, buf, &protobuf.StreamEndResponse{})
	// Free resources.
}
