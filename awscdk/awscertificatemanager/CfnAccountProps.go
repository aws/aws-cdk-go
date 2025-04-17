package awscertificatemanager


// Properties for defining a `CfnAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountProps := &CfnAccountProps{
//   	ExpiryEventsConfiguration: &ExpiryEventsConfigurationProperty{
//   		DaysBeforeExpiry: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-account.html
//
type CfnAccountProps struct {
	// Object containing expiration events options associated with an AWS account .
	//
	// For more information, see [ExpiryEventsConfiguration](https://docs.aws.amazon.com/acm/latest/APIReference/API_ExpiryEventsConfiguration.html) in the API reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-account.html#cfn-certificatemanager-account-expiryeventsconfiguration
	//
	ExpiryEventsConfiguration interface{} `field:"required" json:"expiryEventsConfiguration" yaml:"expiryEventsConfiguration"`
}

