package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AppSync definition.
//
// Specify how you want to define your AppSync API.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//
//   api := appsync.NewGraphqlApi(this, jsii.String("EventBridgeApi"), &GraphqlApiProps{
//   	Name: jsii.String("EventBridgeApi"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.eventbridge.graphql"))),
//   })
//
//   bus := events.NewEventBus(this, jsii.String("DestinationEventBus"), &EventBusProps{
//   })
//
//   dataSource := api.AddEventBridgeDataSource(jsii.String("NoneDS"), bus)
//
//   dataSource.CreateResolver(jsii.String("EventResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("emitEvent"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
type Definition interface {
	// Schema, when AppSync API is created from schema file.
	Schema() ISchema
	// Source APIs for Merged API.
	SourceApiOptions() *SourceApiOptions
}

// The jsii proxy struct for Definition
type jsiiProxy_Definition struct {
	_ byte // padding
}

func (j *jsiiProxy_Definition) Schema() ISchema {
	var returns ISchema
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Definition) SourceApiOptions() *SourceApiOptions {
	var returns *SourceApiOptions
	_jsii_.Get(
		j,
		"sourceApiOptions",
		&returns,
	)
	return returns
}


func NewDefinition_Override(d Definition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Definition",
		nil, // no parameters
		d,
	)
}

// Schema from file, allows schema definition through schema.graphql file.
//
// Returns: API Source with schema from file.
func Definition_FromFile(filePath *string) Definition {
	_init_.Initialize()

	if err := validateDefinition_FromFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns Definition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Definition",
		"fromFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Schema from schema object.
//
// Returns: API Source with schema from file.
func Definition_FromSchema(schema ISchema) Definition {
	_init_.Initialize()

	if err := validateDefinition_FromSchemaParameters(schema); err != nil {
		panic(err)
	}
	var returns Definition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Definition",
		"fromSchema",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Schema from existing AppSync APIs - used for creating a AppSync Merged API.
//
// Returns: API Source with for AppSync Merged API.
func Definition_FromSourceApis(sourceApiOptions *SourceApiOptions) Definition {
	_init_.Initialize()

	if err := validateDefinition_FromSourceApisParameters(sourceApiOptions); err != nil {
		panic(err)
	}
	var returns Definition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Definition",
		"fromSourceApis",
		[]interface{}{sourceApiOptions},
		&returns,
	)

	return returns
}

