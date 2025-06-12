package awsapigatewayv2


// Properties for defining a `CfnRoutingRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoutingRuleProps := &CfnRoutingRuleProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			InvokeApi: &ActionInvokeApiProperty{
//   				ApiId: jsii.String("apiId"),
//   				Stage: jsii.String("stage"),
//
//   				// the properties below are optional
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
type CfnRoutingRuleProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-conditions
	//
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The amazon resource name (ARN) of the domain name resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-domainnamearn
	//
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routingrule.html#cfn-apigatewayv2-routingrule-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

