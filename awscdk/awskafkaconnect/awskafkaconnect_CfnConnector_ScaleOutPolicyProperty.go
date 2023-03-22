package awskafkaconnect


// The scale-out policy for the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleOutPolicyProperty := &scaleOutPolicyProperty{
//   	cpuUtilizationPercentage: jsii.Number(123),
//   }
//
type CfnConnector_ScaleOutPolicyProperty struct {
	// The CPU utilization percentage threshold at which you want connector scale out to be triggered.
	CpuUtilizationPercentage *float64 `field:"required" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

