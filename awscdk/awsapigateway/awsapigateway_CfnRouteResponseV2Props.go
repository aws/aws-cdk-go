package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::RouteResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseModels interface{}
//   var responseParameters interface{}
//
//   cfnRouteResponseV2Props := &cfnRouteResponseV2Props{
//   	apiId: jsii.String("apiId"),
//   	routeId: jsii.String("routeId"),
//   	routeResponseKey: jsii.String("routeResponseKey"),
//
//   	// the properties below are optional
//   	modelSelectionExpression: jsii.String("modelSelectionExpression"),
//   	responseModels: responseModels,
//   	responseParameters: responseParameters,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnRouteResponseV2Props struct {
	// `AWS::ApiGatewayV2::RouteResponse.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::RouteResponse.RouteId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-routeid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
	// `AWS::ApiGatewayV2::RouteResponse.RouteResponseKey`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-routeresponsekey
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RouteResponseKey *string `field:"required" json:"routeResponseKey" yaml:"routeResponseKey"`
	// `AWS::ApiGatewayV2::RouteResponse.ModelSelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-modelselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ModelSelectionExpression *string `field:"optional" json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// `AWS::ApiGatewayV2::RouteResponse.ResponseModels`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-responsemodels
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ResponseModels interface{} `field:"optional" json:"responseModels" yaml:"responseModels"`
	// `AWS::ApiGatewayV2::RouteResponse.ResponseParameters`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-responseparameters
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

