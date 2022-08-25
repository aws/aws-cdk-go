package awsapigateway


// Represents a JSON schema definition of the structure of a REST API model.
//
// Copied from npm module jsonschema.
//
// Example:
//   var api restApi
//
//
//   // We define the JSON Schema for the transformed valid response
//   responseModel := api.addModel(jsii.String("ResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("pollResponse"),
//   		type: apigateway.jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"greeting": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   // We define the JSON Schema for the transformed error response
//   errorResponseModel := api.addModel(jsii.String("ErrorResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ErrorResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.*jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("errorResponse"),
//   		type: apigateway.*jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"message": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
// See: https://github.com/tdegrunt/jsonschema
//
type JsonSchema struct {
	AdditionalItems *[]*JsonSchema `field:"optional" json:"additionalItems" yaml:"additionalItems"`
	AdditionalProperties interface{} `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	AllOf *[]*JsonSchema `field:"optional" json:"allOf" yaml:"allOf"`
	AnyOf *[]*JsonSchema `field:"optional" json:"anyOf" yaml:"anyOf"`
	Contains interface{} `field:"optional" json:"contains" yaml:"contains"`
	// The default value if you use an enum.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	Definitions *map[string]*JsonSchema `field:"optional" json:"definitions" yaml:"definitions"`
	Dependencies *map[string]interface{} `field:"optional" json:"dependencies" yaml:"dependencies"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Enum *[]interface{} `field:"optional" json:"enum" yaml:"enum"`
	ExclusiveMaximum *bool `field:"optional" json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum *bool `field:"optional" json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Format *string `field:"optional" json:"format" yaml:"format"`
	Id *string `field:"optional" json:"id" yaml:"id"`
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	MaxProperties *float64 `field:"optional" json:"maxProperties" yaml:"maxProperties"`
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	MinItems *float64 `field:"optional" json:"minItems" yaml:"minItems"`
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	MinProperties *float64 `field:"optional" json:"minProperties" yaml:"minProperties"`
	MultipleOf *float64 `field:"optional" json:"multipleOf" yaml:"multipleOf"`
	Not **JsonSchema `field:"optional" json:"not" yaml:"not"`
	OneOf *[]*JsonSchema `field:"optional" json:"oneOf" yaml:"oneOf"`
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	PatternProperties *map[string]*JsonSchema `field:"optional" json:"patternProperties" yaml:"patternProperties"`
	Properties *map[string]*JsonSchema `field:"optional" json:"properties" yaml:"properties"`
	PropertyNames **JsonSchema `field:"optional" json:"propertyNames" yaml:"propertyNames"`
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
	Schema JsonSchemaVersion `field:"optional" json:"schema" yaml:"schema"`
	Title *string `field:"optional" json:"title" yaml:"title"`
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	UniqueItems *bool `field:"optional" json:"uniqueItems" yaml:"uniqueItems"`
}

