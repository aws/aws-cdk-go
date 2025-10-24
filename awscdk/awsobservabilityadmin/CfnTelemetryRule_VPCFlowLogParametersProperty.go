package awsobservabilityadmin


// Configuration parameters specific to VPC Flow Logs.
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
	// The format in which VPC Flow Log entries should be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-vpcflowlogparameters.html#cfn-observabilityadmin-telemetryrule-vpcflowlogparameters-logformat
	//
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The maximum interval in seconds between the capture of flow log records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-vpcflowlogparameters.html#cfn-observabilityadmin-telemetryrule-vpcflowlogparameters-maxaggregationinterval
	//
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic to log (ACCEPT, REJECT, or ALL).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-vpcflowlogparameters.html#cfn-observabilityadmin-telemetryrule-vpcflowlogparameters-traffictype
	//
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

