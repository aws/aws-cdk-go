package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var default_ interface{}
//   var enum_ interface{}
//   var jsonSchema_ jsonSchema
//   var restApi restApi
//
//   modelProps := &modelProps{
//   	restApi: restApi,
//   	schema: &jsonSchema{
//   		additionalItems: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		additionalProperties: jsii.Boolean(false),
//   		allOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		anyOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		contains: jsonSchema_,
//   		default: default_,
//   		definitions: map[string]*jsonSchema{
//   			"definitionsKey": jsonSchema_,
//   		},
//   		dependencies: map[string]interface{}{
//   			"dependenciesKey": jsonSchema_,
//   		},
//   		description: jsii.String("description"),
//   		enum: []interface{}{
//   			enum_,
//   		},
//   		exclusiveMaximum: jsii.Boolean(false),
//   		exclusiveMinimum: jsii.Boolean(false),
//   		format: jsii.String("format"),
//   		id: jsii.String("id"),
//   		items: jsonSchema_,
//   		maximum: jsii.Number(123),
//   		maxItems: jsii.Number(123),
//   		maxLength: jsii.Number(123),
//   		maxProperties: jsii.Number(123),
//   		minimum: jsii.Number(123),
//   		minItems: jsii.Number(123),
//   		minLength: jsii.Number(123),
//   		minProperties: jsii.Number(123),
//   		multipleOf: jsii.Number(123),
//   		not: jsonSchema_,
//   		oneOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		pattern: jsii.String("pattern"),
//   		patternProperties: map[string]*jsonSchema{
//   			"patternPropertiesKey": jsonSchema_,
//   		},
//   		properties: map[string]*jsonSchema{
//   			"propertiesKey": jsonSchema_,
//   		},
//   		propertyNames: jsonSchema_,
//   		ref: jsii.String("ref"),
//   		required: []*string{
//   			jsii.String("required"),
//   		},
//   		schema: awscdk.Aws_apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("title"),
//   		type: awscdk.*Aws_apigateway.jsonSchemaType_NULL,
//   		uniqueItems: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	description: jsii.String("description"),
//   	modelName: jsii.String("modelName"),
//   }
//
type ModelProps struct {
	// The schema to use to transform data to one or more output formats.
	//
	// Specify null ({}) if you don't want to specify a schema.
	Schema *JsonSchema `field:"required" json:"schema" yaml:"schema"`
	// The content type for the model.
	//
	// You can also force a
	// content type in the request or response model mapping.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description that identifies this model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// Important
	//   If you specify a name, you cannot perform updates that
	//   require replacement of this resource. You can perform
	//   updates that require no or some interruption. If you
	//   must replace the resource, specify a new name.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The rest API that this model is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

