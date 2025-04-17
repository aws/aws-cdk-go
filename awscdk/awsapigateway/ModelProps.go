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
//   modelProps := &ModelProps{
//   	RestApi: restApi,
//   	Schema: &jsonSchema{
//   		AdditionalItems: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		AdditionalProperties: jsii.Boolean(false),
//   		AllOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		AnyOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		Contains: jsonSchema_,
//   		Default: default_,
//   		Definitions: map[string]*jsonSchema{
//   			"definitionsKey": jsonSchema_,
//   		},
//   		Dependencies: map[string]interface{}{
//   			"dependenciesKey": []interface{}{
//   				jsii.String("dependencies"),
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		Enum: []interface{}{
//   			enum_,
//   		},
//   		ExclusiveMaximum: jsii.Boolean(false),
//   		ExclusiveMinimum: jsii.Boolean(false),
//   		Format: jsii.String("format"),
//   		Id: jsii.String("id"),
//   		Items: jsonSchema_,
//   		Maximum: jsii.Number(123),
//   		MaxItems: jsii.Number(123),
//   		MaxLength: jsii.Number(123),
//   		MaxProperties: jsii.Number(123),
//   		Minimum: jsii.Number(123),
//   		MinItems: jsii.Number(123),
//   		MinLength: jsii.Number(123),
//   		MinProperties: jsii.Number(123),
//   		MultipleOf: jsii.Number(123),
//   		Not: jsonSchema_,
//   		OneOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		Pattern: jsii.String("pattern"),
//   		PatternProperties: map[string]*jsonSchema{
//   			"patternPropertiesKey": jsonSchema_,
//   		},
//   		Properties: map[string]*jsonSchema{
//   			"propertiesKey": jsonSchema_,
//   		},
//   		PropertyNames: jsonSchema_,
//   		Ref: jsii.String("ref"),
//   		Required: []*string{
//   			jsii.String("required"),
//   		},
//   		Schema: awscdk.Aws_apigateway.JsonSchemaVersion_DRAFT4,
//   		Title: jsii.String("title"),
//   		Type: awscdk.*Aws_apigateway.JsonSchemaType_NULL,
//   		UniqueItems: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   	ModelName: jsii.String("modelName"),
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
	// Default: 'application/json'.
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description that identifies this model.
	// Default: None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// Important
	//  If you specify a name, you cannot perform updates that
	//  require replacement of this resource. You can perform
	//  updates that require no or some interruption. If you
	//  must replace the resource, specify a new name.
	// Default: <auto> If you don't specify a name,
	// AWS CloudFormation generates a unique physical ID and
	// uses that ID for the model name. For more information,
	// see Name Type.
	//
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The rest API that this model is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

