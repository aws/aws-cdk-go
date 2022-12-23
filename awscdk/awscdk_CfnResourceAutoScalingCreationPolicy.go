// Version 2 of the AWS Cloud Development Kit library
package awscdk


// For an Auto Scaling group replacement update, specifies how many instances must signal success for the update to succeed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceAutoScalingCreationPolicy := &cfnResourceAutoScalingCreationPolicy{
//   	minSuccessfulInstancesPercent: jsii.Number(123),
//   }
//
type CfnResourceAutoScalingCreationPolicy struct {
	// Specifies the percentage of instances in an Auto Scaling replacement update that must signal success for the update to succeed.
	//
	// You can specify a value from 0 to 100. AWS CloudFormation rounds to the nearest tenth of a percent.
	// For example, if you update five instances with a minimum successful percentage of 50, three instances must signal success.
	// If an instance doesn't send a signal within the time specified by the Timeout property, AWS CloudFormation assumes that the
	// instance wasn't created.
	MinSuccessfulInstancesPercent *float64 `field:"optional" json:"minSuccessfulInstancesPercent" yaml:"minSuccessfulInstancesPercent"`
}

