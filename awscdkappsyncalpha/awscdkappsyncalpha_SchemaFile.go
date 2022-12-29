// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Schema for a GraphQL Api.
//
// If no options are configured, schema will be generated
// code-first.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.addHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &httpDataSourceOptions{
//   	name: jsii.String("httpDsWithStepF"),
//   	description: jsii.String("from appsync to StepFunctions Workflow"),
//   	authorizationConfig: &awsIamConfig{
//   		signingRegion: jsii.String("us-east-1"),
//   		signingServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.createResolver(jsii.String("MutationCallStepFunctionResolver"), &baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("callStepFunction"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type SchemaFile interface {
	ISchema
	// The definition for this schema.
	// Experimental.
	Definition() *string
	// Experimental.
	SetDefinition(val *string)
	// Called when the GraphQL Api is initialized to allow this object to bind to the stack.
	// Experimental.
	Bind(api IGraphqlApi, _options *SchemaBindOptions) ISchemaConfig
}

// The jsii proxy struct for SchemaFile
type jsiiProxy_SchemaFile struct {
	jsiiProxy_ISchema
}

func (j *jsiiProxy_SchemaFile) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}


// Experimental.
func NewSchemaFile(options *SchemaProps) SchemaFile {
	_init_.Initialize()

	if err := validateNewSchemaFileParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SchemaFile{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.SchemaFile",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewSchemaFile_Override(s SchemaFile, options *SchemaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.SchemaFile",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SchemaFile)SetDefinition(val *string) {
	if err := j.validateSetDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

// Generate a Schema from file.
//
// Returns: `SchemaAsset` with immutable schema defintion.
// Experimental.
func SchemaFile_FromAsset(filePath *string) SchemaFile {
	_init_.Initialize()

	if err := validateSchemaFile_FromAssetParameters(filePath); err != nil {
		panic(err)
	}
	var returns SchemaFile

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.SchemaFile",
		"fromAsset",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFile) Bind(api IGraphqlApi, _options *SchemaBindOptions) ISchemaConfig {
	if err := s.validateBindParameters(api, _options); err != nil {
		panic(err)
	}
	var returns ISchemaConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{api, _options},
		&returns,
	)

	return returns
}

