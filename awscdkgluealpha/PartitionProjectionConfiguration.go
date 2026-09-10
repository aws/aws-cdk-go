package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating partition projection configurations.
//
// Example:
//   var myDatabase Database
//
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("data"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []Column{
//   		&Column{
//   			Name: jsii.String("date"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   	PartitionProjection: map[string]PartitionProjectionConfiguration{
//   		"date": glue.PartitionProjectionConfiguration_date(&DatePartitionProjectionConfigurationProps{
//   			"min": jsii.String("2020-01-01"),
//   			"max": jsii.String("2023-12-31"),
//   			"format": jsii.String("yyyy-MM-dd"),
//   			// `step` bundles interval + unit (supply both or neither). Optional at day
//   			// precision or coarser; required when the format is sub-day (e.g. hours).
//   			"step": &DateProjectionStep{
//   				"interval": jsii.Number(1),
//   				"intervalUnit": glue.DateIntervalUnit_DAYS,
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type PartitionProjectionConfiguration interface {
	// The type of partition projection.
	// Experimental.
	Type() PartitionProjectionType
}

// The jsii proxy struct for PartitionProjectionConfiguration
type jsiiProxy_PartitionProjectionConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_PartitionProjectionConfiguration) Type() PartitionProjectionType {
	var returns PartitionProjectionType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a DATE partition projection configuration.
// Experimental.
func PartitionProjectionConfiguration_Date(props *DatePartitionProjectionConfigurationProps) PartitionProjectionConfiguration {
	_init_.Initialize()

	if err := validatePartitionProjectionConfiguration_DateParameters(props); err != nil {
		panic(err)
	}
	var returns PartitionProjectionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.PartitionProjectionConfiguration",
		"date",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an ENUM partition projection configuration.
// Experimental.
func PartitionProjectionConfiguration_Enum(props *EnumPartitionProjectionConfigurationProps) PartitionProjectionConfiguration {
	_init_.Initialize()

	if err := validatePartitionProjectionConfiguration_EnumParameters(props); err != nil {
		panic(err)
	}
	var returns PartitionProjectionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.PartitionProjectionConfiguration",
		"enum",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an INJECTED partition projection configuration.
//
// Partition values are injected at query time through the query statement.
// See: https://docs.aws.amazon.com/athena/latest/ug/partition-projection-supported-types.html#partition-projection-injected-type
//
// Experimental.
func PartitionProjectionConfiguration_Injected() PartitionProjectionConfiguration {
	_init_.Initialize()

	var returns PartitionProjectionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.PartitionProjectionConfiguration",
		"injected",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create an INTEGER partition projection configuration.
// Experimental.
func PartitionProjectionConfiguration_Integer(props *IntegerPartitionProjectionConfigurationProps) PartitionProjectionConfiguration {
	_init_.Initialize()

	if err := validatePartitionProjectionConfiguration_IntegerParameters(props); err != nil {
		panic(err)
	}
	var returns PartitionProjectionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.PartitionProjectionConfiguration",
		"integer",
		[]interface{}{props},
		&returns,
	)

	return returns
}

