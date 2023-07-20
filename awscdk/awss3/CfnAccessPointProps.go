package awss3


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
//
//   	// the properties below are optional
//   	BucketAccountId: jsii.String("bucketAccountId"),
//   	Name: jsii.String("name"),
//   	Policy: policy,
//   	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
//   		BlockPublicAcls: jsii.Boolean(false),
//   		BlockPublicPolicy: jsii.Boolean(false),
//   		IgnorePublicAcls: jsii.Boolean(false),
//   		RestrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html
//
type CfnAccessPointProps struct {
	// The name of the bucket associated with this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account ID associated with the S3 bucket associated with this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-bucketaccountid
	//
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The name of this access point.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The access point policy associated with this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-publicaccessblockconfiguration
	//
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// The Virtual Private Cloud (VPC) configuration for this access point, if one exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

