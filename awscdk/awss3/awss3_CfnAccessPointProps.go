package awss3


// Properties for defining a `CfnAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//   var policyStatus interface{}
//
//   cfnAccessPointProps := &cfnAccessPointProps{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	bucketAccountId: jsii.String("bucketAccountId"),
//   	name: jsii.String("name"),
//   	policy: policy,
//   	policyStatus: policyStatus,
//   	publicAccessBlockConfiguration: &publicAccessBlockConfigurationProperty{
//   		blockPublicAcls: jsii.Boolean(false),
//   		blockPublicPolicy: jsii.Boolean(false),
//   		ignorePublicAcls: jsii.Boolean(false),
//   		restrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		vpcId: jsii.String("vpcId"),
//   	},
//   }
//
type CfnAccessPointProps struct {
	// The name of the bucket associated with this access point.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `AWS::S3::AccessPoint.BucketAccountId`.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The name of this access point.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The access point policy associated with this access point.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The container element for a bucket's policy status.
	PolicyStatus interface{} `field:"optional" json:"policyStatus" yaml:"policyStatus"`
	// The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// The Virtual Private Cloud (VPC) configuration for this access point, if one exists.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

