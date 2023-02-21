package awspinpointemail


// Used to associate a configuration set with a dedicated IP pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryOptionsProperty := &deliveryOptionsProperty{
//   	sendingPoolName: jsii.String("sendingPoolName"),
//   }
//
type CfnConfigurationSet_DeliveryOptionsProperty struct {
	// The name of the dedicated IP pool that you want to associate with the configuration set.
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
}

