// An experiment to bundle the entire CDK into a single module
package awscdk


// Properties for a Dynamic Reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDynamicReferenceProps := &cfnDynamicReferenceProps{
//   	referenceKey: jsii.String("referenceKey"),
//   	service: monocdk.cfnDynamicReferenceService_SSM,
//   }
//
// Experimental.
type CfnDynamicReferenceProps struct {
	// The reference key of the dynamic reference.
	// Experimental.
	ReferenceKey *string `field:"required" json:"referenceKey" yaml:"referenceKey"`
	// The service to retrieve the dynamic reference from.
	// Experimental.
	Service CfnDynamicReferenceService `field:"required" json:"service" yaml:"service"`
}

