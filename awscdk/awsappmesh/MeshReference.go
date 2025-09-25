package awsappmesh


// A reference to a Mesh resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   meshReference := &MeshReference{
//   	MeshArn: jsii.String("meshArn"),
//   	MeshId: jsii.String("meshId"),
//   }
//
type MeshReference struct {
	// The ARN of the Mesh resource.
	MeshArn *string `field:"required" json:"meshArn" yaml:"meshArn"`
	// The Id of the Mesh resource.
	MeshId *string `field:"required" json:"meshId" yaml:"meshId"`
}

