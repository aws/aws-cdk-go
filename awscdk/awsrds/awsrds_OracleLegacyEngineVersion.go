package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the legacy Oracle instance engines (those returned by {@link DatabaseInstanceEngine.oracleSe} and {@link DatabaseInstanceEngine.oracleSe1}). Note: RDS will stop allowing creating new databases with this version in August 2020.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oracleLegacyEngineVersion := awscdk.Aws_rds.oracleLegacyEngineVersion_VER_11_2()
//
// Deprecated: instances can no longer be created with these engine versions. See https://forums.aws.amazon.com/ann.jspa?annID=7341
type OracleLegacyEngineVersion interface {
	// The full version string, for example, "11.2.0.4.v24".
	// Deprecated: instances can no longer be created with these engine versions. See https://forums.aws.amazon.com/ann.jspa?annID=7341
	OracleLegacyFullVersion() *string
	// The major version of the engine, for example, "11.2".
	// Deprecated: instances can no longer be created with these engine versions. See https://forums.aws.amazon.com/ann.jspa?annID=7341
	OracleLegacyMajorVersion() *string
}

// The jsii proxy struct for OracleLegacyEngineVersion
type jsiiProxy_OracleLegacyEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_OracleLegacyEngineVersion) OracleLegacyFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleLegacyFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleLegacyEngineVersion) OracleLegacyMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleLegacyMajorVersion",
		&returns,
	)
	return returns
}


func OracleLegacyEngineVersion_VER_11_2() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_2_V2() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_2_V2",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V1() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V1",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V10() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V10",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V11() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V11",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V12() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V12",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V13() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V13",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V14() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V14",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V15() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V15",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V16() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V16",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V17() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V17",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V18() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V18",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V19() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V19",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V20() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V20",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V21() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V21",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V22() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V22",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V23() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V23",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V24() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V24",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V25() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V25",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V3() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V3",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V4() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V4",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V5() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V5",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V6() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V6",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V7() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V7",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V8() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V8",
		&returns,
	)
	return returns
}

func OracleLegacyEngineVersion_VER_11_2_0_4_V9() OracleLegacyEngineVersion {
	_init_.Initialize()
	var returns OracleLegacyEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.OracleLegacyEngineVersion",
		"VER_11_2_0_4_V9",
		&returns,
	)
	return returns
}

