package interfacesawsmedialive


// A reference to a Node resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeReference := &NodeReference{
//   	NodeArn: jsii.String("nodeArn"),
//   }
//
type NodeReference struct {
	// The Arn of the Node resource.
	NodeArn *string `field:"required" json:"nodeArn" yaml:"nodeArn"`
}

