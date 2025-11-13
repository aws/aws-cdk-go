package interfacesawsses


// A reference to a MailManagerRelay resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerRelayReference := &MailManagerRelayReference{
//   	RelayId: jsii.String("relayId"),
//   }
//
type MailManagerRelayReference struct {
	// The RelayId of the MailManagerRelay resource.
	RelayId *string `field:"required" json:"relayId" yaml:"relayId"`
}

