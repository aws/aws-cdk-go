package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisParametersProperty := &kinesisParametersProperty{
//   	partitionKey: jsii.String("partitionKey"),
//   }
//
type CfnSchedule_KinesisParametersProperty struct {
	// `CfnSchedule.KinesisParametersProperty.PartitionKey`.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

