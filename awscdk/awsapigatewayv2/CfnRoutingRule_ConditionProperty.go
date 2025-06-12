package awsapigatewayv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnRoutingRule_ConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-condition.html#cfn-apigatewayv2-routingrule-condition-matchbasepaths
	//
	MatchBasePaths interface{} `field:"optional" json:"matchBasePaths" yaml:"matchBasePaths"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-condition.html#cfn-apigatewayv2-routingrule-condition-matchheaders
	//
	MatchHeaders interface{} `field:"optional" json:"matchHeaders" yaml:"matchHeaders"`
}

