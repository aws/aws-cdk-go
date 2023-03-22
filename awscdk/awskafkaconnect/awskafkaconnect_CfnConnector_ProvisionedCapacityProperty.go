package awskafkaconnect


// Details about a connector's provisioned capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedCapacityProperty := &provisionedCapacityProperty{
//   	workerCount: jsii.Number(123),
//
//   	// the properties below are optional
//   	mcuCount: jsii.Number(123),
//   }
//
type CfnConnector_ProvisionedCapacityProperty struct {
	// The number of workers that are allocated to the connector.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// The number of microcontroller units (MCUs) allocated to each connector worker.
	//
	// The valid values are 1,2,4,8.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
}

