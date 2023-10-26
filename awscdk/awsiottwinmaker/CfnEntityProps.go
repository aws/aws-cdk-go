package awsiottwinmaker


// Properties for defining a `CfnEntity`.
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
//   cfnEntityProps := &CfnEntityProps{
//   	EntityName: jsii.String("entityName"),
//   	WorkspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	Components: map[string]interface{}{
//   		"componentsKey": &ComponentProperty{
//   			"componentName": jsii.String("componentName"),
//   			"componentTypeId": jsii.String("componentTypeId"),
//   			"definedIn": jsii.String("definedIn"),
//   			"description": jsii.String("description"),
//   			"properties": map[string]interface{}{
//   				"propertiesKey": &PropertyProperty{
//   					"definition": definition,
//   					"value": &dataValueProperty{
//   						"booleanValue": jsii.Boolean(false),
//   						"doubleValue": jsii.Number(123),
//   						"expression": jsii.String("expression"),
//   						"integerValue": jsii.Number(123),
//   						"listValue": []interface{}{
//   							dataValueProperty_,
//   						},
//   						"longValue": jsii.Number(123),
//   						"mapValue": map[string]interface{}{
//   							"mapValueKey": dataValueProperty_,
//   						},
//   						"relationshipValue": relationshipValue,
//   						"stringValue": jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			"propertyGroups": map[string]interface{}{
//   				"propertyGroupsKey": &PropertyGroupProperty{
//   					"groupType": jsii.String("groupType"),
//   					"propertyNames": []*string{
//   						jsii.String("propertyNames"),
//   					},
//   				},
//   			},
//   			"status": &StatusProperty{
//   				"error": error,
//   				"state": jsii.String("state"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EntityId: jsii.String("entityId"),
//   	ParentEntityId: jsii.String("parentEntityId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html
//
type CfnEntityProps struct {
	// The entity name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-entityname
	//
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// The ID of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-workspaceid
	//
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// An object that maps strings to the components in the entity.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the component object see the [component](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_ComponentResponse.html) API reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-components
	//
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// The description of the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-entityid
	//
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The ID of the parent entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-parententityid
	//
	ParentEntityId *string `field:"optional" json:"parentEntityId" yaml:"parentEntityId"`
	// Metadata that you can use to manage the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

