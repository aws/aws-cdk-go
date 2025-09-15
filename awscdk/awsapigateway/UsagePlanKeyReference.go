package awsapigateway


// A reference to a UsagePlanKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   usagePlanKeyReference := &UsagePlanKeyReference{
//   	UsagePlanKeyId: jsii.String("usagePlanKeyId"),
//   }
//
type UsagePlanKeyReference struct {
	// The Id of the UsagePlanKey resource.
	UsagePlanKeyId *string `field:"required" json:"usagePlanKeyId" yaml:"usagePlanKeyId"`
}

