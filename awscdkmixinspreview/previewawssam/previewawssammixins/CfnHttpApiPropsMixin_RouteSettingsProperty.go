package previewawssammixins


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-routesettings.html
//
type CfnHttpApiPropsMixin_RouteSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-routesettings.html#cfn-serverless-httpapi-routesettings-datatraceenabled
	//
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-routesettings.html#cfn-serverless-httpapi-routesettings-detailedmetricsenabled
	//
	DetailedMetricsEnabled interface{} `field:"optional" json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-routesettings.html#cfn-serverless-httpapi-routesettings-logginglevel
	//
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-routesettings.html#cfn-serverless-httpapi-routesettings-throttlingburstlimit
	//
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-routesettings.html#cfn-serverless-httpapi-routesettings-throttlingratelimit
	//
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

