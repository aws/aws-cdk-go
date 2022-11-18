package awsapigateway


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
//   httpIntegrationProps := &httpIntegrationProps{
//   	httpMethod: jsii.String("httpMethod"),
//   	options: &integrationOptions{
//   		cacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		cacheNamespace: jsii.String("cacheNamespace"),
//   		connectionType: awscdk.Aws_apigateway.connectionType_INTERNET,
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
//   	proxy: jsii.Boolean(false),
//   }
//
type HttpIntegrationProps struct {
	// HTTP method to use when invoking the backend URL.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Integration options, such as request/resopnse mapping, content handling, etc.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// Determines whether to use proxy integration or custom integration.
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
}

