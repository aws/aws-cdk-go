package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sendInAppNotificationActionProperty := &SendInAppNotificationActionProperty{
//   	Content: &NotificationContentProperty{
//   		DeDe: jsii.String("deDe"),
//   		EnUs: jsii.String("enUs"),
//   		EsEs: jsii.String("esEs"),
//   		FrFr: jsii.String("frFr"),
//   		IdId: jsii.String("idId"),
//   		ItIt: jsii.String("itIt"),
//   		JaJp: jsii.String("jaJp"),
//   		KoKr: jsii.String("koKr"),
//   		PtBr: jsii.String("ptBr"),
//   		ZhCn: jsii.String("zhCn"),
//   		ZhTw: jsii.String("zhTw"),
//   	},
//   	Recipient: &InAppNotificationRecipientTypeProperty{
//   		UserArns: []*string{
//   			jsii.String("userArns"),
//   		},
//   		UserTags: map[string]*string{
//   			"userTagsKey": jsii.String("userTags"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Exclusion: &InAppNotificationRecipientTypeProperty{
//   		UserArns: []*string{
//   			jsii.String("userArns"),
//   		},
//   		UserTags: map[string]*string{
//   			"userTagsKey": jsii.String("userTags"),
//   		},
//   	},
//   	Priority: jsii.String("priority"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendinappnotificationaction.html
//
type CfnRule_SendInAppNotificationActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendinappnotificationaction.html#cfn-connect-rule-sendinappnotificationaction-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendinappnotificationaction.html#cfn-connect-rule-sendinappnotificationaction-recipient
	//
	Recipient interface{} `field:"required" json:"recipient" yaml:"recipient"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendinappnotificationaction.html#cfn-connect-rule-sendinappnotificationaction-exclusion
	//
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
	// The priority of the notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendinappnotificationaction.html#cfn-connect-rule-sendinappnotificationaction-priority
	//
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
}

