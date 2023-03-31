package awsses


// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryOptionsProperty := &deliveryOptionsProperty{
//   	sendingPoolName: jsii.String("sendingPoolName"),
//   	tlsPolicy: jsii.String("tlsPolicy"),
//   }
//
type CfnConfigurationSet_DeliveryOptionsProperty struct {
	// The name of the dedicated IP pool to associate with the configuration set.
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
	// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	//
	// If the value is `REQUIRE` , messages are only delivered if a TLS connection can be established. If the value is `OPTIONAL` , messages can be delivered in plain text if a TLS connection can't be established.
	//
	// Valid Values: `REQUIRE | OPTIONAL`.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

