package cloudassemblyschema


// Represents a cdk command i.e. `synth`, `deploy`, & `destroy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdkCommand := &CdkCommand{
//   	Enabled: jsii.Boolean(false),
//   	ExpectedMessage: jsii.String("expectedMessage"),
//   	ExpectError: jsii.Boolean(false),
//   }
//
type CdkCommand struct {
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
}

