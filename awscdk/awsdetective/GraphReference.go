package awsdetective


// A reference to a Graph resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   graphReference := &GraphReference{
//   	GraphArn: jsii.String("graphArn"),
//   }
//
type GraphReference struct {
	// The Arn of the Graph resource.
	GraphArn *string `field:"required" json:"graphArn" yaml:"graphArn"`
}

