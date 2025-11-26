package previewawsapigatewayv2mixins


// Represents a collection of route settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routeSettingsProperty := &RouteSettingsProperty{
//   	DataTraceEnabled: jsii.Boolean(false),
//   	DetailedMetricsEnabled: jsii.Boolean(false),
//   	LoggingLevel: jsii.String("loggingLevel"),
//   	ThrottlingBurstLimit: jsii.Number(123),
//   	ThrottlingRateLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-routesettings.html
//
type CfnStagePropsMixin_RouteSettingsProperty struct {
	// Specifies whether ( `true` ) or not ( `false` ) data trace logging is enabled for this route.
	//
	// This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-routesettings.html#cfn-apigatewayv2-stage-routesettings-datatraceenabled
	//
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Specifies whether detailed metrics are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-routesettings.html#cfn-apigatewayv2-stage-routesettings-detailedmetricsenabled
	//
	DetailedMetricsEnabled interface{} `field:"optional" json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// Specifies the logging level for this route: `INFO` , `ERROR` , or `OFF` .
	//
	// This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-routesettings.html#cfn-apigatewayv2-stage-routesettings-logginglevel
	//
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Specifies the throttling burst limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-routesettings.html#cfn-apigatewayv2-stage-routesettings-throttlingburstlimit
	//
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-routesettings.html#cfn-apigatewayv2-stage-routesettings-throttlingratelimit
	//
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

