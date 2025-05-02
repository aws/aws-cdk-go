package awscertificatemanager


// Object containing expiration events options associated with an AWS account .
//
// For more information, see [ExpiryEventsConfiguration](https://docs.aws.amazon.com/acm/latest/APIReference/API_ExpiryEventsConfiguration.html) in the API reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expiryEventsConfigurationProperty := &ExpiryEventsConfigurationProperty{
//   	DaysBeforeExpiry: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-account-expiryeventsconfiguration.html
//
type CfnAccount_ExpiryEventsConfigurationProperty struct {
	// This option specifies the number of days prior to certificate expiration when ACM starts generating `EventBridge` events.
	//
	// ACM sends one event per day per certificate until the certificate expires. By default, accounts receive events starting 45 days before certificate expiration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-account-expiryeventsconfiguration.html#cfn-certificatemanager-account-expiryeventsconfiguration-daysbeforeexpiry
	//
	DaysBeforeExpiry *float64 `field:"optional" json:"daysBeforeExpiry" yaml:"daysBeforeExpiry"`
}

