package awscdk


// When AWS CloudFormation creates the associated resource, configures the number of required success signals and the length of time that AWS CloudFormation waits for those signals.
//
// Example:
//   var resource cfnResource
//
//
//   resource.CfnOptions.CreationPolicy = &CfnCreationPolicy{
//   	ResourceSignal: &CfnResourceSignal{
//   		Count: jsii.Number(3),
//   		Timeout: jsii.String("PR15M"),
//   	},
//   }
//
type CfnResourceSignal struct {
	// The number of success signals AWS CloudFormation must receive before it sets the resource status as CREATE_COMPLETE.
	//
	// If the resource receives a failure signal or doesn't receive the specified number of signals before the timeout period
	// expires, the resource creation fails and AWS CloudFormation rolls the stack back.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The length of time that AWS CloudFormation waits for the number of signals that was specified in the Count property.
	//
	// The timeout period starts after AWS CloudFormation starts creating the resource, and the timeout expires no sooner
	// than the time you specify but can occur shortly thereafter. The maximum time that you can specify is 12 hours.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

