package awspinpoint


// Properties for defining a `CfnPushTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnPushTemplateProps := &cfnPushTemplateProps{
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	adm: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	apns: &aPNSPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	baidu: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	default: &defaultPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	gcm: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   }
//
type CfnPushTemplateProps struct {
	// The name of the message template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The message template to use for the ADM (Amazon Device Messaging) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Adm interface{} `field:"optional" json:"adm" yaml:"adm"`
	// The message template to use for the APNs (Apple Push Notification service) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Apns interface{} `field:"optional" json:"apns" yaml:"apns"`
	// The message template to use for the Baidu (Baidu Cloud Push) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Baidu interface{} `field:"optional" json:"baidu" yaml:"baidu"`
	// The default message template to use for push notification channels.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// The message template to use for the GCM channel, which is used to send notifications through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Gcm interface{} `field:"optional" json:"gcm" yaml:"gcm"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

