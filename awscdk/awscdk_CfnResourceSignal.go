// Version 2 of the AWS Cloud Development Kit library
package awscdk


// When AWS CloudFormation creates the associated resource, configures the number of required success signals and the length of time that AWS CloudFormation waits for those signals.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceSignal := &cfnResourceSignal{
//   	count: jsii.Number(123),
//   	timeout: jsii.String("timeout"),
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

