package interfacesawsservicediscovery


// A reference to a HttpNamespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpNamespaceReference := &HttpNamespaceReference{
//   	HttpNamespaceArn: jsii.String("httpNamespaceArn"),
//   	HttpNamespaceId: jsii.String("httpNamespaceId"),
//   }
//
type HttpNamespaceReference struct {
	// The ARN of the HttpNamespace resource.
	HttpNamespaceArn *string `field:"required" json:"httpNamespaceArn" yaml:"httpNamespaceArn"`
	// The Id of the HttpNamespace resource.
	HttpNamespaceId *string `field:"required" json:"httpNamespaceId" yaml:"httpNamespaceId"`
}

