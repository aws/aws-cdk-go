package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNotifyConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnNotifyConfigurationMixinProps := &CfnNotifyConfigurationMixinProps{
//   	DefaultTemplateId: jsii.String("defaultTemplateId"),
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	DisplayName: jsii.String("displayName"),
//   	EnabledChannels: []*string{
//   		jsii.String("enabledChannels"),
//   	},
//   	EnabledCountries: []*string{
//   		jsii.String("enabledCountries"),
//   	},
//   	PoolId: jsii.String("poolId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseCase: jsii.String("useCase"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html
//
type CfnNotifyConfigurationMixinProps struct {
	// The default template identifier to associate with the notify configuration.
	//
	// If specified, this template is used when sending messages without an explicit template identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-defaulttemplateid
	//
	DefaultTemplateId *string `field:"optional" json:"defaultTemplateId" yaml:"defaultTemplateId"`
	// By default this is set to false.
	//
	// When set to true the notify configuration can't be deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// The display name to associate with the notify configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// An array of channels to enable for the notify configuration.
	//
	// Supported values include SMS and VOICE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-enabledchannels
	//
	EnabledChannels *[]*string `field:"optional" json:"enabledChannels" yaml:"enabledChannels"`
	// An array of two-character ISO country codes, in ISO 3166-1 alpha-2 format, that are enabled for the notify configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-enabledcountries
	//
	EnabledCountries *[]*string `field:"optional" json:"enabledCountries" yaml:"enabledCountries"`
	// The identifier of the pool to associate with the notify configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-poolid
	//
	PoolId *string `field:"optional" json:"poolId" yaml:"poolId"`
	// An array of tags (key and value pairs) associated with the notify configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The use case for the notify configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html#cfn-smsvoice-notifyconfiguration-usecase
	//
	UseCase *string `field:"optional" json:"useCase" yaml:"useCase"`
}

