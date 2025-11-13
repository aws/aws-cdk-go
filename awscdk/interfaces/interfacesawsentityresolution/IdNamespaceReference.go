package interfacesawsentityresolution


// A reference to a IdNamespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idNamespaceReference := &IdNamespaceReference{
//   	IdNamespaceArn: jsii.String("idNamespaceArn"),
//   	IdNamespaceName: jsii.String("idNamespaceName"),
//   }
//
type IdNamespaceReference struct {
	// The ARN of the IdNamespace resource.
	IdNamespaceArn *string `field:"required" json:"idNamespaceArn" yaml:"idNamespaceArn"`
	// The IdNamespaceName of the IdNamespace resource.
	IdNamespaceName *string `field:"required" json:"idNamespaceName" yaml:"idNamespaceName"`
}

