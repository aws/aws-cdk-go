package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceProperty := &ReferenceProperty{
//   	InstanceFamily: jsii.String("instanceFamily"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-reference.html
//
type CfnLaunchTemplate_ReferenceProperty struct {
	// The instance family to refer.
	//
	// Ensure that you specify the correct family name. For example, C6i and C6g are valid values, but C6 is not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-reference.html#cfn-ec2-launchtemplate-reference-instancefamily
	//
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
}

