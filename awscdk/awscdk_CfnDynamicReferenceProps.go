// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Properties for a Dynamic Reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDynamicReferenceProps := &cfnDynamicReferenceProps{
//   	referenceKey: jsii.String("referenceKey"),
//   	service: cdk.cfnDynamicReferenceService_SSM,
//   }
//
type CfnDynamicReferenceProps struct {
	// The reference key of the dynamic reference.
	ReferenceKey *string `field:"required" json:"referenceKey" yaml:"referenceKey"`
	// The service to retrieve the dynamic reference from.
	Service CfnDynamicReferenceService `field:"required" json:"service" yaml:"service"`
}

