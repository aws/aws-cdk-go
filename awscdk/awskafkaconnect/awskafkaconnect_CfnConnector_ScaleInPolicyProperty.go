package awskafkaconnect


// The scale-in policy for the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleInPolicyProperty := &scaleInPolicyProperty{
//   	cpuUtilizationPercentage: jsii.Number(123),
//   }
//
type CfnConnector_ScaleInPolicyProperty struct {
	// Specifies the CPU utilization percentage threshold at which you want connector scale in to be triggered.
	CpuUtilizationPercentage *float64 `field:"required" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

