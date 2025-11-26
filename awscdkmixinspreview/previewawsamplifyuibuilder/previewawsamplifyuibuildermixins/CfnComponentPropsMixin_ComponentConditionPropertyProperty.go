package previewawsamplifyuibuildermixins


// The `ComponentConditionProperty` property specifies a conditional expression for setting a component property.
//
// Use `ComponentConditionProperty` to set a property to different values conditionally, based on the value of another property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var componentConditionPropertyProperty_ ComponentConditionPropertyProperty
//   var componentPropertyProperty_ ComponentPropertyProperty
//
//   componentConditionPropertyProperty := &ComponentConditionPropertyProperty{
//   	Else: &ComponentPropertyProperty{
//   		BindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Field: jsii.String("field"),
//   			Property: jsii.String("property"),
//   		},
//   		Bindings: map[string]interface{}{
//   			"bindingsKey": &FormBindingElementProperty{
//   				"element": jsii.String("element"),
//   				"property": jsii.String("property"),
//   			},
//   		},
//   		CollectionBindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Field: jsii.String("field"),
//   			Property: jsii.String("property"),
//   		},
//   		ComponentName: jsii.String("componentName"),
//   		Concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		Condition: componentConditionPropertyProperty_,
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
//   	Then: &ComponentPropertyProperty{
//   		BindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Field: jsii.String("field"),
//   			Property: jsii.String("property"),
//   		},
//   		Bindings: map[string]interface{}{
//   			"bindingsKey": &FormBindingElementProperty{
//   				"element": jsii.String("element"),
//   				"property": jsii.String("property"),
//   			},
//   		},
//   		CollectionBindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   			Field: jsii.String("field"),
//   			Property: jsii.String("property"),
//   		},
//   		ComponentName: jsii.String("componentName"),
//   		Concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		Condition: componentConditionPropertyProperty_,
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html
//
type CfnComponentPropsMixin_ComponentConditionPropertyProperty struct {
	// The value to assign to the property if the condition is not met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html#cfn-amplifyuibuilder-component-componentconditionproperty-else
	//
	Else interface{} `field:"optional" json:"else" yaml:"else"`
	// The name of a field.
	//
	// Specify this when the property is a data model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html#cfn-amplifyuibuilder-component-componentconditionproperty-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The value of the property to evaluate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html#cfn-amplifyuibuilder-component-componentconditionproperty-operand
	//
	Operand *string `field:"optional" json:"operand" yaml:"operand"`
	// The type of the property to evaluate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html#cfn-amplifyuibuilder-component-componentconditionproperty-operandtype
	//
	OperandType *string `field:"optional" json:"operandType" yaml:"operandType"`
	// The operator to use to perform the evaluation, such as `eq` to represent equals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html#cfn-amplifyuibuilder-component-componentconditionproperty-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The name of the conditional property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html#cfn-amplifyuibuilder-component-componentconditionproperty-property
	//
	Property *string `field:"optional" json:"property" yaml:"property"`
	// The value to assign to the property if the condition is met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentconditionproperty.html#cfn-amplifyuibuilder-component-componentconditionproperty-then
	//
	Then interface{} `field:"optional" json:"then" yaml:"then"`
}

