package awscdkcloudassemblyschema


// Definitions for the integration testing manifest.
type IntegManifest struct {
	// test cases.
	TestCases *map[string]*TestCase `field:"required" json:"testCases" yaml:"testCases"`
	// Version of the manifest.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Enable lookups for this test.
	//
	// If lookups are enabled
	// then `stackUpdateWorkflow` must be set to false.
	// Lookups should only be enabled when you are explicitely testing
	// lookups.
	// Default: false.
	//
	EnableLookups *bool `field:"optional" json:"enableLookups" yaml:"enableLookups"`
	// Additional context to use when performing a synth.
	//
	// Any context provided here will override
	// any default context.
	// Default: - no additional context.
	//
	SynthContext *map[string]*string `field:"optional" json:"synthContext" yaml:"synthContext"`
}

