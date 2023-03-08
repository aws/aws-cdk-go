package awsapigatewayv2


// Properties for defining a `CfnRouteResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseModels interface{}
//   var responseParameters interface{}
//
//   cfnRouteResponseProps := &cfnRouteResponseProps{
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
type CfnRouteResponseProps struct {
	// The API identifier.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The route ID.
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
	// The route response key.
	RouteResponseKey *string `field:"required" json:"routeResponseKey" yaml:"routeResponseKey"`
	// The model selection expression for the route response.
	//
	// Supported only for WebSocket APIs.
	ModelSelectionExpression *string `field:"optional" json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// The response models for the route response.
	ResponseModels interface{} `field:"optional" json:"responseModels" yaml:"responseModels"`
	// The route response parameters.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

