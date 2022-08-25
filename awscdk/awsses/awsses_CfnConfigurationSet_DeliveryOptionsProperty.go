package awsses


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
	// `CfnConfigurationSet.DeliveryOptionsProperty.SendingPoolName`.
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
	// `CfnConfigurationSet.DeliveryOptionsProperty.TlsPolicy`.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

