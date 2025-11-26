package previewawskafkaconnectmixins


// Specifies how the connector scales.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoScalingProperty := &AutoScalingProperty{
//   	MaxWorkerCount: jsii.Number(123),
//   	McuCount: jsii.Number(123),
//   	MinWorkerCount: jsii.Number(123),
//   	ScaleInPolicy: &ScaleInPolicyProperty{
//   		CpuUtilizationPercentage: jsii.Number(123),
//   	},
//   	ScaleOutPolicy: &ScaleOutPolicyProperty{
//   		CpuUtilizationPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html
//
type CfnConnectorPropsMixin_AutoScalingProperty struct {
	// The maximum number of workers allocated to the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-maxworkercount
	//
	MaxWorkerCount *float64 `field:"optional" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// The number of microcontroller units (MCUs) allocated to each connector worker.
	//
	// The valid values are 1,2,4,8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-mcucount
	//
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
	// The minimum number of workers allocated to the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-minworkercount
	//
	MinWorkerCount *float64 `field:"optional" json:"minWorkerCount" yaml:"minWorkerCount"`
	// The sacle-in policy for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-scaleinpolicy
	//
	ScaleInPolicy interface{} `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// The sacle-out policy for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-scaleoutpolicy
	//
	ScaleOutPolicy interface{} `field:"optional" json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}

