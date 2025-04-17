package awspinpoint


// Specifies the content and settings for a push notification that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageProperty := &MessageProperty{
//   	Action: jsii.String("action"),
//   	Body: jsii.String("body"),
//   	ImageIconUrl: jsii.String("imageIconUrl"),
//   	ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   	ImageUrl: jsii.String("imageUrl"),
//   	JsonBody: jsii.String("jsonBody"),
//   	MediaUrl: jsii.String("mediaUrl"),
//   	RawContent: jsii.String("rawContent"),
//   	SilentPush: jsii.Boolean(false),
//   	TimeToLive: jsii.Number(123),
//   	Title: jsii.String("title"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html
//
type CfnCampaign_MessageProperty struct {
	// The action to occur if a recipient taps the push notification. Valid values are:.
	//
	// - `OPEN_APP` – Your app opens or it becomes the foreground app if it was sent to the background. This is the default action.
	// - `DEEP_LINK` – Your app opens and displays a designated user interface in the app. This setting uses the deep-linking features of iOS and Android.
	// - `URL` – The default mobile browser on the recipient's device opens and loads the web page at a URL that you specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The body of the notification message.
	//
	// The maximum number of characters is 200.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The URL of the image to display as the push notification icon, such as the icon for the app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-imageiconurl
	//
	ImageIconUrl *string `field:"optional" json:"imageIconUrl" yaml:"imageIconUrl"`
	// The URL of the image to display as the small, push notification icon, such as a small version of the icon for the app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-imagesmalliconurl
	//
	ImageSmallIconUrl *string `field:"optional" json:"imageSmallIconUrl" yaml:"imageSmallIconUrl"`
	// The URL of an image to display in the push notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-imageurl
	//
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The JSON payload to use for a silent push notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-jsonbody
	//
	JsonBody *string `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// The URL of the image or video to display in the push notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-mediaurl
	//
	MediaUrl *string `field:"optional" json:"mediaUrl" yaml:"mediaUrl"`
	// The raw, JSON-formatted string to use as the payload for the notification message.
	//
	// If specified, this value overrides all other content for the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-rawcontent
	//
	RawContent *string `field:"optional" json:"rawContent" yaml:"rawContent"`
	// Specifies whether the notification is a silent push notification, which is a push notification that doesn't display on a recipient's device.
	//
	// Silent push notifications can be used for cases such as updating an app's configuration, displaying messages in an in-app message center, or supporting phone home functionality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-silentpush
	//
	SilentPush interface{} `field:"optional" json:"silentPush" yaml:"silentPush"`
	// The number of seconds that the push notification service should keep the message, if the service is unable to deliver the notification the first time.
	//
	// This value is converted to an expiration value when it's sent to a push notification service. If this value is `0` , the service treats the notification as if it expires immediately and the service doesn't store or try to deliver the notification again.
	//
	// This value doesn't apply to messages that are sent through the Amazon Device Messaging (ADM) service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-timetolive
	//
	TimeToLive *float64 `field:"optional" json:"timeToLive" yaml:"timeToLive"`
	// The title to display above the notification message on a recipient's device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL to open in a recipient's default mobile browser, if a recipient taps the push notification and the value of the `Action` property is `URL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

