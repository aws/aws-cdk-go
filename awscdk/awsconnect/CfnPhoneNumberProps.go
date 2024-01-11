package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPhoneNumber`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPhoneNumberProps := &CfnPhoneNumberProps{
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	CountryCode: jsii.String("countryCode"),
//   	Description: jsii.String("description"),
//   	Prefix: jsii.String("prefix"),
//   	SourcePhoneNumberArn: jsii.String("sourcePhoneNumberArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html
//
type CfnPhoneNumberProps struct {
	// The Amazon Resource Name (ARN) for Amazon Connect instances or traffic distribution group that phone numbers are claimed to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html#cfn-connect-phonenumber-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The ISO country code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html#cfn-connect-phonenumber-countrycode
	//
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// The description of the phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html#cfn-connect-phonenumber-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The prefix of the phone number. If provided, it must contain `+` as part of the country code.
	//
	// *Pattern* : `^\\+[0-9]{1,15}`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html#cfn-connect-phonenumber-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The claimed phone number ARN that was previously imported from the external service, such as Amazon Pinpoint.
	//
	// If it is from Amazon Pinpoint, it looks like the ARN of the phone number that was imported from Amazon Pinpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html#cfn-connect-phonenumber-sourcephonenumberarn
	//
	SourcePhoneNumberArn *string `field:"optional" json:"sourcePhoneNumberArn" yaml:"sourcePhoneNumberArn"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html#cfn-connect-phonenumber-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-phonenumber.html#cfn-connect-phonenumber-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

