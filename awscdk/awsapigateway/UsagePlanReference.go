package awsapigateway


// A reference to a UsagePlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   usagePlanReference := &UsagePlanReference{
//   	UsagePlanId: jsii.String("usagePlanId"),
//   }
//
type UsagePlanReference struct {
	// The Id of the UsagePlan resource.
	UsagePlanId *string `field:"required" json:"usagePlanId" yaml:"usagePlanId"`
}

