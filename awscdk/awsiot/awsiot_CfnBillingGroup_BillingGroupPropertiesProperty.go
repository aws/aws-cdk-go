package awsiot


// The properties of a billing group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billingGroupPropertiesProperty := &BillingGroupPropertiesProperty{
//   	BillingGroupDescription: jsii.String("billingGroupDescription"),
//   }
//
type CfnBillingGroup_BillingGroupPropertiesProperty struct {
	// The description of the billing group.
	BillingGroupDescription *string `field:"optional" json:"billingGroupDescription" yaml:"billingGroupDescription"`
}

