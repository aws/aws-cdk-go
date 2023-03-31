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
//   componentBindingPropertiesValueProperty := &componentBindingPropertiesValueProperty{
//   	bindingProperties: &componentBindingPropertiesValuePropertiesProperty{
//   		bucket: jsii.String("bucket"),
//   		defaultValue: jsii.String("defaultValue"),
//   		field: jsii.String("field"),
//   		key: jsii.String("key"),
//   		model: jsii.String("model"),
//   		predicates: []interface{}{
//   			&predicateProperty{
//   				and: []interface{}{
//   					predicateProperty_,
//   				},
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operator: jsii.String("operator"),
//   				or: []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   		},
//   		userAttribute: jsii.String("userAttribute"),
//   	},
//   	defaultValue: jsii.String("defaultValue"),
//   	type: jsii.String("type"),
//   }
//
type CfnComponent_ComponentBindingPropertiesValueProperty struct {
	// Describes the properties to customize with data at runtime.
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// The default value of the property.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The property type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

