package awsamplifyuibuilder


// Represents the state configuration when an action modifies a property of another element within the same component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentPropertyProperty_ componentPropertyProperty
//
//   mutationActionSetStateParameterProperty := &MutationActionSetStateParameterProperty{
//   	ComponentName: jsii.String("componentName"),
//   	Property: jsii.String("property"),
//   	Set: &componentPropertyProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-mutationactionsetstateparameter.html
//
type CfnComponent_MutationActionSetStateParameterProperty struct {
	// The name of the component that is being modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-mutationactionsetstateparameter.html#cfn-amplifyuibuilder-component-mutationactionsetstateparameter-componentname
	//
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The name of the component property to apply the state configuration to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-mutationactionsetstateparameter.html#cfn-amplifyuibuilder-component-mutationactionsetstateparameter-property
	//
	Property *string `field:"required" json:"property" yaml:"property"`
	// The state configuration to assign to the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-mutationactionsetstateparameter.html#cfn-amplifyuibuilder-component-mutationactionsetstateparameter-set
	//
	Set interface{} `field:"required" json:"set" yaml:"set"`
}

