package awss3objectlambda


// > Amazon S3 Object Lambda will no longer be open to new customers starting on 11/7/2025.
//
// If you would like to use the service, please sign up prior to 11/7/2025. For capabilities similar to S3 Object Lambda, learn more here - [Amazon S3 Object Lambda availability change](https://docs.aws.amazon.com/AmazonS3/latest/userguide/amazons3-ol-change.html) .
//
// The alias of an Object Lambda Access Point. For more information, see [How to use a bucket-style alias for your S3 bucket Object Lambda Access Point](https://docs.aws.amazon.com/AmazonS3/latest/userguide/olap-use.html#ol-access-points-alias) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasProperty := &AliasProperty{
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-alias.html
//
type CfnAccessPoint_AliasProperty struct {
	// The alias value of the Object Lambda Access Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-alias.html#cfn-s3objectlambda-accesspoint-alias-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
	// The status of the Object Lambda Access Point alias.
	//
	// If the status is `PROVISIONING` , the Object Lambda Access Point is provisioning the alias and the alias is not ready for use yet. If the status is `READY` , the Object Lambda Access Point alias is successfully provisioned and ready for use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-alias.html#cfn-s3objectlambda-accesspoint-alias-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

