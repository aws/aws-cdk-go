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
//   			"interval": jsii.Number(1),
//   			 // optional, defaults to 1
//   			"intervalUnit": glue.DateIntervalUnit_DAYS,
//   		}),
//   	},
//   })
//
// Experimental.
type PartitionProjectionConfiguration interface {
	// Range of partition values for DATE type.
	//
	// Array of [start, end] as date strings.
	// Experimental.
	DateRange() *[]*string
	// Number of digits to pad INTEGER partition values.
	// Experimental.
	Digits() *float64
	// Date format for DATE partition values (Java SimpleDateFormat).
	// Experimental.
	Format() *string
	// Range of partition values for INTEGER type.
	//
	// Array of [min, max] as numbers.
	// Experimental.
	IntegerRange() *[]*float64
	// Interval between partition values.
	// Experimental.
	Interval() *float64
	// Unit for DATE partition interval.
	// Experimental.
	IntervalUnit() DateIntervalUnit
	// The type of partition projection.
	// Experimental.
	Type() PartitionProjectionType
	// Explicit list of values for ENUM partitions.
	// Experimental.
	Values() *[]*string
}

// The jsii proxy struct for PartitionProjectionConfiguration
type jsiiProxy_PartitionProjectionConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_PartitionProjectionConfiguration) DateRange() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dateRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartitionProjectionConfiguration) Digits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartitionProjectionConfiguration) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartitionProjectionConfiguration) IntegerRange() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"integerRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartitionProjectionConfiguration) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PartitionProjectionConfiguration) IntervalUnit() DateIntervalUnit {
	var returns DateIntervalUnit
	_jsii_.Get(
		j,
		"intervalUnit",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_PartitionProjectionConfiguration) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
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

