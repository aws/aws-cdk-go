// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for implementing your own schema.
//
// Useful for providing schema's from sources other than assets.
// Experimental.
type ISchema interface {
	// Binds a schema string to a GraphQlApi.
	//
	// Returns: ISchemaConfig with apiId and schema definition string.
	// Experimental.
	Bind(api IGraphqlApi, options *SchemaBindOptions) ISchemaConfig
}

// The jsii proxy for ISchema
type jsiiProxy_ISchema struct {
	_ byte // padding
}

func (i *jsiiProxy_ISchema) Bind(api IGraphqlApi, options *SchemaBindOptions) ISchemaConfig {
	if err := i.validateBindParameters(api, options); err != nil {
		panic(err)
	}
	var returns ISchemaConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{api, options},
		&returns,
	)

	return returns
}

