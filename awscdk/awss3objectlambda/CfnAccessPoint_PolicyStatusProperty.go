package awss3objectlambda


// Indicates whether this access point policy is public.
//
// For more information about how Amazon S3 evaluates policies to determine whether they are public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyStatusProperty := &PolicyStatusProperty{
//   	IsPublic: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-policystatus.html
//
type CfnAccessPoint_PolicyStatusProperty struct {
	// Specifies whether the Object lambda Access Point Policy is Public or not.
	//
	// Object lambda Access Points are private by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-policystatus.html#cfn-s3objectlambda-accesspoint-policystatus-ispublic
	//
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
}

