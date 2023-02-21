package awscognito


// The SMS configuration type that includes the settings the Cognito User Pool needs to call for the Amazon SNS service to send an SMS message from your AWS account .
//
// The Cognito User Pool makes the request to the Amazon SNS Service by using an IAM role that you provide for your AWS account .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smsConfigurationProperty := &smsConfigurationProperty{
//   	externalId: jsii.String("externalId"),
//   	snsCallerArn: jsii.String("snsCallerArn"),
//   	snsRegion: jsii.String("snsRegion"),
//   }
//
type CfnUserPool_SmsConfigurationProperty struct {
	// The external ID is a value.
	//
	// We recommend you use `ExternalId` to add security to your IAM role, which is used to call Amazon SNS to send SMS messages for your user pool. If you provide an `ExternalId` , the Cognito User Pool uses it when attempting to assume your IAM role. You can also set your roles trust policy to require the `ExternalID` . If you use the Cognito Management Console to create a role for SMS MFA, Cognito creates a role with the required permissions and a trust policy that uses `ExternalId` .
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The Amazon Resource Name (ARN) of the Amazon SNS caller.
	//
	// This is the ARN of the IAM role in your AWS account that Amazon Cognito will use to send SMS messages. SMS messages are subject to a [spending limit](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html) .
	SnsCallerArn *string `field:"optional" json:"snsCallerArn" yaml:"snsCallerArn"`
	// The AWS Region to use with Amazon SNS integration.
	//
	// You can choose the same Region as your user pool, or a supported *Legacy Amazon SNS alternate Region* .
	//
	// Amazon Cognito resources in the Asia Pacific (Seoul) AWS Region must use your Amazon SNS configuration in the Asia Pacific (Tokyo) Region. For more information, see [SMS message settings for Amazon Cognito user pools](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html) .
	SnsRegion *string `field:"optional" json:"snsRegion" yaml:"snsRegion"`
}

