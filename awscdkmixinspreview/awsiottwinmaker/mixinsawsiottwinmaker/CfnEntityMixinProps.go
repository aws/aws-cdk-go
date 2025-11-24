package mixinsawsiottwinmaker


// Properties for CfnEntityPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dataValueProperty_ DataValueProperty
//   var definition interface{}
//   var error interface{}
//   var relationshipValue interface{}
//
//   cfnEntityMixinProps := &CfnEntityMixinProps{
//   	Components: map[string]interface{}{
//   		"componentsKey": &ComponentProperty{
//   			"componentName": jsii.String("componentName"),
//   			"componentTypeId": jsii.String("componentTypeId"),
//   			"definedIn": jsii.String("definedIn"),
//   			"description": jsii.String("description"),
//   			"properties": map[string]interface{}{
//   				"propertiesKey": &PropertyProperty{
//   					"definition": definition,
//   					"value": &DataValueProperty{
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
//   	CompositeComponents: map[string]interface{}{
//   		"compositeComponentsKey": &CompositeComponentProperty{
//   			"componentName": jsii.String("componentName"),
//   			"componentPath": jsii.String("componentPath"),
//   			"componentTypeId": jsii.String("componentTypeId"),
//   			"description": jsii.String("description"),
//   			"properties": map[string]interface{}{
//   				"propertiesKey": &PropertyProperty{
//   					"definition": definition,
//   					"value": &DataValueProperty{
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
//   	EntityName: jsii.String("entityName"),
//   	ParentEntityId: jsii.String("parentEntityId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html
//
type CfnEntityMixinProps struct {
	// An object that maps strings to the components in the entity.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the component object see the [component](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_ComponentResponse.html) API reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-components
	//
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// Maps string to `compositeComponent` updates in the request.
	//
	// Each key of the map represents the `componentPath` of the `compositeComponent` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-compositecomponents
	//
	CompositeComponents interface{} `field:"optional" json:"compositeComponents" yaml:"compositeComponents"`
	// The description of the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-entityid
	//
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The entity name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-entityname
	//
	EntityName *string `field:"optional" json:"entityName" yaml:"entityName"`
	// The ID of the parent entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-parententityid
	//
	ParentEntityId *string `field:"optional" json:"parentEntityId" yaml:"parentEntityId"`
	// Metadata that you can use to manage the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the workspace that contains the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-entity.html#cfn-iottwinmaker-entity-workspaceid
	//
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

