package awsobservabilityadmin


// Telemetry parameters for VPC Flow logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCFlowLogParametersProperty := &VPCFlowLogParametersProperty{
//   	LogFormat: jsii.String("logFormat"),
//   	MaxAggregationInterval: jsii.Number(123),
//   	TrafficType: jsii.String("trafficType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-vpcflowlogparameters.html
//
type CfnTelemetryRule_VPCFlowLogParametersProperty struct {
	// The fields to include in the flow log record.
	//
	// If you omit this parameter, the flow log is created using the default format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-vpcflowlogparameters.html#cfn-observabilityadmin-telemetryrule-vpcflowlogparameters-logformat
	//
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The maximum interval of time, in seconds, during which a flow of packets is captured and aggregated into a flow log record.
	//
	// Default is 600s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-vpcflowlogparameters.html#cfn-observabilityadmin-telemetryrule-vpcflowlogparameters-maxaggregationinterval
	//
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic captured for the flow log.
	//
	// Default is ALL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-vpcflowlogparameters.html#cfn-observabilityadmin-telemetryrule-vpcflowlogparameters-traffictype
	//
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

