package interfacesawsses


// A reference to a MailManagerAddressList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerAddressListReference := &MailManagerAddressListReference{
//   	AddressListArn: jsii.String("addressListArn"),
//   	AddressListId: jsii.String("addressListId"),
//   }
//
type MailManagerAddressListReference struct {
	// The ARN of the MailManagerAddressList resource.
	AddressListArn *string `field:"required" json:"addressListArn" yaml:"addressListArn"`
	// The AddressListId of the MailManagerAddressList resource.
	AddressListId *string `field:"required" json:"addressListId" yaml:"addressListId"`
}

