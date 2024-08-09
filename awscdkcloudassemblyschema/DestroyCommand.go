package awscdkcloudassemblyschema


// Represents a cdk destroy command.
type DestroyCommand struct {
	// Whether or not to run this command as part of the workflow This can be used if you only want to test some of the workflow for example enable `synth` and disable `deploy` & `destroy` in order to limit the test to synthesis.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// This can be used in combination with `expectedError` to validate that a specific message is returned.
	// Default: - do not validate message.
	//
	ExpectedMessage *string `field:"optional" json:"expectedMessage" yaml:"expectedMessage"`
	// If the runner should expect this command to fail.
	// Default: false.
	//
	ExpectError *bool `field:"optional" json:"expectError" yaml:"expectError"`
	// Additional arguments to pass to the command This can be used to test specific CLI functionality.
	// Default: - only default args are used.
	//
	Args *DestroyOptions `field:"optional" json:"args" yaml:"args"`
}

