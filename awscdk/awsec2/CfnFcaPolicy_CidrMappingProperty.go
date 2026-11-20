package awsec2


// A single CIDR range to label mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cidrMappingProperty := &CidrMappingProperty{
//   	CidrLabel: jsii.String("cidrLabel"),
//   	CidrRange: jsii.String("cidrRange"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-cidrmapping.html
//
type CfnFcaPolicy_CidrMappingProperty struct {
	// The label for the CIDR range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-cidrmapping.html#cfn-ec2-fcapolicy-cidrmapping-cidrlabel
	//
	CidrLabel *string `field:"required" json:"cidrLabel" yaml:"cidrLabel"`
	// The CIDR range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-cidrmapping.html#cfn-ec2-fcapolicy-cidrmapping-cidrrange
	//
	CidrRange *string `field:"required" json:"cidrRange" yaml:"cidrRange"`
}

