package awsamplifyuibuilder


// The `ComponentBindingPropertiesValue` property specifies the data binding configuration for a component at runtime.
//
// You can use `ComponentBindingPropertiesValue` to add exposed properties to a component to allow different values to be entered when a component is reused in different places in an app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   componentBindingPropertiesValueProperty := &ComponentBindingPropertiesValueProperty{
//   	BindingProperties: &ComponentBindingPropertiesValuePropertiesProperty{
//   		Bucket: jsii.String("bucket"),
//   		DefaultValue: jsii.String("defaultValue"),
//   		Field: jsii.String("field"),
//   		Key: jsii.String("key"),
//   		Model: jsii.String("model"),
//   		Predicates: []interface{}{
//   			&predicateProperty{
//   				And: []interface{}{
//   					predicateProperty_,
//   				},
//   				Field: jsii.String("field"),
//   				Operand: jsii.String("operand"),
//   				Operator: jsii.String("operator"),
//   				Or: []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   		},
//   		UserAttribute: jsii.String("userAttribute"),
//   	},
//   	DefaultValue: jsii.String("defaultValue"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalue.html
//
type CfnComponent_ComponentBindingPropertiesValueProperty struct {
	// Describes the properties to customize with data at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalue.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalue-bindingproperties
	//
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// The default value of the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalue.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalue-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The property type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalue.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalue-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

