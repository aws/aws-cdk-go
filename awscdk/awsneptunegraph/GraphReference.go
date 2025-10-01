package awsneptunegraph


// A reference to a Graph resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   graphReference := &GraphReference{
//   	GraphArn: jsii.String("graphArn"),
//   	GraphId: jsii.String("graphId"),
//   }
//
type GraphReference struct {
	// The ARN of the Graph resource.
	GraphArn *string `field:"required" json:"graphArn" yaml:"graphArn"`
	// The GraphId of the Graph resource.
	GraphId *string `field:"required" json:"graphId" yaml:"graphId"`
}

