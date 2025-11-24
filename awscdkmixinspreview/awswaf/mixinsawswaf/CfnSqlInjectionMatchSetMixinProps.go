package mixinsawswaf


// Properties for CfnSqlInjectionMatchSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSqlInjectionMatchSetMixinProps := &CfnSqlInjectionMatchSetMixinProps{
//   	Name: jsii.String("name"),
//   	SqlInjectionMatchTuples: []interface{}{
//   		&SqlInjectionMatchTupleProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				Data: jsii.String("data"),
//   				Type: jsii.String("type"),
//   			},
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html
//
type CfnSqlInjectionMatchSetMixinProps struct {
	// The name, if any, of the `SqlInjectionMatchSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for snippets of malicious SQL code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples
	//
	SqlInjectionMatchTuples interface{} `field:"optional" json:"sqlInjectionMatchTuples" yaml:"sqlInjectionMatchTuples"`
}

