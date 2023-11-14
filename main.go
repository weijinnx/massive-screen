package main

import (
	"encoding/json"
	"errors"
	"log"
	"massive-screen/internal/services/clusters"
	"net/http"
)

var (
	svc clusters.Service
)

func init() {
	svc = clusters.New("clusters")
}

func main() {
	statSrv := http.NewServeMux()
	statSrv.HandleFunc("/servers/stat", ServersStat)

	loadSrv := http.NewServeMux()
	loadSrv.HandleFunc("/load", LoadCluster)

	go func() {
		log.Println("Server 'Load' started on: http://localhost:2001")
		_ = http.ListenAndServe("localhost:2001", loadSrv)
	}()

	log.Println("Server 'Stats' started on: http://localhost:2000")
	_ = http.ListenAndServe("localhost:2000", statSrv)
}

// LoadCluster handles GET /load endpoint
func LoadCluster(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var b []byte
	f := func() error {
		if !r.URL.Query().Has("server") {
			return errors.New("server address have to be provided")
		}

		res, err := svc.Load(&clusters.LoadRequest{
			Addr: r.URL.Query().Get("server"),
		})
		if err != nil {
			return err
		}

		b, err = json.Marshal(res)
		if err != nil {
			return err
		}

		return nil
	}
	if err := f(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// ServersStat handles GET /servers/stat endpoint
func ServersStat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var b []byte
	f := func() error {
		if !r.URL.Query().Has("servers") {
			return errors.New("servers addresses have to be provided")
		}

		res, err := svc.Stat(&clusters.StatRequest{
			Servers: r.URL.Query()["servers"],
		})
		if err != nil {
			return err
		}

		b, err = json.Marshal(res)
		if err != nil {
			return err
		}

		return nil
	}
	if err := f(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
