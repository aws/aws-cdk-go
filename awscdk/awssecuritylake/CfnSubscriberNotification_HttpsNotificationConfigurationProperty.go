package awssecuritylake


// Specify the configurations you want to use for HTTPS subscriber notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpsNotificationConfigurationProperty := &HttpsNotificationConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	TargetRoleArn: jsii.String("targetRoleArn"),
//
//   	// the properties below are optional
//   	AuthorizationApiKeyName: jsii.String("authorizationApiKeyName"),
//   	AuthorizationApiKeyValue: jsii.String("authorizationApiKeyValue"),
//   	HttpMethod: jsii.String("httpMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html
//
type CfnSubscriberNotification_HttpsNotificationConfigurationProperty struct {
	// The subscription endpoint in Security Lake .
	//
	// If you prefer notification with an HTTPS endpoint, populate this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
	//
	// For more information about ARNs and how to use them in policies, see [Managing data access](https://docs.aws.amazon.com///security-lake/latest/userguide/subscriber-data-access.html) and [AWS Managed Policies](https://docs.aws.amazon.com//security-lake/latest/userguide/security-iam-awsmanpol.html) in the *Amazon Security Lake User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-targetrolearn
	//
	TargetRoleArn *string `field:"required" json:"targetRoleArn" yaml:"targetRoleArn"`
	// The key name for the notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-authorizationapikeyname
	//
	AuthorizationApiKeyName *string `field:"optional" json:"authorizationApiKeyName" yaml:"authorizationApiKeyName"`
	// The key value for the notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-authorizationapikeyvalue
	//
	AuthorizationApiKeyValue *string `field:"optional" json:"authorizationApiKeyValue" yaml:"authorizationApiKeyValue"`
	// The HTTPS method used for the notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-httpmethod
	//
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
}

