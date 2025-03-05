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
//   cfnPushTemplateProps := &CfnPushTemplateProps{
//   	TemplateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	Adm: &AndroidPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		SmallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Apns: &APNSPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		MediaUrl: jsii.String("mediaUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Baidu: &AndroidPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		SmallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Default: &DefaultPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	DefaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	Gcm: &AndroidPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		SmallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html
//
type CfnPushTemplateProps struct {
	// The name of the message template to use for the message.
	//
	// If specified, this value must match the name of an existing message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatename
	//
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The message template to use for the ADM (Amazon Device Messaging) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-adm
	//
	Adm interface{} `field:"optional" json:"adm" yaml:"adm"`
	// The message template to use for the APNs (Apple Push Notification service) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-apns
	//
	Apns interface{} `field:"optional" json:"apns" yaml:"apns"`
	// The message template to use for the Baidu (Baidu Cloud Push) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-baidu
	//
	Baidu interface{} `field:"optional" json:"baidu" yaml:"baidu"`
	// The default message template to use for push notification channels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-default
	//
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-defaultsubstitutions
	//
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// The message template to use for the GCM channel, which is used to send notifications through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-gcm
	//
	Gcm interface{} `field:"optional" json:"gcm" yaml:"gcm"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatedescription
	//
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

