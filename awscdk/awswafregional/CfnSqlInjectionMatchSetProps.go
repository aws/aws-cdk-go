package awswafregional


// Properties for defining a `CfnSqlInjectionMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSqlInjectionMatchSetProps := &CfnSqlInjectionMatchSetProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	SqlInjectionMatchTuples: []interface{}{
//   		&SqlInjectionMatchTupleProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html
//
type CfnSqlInjectionMatchSetProps struct {
	// The name, if any, of the `SqlInjectionMatchSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for snippets of malicious SQL code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuples
	//
	SqlInjectionMatchTuples interface{} `field:"optional" json:"sqlInjectionMatchTuples" yaml:"sqlInjectionMatchTuples"`
}

