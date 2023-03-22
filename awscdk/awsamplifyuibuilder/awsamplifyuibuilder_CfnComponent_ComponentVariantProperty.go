package awsamplifyuibuilder


// The `ComponentVariant` property specifies the style configuration of a unique variation of a main component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var overrides interface{}
//
//   componentVariantProperty := &componentVariantProperty{
//   	overrides: overrides,
//   	variantValues: map[string]*string{
//   		"variantValuesKey": jsii.String("variantValues"),
//   	},
//   }
//
type CfnComponent_ComponentVariantProperty struct {
	// The properties of the component variant that can be overriden when customizing an instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// The combination of variants that comprise this variant.
	VariantValues interface{} `field:"optional" json:"variantValues" yaml:"variantValues"`
}

