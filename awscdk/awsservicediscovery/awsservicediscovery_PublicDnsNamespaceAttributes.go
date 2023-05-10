package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicDnsNamespaceAttributes := &PublicDnsNamespaceAttributes{
//   	NamespaceArn: jsii.String("namespaceArn"),
//   	NamespaceId: jsii.String("namespaceId"),
//   	NamespaceName: jsii.String("namespaceName"),
//   }
//
// Experimental.
type PublicDnsNamespaceAttributes struct {
	// Namespace ARN for the Namespace.
	// Experimental.
	NamespaceArn *string `field:"required" json:"namespaceArn" yaml:"namespaceArn"`
	// Namespace Id for the Namespace.
	// Experimental.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// A name for the Namespace.
	// Experimental.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
}

