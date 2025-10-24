package awsiottwinmaker


// Properties for defining a `CfnComponentType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ DataTypeProperty
//   var dataValueProperty_ DataValueProperty
//   var relationshipValue interface{}
//
//   cfnComponentTypeProps := &CfnComponentTypeProps{
//   	ComponentTypeId: jsii.String("componentTypeId"),
//   	WorkspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	CompositeComponentTypes: map[string]interface{}{
//   		"compositeComponentTypesKey": &CompositeComponentTypeProperty{
//   			"componentTypeId": jsii.String("componentTypeId"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExtendsFrom: []*string{
//   		jsii.String("extendsFrom"),
//   	},
//   	Functions: map[string]interface{}{
//   		"functionsKey": &FunctionProperty{
//   			"implementedBy": &DataConnectorProperty{
//   				"isNative": jsii.Boolean(false),
//   				"lambda": &LambdaFunctionProperty{
//   					"arn": jsii.String("arn"),
//   				},
//   			},
//   			"requiredProperties": []*string{
//   				jsii.String("requiredProperties"),
//   			},
//   			"scope": jsii.String("scope"),
//   		},
//   	},
//   	IsSingleton: jsii.Boolean(false),
//   	PropertyDefinitions: map[string]interface{}{
//   		"propertyDefinitionsKey": &PropertyDefinitionProperty{
//   			"configurations": map[string]*string{
//   				"configurationsKey": jsii.String("configurations"),
//   			},
//   			"dataType": &DataTypeProperty{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"allowedValues": []interface{}{
//   					&DataValueProperty{
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
//   				"nestedType": dataTypeProperty_,
//   				"relationship": &RelationshipProperty{
//   					"relationshipType": jsii.String("relationshipType"),
//   					"targetComponentTypeId": jsii.String("targetComponentTypeId"),
//   				},
//   				"unitOfMeasure": jsii.String("unitOfMeasure"),
//   			},
//   			"defaultValue": &DataValueProperty{
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
//   			"isExternalId": jsii.Boolean(false),
//   			"isRequiredInEntity": jsii.Boolean(false),
//   			"isStoredExternally": jsii.Boolean(false),
//   			"isTimeSeries": jsii.Boolean(false),
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
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html
//
type CfnComponentTypeProps struct {
	// The ID of the component type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-componenttypeid
	//
	ComponentTypeId *string `field:"required" json:"componentTypeId" yaml:"componentTypeId"`
	// The ID of the workspace that contains the component type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-workspaceid
	//
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Maps strings to `compositeComponentTypes` of the `componentType` .
	//
	// `CompositeComponentType` is referenced by `componentTypeId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-compositecomponenttypes
	//
	CompositeComponentTypes interface{} `field:"optional" json:"compositeComponentTypes" yaml:"compositeComponentTypes"`
	// The description of the component type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parent component type that this component type extends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-extendsfrom
	//
	ExtendsFrom *[]*string `field:"optional" json:"extendsFrom" yaml:"extendsFrom"`
	// An object that maps strings to the functions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the FunctionResponse object see the [FunctionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_FunctionResponse.html) API reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-functions
	//
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
	// A boolean value that specifies whether an entity can have more than one component of this type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-issingleton
	//
	IsSingleton interface{} `field:"optional" json:"isSingleton" yaml:"isSingleton"`
	// An object that maps strings to the property definitions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information about the PropertyDefinitionResponse object, see the [PropertyDefinitionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_PropertyDefinitionResponse.html) API reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-propertydefinitions
	//
	PropertyDefinitions interface{} `field:"optional" json:"propertyDefinitions" yaml:"propertyDefinitions"`
	// An object that maps strings to the property groups in the component type.
	//
	// Each string in the mapping must be unique to this object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-propertygroups
	//
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
	// The ComponentType tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-componenttype.html#cfn-iottwinmaker-componenttype-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

