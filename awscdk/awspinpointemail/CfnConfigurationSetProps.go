package awspinpointemail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConfigurationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetProps := &CfnConfigurationSetProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DeliveryOptions: &DeliveryOptionsProperty{
//   		SendingPoolName: jsii.String("sendingPoolName"),
//   	},
//   	ReputationOptions: &ReputationOptionsProperty{
//   		ReputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	SendingOptions: &SendingOptionsProperty{
//   		SendingEnabled: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackingOptions: &TrackingOptionsProperty{
//   		CustomRedirectDomain: jsii.String("customRedirectDomain"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html
//
type CfnConfigurationSetProps struct {
	// The name of the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html#cfn-pinpointemail-configurationset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html#cfn-pinpointemail-configurationset-deliveryoptions
	//
	DeliveryOptions interface{} `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// An object that defines whether or not Amazon Pinpoint collects reputation metrics for the emails that you send that use the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html#cfn-pinpointemail-configurationset-reputationoptions
	//
	ReputationOptions interface{} `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// An object that defines whether or not Amazon Pinpoint can send email that you send using the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html#cfn-pinpointemail-configurationset-sendingoptions
	//
	SendingOptions interface{} `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// An object that defines the tags (keys and values) that you want to associate with the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html#cfn-pinpointemail-configurationset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An object that defines the open and click tracking options for emails that you send using the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html#cfn-pinpointemail-configurationset-trackingoptions
	//
	TrackingOptions interface{} `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
}

