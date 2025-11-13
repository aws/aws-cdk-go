package interfacesawsservicediscovery


// A reference to a PrivateDnsNamespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateDnsNamespaceReference := &PrivateDnsNamespaceReference{
//   	PrivateDnsNamespaceArn: jsii.String("privateDnsNamespaceArn"),
//   	PrivateDnsNamespaceId: jsii.String("privateDnsNamespaceId"),
//   }
//
type PrivateDnsNamespaceReference struct {
	// The ARN of the PrivateDnsNamespace resource.
	PrivateDnsNamespaceArn *string `field:"required" json:"privateDnsNamespaceArn" yaml:"privateDnsNamespaceArn"`
	// The Id of the PrivateDnsNamespace resource.
	PrivateDnsNamespaceId *string `field:"required" json:"privateDnsNamespaceId" yaml:"privateDnsNamespaceId"`
}

