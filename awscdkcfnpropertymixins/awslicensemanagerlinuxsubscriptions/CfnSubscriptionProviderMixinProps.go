package awslicensemanagerlinuxsubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSubscriptionProviderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSubscriptionProviderMixinProps := &CfnSubscriptionProviderMixinProps{
//   	SecretArn: jsii.String("secretArn"),
//   	SubscriptionProviderSource: jsii.String("subscriptionProviderSource"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerlinuxsubscriptions-subscriptionprovider.html
//
type CfnSubscriptionProviderMixinProps struct {
	// The Amazon Resource Name (ARN) of the Secrets Manager secret that contains the access token for the third-party Linux subscription provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerlinuxsubscriptions-subscriptionprovider.html#cfn-licensemanagerlinuxsubscriptions-subscriptionprovider-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The Linux subscription provider that the registration is for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerlinuxsubscriptions-subscriptionprovider.html#cfn-licensemanagerlinuxsubscriptions-subscriptionprovider-subscriptionprovidersource
	//
	SubscriptionProviderSource *string `field:"optional" json:"subscriptionProviderSource" yaml:"subscriptionProviderSource"`
	// The metadata tags assigned to the registered subscription provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerlinuxsubscriptions-subscriptionprovider.html#cfn-licensemanagerlinuxsubscriptions-subscriptionprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

