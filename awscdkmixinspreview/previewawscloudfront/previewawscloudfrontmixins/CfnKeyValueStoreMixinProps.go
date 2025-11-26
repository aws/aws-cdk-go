package previewawscloudfrontmixins


// Properties for CfnKeyValueStorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKeyValueStoreMixinProps := &CfnKeyValueStoreMixinProps{
//   	Comment: jsii.String("comment"),
//   	ImportSource: &ImportSourceProperty{
//   		SourceArn: jsii.String("sourceArn"),
//   		SourceType: jsii.String("sourceType"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html
//
type CfnKeyValueStoreMixinProps struct {
	// A comment for the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html#cfn-cloudfront-keyvaluestore-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The import source for the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html#cfn-cloudfront-keyvaluestore-importsource
	//
	ImportSource interface{} `field:"optional" json:"importSource" yaml:"importSource"`
	// The name of the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html#cfn-cloudfront-keyvaluestore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

