package previewawsapigatewayv2mixins


// Properties for CfnRoutingRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoutingRuleMixinProps := &CfnRoutingRuleMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			InvokeApi: &ActionInvokeApiProperty{
//   				ApiId: jsii.String("apiId"),
//   				Stage: jsii.String("stage"),
//   				StripBasePath: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			MatchBasePaths: &MatchBasePathsProperty{
//   				AnyOf: []*string{
//   					jsii.String("anyOf"),
//   				},
//   			},
//   			MatchHeaders: &MatchHeadersProperty{
//   				AnyOf: []interface{}{
//   					&MatchHeaderValueProperty{
//   						Header: jsii.String("header"),
//   						ValueGlob: jsii.String("valueGlob"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DomainNameArn: jsii.String("domainNameArn"),
//   	Priority: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html
//
type CfnRoutingRuleMixinProps struct {
	// The resulting action based on matching a routing rules condition.
	//
	// Only InvokeApi is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The conditions of the routing rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The ARN of the domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-domainnamearn
	//
	DomainNameArn *string `field:"optional" json:"domainNameArn" yaml:"domainNameArn"`
	// The order in which API Gateway evaluates a rule.
	//
	// Priority is evaluated from the lowest value to the highest value. Rules can't have the same priority. Priority values 1-1,000,000 are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

