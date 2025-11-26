package previewawscloudfrontmixins


// A key group configuration.
//
// A key group contains a list of public keys that you can use with [CloudFront signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyGroupConfigProperty := &KeyGroupConfigProperty{
//   	Comment: jsii.String("comment"),
//   	Items: []*string{
//   		jsii.String("items"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-keygroup-keygroupconfig.html
//
type CfnKeyGroupPropsMixin_KeyGroupConfigProperty struct {
	// A comment to describe the key group.
	//
	// The comment cannot be longer than 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-keygroup-keygroupconfig.html#cfn-cloudfront-keygroup-keygroupconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A list of the identifiers of the public keys in the key group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-keygroup-keygroupconfig.html#cfn-cloudfront-keygroup-keygroupconfig-items
	//
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
	// A name to identify the key group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-keygroup-keygroupconfig.html#cfn-cloudfront-keygroup-keygroupconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

