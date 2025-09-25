package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProtectConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProtectConfigurationProps := &CfnProtectConfigurationProps{
//   	CountryRuleSet: &CountryRuleSetProperty{
//   		Mms: []interface{}{
//   			&CountryRuleProperty{
//   				CountryCode: jsii.String("countryCode"),
//   				ProtectStatus: jsii.String("protectStatus"),
//   			},
//   		},
//   		Sms: []interface{}{
//   			&CountryRuleProperty{
//   				CountryCode: jsii.String("countryCode"),
//   				ProtectStatus: jsii.String("protectStatus"),
//   			},
//   		},
//   		Voice: []interface{}{
//   			&CountryRuleProperty{
//   				CountryCode: jsii.String("countryCode"),
//   				ProtectStatus: jsii.String("protectStatus"),
//   			},
//   		},
//   	},
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-protectconfiguration.html
//
type CfnProtectConfigurationProps struct {
	// The set of `CountryRules` you specify to control which countries End User Messaging  can send your messages to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-protectconfiguration.html#cfn-smsvoice-protectconfiguration-countryruleset
	//
	CountryRuleSet interface{} `field:"optional" json:"countryRuleSet" yaml:"countryRuleSet"`
	// The status of deletion protection for the protect configuration.
	//
	// When set to true deletion protection is enabled. By default this is set to false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-protectconfiguration.html#cfn-smsvoice-protectconfiguration-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// An array of key and value pair tags that are associated with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-protectconfiguration.html#cfn-smsvoice-protectconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

