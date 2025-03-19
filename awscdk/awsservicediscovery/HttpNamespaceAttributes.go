package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpNamespaceAttributes := &HttpNamespaceAttributes{
//   	NamespaceArn: jsii.String("namespaceArn"),
//   	NamespaceId: jsii.String("namespaceId"),
//   	NamespaceName: jsii.String("namespaceName"),
//   }
//
type HttpNamespaceAttributes struct {
	// Namespace ARN for the Namespace.
	NamespaceArn *string `field:"required" json:"namespaceArn" yaml:"namespaceArn"`
	// Namespace Id for the Namespace.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// A name for the Namespace.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
}

