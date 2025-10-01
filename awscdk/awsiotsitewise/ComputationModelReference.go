package awsiotsitewise


// A reference to a ComputationModel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computationModelReference := &ComputationModelReference{
//   	ComputationModelArn: jsii.String("computationModelArn"),
//   	ComputationModelId: jsii.String("computationModelId"),
//   }
//
type ComputationModelReference struct {
	// The ARN of the ComputationModel resource.
	ComputationModelArn *string `field:"required" json:"computationModelArn" yaml:"computationModelArn"`
	// The ComputationModelId of the ComputationModel resource.
	ComputationModelId *string `field:"required" json:"computationModelId" yaml:"computationModelId"`
}

