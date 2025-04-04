package awskafkaconnect


// Details about a connector's provisioned capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedCapacityProperty := &ProvisionedCapacityProperty{
//   	WorkerCount: jsii.Number(123),
//
//   	// the properties below are optional
//   	McuCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-provisionedcapacity.html
//
type CfnConnector_ProvisionedCapacityProperty struct {
	// The number of workers that are allocated to the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-provisionedcapacity.html#cfn-kafkaconnect-connector-provisionedcapacity-workercount
	//
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// The number of microcontroller units (MCUs) allocated to each connector worker.
	//
	// The valid values are 1,2,4,8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-provisionedcapacity.html#cfn-kafkaconnect-connector-provisionedcapacity-mcucount
	//
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
}

