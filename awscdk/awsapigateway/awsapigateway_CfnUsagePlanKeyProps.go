package awsapigateway


// Properties for defining a `CfnUsagePlanKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlanKeyProps := &cfnUsagePlanKeyProps{
//   	keyId: jsii.String("keyId"),
//   	keyType: jsii.String("keyType"),
//   	usagePlanId: jsii.String("usagePlanId"),
//   }
//
type CfnUsagePlanKeyProps struct {
	// The ID of the usage plan key.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The type of usage plan key.
	//
	// Currently, the only valid key type is `API_KEY` .
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The ID of the usage plan.
	UsagePlanId *string `field:"required" json:"usagePlanId" yaml:"usagePlanId"`
}

