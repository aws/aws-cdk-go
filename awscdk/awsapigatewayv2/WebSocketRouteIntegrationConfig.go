package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Config returned back as a result of the bind.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   webSocketRouteIntegrationConfig := &WebSocketRouteIntegrationConfig{
//   	Type: awscdk.Aws_apigatewayv2.WebSocketIntegrationType_AWS_PROXY,
//   	Uri: jsii.String("uri"),
//
//   	// the properties below are optional
//   	CredentialsRole: role,
//   	Method: jsii.String("method"),
//   	PassthroughBehavior: awscdk.*Aws_apigatewayv2.PassthroughBehavior_WHEN_NO_MATCH,
//   	RequestParameters: map[string]*string{
//   		"requestParametersKey": jsii.String("requestParameters"),
//   	},
//   	RequestTemplates: map[string]*string{
//   		"requestTemplatesKey": jsii.String("requestTemplates"),
//   	},
//   	TemplateSelectionExpression: jsii.String("templateSelectionExpression"),
//   }
//
type WebSocketRouteIntegrationConfig struct {
	// Integration type.
	Type WebSocketIntegrationType `field:"required" json:"type" yaml:"type"`
	// Integration URI.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Credentials role.
	// Default: - No role provided.
	//
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// Integration method.
	// Default: - No integration method.
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Integration passthrough behaviors.
	// Default: - No pass through bahavior.
	//
	PassthroughBehavior PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// Request parameters.
	// Default: - No request parameters provided.
	//
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// Request template.
	// Default: - No request template provided.
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// Template selection expression.
	// Default: - No template selection expression.
	//
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
}

