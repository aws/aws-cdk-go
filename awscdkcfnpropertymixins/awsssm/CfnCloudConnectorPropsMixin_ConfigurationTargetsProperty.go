package awsssm


// The targets for the cloud connector.
//
// If omitted, the entire tenant is targeted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   configurationTargetsProperty := &ConfigurationTargetsProperty{
//   	Subscriptions: []interface{}{
//   		&AzureSubscriptionProperty{
//   			DisplayName: jsii.String("displayName"),
//   			Id: jsii.String("id"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-configurationtargets.html
//
type CfnCloudConnectorPropsMixin_ConfigurationTargetsProperty struct {
	// List of Azure subscriptions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-configurationtargets.html#cfn-ssm-cloudconnector-configurationtargets-subscriptions
	//
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

