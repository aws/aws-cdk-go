package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateIpAddressSpecificationRequestProperty := &PrivateIpAddressSpecificationRequestProperty{
//   	Primary: jsii.Boolean(false),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-privateipaddressspecificationrequest.html
//
type CfnEC2Fleet_PrivateIpAddressSpecificationRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-privateipaddressspecificationrequest.html#cfn-ec2-ec2fleet-privateipaddressspecificationrequest-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-privateipaddressspecificationrequest.html#cfn-ec2-ec2fleet-privateipaddressspecificationrequest-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

