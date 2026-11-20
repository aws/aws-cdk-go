package awsapigateway


// Example:
//   var resource Resource
//
//
//   // Without auto-configuration (manual setup required)
//   resource.AddProxy(&ProxyResourceOptions{
//   	DefaultIntegration: apigateway.NewHttpIntegration(jsii.String("http://example.com/api/{proxy}"), &HttpIntegrationProps{
//   		Options: &IntegrationOptions{
//   			RequestParameters: map[string]*string{
//   				"integration.request.path.proxy": jsii.String("method.request.path.proxy"),
//   			},
//   		},
//   	}),
//   	DefaultMethodOptions: &MethodOptions{
//   		RequestParameters: map[string]*bool{
//   			"method.request.path.proxy": jsii.Boolean(true),
//   		},
//   	},
//   })
//
//   // With auto-configuration (simpler)
//   resource.AddProxy(&ProxyResourceOptions{
//   	AutoConfigurePathParameter: jsii.Boolean(true),
//   	DefaultIntegration: apigateway.NewHttpIntegration(jsii.String("http://example.com/api/{proxy}")),
//   })
//
type HttpIntegrationProps struct {
	// HTTP method to use when invoking the backend URL.
	// Default: GET.
	//
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Integration options, such as request/resopnse mapping, content handling, etc.
	// Default: defaults based on `IntegrationOptions` defaults.
	//
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// Determines whether to use proxy integration or custom integration.
	// Default: true.
	//
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
}

