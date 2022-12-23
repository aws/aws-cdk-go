package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetKinesisStreamParametersProperty := &pipeTargetKinesisStreamParametersProperty{
//   	partitionKey: jsii.String("partitionKey"),
//   }
//
type CfnPipe_PipeTargetKinesisStreamParametersProperty struct {
	// `CfnPipe.PipeTargetKinesisStreamParametersProperty.PartitionKey`.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

