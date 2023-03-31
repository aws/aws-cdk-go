package awss3


// Amazon S3 can send events to Amazon EventBridge whenever certain events happen in your bucket, see [Using EventBridge](https://docs.aws.amazon.com/AmazonS3/latest/userguide/EventBridge.html) in the *Amazon S3 User Guide* .
//
// Unlike other destinations, delivery of events to EventBridge can be either enabled or disabled for a bucket. If enabled, all events will be sent to EventBridge and you can use EventBridge rules to route events to additional targets. For more information, see [What Is Amazon EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html) in the *Amazon EventBridge User Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeConfigurationProperty := &eventBridgeConfigurationProperty{
//   	eventBridgeEnabled: jsii.Boolean(false),
//   }
//
type CfnBucket_EventBridgeConfigurationProperty struct {
	// Enables delivery of events to Amazon EventBridge.
	EventBridgeEnabled interface{} `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
}

