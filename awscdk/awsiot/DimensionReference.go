package awsiot


// A reference to a Dimension resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionReference := &DimensionReference{
//   	DimensionArn: jsii.String("dimensionArn"),
//   	DimensionName: jsii.String("dimensionName"),
//   }
//
type DimensionReference struct {
	// The ARN of the Dimension resource.
	DimensionArn *string `field:"required" json:"dimensionArn" yaml:"dimensionArn"`
	// The Name of the Dimension resource.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
}

