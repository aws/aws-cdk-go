package previewawscloudfrontmixins


// A complex type that contains zero or more `Tag` elements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagsProperty := &TagsProperty{
//   	Items: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-tags.html
//
type CfnAnycastIpListPropsMixin_TagsProperty struct {
	// A complex type that contains `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-tags.html#cfn-cloudfront-anycastiplist-tags-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

