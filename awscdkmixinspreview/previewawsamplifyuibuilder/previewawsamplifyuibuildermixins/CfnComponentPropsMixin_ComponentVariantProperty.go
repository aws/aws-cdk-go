package previewawsamplifyuibuildermixins


// The `ComponentVariant` property specifies the style configuration of a unique variation of a main component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var overrides interface{}
//
//   componentVariantProperty := &ComponentVariantProperty{
//   	Overrides: overrides,
//   	VariantValues: map[string]*string{
//   		"variantValuesKey": jsii.String("variantValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentvariant.html
//
type CfnComponentPropsMixin_ComponentVariantProperty struct {
	// The properties of the component variant that can be overriden when customizing an instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentvariant.html#cfn-amplifyuibuilder-component-componentvariant-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// The combination of variants that comprise this variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentvariant.html#cfn-amplifyuibuilder-component-componentvariant-variantvalues
	//
	VariantValues interface{} `field:"optional" json:"variantValues" yaml:"variantValues"`
}

