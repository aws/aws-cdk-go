// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Additional options for the blue/green deployment.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.additionalOptions} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenAdditionalOptions := &cfnCodeDeployBlueGreenAdditionalOptions{
//   	terminationWaitTimeInMinutes: jsii.Number(123),
//   }
//
type CfnCodeDeployBlueGreenAdditionalOptions struct {
	// Specifies time to wait, in minutes, before terminating the blue resources.
	TerminationWaitTimeInMinutes *float64 `field:"optional" json:"terminationWaitTimeInMinutes" yaml:"terminationWaitTimeInMinutes"`
}

