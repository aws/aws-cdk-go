package awsec2


// The current state of the instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateProperty := &StateProperty{
//   	Code: jsii.String("code"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-state.html
//
type CfnInstance_StateProperty struct {
	// The state of the instance as a 16-bit unsigned integer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-state.html#cfn-ec2-instance-state-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The current state of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-state.html#cfn-ec2-instance-state-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

