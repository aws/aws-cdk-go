package awsamplifyuibuilder


// The `ActionParameters` property specifies the event action configuration for an element of a `Component` or `ComponentChild` .
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
//   actionParametersProperty := &actionParametersProperty{
//   	anchor: &componentPropertyProperty{
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
//   	fields: map[string]interface{}{
//   		"fieldsKey": &componentPropertyProperty{
//   			"bindingProperties": &componentPropertyBindingPropertiesProperty{
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
//   			"collectionBindingProperties": &componentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"componentName": jsii.String("componentName"),
//   			"concat": []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			"condition": &componentConditionPropertyProperty{
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
//   	global: &componentPropertyProperty{
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
//   	id: &componentPropertyProperty{
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
//   	model: jsii.String("model"),
//   	state: &mutationActionSetStateParameterProperty{
//   		componentName: jsii.String("componentName"),
//   		property: jsii.String("property"),
//   		set: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	target: &componentPropertyProperty{
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
//   	type: &componentPropertyProperty{
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
//   	url: &componentPropertyProperty{
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
type CfnComponent_ActionParametersProperty struct {
	// The HTML anchor link to the location to open.
	//
	// Specify this value for a navigation action.
	Anchor interface{} `field:"optional" json:"anchor" yaml:"anchor"`
	// A dictionary of key-value pairs mapping Amplify Studio properties to fields in a data model.
	//
	// Use when the action performs an operation on an Amplify DataStore model.
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Specifies whether the user should be signed out globally.
	//
	// Specify this value for an auth sign out action.
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// The unique ID of the component that the `ActionParameters` apply to.
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// The name of the data model.
	//
	// Use when the action performs an operation on an Amplify DataStore model.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// A key-value pair that specifies the state property name and its initial value.
	State interface{} `field:"optional" json:"state" yaml:"state"`
	// The element within the same component to modify when the action occurs.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
	// The type of navigation action.
	//
	// Valid values are `url` and `anchor` . This value is required for a navigation action.
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// The URL to the location to open.
	//
	// Specify this value for a navigation action.
	Url interface{} `field:"optional" json:"url" yaml:"url"`
}

