package interfacesawskendra


// A reference to a Index resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   indexReference := &IndexReference{
//   	IndexArn: jsii.String("indexArn"),
//   	IndexId: jsii.String("indexId"),
//   }
//
type IndexReference struct {
	// The ARN of the Index resource.
	IndexArn *string `field:"required" json:"indexArn" yaml:"indexArn"`
	// The Id of the Index resource.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
}

