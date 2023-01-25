package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeSettingsProperty := &routeSettingsProperty{
//   	dataTraceEnabled: jsii.Boolean(false),
//   	detailedMetricsEnabled: jsii.Boolean(false),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   }
//
type CfnFunction_RouteSettingsProperty struct {
	// `CfnFunction.RouteSettingsProperty.DataTraceEnabled`.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// `CfnFunction.RouteSettingsProperty.DetailedMetricsEnabled`.
	DetailedMetricsEnabled interface{} `field:"optional" json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// `CfnFunction.RouteSettingsProperty.LoggingLevel`.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// `CfnFunction.RouteSettingsProperty.ThrottlingBurstLimit`.
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// `CfnFunction.RouteSettingsProperty.ThrottlingRateLimit`.
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

