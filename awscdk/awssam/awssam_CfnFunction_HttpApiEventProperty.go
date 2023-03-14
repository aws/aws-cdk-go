package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiEventProperty := &httpApiEventProperty{
//   	apiId: jsii.String("apiId"),
//   	auth: &httpApiFunctionAuthProperty{
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizer: jsii.String("authorizer"),
//   	},
//   	method: jsii.String("method"),
//   	path: jsii.String("path"),
//   	payloadFormatVersion: jsii.String("payloadFormatVersion"),
//   	routeSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	timeoutInMillis: jsii.Number(123),
//   }
//
type CfnFunction_HttpApiEventProperty struct {
	// `CfnFunction.HttpApiEventProperty.ApiId`.
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// `CfnFunction.HttpApiEventProperty.Auth`.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// `CfnFunction.HttpApiEventProperty.Method`.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// `CfnFunction.HttpApiEventProperty.Path`.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// `CfnFunction.HttpApiEventProperty.PayloadFormatVersion`.
	PayloadFormatVersion *string `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// `CfnFunction.HttpApiEventProperty.RouteSettings`.
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// `CfnFunction.HttpApiEventProperty.TimeoutInMillis`.
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
}

