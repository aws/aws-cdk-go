package awss3


// The PublicAccessBlock configuration that you want to apply to this Amazon S3 Multi-Region Access Point.
//
// You can enable the configuration options in any combination. For more information about when Amazon S3 considers an object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessBlockConfigurationProperty := &publicAccessBlockConfigurationProperty{
//   	blockPublicAcls: jsii.Boolean(false),
//   	blockPublicPolicy: jsii.Boolean(false),
//   	ignorePublicAcls: jsii.Boolean(false),
//   	restrictPublicBuckets: jsii.Boolean(false),
//   }
//
type CfnMultiRegionAccessPoint_PublicAccessBlockConfigurationProperty struct {
	// Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes the following behavior:
	//
	// - PUT Bucket ACL and PUT Object ACL calls fail if the specified ACL is public.
	// - PUT Object calls fail if the request includes a public ACL.
	// - PUT Bucket calls fail if the request includes a public ACL.
	//
	// Enabling this setting doesn't affect existing policies or ACLs.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Specifies whether Amazon S3 should block public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	//
	// Enabling this setting doesn't affect existing bucket policies.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.
	//
	// Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Specifies whether Amazon S3 should restrict public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` restricts access to this bucket to only AWS service principals and authorized users within this account if the bucket has a public policy.
	//
	// Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

