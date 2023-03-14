package awss3outposts


// Properties for defining a `CfnAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnAccessPointProps := &cfnAccessPointProps{
//   	bucket: jsii.String("bucket"),
//   	name: jsii.String("name"),
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		vpcId: jsii.String("vpcId"),
//   	},
//
//   	// the properties below are optional
//   	policy: policy,
//   }
//
type CfnAccessPointProps struct {
	// The Amazon Resource Name (ARN) of the S3 on Outposts bucket that is associated with this access point.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of this access point.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The virtual private cloud (VPC) configuration for this access point, if one exists.
	VpcConfiguration interface{} `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// The access point policy associated with this access point.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

