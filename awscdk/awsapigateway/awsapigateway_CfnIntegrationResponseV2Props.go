package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::IntegrationResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseParameters interface{}
//   var responseTemplates interface{}
//
//   cfnIntegrationResponseV2Props := &cfnIntegrationResponseV2Props{
//   	apiId: jsii.String("apiId"),
//   	integrationId: jsii.String("integrationId"),
//   	integrationResponseKey: jsii.String("integrationResponseKey"),
//
//   	// the properties below are optional
//   	contentHandlingStrategy: jsii.String("contentHandlingStrategy"),
//   	responseParameters: responseParameters,
//   	responseTemplates: responseTemplates,
//   	templateSelectionExpression: jsii.String("templateSelectionExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnIntegrationResponseV2Props struct {
	// `AWS::ApiGatewayV2::IntegrationResponse.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::IntegrationResponse.IntegrationId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-integrationid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// `AWS::ApiGatewayV2::IntegrationResponse.IntegrationResponseKey`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-integrationresponsekey
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IntegrationResponseKey *string `field:"required" json:"integrationResponseKey" yaml:"integrationResponseKey"`
	// `AWS::ApiGatewayV2::IntegrationResponse.ContentHandlingStrategy`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-contenthandlingstrategy
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ContentHandlingStrategy *string `field:"optional" json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// `AWS::ApiGatewayV2::IntegrationResponse.ResponseParameters`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-responseparameters
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// `AWS::ApiGatewayV2::IntegrationResponse.ResponseTemplates`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-responsetemplates
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// `AWS::ApiGatewayV2::IntegrationResponse.TemplateSelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-templateselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
}

