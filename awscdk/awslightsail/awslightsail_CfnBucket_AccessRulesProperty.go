package awslightsail


// `AccessRules` is a property of the [AWS::Lightsail::Bucket](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html) resource. It describes access rules for a bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessRulesProperty := &accessRulesProperty{
//   	allowPublicOverrides: jsii.Boolean(false),
//   	objectAccess: jsii.String("objectAccess"),
//   }
//
type CfnBucket_AccessRulesProperty struct {
	// A Boolean value indicating whether the access control list (ACL) permissions that are applied to individual objects override the `GetObject` option that is currently specified.
	//
	// When this is true, you can use the [PutObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectAcl.html) Amazon S3 API operation to set individual objects to public (read-only) or private, using either the `public-read` ACL or the `private` ACL.
	AllowPublicOverrides interface{} `field:"optional" json:"allowPublicOverrides" yaml:"allowPublicOverrides"`
	// Specifies the anonymous access to all objects in a bucket.
	//
	// The following options can be specified:
	//
	// - `public` - Sets all objects in the bucket to public (read-only), making them readable by everyone on the internet.
	//
	// If the `GetObject` value is set to `public` , then all objects in the bucket default to public regardless of the `allowPublicOverrides` value.
	// - `private` - Sets all objects in the bucket to private, making them readable only by you and anyone that you grant access to.
	//
	// If the `GetObject` value is set to `private` , and the `allowPublicOverrides` value is set to `true` , then all objects in the bucket default to private unless they are configured with a `public-read` ACL. Individual objects with a `public-read` ACL are readable by everyone on the internet.
	ObjectAccess *string `field:"optional" json:"objectAccess" yaml:"objectAccess"`
}

