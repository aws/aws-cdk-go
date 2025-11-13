package interfacesawsiot


// A reference to a BillingGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billingGroupReference := &BillingGroupReference{
//   	BillingGroupArn: jsii.String("billingGroupArn"),
//   	BillingGroupName: jsii.String("billingGroupName"),
//   }
//
type BillingGroupReference struct {
	// The ARN of the BillingGroup resource.
	BillingGroupArn *string `field:"required" json:"billingGroupArn" yaml:"billingGroupArn"`
	// The BillingGroupName of the BillingGroup resource.
	BillingGroupName *string `field:"required" json:"billingGroupName" yaml:"billingGroupName"`
}

