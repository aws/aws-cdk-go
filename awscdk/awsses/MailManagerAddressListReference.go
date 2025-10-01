package awsses


// A reference to a MailManagerAddressList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerAddressListReference := &MailManagerAddressListReference{
//   	AddressListId: jsii.String("addressListId"),
//   }
//
type MailManagerAddressListReference struct {
	// The AddressListId of the MailManagerAddressList resource.
	AddressListId *string `field:"required" json:"addressListId" yaml:"addressListId"`
}

