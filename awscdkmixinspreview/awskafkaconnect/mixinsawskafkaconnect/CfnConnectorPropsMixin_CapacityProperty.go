package mixinsawskafkaconnect


// Information about the capacity of the connector, whether it is auto scaled or provisioned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityProperty := &CapacityProperty{
//   	AutoScaling: &AutoScalingProperty{
//   		MaxWorkerCount: jsii.Number(123),
//   		McuCount: jsii.Number(123),
//   		MinWorkerCount: jsii.Number(123),
//   		ScaleInPolicy: &ScaleInPolicyProperty{
//   			CpuUtilizationPercentage: jsii.Number(123),
//   		},
//   		ScaleOutPolicy: &ScaleOutPolicyProperty{
//   			CpuUtilizationPercentage: jsii.Number(123),
//   		},
//   	},
//   	ProvisionedCapacity: &ProvisionedCapacityProperty{
//   		McuCount: jsii.Number(123),
//   		WorkerCount: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-capacity.html
//
type CfnConnectorPropsMixin_CapacityProperty struct {
	// Information about the auto scaling parameters for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-capacity.html#cfn-kafkaconnect-connector-capacity-autoscaling
	//
	AutoScaling interface{} `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// Details about a fixed capacity allocated to a connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-capacity.html#cfn-kafkaconnect-connector-capacity-provisionedcapacity
	//
	ProvisionedCapacity interface{} `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

