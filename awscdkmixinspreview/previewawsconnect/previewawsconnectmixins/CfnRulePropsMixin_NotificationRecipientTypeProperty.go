package previewawsconnectmixins


// The type of notification recipient.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notificationRecipientTypeProperty := &NotificationRecipientTypeProperty{
//   	UserArns: []*string{
//   		jsii.String("userArns"),
//   	},
//   	UserTags: map[string]*string{
//   		"userTagsKey": jsii.String("userTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-notificationrecipienttype.html
//
type CfnRulePropsMixin_NotificationRecipientTypeProperty struct {
	// The Amazon Resource Name (ARN) of the user account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-notificationrecipienttype.html#cfn-connect-rule-notificationrecipienttype-userarns
	//
	UserArns *[]*string `field:"optional" json:"userArns" yaml:"userArns"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }. Amazon Connect users with the specified tags will be notified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-notificationrecipienttype.html#cfn-connect-rule-notificationrecipienttype-usertags
	//
	UserTags interface{} `field:"optional" json:"userTags" yaml:"userTags"`
}

