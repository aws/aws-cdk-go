package awsqbusiness


// Subscription configuration information for an Amazon Q Business application using IAM identity federation for user management.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoSubscriptionConfigurationProperty := &AutoSubscriptionConfigurationProperty{
//   	AutoSubscribe: jsii.String("autoSubscribe"),
//
//   	// the properties below are optional
//   	DefaultSubscriptionType: jsii.String("defaultSubscriptionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-autosubscriptionconfiguration.html
//
type CfnApplication_AutoSubscriptionConfigurationProperty struct {
	// Describes whether automatic subscriptions are enabled for an Amazon Q Business application using IAM identity federation for user management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-autosubscriptionconfiguration.html#cfn-qbusiness-application-autosubscriptionconfiguration-autosubscribe
	//
	AutoSubscribe *string `field:"required" json:"autoSubscribe" yaml:"autoSubscribe"`
	// Describes the default subscription type assigned to an Amazon Q Business application using IAM identity federation for user management.
	//
	// If the value for `autoSubscribe` is set to `ENABLED` you must select a value for this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-autosubscriptionconfiguration.html#cfn-qbusiness-application-autosubscriptionconfiguration-defaultsubscriptiontype
	//
	DefaultSubscriptionType *string `field:"optional" json:"defaultSubscriptionType" yaml:"defaultSubscriptionType"`
}

