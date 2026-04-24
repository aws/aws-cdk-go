package awsdevopsagent


// Azure subscription integration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureConfigurationProperty := &AzureConfigurationProperty{
//   	SubscriptionId: jsii.String("subscriptionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-azureconfiguration.html
//
type CfnAssociation_AzureConfigurationProperty struct {
	// Azure subscription ID corresponding to provided resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-azureconfiguration.html#cfn-devopsagent-association-azureconfiguration-subscriptionid
	//
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
}

