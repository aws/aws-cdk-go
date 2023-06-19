package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::Route`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var requestModels interface{}
//   var requestParameters interface{}
//
//   cfnRouteV2Props := &CfnRouteV2Props{
//   	ApiId: jsii.String("apiId"),
//   	RouteKey: jsii.String("routeKey"),
//
//   	// the properties below are optional
//   	ApiKeyRequired: jsii.Boolean(false),
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	AuthorizationType: jsii.String("authorizationType"),
//   	AuthorizerId: jsii.String("authorizerId"),
//   	ModelSelectionExpression: jsii.String("modelSelectionExpression"),
//   	OperationName: jsii.String("operationName"),
//   	RequestModels: requestModels,
//   	RequestParameters: requestParameters,
//   	RouteResponseSelectionExpression: jsii.String("routeResponseSelectionExpression"),
//   	Target: jsii.String("target"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnRouteV2Props struct {
	// `AWS::ApiGatewayV2::Route.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::Route.RouteKey`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-routekey
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RouteKey *string `field:"required" json:"routeKey" yaml:"routeKey"`
	// `AWS::ApiGatewayV2::Route.ApiKeyRequired`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-apikeyrequired
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// `AWS::ApiGatewayV2::Route.AuthorizationScopes`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-authorizationscopes
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// `AWS::ApiGatewayV2::Route.AuthorizationType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-authorizationtype
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// `AWS::ApiGatewayV2::Route.AuthorizerId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-authorizerid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// `AWS::ApiGatewayV2::Route.ModelSelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-modelselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ModelSelectionExpression *string `field:"optional" json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// `AWS::ApiGatewayV2::Route.OperationName`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-operationname
	//
	// Deprecated: moved to package aws-apigatewayv2.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// `AWS::ApiGatewayV2::Route.RequestModels`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-requestmodels
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RequestModels interface{} `field:"optional" json:"requestModels" yaml:"requestModels"`
	// `AWS::ApiGatewayV2::Route.RequestParameters`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-requestparameters
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// `AWS::ApiGatewayV2::Route.RouteResponseSelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-routeresponseselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RouteResponseSelectionExpression *string `field:"optional" json:"routeResponseSelectionExpression" yaml:"routeResponseSelectionExpression"`
	// `AWS::ApiGatewayV2::Route.Target`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html#cfn-apigatewayv2-route-target
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

