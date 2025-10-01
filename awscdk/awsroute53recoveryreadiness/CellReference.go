package awsroute53recoveryreadiness


// A reference to a Cell resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cellReference := &CellReference{
//   	CellArn: jsii.String("cellArn"),
//   	CellName: jsii.String("cellName"),
//   }
//
type CellReference struct {
	// The ARN of the Cell resource.
	CellArn *string `field:"required" json:"cellArn" yaml:"cellArn"`
	// The CellName of the Cell resource.
	CellName *string `field:"required" json:"cellName" yaml:"cellName"`
}

