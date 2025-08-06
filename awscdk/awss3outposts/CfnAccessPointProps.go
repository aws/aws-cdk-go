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
//   cfnAccessPointProps := &CfnAccessPointProps{
//   	Bucket: jsii.String("bucket"),
//   	Name: jsii.String("name"),
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		VpcId: jsii.String("vpcId"),
//   	},
//
//   	// the properties below are optional
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html
//
type CfnAccessPointProps struct {
	// The Amazon Resource Name (ARN) of the S3 on Outposts bucket that is associated with this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The virtual private cloud (VPC) configuration for this access point, if one exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// The access point policy associated with this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

