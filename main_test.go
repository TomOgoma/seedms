package main

import (
	"testing"
	"time"
	"flag"
	// TODO SEEDMS change all references to github.com/tomogoma/seedms
	// to the new import path
	"github.com/tomogoma/seedms/config"
	confhelper "github.com/tomogoma/go-commons/config"
	"github.com/limetext/log4go"
)

func init() {
	flag.Parse()
}

func Test_bootstrap(t *testing.T) {
	conf := config.Config{}
	err := confhelper.ReadYamlConfig(*confFilePath, &conf)
	if err != nil {
		t.Fatalf("Error setting up: %s", err)
	}
	log := log4go.Logger{}
	srvErrCh := make(chan error)
	go func() {
		srvErrCh <- bootstrap(log, conf)
	}()
	timeoutTimer := time.NewTimer(1 * time.Second)
	select {
	case err := <-srvErrCh:
		t.Errorf("bootstrap(): %v", err)
	case <-timeoutTimer.C: // server started on time
	}
}

func Test_bootstrap_invalidConfig(t *testing.T) {
	conf := config.Config{}
	log := log4go.Logger{}
	srvErrCh := make(chan error)
	go func() {
		srvErrCh <- bootstrap(log, conf)
	}()
	timeoutTimer := time.NewTimer(1 * time.Second)
	select {
	case err := <-srvErrCh:
		if err == nil {
			t.Error("expected an error but got nil")
		}
	case <-timeoutTimer.C:
		t.Error("expected an error but got nil")
	}
}
