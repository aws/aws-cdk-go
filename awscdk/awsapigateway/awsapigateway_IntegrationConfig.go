package awsapigateway


// Result of binding an Integration to a Method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var vpcLink vpcLink
//
//   integrationConfig := &integrationConfig{
//   	type: awscdk.Aws_apigateway.integrationType_AWS,
//
//   	// the properties below are optional
//   	deploymentToken: jsii.String("deploymentToken"),
//   	integrationHttpMethod: jsii.String("integrationHttpMethod"),
//   	options: &integrationOptions{
//   		cacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		cacheNamespace: jsii.String("cacheNamespace"),
//   		connectionType: awscdk.*Aws_apigateway.connectionType_INTERNET,
//   		contentHandling: awscdk.*Aws_apigateway.contentHandling_CONVERT_TO_BINARY,
//   		credentialsPassthrough: jsii.Boolean(false),
//   		credentialsRole: role,
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentHandling: awscdk.*Aws_apigateway.*contentHandling_CONVERT_TO_BINARY,
//   				responseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				responseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				selectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.*Aws_apigateway.passthroughBehavior_WHEN_NO_MATCH,
//   		requestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		requestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		timeout: cdk.duration.minutes(jsii.Number(30)),
//   		vpcLink: vpcLink,
//   	},
//   	uri: jsii.String("uri"),
//   }
//
type IntegrationConfig struct {
	// Specifies an API method integration type.
	Type IntegrationType `field:"required" json:"type" yaml:"type"`
	// This value is included in computing the Deployment's fingerprint.
	//
	// When the fingerprint
	// changes, a new deployment is triggered.
	// This property should contain values associated with the Integration that upon changing
	// should trigger a fresh the Deployment needs to be refreshed.
	DeploymentToken *string `field:"optional" json:"deploymentToken" yaml:"deploymentToken"`
	// The integration's HTTP method type.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Integration options.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// The Uniform Resource Identifier (URI) for the integration.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

