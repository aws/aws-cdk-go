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
//   messageConfigurationProperty := &messageConfigurationProperty{
//   	admMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	apnsMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	baiduMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	customMessage: &campaignCustomMessageProperty{
//   		data: jsii.String("data"),
//   	},
//   	defaultMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	emailMessage: &campaignEmailMessageProperty{
//   		body: jsii.String("body"),
//   		fromAddress: jsii.String("fromAddress"),
//   		htmlBody: jsii.String("htmlBody"),
//   		title: jsii.String("title"),
//   	},
//   	gcmMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	inAppMessage: &campaignInAppMessageProperty{
//   		content: []interface{}{
//   			&inAppMessageContentProperty{
//   				backgroundColor: jsii.String("backgroundColor"),
//   				bodyConfig: &inAppMessageBodyConfigProperty{
//   					alignment: jsii.String("alignment"),
//   					body: jsii.String("body"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				headerConfig: &inAppMessageHeaderConfigProperty{
//   					alignment: jsii.String("alignment"),
//   					header: jsii.String("header"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				imageUrl: jsii.String("imageUrl"),
//   				primaryBtn: &inAppMessageButtonProperty{
//   					android: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					defaultConfig: &defaultButtonConfigurationProperty{
//   						backgroundColor: jsii.String("backgroundColor"),
//   						borderRadius: jsii.Number(123),
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   						text: jsii.String("text"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					ios: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					web: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   				},
//   				secondaryBtn: &inAppMessageButtonProperty{
//   					android: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					defaultConfig: &defaultButtonConfigurationProperty{
//   						backgroundColor: jsii.String("backgroundColor"),
//   						borderRadius: jsii.Number(123),
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   						text: jsii.String("text"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					ios: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					web: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   				},
//   			},
//   		},
//   		customConfig: customConfig,
//   		layout: jsii.String("layout"),
//   	},
//   	smsMessage: &campaignSmsMessageProperty{
//   		body: jsii.String("body"),
//   		entityId: jsii.String("entityId"),
//   		messageType: jsii.String("messageType"),
//   		originationNumber: jsii.String("originationNumber"),
//   		senderId: jsii.String("senderId"),
//   		templateId: jsii.String("templateId"),
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

