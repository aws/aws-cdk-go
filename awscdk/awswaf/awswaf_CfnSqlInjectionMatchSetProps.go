package awswaf


// Properties for defining a `CfnSqlInjectionMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSqlInjectionMatchSetProps := &cfnSqlInjectionMatchSetProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	sqlInjectionMatchTuples: []interface{}{
//   		&sqlInjectionMatchTupleProperty{
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
type CfnSqlInjectionMatchSetProps struct {
	// The name, if any, of the `SqlInjectionMatchSet` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for snippets of malicious SQL code.
	SqlInjectionMatchTuples interface{} `field:"optional" json:"sqlInjectionMatchTuples" yaml:"sqlInjectionMatchTuples"`
}

