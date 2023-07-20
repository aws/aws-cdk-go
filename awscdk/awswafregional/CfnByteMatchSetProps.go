package awswafregional


// Properties for defining a `CfnByteMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnByteMatchSetProps := &CfnByteMatchSetProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ByteMatchTuples: []interface{}{
//   		&ByteMatchTupleProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Data: jsii.String("data"),
//   			},
//   			PositionalConstraint: jsii.String("positionalConstraint"),
//   			TextTransformation: jsii.String("textTransformation"),
//
//   			// the properties below are optional
//   			TargetString: jsii.String("targetString"),
//   			TargetStringBase64: jsii.String("targetStringBase64"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html
//
type CfnByteMatchSetProps struct {
	// A friendly name or description of the `ByteMatchSet` .
	//
	// You can't change `Name` after you create a `ByteMatchSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-bytematchtuples
	//
	ByteMatchTuples interface{} `field:"optional" json:"byteMatchTuples" yaml:"byteMatchTuples"`
}

