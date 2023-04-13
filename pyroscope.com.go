//go:build with.pyroscope

package main

func configurePyroscope() {
    // These 2 lines are only required if you're using mutex or block profiling
    // Read the explanation below for how to set these rates:
    runtime.SetMutexProfileFraction(5)
    runtime.SetBlockProfileRate(5)
    
    pyroscope.Start(pyroscope.Config{
	ApplicationName: "antlr.testrig",
	
	// replace this with the address of pyroscope server
	ServerAddress: "http://localhost:4040",
	
	// you can disable logging by setting this to nil
	Logger: pyroscope.StandardLogger,
	
	// optionally, if authentication is enabled, specify the API key:
	// AuthToken:    os.Getenv("PYROSCOPE_AUTH_TOKEN"),
	
	// you can provide static tags via a map:
	Tags: map[string]string{"hostname": "jimidle"},
	
	ProfileTypes: []pyroscope.ProfileType{
	    // these profile types are enabled by default:
	    pyroscope.ProfileCPU,
	    pyroscope.ProfileAllocObjects,
	    pyroscope.ProfileAllocSpace,
	    pyroscope.ProfileInuseObjects,
	    pyroscope.ProfileInuseSpace,
	    
	    // these profile types are optional:
	    pyroscope.ProfileGoroutines,
	    pyroscope.ProfileMutexCount,
	    pyroscope.ProfileMutexDuration,
	    pyroscope.ProfileBlockCount,
	    pyroscope.ProfileBlockDuration,
	},
    })
}

func dumpPprof() {
    pyroscope.Stop()
    f, _ := os.Create("heaptest.pprof")
    pprof.Lookup("heap").WriteTo(f, 0)
    _ = f.Close()
}
