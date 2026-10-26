package awscloudfront


// How CloudFront handles query arguments, including the profile paired with each one.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   queryArgProfileConfigProperty := &QueryArgProfileConfigProperty{
//   	ForwardWhenQueryArgProfileIsUnknown: jsii.Boolean(false),
//   	QueryArgProfiles: []interface{}{
//   		&QueryArgProfileProperty{
//   			ProfileId: jsii.String("profileId"),
//   			QueryArg: jsii.String("queryArg"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-queryargprofileconfig.html
//
type CfnFieldLevelEncryptionConfigPropsMixin_QueryArgProfileConfigProperty struct {
	// Whether CloudFront forwards a request when the query argument names a profile that does not exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-queryargprofileconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-queryargprofileconfig-forwardwhenqueryargprofileisunknown
	//
	ForwardWhenQueryArgProfileIsUnknown interface{} `field:"optional" json:"forwardWhenQueryArgProfileIsUnknown" yaml:"forwardWhenQueryArgProfileIsUnknown"`
	// The query arguments paired with a profile.
	//
	// Required because the service returns an empty list rather than preserving absence, so an omitted value cannot round-trip. An empty list is accepted and is preserved.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-fieldlevelencryptionconfig-queryargprofileconfig.html#cfn-cloudfront-fieldlevelencryptionconfig-queryargprofileconfig-queryargprofiles
	//
	QueryArgProfiles interface{} `field:"optional" json:"queryArgProfiles" yaml:"queryArgProfiles"`
}

