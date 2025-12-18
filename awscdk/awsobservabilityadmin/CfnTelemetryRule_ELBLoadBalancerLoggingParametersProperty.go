package awsobservabilityadmin


// Configuration parameters for ELB load balancer logging, including output format and field delimiter settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eLBLoadBalancerLoggingParametersProperty := &ELBLoadBalancerLoggingParametersProperty{
//   	FieldDelimiter: jsii.String("fieldDelimiter"),
//   	OutputFormat: jsii.String("outputFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-elbloadbalancerloggingparameters.html
//
type CfnTelemetryRule_ELBLoadBalancerLoggingParametersProperty struct {
	// The delimiter character used to separate fields in ELB access log entries when using plain text format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-elbloadbalancerloggingparameters.html#cfn-observabilityadmin-telemetryrule-elbloadbalancerloggingparameters-fielddelimiter
	//
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// The format for ELB access log entries (plain text or JSON format).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-elbloadbalancerloggingparameters.html#cfn-observabilityadmin-telemetryrule-elbloadbalancerloggingparameters-outputformat
	//
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

