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
//   componentConditionPropertyProperty := &ComponentConditionPropertyProperty{
//   	Else: &componentPropertyProperty{
//   		BindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Property: jsii.String("property"),
//
//   			// the properties below are optional
//   			Field: jsii.String("field"),
//   		},
//   		Bindings: map[string]interface{}{
//   			"bindingsKey": &FormBindingElementProperty{
//   				"element": jsii.String("element"),
//   				"property": jsii.String("property"),
//   			},
//   		},
//   		CollectionBindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Property: jsii.String("property"),
//
//   			// the properties below are optional
//   			Field: jsii.String("field"),
//   		},
//   		ComponentName: jsii.String("componentName"),
//   		Concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		Condition: &ComponentConditionPropertyProperty{
//   			Else: componentPropertyProperty_,
//   			Field: jsii.String("field"),
//   			Operand: jsii.String("operand"),
//   			OperandType: jsii.String("operandType"),
//   			Operator: jsii.String("operator"),
//   			Property: jsii.String("property"),
//   			Then: componentPropertyProperty_,
//   		},
//   		Configured: jsii.Boolean(false),
//   		DefaultValue: jsii.String("defaultValue"),
//   		Event: jsii.String("event"),
//   		ImportedValue: jsii.String("importedValue"),
//   		Model: jsii.String("model"),
//   		Property: jsii.String("property"),
//   		Type: jsii.String("type"),
//   		UserAttribute: jsii.String("userAttribute"),
//   		Value: jsii.String("value"),
//   	},
//   	Field: jsii.String("field"),
//   	Operand: jsii.String("operand"),
//   	OperandType: jsii.String("operandType"),
//   	Operator: jsii.String("operator"),
//   	Property: jsii.String("property"),
//   	Then: &componentPropertyProperty{
//   		BindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Property: jsii.String("property"),
//
//   			// the properties below are optional
//   			Field: jsii.String("field"),
//   		},
//   		Bindings: map[string]interface{}{
//   			"bindingsKey": &FormBindingElementProperty{
//   				"element": jsii.String("element"),
//   				"property": jsii.String("property"),
//   			},
//   		},
//   		CollectionBindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Property: jsii.String("property"),
//
//   			// the properties below are optional
//   			Field: jsii.String("field"),
//   		},
//   		ComponentName: jsii.String("componentName"),
//   		Concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		Condition: &ComponentConditionPropertyProperty{
//   			Else: componentPropertyProperty_,
//   			Field: jsii.String("field"),
//   			Operand: jsii.String("operand"),
//   			OperandType: jsii.String("operandType"),
//   			Operator: jsii.String("operator"),
//   			Property: jsii.String("property"),
//   			Then: componentPropertyProperty_,
//   		},
//   		Configured: jsii.Boolean(false),
//   		DefaultValue: jsii.String("defaultValue"),
//   		Event: jsii.String("event"),
//   		ImportedValue: jsii.String("importedValue"),
//   		Model: jsii.String("model"),
//   		Property: jsii.String("property"),
//   		Type: jsii.String("type"),
//   		UserAttribute: jsii.String("userAttribute"),
//   		Value: jsii.String("value"),
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

