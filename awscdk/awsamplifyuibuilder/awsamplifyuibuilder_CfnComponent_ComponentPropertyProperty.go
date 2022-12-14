package awsamplifyuibuilder


// The `ComponentProperty` property specifies the configuration for all of a component's properties.
//
// Use `ComponentProperty` to specify the values to render or bind by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentConditionPropertyProperty_ componentConditionPropertyProperty
//   var componentPropertyProperty_ componentPropertyProperty
//
//   componentPropertyProperty := &componentPropertyProperty{
//   	bindingProperties: &componentPropertyBindingPropertiesProperty{
//   		property: jsii.String("property"),
//
//   		// the properties below are optional
//   		field: jsii.String("field"),
//   	},
//   	bindings: map[string]interface{}{
//   		"bindingsKey": &FormBindingElementProperty{
//   			"element": jsii.String("element"),
//   			"property": jsii.String("property"),
//   		},
//   	},
//   	collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   		property: jsii.String("property"),
//
//   		// the properties below are optional
//   		field: jsii.String("field"),
//   	},
//   	componentName: jsii.String("componentName"),
//   	concat: []interface{}{
//   		&componentPropertyProperty{
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
//   	condition: &componentConditionPropertyProperty{
//   		else: &componentPropertyProperty{
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
//   			condition: componentConditionPropertyProperty_,
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
//   		field: jsii.String("field"),
//   		operand: jsii.String("operand"),
//   		operandType: jsii.String("operandType"),
//   		operator: jsii.String("operator"),
//   		property: jsii.String("property"),
//   		then: &componentPropertyProperty{
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
//   			condition: componentConditionPropertyProperty_,
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
//   	configured: jsii.Boolean(false),
//   	defaultValue: jsii.String("defaultValue"),
//   	event: jsii.String("event"),
//   	importedValue: jsii.String("importedValue"),
//   	model: jsii.String("model"),
//   	property: jsii.String("property"),
//   	type: jsii.String("type"),
//   	userAttribute: jsii.String("userAttribute"),
//   	value: jsii.String("value"),
//   }
//
type CfnComponent_ComponentPropertyProperty struct {
	// The information to bind the component property to data at runtime.
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// The information to bind the component property to form data.
	Bindings interface{} `field:"optional" json:"bindings" yaml:"bindings"`
	// The information to bind the component property to data at runtime.
	//
	// Use this for collection components.
	CollectionBindingProperties interface{} `field:"optional" json:"collectionBindingProperties" yaml:"collectionBindingProperties"`
	// The name of the component that is affected by an event.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// A list of component properties to concatenate to create the value to assign to this component property.
	Concat interface{} `field:"optional" json:"concat" yaml:"concat"`
	// The conditional expression to use to assign a value to the component property.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Specifies whether the user configured the property in Amplify Studio after importing it.
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
	// The default value to assign to the component property.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// An event that occurs in your app.
	//
	// Use this for workflow data binding.
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The default value assigned to the property when the component is imported into an app.
	ImportedValue *string `field:"optional" json:"importedValue" yaml:"importedValue"`
	// The data model to use to assign a value to the component property.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The name of the component's property that is affected by an event.
	Property *string `field:"optional" json:"property" yaml:"property"`
	// The component type.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// An authenticated user attribute to use to assign a value to the component property.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
	// The value to assign to the component property.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

