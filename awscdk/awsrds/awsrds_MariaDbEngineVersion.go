package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the MariaDB instance engines (those returned by {@link DatabaseInstanceEngine.mariaDb}).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mariaDbEngineVersion := awscdk.Aws_rds.MariaDbEngineVersion_VER_10_0()
//
// Experimental.
type MariaDbEngineVersion interface {
	// The full version string, for example, "10.5.28".
	// Experimental.
	MariaDbFullVersion() *string
	// The major version of the engine, for example, "10.5".
	// Experimental.
	MariaDbMajorVersion() *string
}

// The jsii proxy struct for MariaDbEngineVersion
type jsiiProxy_MariaDbEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_MariaDbEngineVersion) MariaDbFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariaDbFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariaDbEngineVersion) MariaDbMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariaDbMajorVersion",
		&returns,
	)
	return returns
}


// Create a new MariaDbEngineVersion with an arbitrary version.
// Experimental.
func MariaDbEngineVersion_Of(mariaDbFullVersion *string, mariaDbMajorVersion *string) MariaDbEngineVersion {
	_init_.Initialize()

	if err := validateMariaDbEngineVersion_OfParameters(mariaDbFullVersion, mariaDbMajorVersion); err != nil {
		panic(err)
	}
	var returns MariaDbEngineVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"of",
		[]interface{}{mariaDbFullVersion, mariaDbMajorVersion},
		&returns,
	)

	return returns
}

func MariaDbEngineVersion_VER_10_0() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_17() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0_17",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_24() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0_24",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_28() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0_28",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_31() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0_31",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_32() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0_32",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_34() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0_34",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_35() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_0_35",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_1",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_14() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_1_14",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_19() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_1_19",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_23() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_1_23",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_26() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_1_26",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_31() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_1_31",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_34() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_1_34",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_11() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_11",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_12() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_12",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_15() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_15",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_21() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_21",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_32() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_32",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_37() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_37",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_39() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_39",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_40() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_40",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_41() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_2_41",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_13() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3_13",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_20() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3_20",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_23() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3_23",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_28() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3_28",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_31() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3_31",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_32() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3_32",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_8() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_3_8",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_4",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_13() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_4_13",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_18() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_4_18",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_21() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_4_21",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_22() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_4_22",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_8() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_4_8",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_5",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_12() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_5_12",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_13() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_5_13",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_8() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_5_8",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_9() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MariaDbEngineVersion",
		"VER_10_5_9",
		&returns,
	)
	return returns
}

