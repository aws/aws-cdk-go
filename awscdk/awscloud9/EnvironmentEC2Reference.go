package awscloud9


// A reference to a EnvironmentEC2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentEC2Reference := &EnvironmentEC2Reference{
//   	EnvironmentEc2Arn: jsii.String("environmentEc2Arn"),
//   	EnvironmentEc2Id: jsii.String("environmentEc2Id"),
//   }
//
type EnvironmentEC2Reference struct {
	// The ARN of the EnvironmentEC2 resource.
	EnvironmentEc2Arn *string `field:"required" json:"environmentEc2Arn" yaml:"environmentEc2Arn"`
	// The Id of the EnvironmentEC2 resource.
	EnvironmentEc2Id *string `field:"required" json:"environmentEc2Id" yaml:"environmentEc2Id"`
}

