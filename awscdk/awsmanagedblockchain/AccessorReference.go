package awsmanagedblockchain


// A reference to a Accessor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessorReference := &AccessorReference{
//   	AccessorArn: jsii.String("accessorArn"),
//   	AccessorId: jsii.String("accessorId"),
//   }
//
type AccessorReference struct {
	// The ARN of the Accessor resource.
	AccessorArn *string `field:"required" json:"accessorArn" yaml:"accessorArn"`
	// The Id of the Accessor resource.
	AccessorId *string `field:"required" json:"accessorId" yaml:"accessorId"`
}

