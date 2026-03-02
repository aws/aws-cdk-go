package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNotificationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNotificationMixinProps := &CfnNotificationMixinProps{
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
//   	ExpiresAt: jsii.String("expiresAt"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Priority: jsii.String("priority"),
//   	Recipients: []*string{
//   		jsii.String("recipients"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-notification.html
//
type CfnNotificationMixinProps struct {
	// The content of a notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-notification.html#cfn-connect-notification-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// The time a notification will expire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-notification.html#cfn-connect-notification-expiresat
	//
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-notification.html#cfn-connect-notification-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The priority of notification.
	//
	// In the Amazon Connect console, when you create a notification, you are prompted to assign one of the following priorities: High (HIGH) or LOW (LOW).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-notification.html#cfn-connect-notification-priority
	//
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// The recipients of the notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-notification.html#cfn-connect-notification-recipients
	//
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
	// One or more tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-notification.html#cfn-connect-notification-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

