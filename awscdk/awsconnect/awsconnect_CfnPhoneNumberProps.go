package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPhoneNumber`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPhoneNumberProps := &CfnPhoneNumberProps{
//   	CountryCode: jsii.String("countryCode"),
//   	TargetArn: jsii.String("targetArn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Prefix: jsii.String("prefix"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPhoneNumberProps struct {
	// The ISO country code.
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// The Amazon Resource Name (ARN) for Amazon Connect instances or traffic distribution group that phone numbers are claimed to.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The type of phone number.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the phone number.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The prefix of the phone number. If provided, it must contain `+` as part of the country code.
	//
	// *Pattern* : `^\\+[0-9]{1,15}`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

