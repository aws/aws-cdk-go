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
// Experimental.
type JsonSchema struct {
	// Experimental.
	AdditionalItems *[]*JsonSchema `field:"optional" json:"additionalItems" yaml:"additionalItems"`
	// Experimental.
	AdditionalProperties interface{} `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Experimental.
	AllOf *[]*JsonSchema `field:"optional" json:"allOf" yaml:"allOf"`
	// Experimental.
	AnyOf *[]*JsonSchema `field:"optional" json:"anyOf" yaml:"anyOf"`
	// Experimental.
	Contains interface{} `field:"optional" json:"contains" yaml:"contains"`
	// The default value if you use an enum.
	// Experimental.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Experimental.
	Definitions *map[string]*JsonSchema `field:"optional" json:"definitions" yaml:"definitions"`
	// Experimental.
	Dependencies *map[string]interface{} `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	Enum *[]interface{} `field:"optional" json:"enum" yaml:"enum"`
	// Experimental.
	ExclusiveMaximum *bool `field:"optional" json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	// Experimental.
	ExclusiveMinimum *bool `field:"optional" json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	// Experimental.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Experimental.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Experimental.
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// Experimental.
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Experimental.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Experimental.
	MaxProperties *float64 `field:"optional" json:"maxProperties" yaml:"maxProperties"`
	// Experimental.
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// Experimental.
	MinItems *float64 `field:"optional" json:"minItems" yaml:"minItems"`
	// Experimental.
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// Experimental.
	MinProperties *float64 `field:"optional" json:"minProperties" yaml:"minProperties"`
	// Experimental.
	MultipleOf *float64 `field:"optional" json:"multipleOf" yaml:"multipleOf"`
	// Experimental.
	Not **JsonSchema `field:"optional" json:"not" yaml:"not"`
	// Experimental.
	OneOf *[]*JsonSchema `field:"optional" json:"oneOf" yaml:"oneOf"`
	// Experimental.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Experimental.
	PatternProperties *map[string]*JsonSchema `field:"optional" json:"patternProperties" yaml:"patternProperties"`
	// Experimental.
	Properties *map[string]*JsonSchema `field:"optional" json:"properties" yaml:"properties"`
	// Experimental.
	PropertyNames **JsonSchema `field:"optional" json:"propertyNames" yaml:"propertyNames"`
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// Experimental.
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
	// Experimental.
	Schema JsonSchemaVersion `field:"optional" json:"schema" yaml:"schema"`
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Experimental.
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// Experimental.
	UniqueItems *bool `field:"optional" json:"uniqueItems" yaml:"uniqueItems"`
}

