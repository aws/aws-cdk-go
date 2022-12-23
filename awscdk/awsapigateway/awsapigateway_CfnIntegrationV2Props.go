package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::Integration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var requestParameters interface{}
//   var requestTemplates interface{}
//
//   cfnIntegrationV2Props := &cfnIntegrationV2Props{
//   	apiId: jsii.String("apiId"),
//   	integrationType: jsii.String("integrationType"),
//
//   	// the properties below are optional
//   	connectionType: jsii.String("connectionType"),
//   	contentHandlingStrategy: jsii.String("contentHandlingStrategy"),
//   	credentialsArn: jsii.String("credentialsArn"),
//   	description: jsii.String("description"),
//   	integrationMethod: jsii.String("integrationMethod"),
//   	integrationUri: jsii.String("integrationUri"),
//   	passthroughBehavior: jsii.String("passthroughBehavior"),
//   	payloadFormatVersion: jsii.String("payloadFormatVersion"),
//   	requestParameters: requestParameters,
//   	requestTemplates: requestTemplates,
//   	templateSelectionExpression: jsii.String("templateSelectionExpression"),
//   	timeoutInMillis: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnIntegrationV2Props struct {
	// `AWS::ApiGatewayV2::Integration.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::Integration.IntegrationType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-integrationtype
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// `AWS::ApiGatewayV2::Integration.ConnectionType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-connectiontype
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// `AWS::ApiGatewayV2::Integration.ContentHandlingStrategy`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-contenthandlingstrategy
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ContentHandlingStrategy *string `field:"optional" json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// `AWS::ApiGatewayV2::Integration.CredentialsArn`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-credentialsarn
	//
	// Deprecated: moved to package aws-apigatewayv2.
	CredentialsArn *string `field:"optional" json:"credentialsArn" yaml:"credentialsArn"`
	// `AWS::ApiGatewayV2::Integration.Description`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-description
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::ApiGatewayV2::Integration.IntegrationMethod`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-integrationmethod
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IntegrationMethod *string `field:"optional" json:"integrationMethod" yaml:"integrationMethod"`
	// `AWS::ApiGatewayV2::Integration.IntegrationUri`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-integrationuri
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IntegrationUri *string `field:"optional" json:"integrationUri" yaml:"integrationUri"`
	// `AWS::ApiGatewayV2::Integration.PassthroughBehavior`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-passthroughbehavior
	//
	// Deprecated: moved to package aws-apigatewayv2.
	PassthroughBehavior *string `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// `AWS::ApiGatewayV2::Integration.PayloadFormatVersion`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-payloadformatversion
	//
	// Deprecated: moved to package aws-apigatewayv2.
	PayloadFormatVersion *string `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// `AWS::ApiGatewayV2::Integration.RequestParameters`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-requestparameters
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// `AWS::ApiGatewayV2::Integration.RequestTemplates`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-requesttemplates
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RequestTemplates interface{} `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// `AWS::ApiGatewayV2::Integration.TemplateSelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-templateselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
	// `AWS::ApiGatewayV2::Integration.TimeoutInMillis`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html#cfn-apigatewayv2-integration-timeoutinmillis
	//
	// Deprecated: moved to package aws-apigatewayv2.
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
}

