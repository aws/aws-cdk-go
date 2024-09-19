package awssecuritylake


// Specify the configurations you want to use for subscriber notification.
//
// The subscriber is notified when new data is written to the data lake for sources that the subscriber consumes in Security Lake .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sqsNotificationConfiguration interface{}
//
//   notificationConfigurationProperty := &NotificationConfigurationProperty{
//   	HttpsNotificationConfiguration: &HttpsNotificationConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		TargetRoleArn: jsii.String("targetRoleArn"),
//
//   		// the properties below are optional
//   		AuthorizationApiKeyName: jsii.String("authorizationApiKeyName"),
//   		AuthorizationApiKeyValue: jsii.String("authorizationApiKeyValue"),
//   		HttpMethod: jsii.String("httpMethod"),
//   	},
//   	SqsNotificationConfiguration: sqsNotificationConfiguration,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-notificationconfiguration.html
//
type CfnSubscriberNotification_NotificationConfigurationProperty struct {
	// The configurations used for HTTPS subscriber notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-notificationconfiguration.html#cfn-securitylake-subscribernotification-notificationconfiguration-httpsnotificationconfiguration
	//
	HttpsNotificationConfiguration interface{} `field:"optional" json:"httpsNotificationConfiguration" yaml:"httpsNotificationConfiguration"`
	// The configurations for SQS subscriber notification.
	//
	// The members of this structure are context-dependent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-notificationconfiguration.html#cfn-securitylake-subscribernotification-notificationconfiguration-sqsnotificationconfiguration
	//
	SqsNotificationConfiguration interface{} `field:"optional" json:"sqsNotificationConfiguration" yaml:"sqsNotificationConfiguration"`
}

