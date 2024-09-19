package awssecuritylake


// Properties for defining a `CfnSubscriberNotification`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sqsNotificationConfiguration interface{}
//
//   cfnSubscriberNotificationProps := &CfnSubscriberNotificationProps{
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		HttpsNotificationConfiguration: &HttpsNotificationConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			TargetRoleArn: jsii.String("targetRoleArn"),
//
//   			// the properties below are optional
//   			AuthorizationApiKeyName: jsii.String("authorizationApiKeyName"),
//   			AuthorizationApiKeyValue: jsii.String("authorizationApiKeyValue"),
//   			HttpMethod: jsii.String("httpMethod"),
//   		},
//   		SqsNotificationConfiguration: sqsNotificationConfiguration,
//   	},
//   	SubscriberArn: jsii.String("subscriberArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscribernotification.html
//
type CfnSubscriberNotificationProps struct {
	// Specify the configurations you want to use for subscriber notification.
	//
	// The subscriber is notified when new data is written to the data lake for sources that the subscriber consumes in Security Lake .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscribernotification.html#cfn-securitylake-subscribernotification-notificationconfiguration
	//
	NotificationConfiguration interface{} `field:"required" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// The Amazon Resource Name (ARN) of the Security Lake subscriber.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscribernotification.html#cfn-securitylake-subscribernotification-subscriberarn
	//
	SubscriberArn *string `field:"required" json:"subscriberArn" yaml:"subscriberArn"`
}

