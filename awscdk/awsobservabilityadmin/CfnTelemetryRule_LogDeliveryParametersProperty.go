package awsobservabilityadmin


// Configuration parameters for Amazon Bedrock AgentCore logging, including `logType` settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDeliveryParametersProperty := &LogDeliveryParametersProperty{
//   	LogTypes: []*string{
//   		jsii.String("logTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-logdeliveryparameters.html
//
type CfnTelemetryRule_LogDeliveryParametersProperty struct {
	// The type of log that the source is sending.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-logdeliveryparameters.html#cfn-observabilityadmin-telemetryrule-logdeliveryparameters-logtypes
	//
	LogTypes *[]*string `field:"optional" json:"logTypes" yaml:"logTypes"`
}

