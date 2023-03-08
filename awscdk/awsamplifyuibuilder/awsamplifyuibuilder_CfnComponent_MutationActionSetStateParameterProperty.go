package awsamplifyuibuilder


// The `MutationActionSetStateParameter` property specifies the state configuration when an action modifies a property of another element within the same component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentPropertyProperty_ componentPropertyProperty
//
//   mutationActionSetStateParameterProperty := &mutationActionSetStateParameterProperty{
//   	componentName: jsii.String("componentName"),
//   	property: jsii.String("property"),
//   	set: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: map[string]interface{}{
//   			"bindingsKey": &FormBindingElementProperty{
//   				"element": jsii.String("element"),
//   				"property": jsii.String("property"),
//   			},
//   		},
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnComponent_MutationActionSetStateParameterProperty struct {
	// The name of the component that is being modified.
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The name of the component property to apply the state configuration to.
	Property *string `field:"required" json:"property" yaml:"property"`
	// The state configuration to assign to the property.
	Set interface{} `field:"required" json:"set" yaml:"set"`
}

