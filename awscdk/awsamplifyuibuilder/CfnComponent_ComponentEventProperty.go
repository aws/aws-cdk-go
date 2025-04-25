package awsamplifyuibuilder


// The `ComponentEvent` property specifies the configuration of an event.
//
// You can bind an event and a corresponding action to a `Component` or a `ComponentChild` . A button click is an example of an event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentPropertyProperty_ componentPropertyProperty
//
//   componentEventProperty := &ComponentEventProperty{
//   	Action: jsii.String("action"),
//   	BindingEvent: jsii.String("bindingEvent"),
//   	Parameters: &ActionParametersProperty{
//   		Anchor: &componentPropertyProperty{
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
//   		Fields: map[string]interface{}{
//   			"fieldsKey": &componentPropertyProperty{
//   				"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   					"property": jsii.String("property"),
//
//   					// the properties below are optional
//   					"field": jsii.String("field"),
//   				},
//   				"bindings": map[string]interface{}{
//   					"bindingsKey": &FormBindingElementProperty{
//   						"element": jsii.String("element"),
//   						"property": jsii.String("property"),
//   					},
//   				},
//   				"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   					"property": jsii.String("property"),
//
//   					// the properties below are optional
//   					"field": jsii.String("field"),
//   				},
//   				"componentName": jsii.String("componentName"),
//   				"concat": []interface{}{
//   					componentPropertyProperty_,
//   				},
//   				"condition": &ComponentConditionPropertyProperty{
//   					"else": componentPropertyProperty_,
//   					"field": jsii.String("field"),
//   					"operand": jsii.String("operand"),
//   					"operandType": jsii.String("operandType"),
//   					"operator": jsii.String("operator"),
//   					"property": jsii.String("property"),
//   					"then": componentPropertyProperty_,
//   				},
//   				"configured": jsii.Boolean(false),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"event": jsii.String("event"),
//   				"importedValue": jsii.String("importedValue"),
//   				"model": jsii.String("model"),
//   				"property": jsii.String("property"),
//   				"type": jsii.String("type"),
//   				"userAttribute": jsii.String("userAttribute"),
//   				"value": jsii.String("value"),
//   			},
//   		},
//   		Global: &componentPropertyProperty{
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
//   		Id: &componentPropertyProperty{
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
//   		Model: jsii.String("model"),
//   		State: &MutationActionSetStateParameterProperty{
//   			ComponentName: jsii.String("componentName"),
//   			Property: jsii.String("property"),
//   			Set: &componentPropertyProperty{
//   				BindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   					Property: jsii.String("property"),
//
//   					// the properties below are optional
//   					Field: jsii.String("field"),
//   				},
//   				Bindings: map[string]interface{}{
//   					"bindingsKey": &FormBindingElementProperty{
//   						"element": jsii.String("element"),
//   						"property": jsii.String("property"),
//   					},
//   				},
//   				CollectionBindingProperties: &ComponentPropertyBindingPropertiesProperty{
//   					Property: jsii.String("property"),
//
//   					// the properties below are optional
//   					Field: jsii.String("field"),
//   				},
//   				ComponentName: jsii.String("componentName"),
//   				Concat: []interface{}{
//   					componentPropertyProperty_,
//   				},
//   				Condition: &ComponentConditionPropertyProperty{
//   					Else: componentPropertyProperty_,
//   					Field: jsii.String("field"),
//   					Operand: jsii.String("operand"),
//   					OperandType: jsii.String("operandType"),
//   					Operator: jsii.String("operator"),
//   					Property: jsii.String("property"),
//   					Then: componentPropertyProperty_,
//   				},
//   				Configured: jsii.Boolean(false),
//   				DefaultValue: jsii.String("defaultValue"),
//   				Event: jsii.String("event"),
//   				ImportedValue: jsii.String("importedValue"),
//   				Model: jsii.String("model"),
//   				Property: jsii.String("property"),
//   				Type: jsii.String("type"),
//   				UserAttribute: jsii.String("userAttribute"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Target: &componentPropertyProperty{
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
//   		Type: &componentPropertyProperty{
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
//   		Url: &componentPropertyProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentevent.html
//
type CfnComponent_ComponentEventProperty struct {
	// The action to perform when a specific event is raised.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentevent.html#cfn-amplifyuibuilder-component-componentevent-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Binds an event to an action on a component.
	//
	// When you specify a `bindingEvent` , the event is called when the action is performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentevent.html#cfn-amplifyuibuilder-component-componentevent-bindingevent
	//
	BindingEvent *string `field:"optional" json:"bindingEvent" yaml:"bindingEvent"`
	// Describes information about the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentevent.html#cfn-amplifyuibuilder-component-componentevent-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

