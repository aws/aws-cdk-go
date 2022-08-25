package cloudassemblyschema


// Artifact properties for the Construct Tree Artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   treeArtifactProperties := &treeArtifactProperties{
//   	file: jsii.String("file"),
//   }
//
type TreeArtifactProperties struct {
	// Filename of the tree artifact.
	File *string `field:"required" json:"file" yaml:"file"`
}

