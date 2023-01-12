package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the Aurora PostgreSQL cluster engine (those returned by `DatabaseClusterEngine.auroraPostgres`).
//
// https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Updates.html
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraPostgresEngineVersion := awscdk.Aws_rds.auroraPostgresEngineVersion_VER_10_11()
//
type AuroraPostgresEngineVersion interface {
	// The full version string, for example, "9.6.25.1".
	AuroraPostgresFullVersion() *string
	// The major version of the engine, for example, "9.6".
	AuroraPostgresMajorVersion() *string
}

// The jsii proxy struct for AuroraPostgresEngineVersion
type jsiiProxy_AuroraPostgresEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AuroraPostgresEngineVersion) AuroraPostgresFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraPostgresFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraPostgresEngineVersion) AuroraPostgresMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraPostgresMajorVersion",
		&returns,
	)
	return returns
}


// Create a new AuroraPostgresEngineVersion with an arbitrary version.
func AuroraPostgresEngineVersion_Of(auroraPostgresFullVersion *string, auroraPostgresMajorVersion *string, auroraPostgresFeatures *AuroraPostgresEngineFeatures) AuroraPostgresEngineVersion {
	_init_.Initialize()

	if err := validateAuroraPostgresEngineVersion_OfParameters(auroraPostgresFullVersion, auroraPostgresMajorVersion, auroraPostgresFeatures); err != nil {
		panic(err)
	}
	var returns AuroraPostgresEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"of",
		[]interface{}{auroraPostgresFullVersion, auroraPostgresMajorVersion, auroraPostgresFeatures},
		&returns,
	)

	return returns
}

func AuroraPostgresEngineVersion_VER_10_11() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_11",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_12() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_12",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_13() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_13",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_14() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_14",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_16() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_16",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_17() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_17",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_18() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_18",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_19() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_19",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_20() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_20",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_21() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_21",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_4() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_4",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_5() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_5",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_6() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_6",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_10_7() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_10_7",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_11() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_11",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_12() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_12",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_13() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_13",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_14() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_14",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_15() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_15",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_16() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_16",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_4() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_4",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_6() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_6",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_7() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_7",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_8() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_8",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_9() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_9",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_10() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_10",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_11() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_11",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_4() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_4",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_6() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_6",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_7() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_7",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_8() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_8",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_9() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_9",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_3() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_3",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_4() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_4",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_5() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_5",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_6() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_6",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_7() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_7",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_14_3() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_3",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_14_4() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_4",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_14_5() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_5",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_11() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_11",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_12() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_12",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_16() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_16",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_17() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_17",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_18() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_18",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_19() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_19",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_8() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_8",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_9_6_9() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_9",
		&returns,
	)
	return returns
}

