package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
)

// Interface for implementing your own schema.
//
// Useful for providing schema's from sources other than assets.
type ISchema interface {
	// Binds a schema string to a GraphQlApi.
	//
	// Returns: ISchemaConfig with apiId and schema definition string.
	Bind(api interfacesawsappsync.IGraphQLApiRef, options *SchemaBindOptions) ISchemaConfig
}

// The jsii proxy for ISchema
type jsiiProxy_ISchema struct {
	_ byte // padding
}

func (i *jsiiProxy_ISchema) Bind(api interfacesawsappsync.IGraphQLApiRef, options *SchemaBindOptions) ISchemaConfig {
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

