package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines the input/output formats and ser/de for a single DataFormat.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &TableProps{
//   	Database: myDatabase,
//   	TableName: jsii.String("my_table"),
//   	Columns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []*column{
//   		&column{
//   			Name: jsii.String("year"),
//   			Type: glue.Schema_SMALL_INT(),
//   		},
//   		&column{
//   			Name: jsii.String("month"),
//   			Type: glue.Schema_SMALL_INT(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
//
// Experimental.
type DataFormat interface {
	// Classification string given to tables with this data format.
	// Experimental.
	ClassificationString() ClassificationString
	// `InputFormat` for this data format.
	// Experimental.
	InputFormat() InputFormat
	// `OutputFormat` for this data format.
	// Experimental.
	OutputFormat() OutputFormat
	// Serialization library for this data format.
	// Experimental.
	SerializationLibrary() SerializationLibrary
}

// The jsii proxy struct for DataFormat
type jsiiProxy_DataFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_DataFormat) ClassificationString() ClassificationString {
	var returns ClassificationString
	_jsii_.Get(
		j,
		"classificationString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormat) InputFormat() InputFormat {
	var returns InputFormat
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormat) OutputFormat() OutputFormat {
	var returns OutputFormat
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormat) SerializationLibrary() SerializationLibrary {
	var returns SerializationLibrary
	_jsii_.Get(
		j,
		"serializationLibrary",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataFormat(props *DataFormatProps) DataFormat {
	_init_.Initialize()

	if err := validateNewDataFormatParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFormat{}

	_jsii_.Create(
		"monocdk.aws_glue.DataFormat",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewDataFormat_Override(d DataFormat, props *DataFormatProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_glue.DataFormat",
		[]interface{}{props},
		d,
	)
}

func DataFormat_APACHE_LOGS() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"APACHE_LOGS",
		&returns,
	)
	return returns
}

func DataFormat_AVRO() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"AVRO",
		&returns,
	)
	return returns
}

func DataFormat_CLOUDTRAIL_LOGS() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"CLOUDTRAIL_LOGS",
		&returns,
	)
	return returns
}

func DataFormat_CSV() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"CSV",
		&returns,
	)
	return returns
}

func DataFormat_JSON() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"JSON",
		&returns,
	)
	return returns
}

func DataFormat_LOGSTASH() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"LOGSTASH",
		&returns,
	)
	return returns
}

func DataFormat_ORC() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"ORC",
		&returns,
	)
	return returns
}

func DataFormat_PARQUET() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"PARQUET",
		&returns,
	)
	return returns
}

func DataFormat_TSV() DataFormat {
	_init_.Initialize()
	var returns DataFormat
	_jsii_.StaticGet(
		"monocdk.aws_glue.DataFormat",
		"TSV",
		&returns,
	)
	return returns
}

