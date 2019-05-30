/* Copyright 2019 mingzhuyang All Rights Reserved. */

/* log_test.go - log package test file */
/*
modification history
--------------------
2019-05-13   mingzhuyang   create


*/

package log

import "testing"

//imports

//Constants

//Typedefs

//Globals

//functions

func TestLogFactory_Build_Console(t *testing.T) {
	logConf := &LogConfiguration{
		Level:   "debug",
		Console: true,
		Service: "log_test",
	}
	factory := NewLogFactory(logConf)
	err := factory.Build()
	if err != nil {
		t.Fatal(err.Error())
	}
	Warn("这是测试log %s", "2233333")
}

func TestLogFactory_Build_File(t *testing.T) {
	logConf := &LogConfiguration{
		Level:   "debug",
		File:    true,
		Service: "log_test",
		FileConfig: &LogFileConfig{
			Name:  "test",
			Path:  "./log",
			Split: true,
		},
	}
	factory := NewLogFactory(logConf)
	err := factory.Build()
	if err != nil {
		t.Fatal(err.Error())
	}
	Warn("这是测试log %s", "2233333")
}

func init() {
	logConf := &LogConfiguration{
		Level:   "debug",
		File:    true,
		Console:true,
		Service: "log_benchmark_test",
		FileConfig: &LogFileConfig{
			Name:  "test",
			Path:  "./log",
			Split: true,
		},
	}
	factory := NewLogFactory(logConf)
	err := factory.Build()
	if err != nil {
		panic(err.Error())
	}
}

func BenchmarkDebug(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Debug("这是测试log %s，%d", "2233333", 1233333)
	}
}