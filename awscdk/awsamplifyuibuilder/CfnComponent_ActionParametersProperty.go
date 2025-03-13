package awsamplifyuibuilder


// Represents the event action configuration for an element of a `Component` or `ComponentChild` .
//
// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components. `ActionParameters` defines the action that is performed when an event occurs on the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentPropertyProperty_ componentPropertyProperty
//
//   actionParametersProperty := &ActionParametersProperty{
//   	Anchor: &componentPropertyProperty{
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
//   	Fields: map[string]interface{}{
//   		"fieldsKey": &componentPropertyProperty{
//   			"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"bindings": map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"componentName": jsii.String("componentName"),
//   			"concat": []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			"condition": &ComponentConditionPropertyProperty{
//   				"else": componentPropertyProperty_,
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operandType": jsii.String("operandType"),
//   				"operator": jsii.String("operator"),
//   				"property": jsii.String("property"),
//   				"then": componentPropertyProperty_,
//   			},
//   			"configured": jsii.Boolean(false),
//   			"defaultValue": jsii.String("defaultValue"),
//   			"event": jsii.String("event"),
//   			"importedValue": jsii.String("importedValue"),
//   			"model": jsii.String("model"),
//   			"property": jsii.String("property"),
//   			"type": jsii.String("type"),
//   			"userAttribute": jsii.String("userAttribute"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	Global: &componentPropertyProperty{
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
//   	Id: &componentPropertyProperty{
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
//   	Model: jsii.String("model"),
//   	State: &MutationActionSetStateParameterProperty{
//   		ComponentName: jsii.String("componentName"),
//   		Property: jsii.String("property"),
//   		Set: &componentPropertyProperty{
//   			BindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   				Property: jsii.String("property"),
//
//   				// the properties below are optional
//   				Field: jsii.String("field"),
//   			},
//   			Bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			CollectionBindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   				Property: jsii.String("property"),
//
//   				// the properties below are optional
//   				Field: jsii.String("field"),
//   			},
//   			ComponentName: jsii.String("componentName"),
//   			Concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			Condition: &ComponentConditionPropertyProperty{
//   				Else: componentPropertyProperty_,
//   				Field: jsii.String("field"),
//   				Operand: jsii.String("operand"),
//   				OperandType: jsii.String("operandType"),
//   				Operator: jsii.String("operator"),
//   				Property: jsii.String("property"),
//   				Then: componentPropertyProperty_,
//   			},
//   			Configured: jsii.Boolean(false),
//   			DefaultValue: jsii.String("defaultValue"),
//   			Event: jsii.String("event"),
//   			ImportedValue: jsii.String("importedValue"),
//   			Model: jsii.String("model"),
//   			Property: jsii.String("property"),
//   			Type: jsii.String("type"),
//   			UserAttribute: jsii.String("userAttribute"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Target: &componentPropertyProperty{
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
//   	Type: &componentPropertyProperty{
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
//   	Url: &componentPropertyProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html
//
type CfnComponent_ActionParametersProperty struct {
	// The HTML anchor link to the location to open.
	//
	// Specify this value for a navigation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-anchor
	//
	Anchor interface{} `field:"optional" json:"anchor" yaml:"anchor"`
	// A dictionary of key-value pairs mapping Amplify Studio properties to fields in a data model.
	//
	// Use when the action performs an operation on an Amplify DataStore model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Specifies whether the user should be signed out globally.
	//
	// Specify this value for an auth sign out action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-global
	//
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// The unique ID of the component that the `ActionParameters` apply to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-id
	//
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// The name of the data model.
	//
	// Use when the action performs an operation on an Amplify DataStore model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// A key-value pair that specifies the state property name and its initial value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-state
	//
	State interface{} `field:"optional" json:"state" yaml:"state"`
	// The element within the same component to modify when the action occurs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
	// The type of navigation action.
	//
	// Valid values are `url` and `anchor` . This value is required for a navigation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-type
	//
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// The URL to the location to open.
	//
	// Specify this value for a navigation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-url
	//
	Url interface{} `field:"optional" json:"url" yaml:"url"`
}

