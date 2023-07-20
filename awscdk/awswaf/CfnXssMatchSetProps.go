package awswaf


// Properties for defining a `CfnXssMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnXssMatchSetProps := &CfnXssMatchSetProps{
//   	Name: jsii.String("name"),
//   	XssMatchTuples: []interface{}{
//   		&XssMatchTupleProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Data: jsii.String("data"),
//   			},
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html
//
type CfnXssMatchSetProps struct {
	// The name, if any, of the `XssMatchSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for cross-site scripting attacks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-xssmatchtuples
	//
	XssMatchTuples interface{} `field:"required" json:"xssMatchTuples" yaml:"xssMatchTuples"`
}

