package awslambda


// A reference to a LayerVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layerVersionReference := &LayerVersionReference{
//   	LayerVersionArn: jsii.String("layerVersionArn"),
//   }
//
type LayerVersionReference struct {
	// The LayerVersionArn of the LayerVersion resource.
	LayerVersionArn *string `field:"required" json:"layerVersionArn" yaml:"layerVersionArn"`
}

