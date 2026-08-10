package interfacesawsscn


// A reference to a Namespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceReference := &NamespaceReference{
//   	NamespaceArn: jsii.String("namespaceArn"),
//   }
//
type NamespaceReference struct {
	// The Arn of the Namespace resource.
	NamespaceArn *string `field:"required" json:"namespaceArn" yaml:"namespaceArn"`
}

