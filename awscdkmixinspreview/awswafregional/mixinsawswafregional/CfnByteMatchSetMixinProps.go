package mixinsawswafregional


// Properties for CfnByteMatchSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnByteMatchSetMixinProps := &CfnByteMatchSetMixinProps{
//   	ByteMatchTuples: []interface{}{
//   		&ByteMatchTupleProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				Data: jsii.String("data"),
//   				Type: jsii.String("type"),
//   			},
//   			PositionalConstraint: jsii.String("positionalConstraint"),
//   			TargetString: jsii.String("targetString"),
//   			TargetStringBase64: jsii.String("targetStringBase64"),
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html
//
type CfnByteMatchSetMixinProps struct {
	// Specifies the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-bytematchtuples
	//
	ByteMatchTuples interface{} `field:"optional" json:"byteMatchTuples" yaml:"byteMatchTuples"`
	// A friendly name or description of the `ByteMatchSet` .
	//
	// You can't change `Name` after you create a `ByteMatchSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

