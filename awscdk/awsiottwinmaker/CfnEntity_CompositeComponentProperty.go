package awsiottwinmaker


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
//   compositeComponentProperty := &CompositeComponentProperty{
//   	ComponentName: jsii.String("componentName"),
//   	ComponentPath: jsii.String("componentPath"),
//   	ComponentTypeId: jsii.String("componentTypeId"),
//   	Description: jsii.String("description"),
//   	Properties: map[string]interface{}{
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
//   	PropertyGroups: map[string]interface{}{
//   		"propertyGroupsKey": &PropertyGroupProperty{
//   			"groupType": jsii.String("groupType"),
//   			"propertyNames": []*string{
//   				jsii.String("propertyNames"),
//   			},
//   		},
//   	},
//   	Status: &StatusProperty{
//   		Error: error,
//   		State: jsii.String("state"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html
//
type CfnEntity_CompositeComponentProperty struct {
	// The name of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html#cfn-iottwinmaker-entity-compositecomponent-componentname
	//
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The path of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html#cfn-iottwinmaker-entity-compositecomponent-componentpath
	//
	ComponentPath *string `field:"optional" json:"componentPath" yaml:"componentPath"`
	// The ID of the component type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html#cfn-iottwinmaker-entity-compositecomponent-componenttypeid
	//
	ComponentTypeId *string `field:"optional" json:"componentTypeId" yaml:"componentTypeId"`
	// The description of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html#cfn-iottwinmaker-entity-compositecomponent-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An object that maps strings to the properties to set in the component type.
	//
	// Each string in the mapping must be unique to this object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html#cfn-iottwinmaker-entity-compositecomponent-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// An object that maps strings to the property groups to set in the component type.
	//
	// Each string in the mapping must be unique to this object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html#cfn-iottwinmaker-entity-compositecomponent-propertygroups
	//
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-compositecomponent.html#cfn-iottwinmaker-entity-compositecomponent-status
	//
	Status interface{} `field:"optional" json:"status" yaml:"status"`
}

