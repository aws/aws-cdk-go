package awsredshiftserverless


// A reference to a Namespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceReference := &NamespaceReference{
//   	NamespaceName: jsii.String("namespaceName"),
//   }
//
type NamespaceReference struct {
	// The NamespaceName of the Namespace resource.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
}

