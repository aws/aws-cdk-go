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
//   integrationConfig := &IntegrationConfig{
//   	Type: awscdk.Aws_apigateway.IntegrationType_AWS,
//
//   	// the properties below are optional
//   	DeploymentToken: jsii.String("deploymentToken"),
//   	IntegrationHttpMethod: jsii.String("integrationHttpMethod"),
//   	Options: &IntegrationOptions{
//   		CacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		CacheNamespace: jsii.String("cacheNamespace"),
//   		ConnectionType: awscdk.*Aws_apigateway.ConnectionType_INTERNET,
//   		ContentHandling: awscdk.*Aws_apigateway.ContentHandling_CONVERT_TO_BINARY,
//   		CredentialsPassthrough: jsii.Boolean(false),
//   		CredentialsRole: role,
//   		IntegrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				StatusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				ContentHandling: awscdk.*Aws_apigateway.ContentHandling_CONVERT_TO_BINARY,
//   				ResponseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				ResponseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				SelectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		PassthroughBehavior: awscdk.*Aws_apigateway.PassthroughBehavior_WHEN_NO_MATCH,
//   		RequestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		RequestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   		VpcLink: vpcLink,
//   	},
//   	Uri: jsii.String("uri"),
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
	// Default: undefined deployments are not triggered for any change to this integration.
	//
	DeploymentToken *string `field:"optional" json:"deploymentToken" yaml:"deploymentToken"`
	// The integration's HTTP method type.
	//
	// Required unless you use a MOCK integration.
	// Default: - no integration method specified.
	//
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Integration options.
	// Default: - no integration options.
	//
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// The Uniform Resource Identifier (URI) for the integration.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#uri
	//
	// Default: - no URI. Usually applies to MOCK integration
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

