package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Serialization library to use when serializing/deserializing (SerDe) table records.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   serializationLibrary := glue_alpha.SerializationLibrary_AVRO()
//
// See: https://cwiki.apache.org/confluence/display/Hive/SerDe
//
// Experimental.
type SerializationLibrary interface {
	// Experimental.
	ClassName() *string
}

// The jsii proxy struct for SerializationLibrary
type jsiiProxy_SerializationLibrary struct {
	_ byte // padding
}

func (j *jsiiProxy_SerializationLibrary) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}


// Experimental.
func NewSerializationLibrary(className *string) SerializationLibrary {
	_init_.Initialize()

	if err := validateNewSerializationLibraryParameters(className); err != nil {
		panic(err)
	}
	j := jsiiProxy_SerializationLibrary{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		[]interface{}{className},
		&j,
	)

	return &j
}

// Experimental.
func NewSerializationLibrary_Override(s SerializationLibrary, className *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		[]interface{}{className},
		s,
	)
}

func SerializationLibrary_AVRO() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"AVRO",
		&returns,
	)
	return returns
}

func SerializationLibrary_CLOUDTRAIL() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"CLOUDTRAIL",
		&returns,
	)
	return returns
}

func SerializationLibrary_GROK() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"GROK",
		&returns,
	)
	return returns
}

func SerializationLibrary_HIVE_JSON() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"HIVE_JSON",
		&returns,
	)
	return returns
}

func SerializationLibrary_LAZY_SIMPLE() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"LAZY_SIMPLE",
		&returns,
	)
	return returns
}

func SerializationLibrary_OPEN_CSV() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"OPEN_CSV",
		&returns,
	)
	return returns
}

func SerializationLibrary_OPENX_JSON() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"OPENX_JSON",
		&returns,
	)
	return returns
}

func SerializationLibrary_ORC() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"ORC",
		&returns,
	)
	return returns
}

func SerializationLibrary_PARQUET() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"PARQUET",
		&returns,
	)
	return returns
}

func SerializationLibrary_REGEXP() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		"REGEXP",
		&returns,
	)
	return returns
}

