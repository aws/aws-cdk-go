package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the MySQL instance engines (those returned by `DatabaseInstanceEngine.mysql`).
//
// Example:
//   var vpc vpc
//
//
//   iopsInstance := rds.NewDatabaseInstance(this, jsii.String("IopsInstance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_30(),
//   	}),
//   	Vpc: Vpc,
//   	StorageType: rds.StorageType_IO1,
//   	Iops: jsii.Number(5000),
//   })
//
//   gp3Instance := rds.NewDatabaseInstance(this, jsii.String("Gp3Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_*Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_30(),
//   	}),
//   	Vpc: Vpc,
//   	AllocatedStorage: jsii.Number(500),
//   	StorageType: rds.StorageType_GP3,
//   	StorageThroughput: jsii.Number(500),
//   })
//
type MysqlEngineVersion interface {
	// The full version string, for example, "10.5.28".
	MysqlFullVersion() *string
	// The major version of the engine, for example, "10.5".
	MysqlMajorVersion() *string
}

// The jsii proxy struct for MysqlEngineVersion
type jsiiProxy_MysqlEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_MysqlEngineVersion) MysqlFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlEngineVersion) MysqlMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlMajorVersion",
		&returns,
	)
	return returns
}


// Create a new MysqlEngineVersion with an arbitrary version.
func MysqlEngineVersion_Of(mysqlFullVersion *string, mysqlMajorVersion *string) MysqlEngineVersion {
	_init_.Initialize()

	if err := validateMysqlEngineVersion_OfParameters(mysqlFullVersion, mysqlMajorVersion); err != nil {
		panic(err)
	}
	var returns MysqlEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"of",
		[]interface{}{mysqlFullVersion, mysqlMajorVersion},
		&returns,
	)

	return returns
}

func MysqlEngineVersion_VER_5_5_54() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_5_54",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_16() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_16",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_17() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_17",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_19() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_19",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_21() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_21",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_22() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_22",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_23() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_23",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_24() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_24",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_25() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_25",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_26() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_26",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_28() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_28",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_30() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_30",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_31() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_31",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_33() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_33",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_34() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_34",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_35() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_35",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_36() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_36",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_37() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_37",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_38() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_38",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_39() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_39",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_40() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_40",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_41() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_41",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_42() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_42",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_43() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_43",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_44() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_7_44",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_11() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_11",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_13() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_13",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_15() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_15",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_16() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_16",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_17() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_17",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_19() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_19",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_20() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_20",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_21() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_21",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_23() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_23",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_25() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_25",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_26() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_26",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_27() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_27",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_28() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_28",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_29() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_29",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_30() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_30",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_31() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_31",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_32() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_32",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_33() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_33",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_34() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_34",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_35() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_35",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_36() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_8_0_36",
		&returns,
	)
	return returns
}

