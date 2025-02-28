package awskafkaconnect


// The configuration of the workers, which are the processes that run the connector logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerConfigurationProperty := &WorkerConfigurationProperty{
//   	Revision: jsii.Number(123),
//   	WorkerConfigurationArn: jsii.String("workerConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerconfiguration.html
//
type CfnConnector_WorkerConfigurationProperty struct {
	// The revision of the worker configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerconfiguration.html#cfn-kafkaconnect-connector-workerconfiguration-revision
	//
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
	// The Amazon Resource Name (ARN) of the worker configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerconfiguration.html#cfn-kafkaconnect-connector-workerconfiguration-workerconfigurationarn
	//
	WorkerConfigurationArn *string `field:"required" json:"workerConfigurationArn" yaml:"workerConfigurationArn"`
}

