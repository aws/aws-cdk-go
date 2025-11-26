package previewawskafkaconnectmixins


// The scale-out policy for the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scaleOutPolicyProperty := &ScaleOutPolicyProperty{
//   	CpuUtilizationPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-scaleoutpolicy.html
//
type CfnConnectorPropsMixin_ScaleOutPolicyProperty struct {
	// The CPU utilization percentage threshold at which you want connector scale out to be triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-scaleoutpolicy.html#cfn-kafkaconnect-connector-scaleoutpolicy-cpuutilizationpercentage
	//
	CpuUtilizationPercentage *float64 `field:"optional" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

