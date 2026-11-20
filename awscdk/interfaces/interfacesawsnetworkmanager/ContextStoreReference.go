package interfacesawsnetworkmanager


// A reference to a ContextStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextStoreReference := &ContextStoreReference{
//   	ContextStoreArn: jsii.String("contextStoreArn"),
//   	ContextStoreId: jsii.String("contextStoreId"),
//   }
//
type ContextStoreReference struct {
	// The ARN of the ContextStore resource.
	ContextStoreArn *string `field:"required" json:"contextStoreArn" yaml:"contextStoreArn"`
	// The ContextStoreId of the ContextStore resource.
	ContextStoreId *string `field:"required" json:"contextStoreId" yaml:"contextStoreId"`
}

