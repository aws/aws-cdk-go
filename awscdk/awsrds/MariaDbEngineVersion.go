package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the MariaDB instance engines (those returned by `DatabaseInstanceEngine.mariaDb`).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mariaDbEngineVersion := awscdk.Aws_rds.MariaDbEngineVersion_VER_10_11()
//
type MariaDbEngineVersion interface {
	// The full version string, for example, "10.5.28".
	MariaDbFullVersion() *string
	// The major version of the engine, for example, "10.5".
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
func MariaDbEngineVersion_Of(mariaDbFullVersion *string, mariaDbMajorVersion *string) MariaDbEngineVersion {
	_init_.Initialize()

	if err := validateMariaDbEngineVersion_OfParameters(mariaDbFullVersion, mariaDbMajorVersion); err != nil {
		panic(err)
	}
	var returns MariaDbEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"of",
		[]interface{}{mariaDbFullVersion, mariaDbMajorVersion},
		&returns,
	)

	return returns
}

func MariaDbEngineVersion_VER_10_11() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_11",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_11_4() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_11_4",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_11_5() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_11_5",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_11_6() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_11_6",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_11_7() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_11_7",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_11() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_11",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_12() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_12",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_15() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_15",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_21() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_21",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_32() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_32",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_37() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_37",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_39() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_39",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_40() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_40",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_41() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_41",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_43() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_43",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_2_44() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_2_44",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_13() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_13",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_20() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_20",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_23() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_23",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_28() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_28",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_31() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_31",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_32() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_32",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_34() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_34",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_35() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_35",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_36() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_36",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_37() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_37",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_38() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_38",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_39() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_39",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_3_8() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_3_8",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_13() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_13",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_18() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_18",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_21() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_21",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_22() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_22",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_24() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_24",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_25() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_25",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_26() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_26",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_27() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_27",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_28() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_28",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_29() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_29",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_30() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_30",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_31() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_31",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_32() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_32",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_33() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_33",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_4_8() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_4_8",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_12() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_12",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_13() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_13",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_15() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_15",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_16() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_16",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_17() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_17",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_18() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_18",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_19() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_19",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_20() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_20",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_21() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_21",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_22() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_22",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_23() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_23",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_24() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_24",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_8() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_8",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_5_9() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_5_9",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_10() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_10",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_11() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_11",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_12() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_12",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_13() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_13",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_14() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_14",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_15() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_15",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_16() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_16",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_17() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_17",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_5() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_5",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_7() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_7",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_6_8() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_6_8",
		&returns,
	)
	return returns
}

