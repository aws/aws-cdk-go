package awswafregional


// Properties for defining a `CfnByteMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnByteMatchSetProps := &cfnByteMatchSetProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	byteMatchTuples: []interface{}{
//   		&byteMatchTupleProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				data: jsii.String("data"),
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformation: jsii.String("textTransformation"),
//
//   			// the properties below are optional
//   			targetString: jsii.String("targetString"),
//   			targetStringBase64: jsii.String("targetStringBase64"),
//   		},
//   	},
//   }
//
type CfnByteMatchSetProps struct {
	// A friendly name or description of the `ByteMatchSet` .
	//
	// You can't change `Name` after you create a `ByteMatchSet` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings.
	ByteMatchTuples interface{} `field:"optional" json:"byteMatchTuples" yaml:"byteMatchTuples"`
}

