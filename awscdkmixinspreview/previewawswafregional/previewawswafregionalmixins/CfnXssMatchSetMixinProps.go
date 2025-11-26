package previewawswafregionalmixins


// Properties for CfnXssMatchSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnXssMatchSetMixinProps := &CfnXssMatchSetMixinProps{
//   	Name: jsii.String("name"),
//   	XssMatchTuples: []interface{}{
//   		&XssMatchTupleProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				Data: jsii.String("data"),
//   				Type: jsii.String("type"),
//   			},
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html
//
type CfnXssMatchSetMixinProps struct {
	// The name, if any, of the `XssMatchSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for cross-site scripting attacks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-xssmatchtuples
	//
	XssMatchTuples interface{} `field:"optional" json:"xssMatchTuples" yaml:"xssMatchTuples"`
}

