package awscloudfront


// A content type paired with the field-level encryption profile that encrypts it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   contentTypeProfileProperty := &ContentTypeProfileProperty{
//   	ContentType: jsii.String("contentType"),
//   	Format: jsii.String("format"),
//   	ProfileId: jsii.String("profileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-contenttypeprofile.html
//
type CfnFieldLevelEncryptionConfigPropsMixin_ContentTypeProfileProperty struct {
	// The content type CloudFront encrypts.
	//
	// The service accepts only application/x-www-form-urlencoded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-contenttypeprofile.html#cfn-cloudfront-fieldlevelencryptionconfig-contenttypeprofile-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The encryption format CloudFront uses for this content type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-contenttypeprofile.html#cfn-cloudfront-fieldlevelencryptionconfig-contenttypeprofile-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The identifier of the field-level encryption profile CloudFront uses for this content type.
	//
	// Required, so that it cannot be omitted: the service accepts a content-type profile with the member absent on create and refuses one on update, and an update sends the whole profile list, so a configuration created that way could never be updated again. Supply an empty string to configure a content type and forward it WITHOUT field-level encryption. That mode is available through this resource and round-trips: create, update and read all accept and return the empty string, and it is the service's own encoding for the state -- asked for a content-type list with no items, CloudFront returns one item whose ProfileId is empty. Note that this is not the same as omitting the content type from the configuration, because ForwardWhenContentTypeIsUnknown governs only content types absent from it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-contenttypeprofile.html#cfn-cloudfront-fieldlevelencryptionconfig-contenttypeprofile-profileid
	//
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
}

