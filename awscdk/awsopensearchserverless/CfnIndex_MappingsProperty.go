package awsopensearchserverless


// Index Mappings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var propertyMappingProperty_ propertyMappingProperty
//
//   mappingsProperty := &MappingsProperty{
//   	Properties: map[string]interface{}{
//   		"propertiesKey": &propertyMappingProperty{
//   			"type": jsii.String("type"),
//
//   			// the properties below are optional
//   			"dimension": jsii.Number(123),
//   			"index": jsii.Boolean(false),
//   			"method": &MethodProperty{
//   				"engine": jsii.String("engine"),
//   				"name": jsii.String("name"),
//
//   				// the properties below are optional
//   				"parameters": &ParametersProperty{
//   					"efConstruction": jsii.Number(123),
//   					"m": jsii.Number(123),
//   				},
//   				"spaceType": jsii.String("spaceType"),
//   			},
//   			"properties": map[string]interface{}{
//   				"propertiesKey": propertyMappingProperty_,
//   			},
//   			"value": jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-mappings.html
//
type CfnIndex_MappingsProperty struct {
	// Defines the fields within the mapping, including their types and configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-mappings.html#cfn-opensearchserverless-index-mappings-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

