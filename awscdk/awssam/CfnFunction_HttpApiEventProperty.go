package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiEventProperty := &HttpApiEventProperty{
//   	ApiId: jsii.String("apiId"),
//   	Auth: &HttpApiFunctionAuthProperty{
//   		AuthorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		Authorizer: jsii.String("authorizer"),
//   	},
//   	Method: jsii.String("method"),
//   	Path: jsii.String("path"),
//   	PayloadFormatVersion: jsii.String("payloadFormatVersion"),
//   	RouteSettings: &RouteSettingsProperty{
//   		DataTraceEnabled: jsii.Boolean(false),
//   		DetailedMetricsEnabled: jsii.Boolean(false),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   	},
//   	TimeoutInMillis: jsii.Number(123),
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

