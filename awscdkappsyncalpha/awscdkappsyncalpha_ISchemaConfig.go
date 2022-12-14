// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for bound graphql schema.
//
// Returned from ISchema.bind allowing late binding of schemas to graphqlapi-base
// Experimental.
type ISchemaConfig interface {
	// The ID of the api the schema is bound to.
	// Experimental.
	ApiId() *string
	// Experimental.
	SetApiId(a *string)
	// The schema definition string.
	// Experimental.
	Definition() *string
	// Experimental.
	SetDefinition(d *string)
}

// The jsii proxy for ISchemaConfig
type jsiiProxy_ISchemaConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_ISchemaConfig) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchemaConfig)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_ISchemaConfig) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchemaConfig)SetDefinition(val *string) {
	if err := j.validateSetDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

