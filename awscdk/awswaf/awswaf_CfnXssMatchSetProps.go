package awswaf


// Properties for defining a `CfnXssMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnXssMatchSetProps := &cfnXssMatchSetProps{
//   	name: jsii.String("name"),
//   	xssMatchTuples: []interface{}{
//   		&xssMatchTupleProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				data: jsii.String("data"),
//   			},
//   			textTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }
//
type CfnXssMatchSetProps struct {
	// The name, if any, of the `XssMatchSet` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples interface{} `field:"required" json:"xssMatchTuples" yaml:"xssMatchTuples"`
}

