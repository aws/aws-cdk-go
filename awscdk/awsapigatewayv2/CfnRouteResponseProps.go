package awsapigatewayv2


// Properties for defining a `CfnRouteResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseModels interface{}
//
//   cfnRouteResponseProps := &CfnRouteResponseProps{
//   	ApiId: jsii.String("apiId"),
//   	RouteId: jsii.String("routeId"),
//   	RouteResponseKey: jsii.String("routeResponseKey"),
//
//   	// the properties below are optional
//   	ModelSelectionExpression: jsii.String("modelSelectionExpression"),
//   	ResponseModels: responseModels,
//   	ResponseParameters: map[string]interface{}{
//   		"responseParametersKey": &ParameterConstraintsProperty{
//   			"required": jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html
//
type CfnRouteResponseProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-apiid
	//
	ApiId interface{} `field:"required" json:"apiId" yaml:"apiId"`
	// The route ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-routeid
	//
	RouteId interface{} `field:"required" json:"routeId" yaml:"routeId"`
	// The route response key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-routeresponsekey
	//
	RouteResponseKey *string `field:"required" json:"routeResponseKey" yaml:"routeResponseKey"`
	// The model selection expression for the route response.
	//
	// Supported only for WebSocket APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-modelselectionexpression
	//
	ModelSelectionExpression *string `field:"optional" json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// The response models for the route response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-responsemodels
	//
	ResponseModels interface{} `field:"optional" json:"responseModels" yaml:"responseModels"`
	// The route response parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html#cfn-apigatewayv2-routeresponse-responseparameters
	//
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

