package awsmediaconnectalpha


// Properties for flow output.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//
//
//   // SRT Caller output with encryption
//   output := awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("EncryptedOutput"), &FlowOutputProps{
//   	Flow: flow,
//   	Description: jsii.String("Encrypted SRT output"),
//   	Output: awsmediaconnectalpha.OutputConfiguration_SrtCaller(&SrtCallerOutputConfig{
//   		Destination: jsii.String("203.0.113.100"),
//   		Port: jsii.Number(7000),
//   		Encryption: &SrtPasswordEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
// Experimental.
type FlowOutputProps struct {
	// The flow this output is attached to.
	// Experimental.
	Flow IFlow `field:"required" json:"flow" yaml:"flow"`
	// Output configuration.
	// Experimental.
	Output OutputConfiguration `field:"required" json:"output" yaml:"output"`
	// A description of the output.
	//
	// This description appears only on the MediaConnect console and will not be seen by the end user.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the output.
	// Default: - auto-generated.
	//
	// Experimental.
	FlowOutputName *string `field:"optional" json:"flowOutputName" yaml:"flowOutputName"`
	// An indication of whether the output should transmit data or not.
	// Default: - undefined; when omitted, MediaConnect enables the output (ENABLED) at deploy time.
	//
	// Experimental.
	OutputStatus State `field:"optional" json:"outputStatus" yaml:"outputStatus"`
}

