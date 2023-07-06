package awsapigateway


// Properties for defining a `CfnUsagePlanKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlanKeyProps := &CfnUsagePlanKeyProps{
//   	KeyId: jsii.String("keyId"),
//   	KeyType: jsii.String("keyType"),
//   	UsagePlanId: jsii.String("usagePlanId"),
//   }
//
type CfnUsagePlanKeyProps struct {
	// The Id of the UsagePlanKey resource.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The type of a UsagePlanKey resource for a plan customer.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The Id of the UsagePlan resource representing the usage plan containing the UsagePlanKey resource representing a plan customer.
	UsagePlanId *string `field:"required" json:"usagePlanId" yaml:"usagePlanId"`
}

