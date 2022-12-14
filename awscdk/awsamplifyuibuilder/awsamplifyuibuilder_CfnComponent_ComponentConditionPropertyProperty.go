package awsamplifyuibuilder


// The `ComponentConditionProperty` property specifies a conditional expression for setting a component property.
//
// Use `ComponentConditionProperty` to set a property to different values conditionally, based on the value of another property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentPropertyProperty_ componentPropertyProperty
//
//   componentConditionPropertyProperty := &componentConditionPropertyProperty{
//   	else: &componentPropertyProperty{
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
//   	field: jsii.String("field"),
//   	operand: jsii.String("operand"),
//   	operandType: jsii.String("operandType"),
//   	operator: jsii.String("operator"),
//   	property: jsii.String("property"),
//   	then: &componentPropertyProperty{
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
type CfnComponent_ComponentConditionPropertyProperty struct {
	// The value to assign to the property if the condition is not met.
	Else interface{} `field:"optional" json:"else" yaml:"else"`
	// The name of a field.
	//
	// Specify this when the property is a data model.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The value of the property to evaluate.
	Operand *string `field:"optional" json:"operand" yaml:"operand"`
	// The type of the property to evaluate.
	OperandType *string `field:"optional" json:"operandType" yaml:"operandType"`
	// The operator to use to perform the evaluation, such as `eq` to represent equals.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The name of the conditional property.
	Property *string `field:"optional" json:"property" yaml:"property"`
	// The value to assign to the property if the condition is met.
	Then interface{} `field:"optional" json:"then" yaml:"then"`
}

