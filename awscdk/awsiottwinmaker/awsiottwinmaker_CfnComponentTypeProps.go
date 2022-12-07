package awsiottwinmaker


// Properties for defining a `CfnComponentType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   cfnComponentTypeProps := &cfnComponentTypeProps{
//   	componentTypeId: jsii.String("componentTypeId"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	extendsFrom: []*string{
//   		jsii.String("extendsFrom"),
//   	},
//   	functions: map[string]interface{}{
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
//   	isSingleton: jsii.Boolean(false),
//   	propertyDefinitions: map[string]interface{}{
//   		"propertyDefinitionsKey": &PropertyDefinitionProperty{
//   			"configurations": map[string]*string{
//   				"configurationsKey": jsii.String("configurations"),
//   			},
//   			"dataType": &dataTypeProperty{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"allowedValues": []interface{}{
//   					&dataValueProperty{
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
//   			"defaultValue": &dataValueProperty{
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
//   	propertyGroups: map[string]interface{}{
//   		"propertyGroupsKey": &PropertyGroupProperty{
//   			"groupType": jsii.String("groupType"),
//   			"propertyNames": []*string{
//   				jsii.String("propertyNames"),
//   			},
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnComponentTypeProps struct {
	// The ID of the component type.
	ComponentTypeId *string `field:"required" json:"componentTypeId" yaml:"componentTypeId"`
	// The ID of the workspace.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// The description of the component type.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parent component type that this component type extends.
	ExtendsFrom *[]*string `field:"optional" json:"extendsFrom" yaml:"extendsFrom"`
	// An object that maps strings to the functions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the FunctionResponse object see the [FunctionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_FunctionResponse.html) API reference.
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
	// A boolean value that specifies whether an entity can have more than one component of this type.
	IsSingleton interface{} `field:"optional" json:"isSingleton" yaml:"isSingleton"`
	// An object that maps strings to the property definitions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information about the PropertyDefinitionResponse object, see the [PropertyDefinitionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_PropertyDefinitionResponse.html) API reference.
	PropertyDefinitions interface{} `field:"optional" json:"propertyDefinitions" yaml:"propertyDefinitions"`
	// `AWS::IoTTwinMaker::ComponentType.PropertyGroups`.
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
	// The ComponentType tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

