package awsapigateway


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
type JsonSchemaType string

const (
	JsonSchemaType_NULL JsonSchemaType = "NULL"
	JsonSchemaType_BOOLEAN JsonSchemaType = "BOOLEAN"
	JsonSchemaType_OBJECT JsonSchemaType = "OBJECT"
	JsonSchemaType_ARRAY JsonSchemaType = "ARRAY"
	JsonSchemaType_NUMBER JsonSchemaType = "NUMBER"
	JsonSchemaType_INTEGER JsonSchemaType = "INTEGER"
	JsonSchemaType_STRING JsonSchemaType = "STRING"
)

