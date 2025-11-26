package previewawswafv2mixins


// Application details defined during the web ACL creation process.
//
// Application attributes help AWS WAF give recommendations for protection packs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationAttributeProperty := &ApplicationAttributeProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-applicationattribute.html
//
type CfnWebACLPropsMixin_ApplicationAttributeProperty struct {
	// Specifies the attribute name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-applicationattribute.html#cfn-wafv2-webacl-applicationattribute-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the attribute value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-applicationattribute.html#cfn-wafv2-webacl-applicationattribute-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

