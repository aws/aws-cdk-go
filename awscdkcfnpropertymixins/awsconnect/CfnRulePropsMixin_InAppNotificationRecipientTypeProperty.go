package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inAppNotificationRecipientTypeProperty := &InAppNotificationRecipientTypeProperty{
//   	UserArns: []*string{
//   		jsii.String("userArns"),
//   	},
//   	UserTags: map[string]*string{
//   		"userTagsKey": jsii.String("userTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-inappnotificationrecipienttype.html
//
type CfnRulePropsMixin_InAppNotificationRecipientTypeProperty struct {
	// The list of recipients by user arns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-inappnotificationrecipienttype.html#cfn-connect-rule-inappnotificationrecipienttype-userarns
	//
	UserArns *[]*string `field:"optional" json:"userArns" yaml:"userArns"`
	// The collection of recipients who are identified by user tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-inappnotificationrecipienttype.html#cfn-connect-rule-inappnotificationrecipienttype-usertags
	//
	UserTags interface{} `field:"optional" json:"userTags" yaml:"userTags"`
}

