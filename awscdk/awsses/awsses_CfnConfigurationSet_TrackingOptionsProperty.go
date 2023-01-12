package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trackingOptionsProperty := &trackingOptionsProperty{
//   	customRedirectDomain: jsii.String("customRedirectDomain"),
//   }
//
type CfnConfigurationSet_TrackingOptionsProperty struct {
	// `CfnConfigurationSet.TrackingOptionsProperty.CustomRedirectDomain`.
	CustomRedirectDomain *string `field:"optional" json:"customRedirectDomain" yaml:"customRedirectDomain"`
}

