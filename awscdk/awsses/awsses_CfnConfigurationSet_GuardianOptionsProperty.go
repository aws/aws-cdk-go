package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardianOptionsProperty := &guardianOptionsProperty{
//   	optimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   }
//
type CfnConfigurationSet_GuardianOptionsProperty struct {
	// `CfnConfigurationSet.GuardianOptionsProperty.OptimizedSharedDelivery`.
	OptimizedSharedDelivery *string `field:"required" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

