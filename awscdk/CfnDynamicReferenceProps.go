package awscdk


// Properties for a Dynamic Reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDynamicReferenceProps := &CfnDynamicReferenceProps{
//   	ReferenceKey: jsii.String("referenceKey"),
//   	Service: cdk.CfnDynamicReferenceService_SSM,
//   }
//
type CfnDynamicReferenceProps struct {
	// The reference key of the dynamic reference.
	ReferenceKey *string `field:"required" json:"referenceKey" yaml:"referenceKey"`
	// The service to retrieve the dynamic reference from.
	Service CfnDynamicReferenceService `field:"required" json:"service" yaml:"service"`
}

