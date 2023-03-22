package awspinpoint


// Specifies channel-specific content and settings for a message template that can be used in push notifications that are sent through the ADM (Amazon Device Messaging), Baidu (Baidu Cloud Push), or GCM (Firebase Cloud Messaging, formerly Google Cloud Messaging) channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   androidPushNotificationTemplateProperty := &androidPushNotificationTemplateProperty{
//   	action: jsii.String("action"),
//   	body: jsii.String("body"),
//   	imageIconUrl: jsii.String("imageIconUrl"),
//   	imageUrl: jsii.String("imageUrl"),
//   	smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   	sound: jsii.String("sound"),
//   	title: jsii.String("title"),
//   	url: jsii.String("url"),
//   }
//
type CfnPushTemplate_AndroidPushNotificationTemplateProperty struct {
	// The action to occur if a recipient taps a push notification that's based on the message template.
	//
	// Valid values are:
	//
	// - `OPEN_APP` – Your app opens or it becomes the foreground app if it was sent to the background. This is the default action.
	// - `DEEP_LINK` – Your app opens and displays a designated user interface in the app. This action uses the deep-linking features of the Android platform.
	// - `URL` – The default mobile browser on the recipient's device opens and loads the web page at a URL that you specify.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The message body to use in a push notification that's based on the message template.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The URL of the large icon image to display in the content view of a push notification that's based on the message template.
	ImageIconUrl *string `field:"optional" json:"imageIconUrl" yaml:"imageIconUrl"`
	// The URL of an image to display in a push notification that's based on the message template.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The URL of the small icon image to display in the status bar and the content view of a push notification that's based on the message template.
	SmallImageIconUrl *string `field:"optional" json:"smallImageIconUrl" yaml:"smallImageIconUrl"`
	// The sound to play when a recipient receives a push notification that's based on the message template.
	//
	// You can use the default stream or specify the file name of a sound resource that's bundled in your app. On an Android platform, the sound file must reside in `/res/raw/` .
	Sound *string `field:"optional" json:"sound" yaml:"sound"`
	// The title to use in a push notification that's based on the message template.
	//
	// This title appears above the notification message on a recipient's device.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL to open in a recipient's default mobile browser, if a recipient taps a push notification that's based on the message template and the value of the `Action` property is `URL` .
	Url *string `field:"optional" json:"url" yaml:"url"`
}

