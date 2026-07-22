package interfacesawscodebuild


// A reference to a BuildBatch resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buildBatchReference := &BuildBatchReference{
//   	BuildBatchArn: jsii.String("buildBatchArn"),
//   	BuildBatchId: jsii.String("buildBatchId"),
//   }
//
type BuildBatchReference struct {
	// The ARN of the BuildBatch resource.
	BuildBatchArn *string `field:"required" json:"buildBatchArn" yaml:"buildBatchArn"`
	// The Id of the BuildBatch resource.
	BuildBatchId *string `field:"required" json:"buildBatchId" yaml:"buildBatchId"`
}

