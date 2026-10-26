package awscloudfront


// Properties for defining a `CfnFieldLevelEncryptionConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFieldLevelEncryptionConfigProps := &CfnFieldLevelEncryptionConfigProps{
//   	CallerReference: jsii.String("callerReference"),
//   	ContentTypeProfileConfig: &ContentTypeProfileConfigProperty{
//   		ContentTypeProfiles: []interface{}{
//   			&ContentTypeProfileProperty{
//   				ContentType: jsii.String("contentType"),
//   				Format: jsii.String("format"),
//   				ProfileId: jsii.String("profileId"),
//   			},
//   		},
//   		ForwardWhenContentTypeIsUnknown: jsii.Boolean(false),
//   	},
//   	QueryArgProfileConfig: &QueryArgProfileConfigProperty{
//   		ForwardWhenQueryArgProfileIsUnknown: jsii.Boolean(false),
//   		QueryArgProfiles: []interface{}{
//   			&QueryArgProfileProperty{
//   				ProfileId: jsii.String("profileId"),
//   				QueryArg: jsii.String("queryArg"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionconfig.html
//
type CfnFieldLevelEncryptionConfigProps struct {
	// A unique value that identifies the create request.
	//
	// Any string unique within the account works; CloudFront refuses to change it once set, so supplying a different value replaces the configuration. It is also the only value CloudFront returns that identifies a configuration as this stack's, so cleaning up after a failed create depends on the template keeping it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-callerreference
	//
	CallerReference *string `field:"required" json:"callerReference" yaml:"callerReference"`
	// How CloudFront handles content types, including the profile paired with each one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-contenttypeprofileconfig
	//
	ContentTypeProfileConfig interface{} `field:"required" json:"contentTypeProfileConfig" yaml:"contentTypeProfileConfig"`
	// How CloudFront handles query arguments, including the profile paired with each one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-queryargprofileconfig
	//
	QueryArgProfileConfig interface{} `field:"required" json:"queryArgProfileConfig" yaml:"queryArgProfileConfig"`
	// An optional comment describing the field-level encryption configuration.
	//
	// CloudFront stores an omitted comment and an explicitly empty one identically as an empty string, and Read normalizes that empty string back to absent -- so omitting Comment round-trips, while asking for an empty string does not and would report perpetual drift. `minLength: 1` refuses the empty string at template validation; omission is the way to express no comment. The service documents a maximum of 128 characters, and `maxLength` moves that refusal from the API to template validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-fieldlevelencryptionconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

