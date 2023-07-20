package awspinpointemail


// Used to associate a configuration set with a dedicated IP pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryOptionsProperty := &DeliveryOptionsProperty{
//   	SendingPoolName: jsii.String("sendingPoolName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-deliveryoptions.html
//
type CfnConfigurationSet_DeliveryOptionsProperty struct {
	// The name of the dedicated IP pool that you want to associate with the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-deliveryoptions.html#cfn-pinpointemail-configurationset-deliveryoptions-sendingpoolname
	//
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
}

