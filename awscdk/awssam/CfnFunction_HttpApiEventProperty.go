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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html
//
type CfnFunction_HttpApiEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html#cfn-serverless-function-httpapievent-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html#cfn-serverless-function-httpapievent-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html#cfn-serverless-function-httpapievent-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html#cfn-serverless-function-httpapievent-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html#cfn-serverless-function-httpapievent-payloadformatversion
	//
	PayloadFormatVersion *string `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html#cfn-serverless-function-httpapievent-routesettings
	//
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapievent.html#cfn-serverless-function-httpapievent-timeoutinmillis
	//
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
}

