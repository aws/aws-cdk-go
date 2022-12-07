package awsiottwinmaker


// The entity componenet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var definition interface{}
//   var error interface{}
//   var relationshipValue interface{}
//
//   componentProperty := &componentProperty{
//   	componentName: jsii.String("componentName"),
//   	componentTypeId: jsii.String("componentTypeId"),
//   	definedIn: jsii.String("definedIn"),
//   	description: jsii.String("description"),
//   	properties: map[string]interface{}{
//   		"propertiesKey": &PropertyProperty{
//   			"definition": definition,
//   			"value": &dataValueProperty{
//   				"booleanValue": jsii.Boolean(false),
//   				"doubleValue": jsii.Number(123),
//   				"expression": jsii.String("expression"),
//   				"integerValue": jsii.Number(123),
//   				"listValue": []interface{}{
//   					dataValueProperty_,
//   				},
//   				"longValue": jsii.Number(123),
//   				"mapValue": map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				"relationshipValue": relationshipValue,
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	propertyGroups: map[string]interface{}{
//   		"propertyGroupsKey": &PropertyGroupProperty{
//   			"groupType": jsii.String("groupType"),
//   			"propertyNames": []*string{
//   				jsii.String("propertyNames"),
//   			},
//   		},
//   	},
//   	status: &statusProperty{
//   		error: error,
//   		state: jsii.String("state"),
//   	},
//   }
//
type CfnEntity_ComponentProperty struct {
	// The name of the component.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The ID of the ComponentType.
	ComponentTypeId *string `field:"optional" json:"componentTypeId" yaml:"componentTypeId"`
	// The name of the property definition set in the request.
	DefinedIn *string `field:"optional" json:"definedIn" yaml:"definedIn"`
	// The description of the component.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An object that maps strings to the properties to set in the component type.
	//
	// Each string in the mapping must be unique to this object.
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// `CfnEntity.ComponentProperty.PropertyGroups`.
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
	// The status of the component.
	Status interface{} `field:"optional" json:"status" yaml:"status"`
}

