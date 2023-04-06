package awspinpoint


// Specifies the message configuration settings for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//
//   messageConfigurationProperty := &MessageConfigurationProperty{
//   	AdmMessage: &MessageProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		JsonBody: jsii.String("jsonBody"),
//   		MediaUrl: jsii.String("mediaUrl"),
//   		RawContent: jsii.String("rawContent"),
//   		SilentPush: jsii.Boolean(false),
//   		TimeToLive: jsii.Number(123),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	ApnsMessage: &MessageProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		JsonBody: jsii.String("jsonBody"),
//   		MediaUrl: jsii.String("mediaUrl"),
//   		RawContent: jsii.String("rawContent"),
//   		SilentPush: jsii.Boolean(false),
//   		TimeToLive: jsii.Number(123),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	BaiduMessage: &MessageProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		JsonBody: jsii.String("jsonBody"),
//   		MediaUrl: jsii.String("mediaUrl"),
//   		RawContent: jsii.String("rawContent"),
//   		SilentPush: jsii.Boolean(false),
//   		TimeToLive: jsii.Number(123),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	CustomMessage: &CampaignCustomMessageProperty{
//   		Data: jsii.String("data"),
//   	},
//   	DefaultMessage: &MessageProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		JsonBody: jsii.String("jsonBody"),
//   		MediaUrl: jsii.String("mediaUrl"),
//   		RawContent: jsii.String("rawContent"),
//   		SilentPush: jsii.Boolean(false),
//   		TimeToLive: jsii.Number(123),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	EmailMessage: &CampaignEmailMessageProperty{
//   		Body: jsii.String("body"),
//   		FromAddress: jsii.String("fromAddress"),
//   		HtmlBody: jsii.String("htmlBody"),
//   		Title: jsii.String("title"),
//   	},
//   	GcmMessage: &MessageProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		JsonBody: jsii.String("jsonBody"),
//   		MediaUrl: jsii.String("mediaUrl"),
//   		RawContent: jsii.String("rawContent"),
//   		SilentPush: jsii.Boolean(false),
//   		TimeToLive: jsii.Number(123),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	InAppMessage: &CampaignInAppMessageProperty{
//   		Content: []interface{}{
//   			&InAppMessageContentProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				BodyConfig: &InAppMessageBodyConfigProperty{
//   					Alignment: jsii.String("alignment"),
//   					Body: jsii.String("body"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				HeaderConfig: &InAppMessageHeaderConfigProperty{
//   					Alignment: jsii.String("alignment"),
//   					Header: jsii.String("header"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				ImageUrl: jsii.String("imageUrl"),
//   				PrimaryBtn: &InAppMessageButtonProperty{
//   					Android: &OverrideButtonConfigurationProperty{
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   					},
//   					DefaultConfig: &DefaultButtonConfigurationProperty{
//   						BackgroundColor: jsii.String("backgroundColor"),
//   						BorderRadius: jsii.Number(123),
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   						Text: jsii.String("text"),
//   						TextColor: jsii.String("textColor"),
//   					},
//   					Ios: &OverrideButtonConfigurationProperty{
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   					},
//   					Web: &OverrideButtonConfigurationProperty{
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   					},
//   				},
//   				SecondaryBtn: &InAppMessageButtonProperty{
//   					Android: &OverrideButtonConfigurationProperty{
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   					},
//   					DefaultConfig: &DefaultButtonConfigurationProperty{
//   						BackgroundColor: jsii.String("backgroundColor"),
//   						BorderRadius: jsii.Number(123),
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   						Text: jsii.String("text"),
//   						TextColor: jsii.String("textColor"),
//   					},
//   					Ios: &OverrideButtonConfigurationProperty{
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   					},
//   					Web: &OverrideButtonConfigurationProperty{
//   						ButtonAction: jsii.String("buttonAction"),
//   						Link: jsii.String("link"),
//   					},
//   				},
//   			},
//   		},
//   		CustomConfig: customConfig,
//   		Layout: jsii.String("layout"),
//   	},
//   	SmsMessage: &CampaignSmsMessageProperty{
//   		Body: jsii.String("body"),
//   		EntityId: jsii.String("entityId"),
//   		MessageType: jsii.String("messageType"),
//   		OriginationNumber: jsii.String("originationNumber"),
//   		SenderId: jsii.String("senderId"),
//   		TemplateId: jsii.String("templateId"),
//   	},
//   }
//
type CfnCampaign_MessageConfigurationProperty struct {
	// The message that the campaign sends through the ADM (Amazon Device Messaging) channel.
	//
	// If specified, this message overrides the default message.
	AdmMessage interface{} `field:"optional" json:"admMessage" yaml:"admMessage"`
	// The message that the campaign sends through the APNs (Apple Push Notification service) channel.
	//
	// If specified, this message overrides the default message.
	ApnsMessage interface{} `field:"optional" json:"apnsMessage" yaml:"apnsMessage"`
	// The message that the campaign sends through the Baidu (Baidu Cloud Push) channel.
	//
	// If specified, this message overrides the default message.
	BaiduMessage interface{} `field:"optional" json:"baiduMessage" yaml:"baiduMessage"`
	// The message that the campaign sends through a custom channel, as specified by the delivery configuration ( `CustomDeliveryConfiguration` ) settings for the campaign.
	//
	// If specified, this message overrides the default message.
	CustomMessage interface{} `field:"optional" json:"customMessage" yaml:"customMessage"`
	// The default message that the campaign sends through all the channels that are configured for the campaign.
	DefaultMessage interface{} `field:"optional" json:"defaultMessage" yaml:"defaultMessage"`
	// The message that the campaign sends through the email channel.
	//
	// If specified, this message overrides the default message.
	EmailMessage interface{} `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// The message that the campaign sends through the GCM channel, which enables Amazon Pinpoint to send push notifications through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// If specified, this message overrides the default message.
	GcmMessage interface{} `field:"optional" json:"gcmMessage" yaml:"gcmMessage"`
	// The default message for the in-app messaging channel.
	//
	// This message overrides the default message ( `DefaultMessage` ).
	InAppMessage interface{} `field:"optional" json:"inAppMessage" yaml:"inAppMessage"`
	// The message that the campaign sends through the SMS channel.
	//
	// If specified, this message overrides the default message.
	SmsMessage interface{} `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

