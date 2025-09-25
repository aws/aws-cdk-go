package awskafkaconnect


// A reference to a WorkerConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerConfigurationReference := &WorkerConfigurationReference{
//   	WorkerConfigurationArn: jsii.String("workerConfigurationArn"),
//   }
//
type WorkerConfigurationReference struct {
	// The WorkerConfigurationArn of the WorkerConfiguration resource.
	WorkerConfigurationArn *string `field:"required" json:"workerConfigurationArn" yaml:"workerConfigurationArn"`
}

