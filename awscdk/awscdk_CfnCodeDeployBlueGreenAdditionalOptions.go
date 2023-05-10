// An experiment to bundle the entire CDK into a single module
package awscdk


// Additional options for the blue/green deployment.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.additionalOptions} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenAdditionalOptions := &CfnCodeDeployBlueGreenAdditionalOptions{
//   	TerminationWaitTimeInMinutes: jsii.Number(123),
//   }
//
// Experimental.
type CfnCodeDeployBlueGreenAdditionalOptions struct {
	// Specifies time to wait, in minutes, before terminating the blue resources.
	// Experimental.
	TerminationWaitTimeInMinutes *float64 `field:"optional" json:"terminationWaitTimeInMinutes" yaml:"terminationWaitTimeInMinutes"`
}

