package main

import (
    "testing"
)

func Test_mvbcheck(t *testing.T) {
    
    // Does nothing without with.pyroscope build tag
    configurePyroscope()
    
    type args struct {
	inf string
    }
    tests := []struct {
	name string
	args args
    }{
	{
	    name: "Speed test,",
	    args: args{
		inf: "input",
	    },
	},
    }
    
    for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
	    
	    // Can add in trace options here
	    // antlr.ConfigureRuntime(antlr.WithParserATNSimulatorTraceATNSim(true))
	    
	    testRun(tt.args.inf)
	})
    }
    
    dumpPprof()
}
