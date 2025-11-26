package previewawspinpointmixins


// Specifies channel-specific content and settings for a message template that can be used in push notifications that are sent through the APNs (Apple Push Notification service) channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aPNSPushNotificationTemplateProperty := &APNSPushNotificationTemplateProperty{
//   	Action: jsii.String("action"),
//   	Body: jsii.String("body"),
//   	MediaUrl: jsii.String("mediaUrl"),
//   	Sound: jsii.String("sound"),
//   	Title: jsii.String("title"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.html
//
type CfnPushTemplatePropsMixin_APNSPushNotificationTemplateProperty struct {
	// The action to occur if a recipient taps a push notification that's based on the message template.
	//
	// Valid values are:
	//
	// - `OPEN_APP` – Your app opens or it becomes the foreground app if it was sent to the background. This is the default action.
	// - `DEEP_LINK` – Your app opens and displays a designated user interface in the app. This setting uses the deep-linking features of the iOS platform.
	// - `URL` – The default mobile browser on the recipient's device opens and loads the web page at a URL that you specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.html#cfn-pinpoint-pushtemplate-apnspushnotificationtemplate-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The message body to use in push notifications that are based on the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.html#cfn-pinpoint-pushtemplate-apnspushnotificationtemplate-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The URL of an image or video to display in push notifications that are based on the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.html#cfn-pinpoint-pushtemplate-apnspushnotificationtemplate-mediaurl
	//
	MediaUrl *string `field:"optional" json:"mediaUrl" yaml:"mediaUrl"`
	// The key for the sound to play when the recipient receives a push notification that's based on the message template.
	//
	// The value for this key is the name of a sound file in your app's main bundle or the `Library/Sounds` folder in your app's data container. If the sound file can't be found or you specify `default` for the value, the system plays the default alert sound.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.html#cfn-pinpoint-pushtemplate-apnspushnotificationtemplate-sound
	//
	Sound *string `field:"optional" json:"sound" yaml:"sound"`
	// The title to use in push notifications that are based on the message template.
	//
	// This title appears above the notification message on a recipient's device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.html#cfn-pinpoint-pushtemplate-apnspushnotificationtemplate-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL to open in the recipient's default mobile browser, if a recipient taps a push notification that's based on the message template and the value of the `Action` property is `URL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.html#cfn-pinpoint-pushtemplate-apnspushnotificationtemplate-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

