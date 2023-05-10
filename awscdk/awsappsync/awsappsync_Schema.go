package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Schema for a GraphQL Api.
//
// If no options are configured, schema will be generated
// code-first.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Schema: appsync.Schema_FromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.AddHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &HttpDataSourceOptions{
//   	Name: jsii.String("httpDsWithStepF"),
//   	Description: jsii.String("from appsync to StepFunctions Workflow"),
//   	AuthorizationConfig: &AwsIamConfig{
//   		SigningRegion: jsii.String("us-east-1"),
//   		SigningServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.CreateResolver(&BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("callStepFunction"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type Schema interface {
	// The definition for this schema.
	// Experimental.
	Definition() *string
	// Experimental.
	SetDefinition(val *string)
	// Add a mutation field to the schema's Mutation. CDK will create an Object Type called 'Mutation'. For example,.
	//
	// type Mutation {
	//    fieldName: Field.returnType
	// }.
	// Experimental.
	AddMutation(fieldName *string, field ResolvableField) ObjectType
	// Add a query field to the schema's Query. CDK will create an Object Type called 'Query'. For example,.
	//
	// type Query {
	//    fieldName: Field.returnType
	// }.
	// Experimental.
	AddQuery(fieldName *string, field ResolvableField) ObjectType
	// Add a subscription field to the schema's Subscription. CDK will create an Object Type called 'Subscription'. For example,.
	//
	// type Subscription {
	//    fieldName: Field.returnType
	// }.
	// Experimental.
	AddSubscription(fieldName *string, field Field) ObjectType
	// Escape hatch to add to Schema as desired.
	//
	// Will always result
	// in a newline.
	// Experimental.
	AddToSchema(addition *string, delimiter *string)
	// Add type to the schema.
	// Experimental.
	AddType(type_ IIntermediateType) IIntermediateType
	// Called when the GraphQL Api is initialized to allow this object to bind to the stack.
	// Experimental.
	Bind(api GraphqlApi) CfnGraphQLSchema
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	_ byte // padding
}

func (j *jsiiProxy_Schema) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}


// Experimental.
func NewSchema(options *SchemaOptions) Schema {
	_init_.Initialize()

	if err := validateNewSchemaParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Schema{}

	_jsii_.Create(
		"monocdk.aws_appsync.Schema",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewSchema_Override(s Schema, options *SchemaOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.Schema",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_Schema)SetDefinition(val *string) {
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
func Schema_FromAsset(filePath *string) Schema {
	_init_.Initialize()

	if err := validateSchema_FromAssetParameters(filePath); err != nil {
		panic(err)
	}
	var returns Schema

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Schema",
		"fromAsset",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddMutation(fieldName *string, field ResolvableField) ObjectType {
	if err := s.validateAddMutationParameters(fieldName, field); err != nil {
		panic(err)
	}
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addMutation",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddQuery(fieldName *string, field ResolvableField) ObjectType {
	if err := s.validateAddQueryParameters(fieldName, field); err != nil {
		panic(err)
	}
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addQuery",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddSubscription(fieldName *string, field Field) ObjectType {
	if err := s.validateAddSubscriptionParameters(fieldName, field); err != nil {
		panic(err)
	}
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addSubscription",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddToSchema(addition *string, delimiter *string) {
	if err := s.validateAddToSchemaParameters(addition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToSchema",
		[]interface{}{addition, delimiter},
	)
}

func (s *jsiiProxy_Schema) AddType(type_ IIntermediateType) IIntermediateType {
	if err := s.validateAddTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns IIntermediateType

	_jsii_.Invoke(
		s,
		"addType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) Bind(api GraphqlApi) CfnGraphQLSchema {
	if err := s.validateBindParameters(api); err != nil {
		panic(err)
	}
	var returns CfnGraphQLSchema

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{api},
		&returns,
	)

	return returns
}

