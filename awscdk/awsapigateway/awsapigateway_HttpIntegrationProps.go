package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var role role
//   var vpcLink vpcLink
//
//   httpIntegrationProps := &HttpIntegrationProps{
//   	HttpMethod: jsii.String("httpMethod"),
//   	Options: &IntegrationOptions{
//   		CacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		CacheNamespace: jsii.String("cacheNamespace"),
//   		ConnectionType: awscdk.Aws_apigateway.ConnectionType_INTERNET,
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
//   		Timeout: duration,
//   		VpcLink: vpcLink,
//   	},
//   	Proxy: jsii.Boolean(false),
//   }
//
// Experimental.
type HttpIntegrationProps struct {
	// HTTP method to use when invoking the backend URL.
	// Experimental.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Integration options, such as request/resopnse mapping, content handling, etc.
	// Experimental.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// Determines whether to use proxy integration or custom integration.
	// Experimental.
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
}

