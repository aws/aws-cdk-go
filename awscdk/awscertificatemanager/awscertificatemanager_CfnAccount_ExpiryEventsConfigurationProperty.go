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
//   expiryEventsConfigurationProperty := &expiryEventsConfigurationProperty{
//   	daysBeforeExpiry: jsii.Number(123),
//   }
//
type CfnAccount_ExpiryEventsConfigurationProperty struct {
	// This option specifies the number of days prior to certificate expiration when ACM starts generating `EventBridge` events.
	//
	// ACM sends one event per day per certificate until the certificate expires. By default, accounts receive events starting 45 days before certificate expiration.
	DaysBeforeExpiry *float64 `field:"optional" json:"daysBeforeExpiry" yaml:"daysBeforeExpiry"`
}

