package awscertificatemanager


// Properties for defining a `CfnAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountProps := &cfnAccountProps{
//   	expiryEventsConfiguration: &expiryEventsConfigurationProperty{
//   		daysBeforeExpiry: jsii.Number(123),
//   	},
//   }
//
type CfnAccountProps struct {
	// Object containing expiration events options associated with an AWS account .
	//
	// For more information, see [ExpiryEventsConfiguration](https://docs.aws.amazon.com/acm/latest/APIReference/API_ExpiryEventsConfiguration.html) in the API reference.
	ExpiryEventsConfiguration interface{} `field:"required" json:"expiryEventsConfiguration" yaml:"expiryEventsConfiguration"`
}

