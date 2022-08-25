package awscustomerprofiles


// A map in which each key is an event type from an external application such as Segment or Shopify, and each value is an `ObjectTypeName` (template) used to ingest the event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectTypeMappingProperty := &objectTypeMappingProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnIntegration_ObjectTypeMappingProperty struct {
	// The key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

