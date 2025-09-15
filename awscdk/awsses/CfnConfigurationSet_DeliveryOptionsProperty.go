package awsses


// Specifies the name of the dedicated IP pool to associate with the configuration set and whether messages that use the configuration set are required to use Transport Layer Security (TLS).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryOptionsProperty := &DeliveryOptionsProperty{
//   	MaxDeliverySeconds: jsii.Number(123),
//   	SendingPoolName: jsii.String("sendingPoolName"),
//   	TlsPolicy: jsii.String("tlsPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-deliveryoptions.html
//
type CfnConfigurationSet_DeliveryOptionsProperty struct {
	// The name of the configuration set used when sent through a configuration set with archiving enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-deliveryoptions.html#cfn-ses-configurationset-deliveryoptions-maxdeliveryseconds
	//
	MaxDeliverySeconds *float64 `field:"optional" json:"maxDeliverySeconds" yaml:"maxDeliverySeconds"`
	// The name of the dedicated IP pool to associate with the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-deliveryoptions.html#cfn-ses-configurationset-deliveryoptions-sendingpoolname
	//
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
	// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	//
	// If the value is `REQUIRE` , messages are only delivered if a TLS connection can be established. If the value is `OPTIONAL` , messages can be delivered in plain text if a TLS connection can't be established.
	//
	// Valid Values: `REQUIRE | OPTIONAL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-deliveryoptions.html#cfn-ses-configurationset-deliveryoptions-tlspolicy
	//
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

