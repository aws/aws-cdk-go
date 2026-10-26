package awscloudfront


// How CloudFront handles content types, including the profile paired with each one.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentTypeProfileConfigProperty := &ContentTypeProfileConfigProperty{
//   	ContentTypeProfiles: []interface{}{
//   		&ContentTypeProfileProperty{
//   			ContentType: jsii.String("contentType"),
//   			Format: jsii.String("format"),
//   			ProfileId: jsii.String("profileId"),
//   		},
//   	},
//   	ForwardWhenContentTypeIsUnknown: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-contenttypeprofileconfig.html
//
type CfnFieldLevelEncryptionConfig_ContentTypeProfileConfigProperty struct {
	// The content types paired with a profile.
	//
	// Required, and at least one item, because the service synthesizes one default item when the list is absent or empty, so neither state round-trips.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-contenttypeprofileconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-contenttypeprofileconfig-contenttypeprofiles
	//
	ContentTypeProfiles interface{} `field:"required" json:"contentTypeProfiles" yaml:"contentTypeProfiles"`
	// Whether CloudFront forwards a request when the content type names a profile that does not exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-contenttypeprofileconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-contenttypeprofileconfig-forwardwhencontenttypeisunknown
	//
	ForwardWhenContentTypeIsUnknown interface{} `field:"required" json:"forwardWhenContentTypeIsUnknown" yaml:"forwardWhenContentTypeIsUnknown"`
}

