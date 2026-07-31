package awsmediaconnectalpha


// Options for Router Source.
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
type RouterSource struct {
	// Options to decrypt incoming feed.
	// Default: - no decryption.
	//
	// Experimental.
	Decryption *TransitEncryption `field:"optional" json:"decryption" yaml:"decryption"`
	// The router output that feeds this flow source.
	// Default: - no router output connected.
	//
	// Experimental.
	RouterOutput IRouterOutput `field:"optional" json:"routerOutput" yaml:"routerOutput"`
}

