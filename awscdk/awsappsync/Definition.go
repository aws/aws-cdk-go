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
//   sourceApi := appsync.NewGraphqlApi(this, jsii.String("FirstSourceAPI"), &GraphqlApiProps{
//   	Name: jsii.String("FirstSourceAPI"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.merged-api-1.graphql"))),
//   })
//
//   importedMergedApi := appsync.GraphqlApi_FromGraphqlApiAttributes(this, jsii.String("ImportedMergedApi"), &GraphqlApiAttributes{
//   	GraphqlApiId: jsii.String("MyApiId"),
//   	GraphqlApiArn: jsii.String("MyApiArn"),
//   })
//
//   importedExecutionRole := iam.Role_FromRoleArn(this, jsii.String("ExecutionRole"), jsii.String("arn:aws:iam::ACCOUNT:role/MyExistingRole"))
//   appsync.NewSourceApiAssociation(this, jsii.String("SourceApiAssociation2"), &SourceApiAssociationProps{
//   	SourceApi: sourceApi,
//   	MergedApi: importedMergedApi,
//   	MergeType: appsync.MergeType_MANUAL_MERGE,
//   	MergedApiExecutionRole: importedExecutionRole,
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

