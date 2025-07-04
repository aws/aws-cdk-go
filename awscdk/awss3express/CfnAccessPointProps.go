package awss3express


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
//   	Scope: &ScopeProperty{
//   		Permissions: []*string{
//   			jsii.String("permissions"),
//   		},
//   		Prefixes: []*string{
//   			jsii.String("prefixes"),
//   		},
//   	},
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html
//
type CfnAccessPointProps struct {
	// The name of the bucket that you want to associate the access point with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html#cfn-s3express-accesspoint-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account ID that owns the bucket associated with this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html#cfn-s3express-accesspoint-bucketaccountid
	//
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// An access point name consists of a base name you provide, followed by the zoneID ( AWS Local Zone) followed by the prefix `--xa-s3` .
	//
	// For example, accesspointname--zoneID--xa-s3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html#cfn-s3express-accesspoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The access point policy associated with the specified access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html#cfn-s3express-accesspoint-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// Public access is blocked by default to access points for directory buckets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html#cfn-s3express-accesspoint-publicaccessblockconfiguration
	//
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// You can use the access point scope to restrict access to specific prefixes, API operations, or a combination of both.
	//
	// For more information, see [Manage the scope of your access points for directory buckets.](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-manage-scope.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html#cfn-s3express-accesspoint-scope
	//
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// If you include this field, Amazon S3 restricts access to this access point to requests from the specified virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html#cfn-s3express-accesspoint-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

