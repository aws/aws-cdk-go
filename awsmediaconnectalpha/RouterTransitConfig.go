package awsmediaconnectalpha


// Output configuration to Router.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//   var existingRouterOutput RouterOutput
//
//
//   // Flow output to router with transit encryption
//   routerOutput := awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("RouterOutput"), &FlowOutputProps{
//   	Flow: flow,
//   	Output: awsmediaconnectalpha.OutputConfiguration_Router(&RouterTransitConfig{
//   		Encryption: &TransitEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
//   // Flow source from router with transit encryption
//   flowFromRouter := awsmediaconnectalpha.NewFlow(stack, jsii.String("FlowFromRouter"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Router(&RouterSource{
//   		RouterOutput: existingRouterOutput,
//   		Decryption: &TransitEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
// Experimental.
type RouterTransitConfig struct {
	// Specifies whether encryption is to be used.
	// Default: no encryption.
	//
	// Experimental.
	Encryption *TransitEncryption `field:"optional" json:"encryption" yaml:"encryption"`
}

