package awskafkaconnect


// The scale-in policy for the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleInPolicyProperty := &ScaleInPolicyProperty{
//   	CpuUtilizationPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-scaleinpolicy.html
//
type CfnConnector_ScaleInPolicyProperty struct {
	// Specifies the CPU utilization percentage threshold at which you want connector scale in to be triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-scaleinpolicy.html#cfn-kafkaconnect-connector-scaleinpolicy-cpuutilizationpercentage
	//
	CpuUtilizationPercentage *float64 `field:"required" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

