package awskafkaconnect


// The configuration of the workers, which are the processes that run the connector logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerConfigurationProperty := &workerConfigurationProperty{
//   	revision: jsii.Number(123),
//   	workerConfigurationArn: jsii.String("workerConfigurationArn"),
//   }
//
type CfnConnector_WorkerConfigurationProperty struct {
	// The revision of the worker configuration.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
	// The Amazon Resource Name (ARN) of the worker configuration.
	WorkerConfigurationArn *string `field:"required" json:"workerConfigurationArn" yaml:"workerConfigurationArn"`
}

