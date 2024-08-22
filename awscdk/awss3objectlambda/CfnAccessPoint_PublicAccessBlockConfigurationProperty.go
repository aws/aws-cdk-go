package awss3objectlambda


// The `PublicAccessBlock` configuration that you want to apply to this Amazon S3 account.
//
// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
//
// This data type is not supported for Amazon S3 on Outposts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessBlockConfigurationProperty := &PublicAccessBlockConfigurationProperty{
//   	BlockPublicAcls: jsii.Boolean(false),
//   	BlockPublicPolicy: jsii.Boolean(false),
//   	IgnorePublicAcls: jsii.Boolean(false),
//   	RestrictPublicBuckets: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-publicaccessblockconfiguration.html
//
type CfnAccessPoint_PublicAccessBlockConfigurationProperty struct {
	// Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account.
	//
	// Setting this element to `TRUE` causes the following behavior:
	//
	// - `PutBucketAcl` and `PutObjectAcl` calls fail if the specified ACL is public.
	// - PUT Object calls fail if the request includes a public ACL.
	// - PUT Bucket calls fail if the request includes a public ACL.
	//
	// Enabling this setting doesn't affect existing policies or ACLs.
	//
	// This property is not supported for Amazon S3 on Outposts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-publicaccessblockconfiguration.html#cfn-s3objectlambda-accesspoint-publicaccessblockconfiguration-blockpublicacls
	//
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Specifies whether Amazon S3 should block public bucket policies for buckets in this account.
	//
	// Setting this element to `TRUE` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	//
	// Enabling this setting doesn't affect existing bucket policies.
	//
	// This property is not supported for Amazon S3 on Outposts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-publicaccessblockconfiguration.html#cfn-s3objectlambda-accesspoint-publicaccessblockconfiguration-blockpublicpolicy
	//
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Specifies whether Amazon S3 should ignore public ACLs for buckets in this account.
	//
	// Setting this element to `TRUE` causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain.
	//
	// Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.
	//
	// This property is not supported for Amazon S3 on Outposts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-publicaccessblockconfiguration.html#cfn-s3objectlambda-accesspoint-publicaccessblockconfiguration-ignorepublicacls
	//
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Specifies whether Amazon S3 should restrict public bucket policies for buckets in this account.
	//
	// Setting this element to `TRUE` restricts access to buckets with public policies to only AWS-service principals and authorized users within this account.
	//
	// Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.
	//
	// This property is not supported for Amazon S3 on Outposts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-publicaccessblockconfiguration.html#cfn-s3objectlambda-accesspoint-publicaccessblockconfiguration-restrictpublicbuckets
	//
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

