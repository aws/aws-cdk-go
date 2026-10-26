package awscloudfront


// A query argument paired with the field-level encryption profile that encrypts it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryArgProfileProperty := &QueryArgProfileProperty{
//   	ProfileId: jsii.String("profileId"),
//   	QueryArg: jsii.String("queryArg"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-queryargprofile.html
//
type CfnFieldLevelEncryptionConfig_QueryArgProfileProperty struct {
	// The identifier of the field-level encryption profile CloudFront uses for this query argument.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-queryargprofile.html#cfn-cloudfront-fieldlevelencryptionconfig-queryargprofile-profileid
	//
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The query argument that CloudFront encrypts with the paired profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-queryargprofile.html#cfn-cloudfront-fieldlevelencryptionconfig-queryargprofile-queryarg
	//
	QueryArg *string `field:"required" json:"queryArg" yaml:"queryArg"`
}

