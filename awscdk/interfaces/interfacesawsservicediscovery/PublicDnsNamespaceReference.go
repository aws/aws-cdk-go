package interfacesawsservicediscovery


// A reference to a PublicDnsNamespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicDnsNamespaceReference := &PublicDnsNamespaceReference{
//   	PublicDnsNamespaceArn: jsii.String("publicDnsNamespaceArn"),
//   	PublicDnsNamespaceId: jsii.String("publicDnsNamespaceId"),
//   }
//
type PublicDnsNamespaceReference struct {
	// The ARN of the PublicDnsNamespace resource.
	PublicDnsNamespaceArn *string `field:"required" json:"publicDnsNamespaceArn" yaml:"publicDnsNamespaceArn"`
	// The Id of the PublicDnsNamespace resource.
	PublicDnsNamespaceId *string `field:"required" json:"publicDnsNamespaceId" yaml:"publicDnsNamespaceId"`
}

