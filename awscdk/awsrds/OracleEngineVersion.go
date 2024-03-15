package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the Oracle instance engines.
//
// Those returned by the following list.
// - `DatabaseInstanceEngine.oracleSe2`
// - `DatabaseInstanceEngine.oracleSe2Cdb`
// - `DatabaseInstanceEngine.oracleEe`
// - `DatabaseInstanceEngine.oracleEeCdb`.
//
// Example:
//   var vpc vpc
//
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_OracleSe2(&OracleSe2InstanceEngineProps{
//   		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	// optional, defaults to m5.large
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("syscdk")),
//   	 // Optional - will default to 'admin' username and generated password
//   	Vpc: Vpc,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   })
//
type OracleEngineVersion interface {
	// The full version string, for example, "19.0.0.0.ru-2019-10.rur-2019-10.r1".
	OracleFullVersion() *string
	// The major version of the engine, for example, "19".
	OracleMajorVersion() *string
}

// The jsii proxy struct for OracleEngineVersion
type jsiiProxy_OracleEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_OracleEngineVersion) OracleFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleEngineVersion) OracleMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleMajorVersion",
		&returns,
	)
	return returns
}


// Creates a new OracleEngineVersion with an arbitrary version.
func OracleEngineVersion_Of(oracleFullVersion *string, oracleMajorVersion *string) OracleEngineVersion {
	_init_.Initialize()

	if err := validateOracleEngineVersion_OfParameters(oracleFullVersion, oracleMajorVersion); err != nil {
		panic(err)
	}
	var returns OracleEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"of",
		[]interface{}{oracleFullVersion, oracleMajorVersion},
		&returns,
	)

	return returns
}

func OracleEngineVersion_VER_12_1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V10() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V10",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V11() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V11",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V12() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V12",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V13() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V13",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V14() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V14",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V15() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V15",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V16() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V16",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V17() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V17",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V18() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V18",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V19() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V19",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V2() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V2",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V20() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V20",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V21() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V21",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V22() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V22",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V23() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V23",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V24() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V24",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V25() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V25",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V26() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V26",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V27() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V27",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V28() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V28",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V29() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V29",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V3() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V3",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V4() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V4",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V5() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V5",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V6() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V6",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V7() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V7",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V8() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V8",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_1_0_2_V9() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_1_0_2_V9",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2018_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2018_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2019_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2019_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2019_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2019_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2019_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2019_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2019_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2019_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2020_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2020_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2020_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2020_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2020_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2020_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2020_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2020_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2021_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2021_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2021_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2021_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2021_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2021_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2021_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2021_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_12_2_0_1_2022_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_12_2_0_1_2022_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2019_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2019_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2019_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2019_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2020_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2020_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2020_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2020_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2020_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2020_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2020_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2020_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2021_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2021_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_18_0_0_0_2021_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_18_0_0_0_2021_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2019_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2019_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2019_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2019_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2020_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2020_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2020_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2020_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2020_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2020_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2020_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2020_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2021_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2021_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2021_01_R2() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2021_01_R2",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2021_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2021_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2021_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2021_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2021_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2021_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2022_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2022_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2022_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2022_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2022_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2022_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2022_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2022_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2023_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2023_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2023_01_R2() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2023_01_R2",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2023_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2023_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2023_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2023_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2023_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2023_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_19_0_0_0_2024_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_19_0_0_0_2024_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2022_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2022_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2022_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2022_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2022_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2022_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2022_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2022_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2023_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2023_01_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2023_01_R2() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2023_01_R2",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2023_04_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2023_04_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2023_07_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2023_07_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2023_10_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2023_10_R1",
		&returns,
	)
	return returns
}

func OracleEngineVersion_VER_21_0_0_0_2024_01_R1() OracleEngineVersion {
	_init_.Initialize()
	var returns OracleEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		"VER_21_0_0_0_2024_01_R1",
		&returns,
	)
	return returns
}

