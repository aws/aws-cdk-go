package awsstepfunctions


// Properties for defining a `CfnActivity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnActivityProps := &CfnActivityProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		KmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html
//
type CfnActivityProps struct {
	// The name of the activity.
	//
	// A name must *not* contain:
	//
	// - white space
	// - brackets `< > { } [ ]`
	// - wildcard characters `? *`
	// - special characters `" # % \ ^ | ~ ` $ & , ; : /`
	// - control characters ( `U+0000-001F` , `U+007F-009F` , `U+FFFE-FFFF` )
	// - surrogates ( `U+D800-DFFF` )
	// - invalid characters ( `U+10FFFF` )
	//
	// To enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and _.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html#cfn-stepfunctions-activity-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Encryption configuration for the activity.
	//
	// Activity configuration is immutable, and resource names must be unique. To set customer managed keys for encryption, you must create a *new Activity* . If you attempt to change the configuration in your CFN template for an existing activity, you will receive an `ActivityAlreadyExists` exception.
	//
	// To update your activity to include customer managed keys, set a new activity name within your CloudFormation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html#cfn-stepfunctions-activity-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The list of tags to add to a resource.
	//
	// Tags may only contain Unicode letters, digits, white space, or these symbols: `_ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html#cfn-stepfunctions-activity-tags
	//
	Tags *[]*CfnActivity_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

