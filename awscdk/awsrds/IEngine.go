package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A common interface for database engines.
//
// Don't implement this interface directly,
// instead implement one of the known sub-interfaces,
// like IClusterEngine and IInstanceEngine.
type IEngine interface {
	// The default name of the master database user if one was not provided explicitly.
	//
	// The global default of 'admin' will be used if this is `undefined`.
	// Note that 'admin' is a reserved word in PostgreSQL and cannot be used.
	DefaultUsername() *string
	// The family this engine belongs to, like "MYSQL", or "POSTGRESQL".
	//
	// This property is used when creating a Database Proxy.
	// Most engines don't belong to any family
	// (and because of that, you can't create Database Proxies for their Clusters or Instances).
	EngineFamily() *string
	// The type of the engine, for example "mysql".
	EngineType() *string
	// The exact version of the engine that is used, for example "5.1.42".
	EngineVersion() *EngineVersion
	// The family to use for ParameterGroups using this engine.
	//
	// This is usually equal to "<engineType><engineMajorVersion>",
	// but can sometimes be a variation of that.
	// You can pass this property when creating new ParameterGroup.
	ParameterGroupFamily() *string
}

// The jsii proxy for IEngine
type jsiiProxy_IEngine struct {
	_ byte // padding
}

func (j *jsiiProxy_IEngine) DefaultUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) EngineFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) EngineVersion() *EngineVersion {
	var returns *EngineVersion
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) ParameterGroupFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamily",
		&returns,
	)
	return returns
}

