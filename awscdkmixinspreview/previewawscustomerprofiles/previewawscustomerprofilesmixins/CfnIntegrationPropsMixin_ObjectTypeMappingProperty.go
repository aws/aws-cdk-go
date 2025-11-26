package previewawscustomerprofilesmixins


// A map in which each key is an event type from an external application such as Segment or Shopify, and each value is an `ObjectTypeName` (template) used to ingest the event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectTypeMappingProperty := &ObjectTypeMappingProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-objecttypemapping.html
//
type CfnIntegrationPropsMixin_ObjectTypeMappingProperty struct {
	// The key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-objecttypemapping.html#cfn-customerprofiles-integration-objecttypemapping-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-objecttypemapping.html#cfn-customerprofiles-integration-objecttypemapping-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

