package previewawssecuritylakemixins


// Specify the configurations you want to use for HTTPS subscriber notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpsNotificationConfigurationProperty := &HttpsNotificationConfigurationProperty{
//   	AuthorizationApiKeyName: jsii.String("authorizationApiKeyName"),
//   	AuthorizationApiKeyValue: jsii.String("authorizationApiKeyValue"),
//   	Endpoint: jsii.String("endpoint"),
//   	HttpMethod: jsii.String("httpMethod"),
//   	TargetRoleArn: jsii.String("targetRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html
//
type CfnSubscriberNotificationPropsMixin_HttpsNotificationConfigurationProperty struct {
	// The key name for the notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-authorizationapikeyname
	//
	AuthorizationApiKeyName *string `field:"optional" json:"authorizationApiKeyName" yaml:"authorizationApiKeyName"`
	// The key value for the notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-authorizationapikeyvalue
	//
	AuthorizationApiKeyValue *string `field:"optional" json:"authorizationApiKeyValue" yaml:"authorizationApiKeyValue"`
	// The subscription endpoint in Security Lake .
	//
	// If you prefer notification with an HTTPS endpoint, populate this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The HTTPS method used for the notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-httpmethod
	//
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
	//
	// For more information about ARNs and how to use them in policies, see [Managing data access](https://docs.aws.amazon.com///security-lake/latest/userguide/subscriber-data-access.html) and [AWS Managed Policies](https://docs.aws.amazon.com//security-lake/latest/userguide/security-iam-awsmanpol.html) in the *Amazon Security Lake User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.html#cfn-securitylake-subscribernotification-httpsnotificationconfiguration-targetrolearn
	//
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

