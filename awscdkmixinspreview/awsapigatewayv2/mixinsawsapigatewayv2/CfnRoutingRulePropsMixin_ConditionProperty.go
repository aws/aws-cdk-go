package mixinsawsapigatewayv2


// Represents a condition.
//
// Conditions can contain up to two `matchHeaders` conditions and one `matchBasePaths` conditions. API Gateway evaluates header conditions and base path conditions together. You can only use AND between header and base path conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionProperty := &ConditionProperty{
//   	MatchBasePaths: &MatchBasePathsProperty{
//   		AnyOf: []*string{
//   			jsii.String("anyOf"),
//   		},
//   	},
//   	MatchHeaders: &MatchHeadersProperty{
//   		AnyOf: []interface{}{
//   			&MatchHeaderValueProperty{
//   				Header: jsii.String("header"),
//   				ValueGlob: jsii.String("valueGlob"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-condition.html
//
type CfnRoutingRulePropsMixin_ConditionProperty struct {
	// The base path to be matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-condition.html#cfn-apigatewayv2-routingrule-condition-matchbasepaths
	//
	MatchBasePaths interface{} `field:"optional" json:"matchBasePaths" yaml:"matchBasePaths"`
	// The headers to be matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-condition.html#cfn-apigatewayv2-routingrule-condition-matchheaders
	//
	MatchHeaders interface{} `field:"optional" json:"matchHeaders" yaml:"matchHeaders"`
}

