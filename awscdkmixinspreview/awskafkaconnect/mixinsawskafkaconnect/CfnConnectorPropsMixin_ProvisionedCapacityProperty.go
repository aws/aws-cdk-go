package mixinsawskafkaconnect


// Details about a connector's provisioned capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   provisionedCapacityProperty := &ProvisionedCapacityProperty{
//   	McuCount: jsii.Number(123),
//   	WorkerCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-provisionedcapacity.html
//
type CfnConnectorPropsMixin_ProvisionedCapacityProperty struct {
	// The number of microcontroller units (MCUs) allocated to each connector worker.
	//
	// The valid values are 1,2,4,8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-provisionedcapacity.html#cfn-kafkaconnect-connector-provisionedcapacity-mcucount
	//
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
	// The number of workers that are allocated to the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-provisionedcapacity.html#cfn-kafkaconnect-connector-provisionedcapacity-workercount
	//
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
}

