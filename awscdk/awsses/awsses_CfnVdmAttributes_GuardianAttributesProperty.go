package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardianAttributesProperty := &guardianAttributesProperty{
//   	optimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   }
//
type CfnVdmAttributes_GuardianAttributesProperty struct {
	// `CfnVdmAttributes.GuardianAttributesProperty.OptimizedSharedDelivery`.
	OptimizedSharedDelivery *string `field:"optional" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

