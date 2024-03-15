package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the SQL Server instance engines (those returned by `DatabaseInstanceEngine.sqlServerSe`, `DatabaseInstanceEngine.sqlServerEx`, `DatabaseInstanceEngine.sqlServerWeb` and `DatabaseInstanceEngine.sqlServerEe`).
//
// Example:
//   var vpc vpc
//
//
//   parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &ParameterGroupProps{
//   	Engine: rds.DatabaseInstanceEngine_SqlServerEe(&SqlServerEeInstanceEngineProps{
//   		Version: rds.SqlServerEngineVersion_VER_11(),
//   	}),
//   	Parameters: map[string]*string{
//   		"locks": jsii.String("100"),
//   	},
//   })
//
//   rds.NewDatabaseInstance(this, jsii.String("Database"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_SQL_SERVER_EE(),
//   	Vpc: Vpc,
//   	ParameterGroup: ParameterGroup,
//   })
//
type SqlServerEngineVersion interface {
	// The full version string, for example, "15.00.3049.1.v1".
	SqlServerFullVersion() *string
	// The major version of the engine, for example, "15.00".
	SqlServerMajorVersion() *string
}

// The jsii proxy struct for SqlServerEngineVersion
type jsiiProxy_SqlServerEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_SqlServerEngineVersion) SqlServerFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServerFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlServerEngineVersion) SqlServerMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServerMajorVersion",
		&returns,
	)
	return returns
}


// Create a new SqlServerEngineVersion with an arbitrary version.
func SqlServerEngineVersion_Of(sqlServerFullVersion *string, sqlServerMajorVersion *string) SqlServerEngineVersion {
	_init_.Initialize()

	if err := validateSqlServerEngineVersion_OfParameters(sqlServerFullVersion, sqlServerMajorVersion); err != nil {
		panic(err)
	}
	var returns SqlServerEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"of",
		[]interface{}{sqlServerFullVersion, sqlServerMajorVersion},
		&returns,
	)

	return returns
}

func SqlServerEngineVersion_VER_11() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_11",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_11_00_5058_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_11_00_5058_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_11_00_6020_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_11_00_6020_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_11_00_6594_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_11_00_6594_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_11_00_7462_6_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_11_00_7462_6_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_11_00_7493_4_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_11_00_7493_4_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_4422_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_4422_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_5000_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_5000_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_5546_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_5546_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_5571_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_5571_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_6293_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_6293_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_6329_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_6329_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_6433_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_6433_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_6439_10_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_6439_10_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_6444_4_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_6444_4_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_12_00_6449_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_12_00_6449_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_2164_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_2164_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_4422_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_4422_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_4451_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_4451_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_4466_4_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_4466_4_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_4522_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_4522_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5216_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5216_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5292_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5292_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5366_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5366_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5426_0_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5426_0_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5598_27_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5598_27_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5820_21_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5820_21_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5850_14_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5850_14_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_5882_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_5882_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_6300_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_6300_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_6419_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_6419_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_6430_49_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_6430_49_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_13_00_6435_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_13_00_6435_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_1000_169_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_1000_169_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3015_40_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3015_40_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3035_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3035_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3049_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3049_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3192_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3192_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3223_3_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3223_3_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3281_6_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3281_6_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3294_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3294_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3356_20_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3356_20_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3381_3_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3381_3_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3401_7_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3401_7_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3421_10_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3421_10_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3451_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3451_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3460_9_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3460_9_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_14_00_3465_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_14_00_3465_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4043_16_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4043_16_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4073_23_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4073_23_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4153_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4153_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4198_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4198_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4236_7_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4236_7_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4312_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4312_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4316_3_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4316_3_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4322_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4322_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4335_1_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4335_1_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4345_5_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4345_5_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_15_00_4355_3_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_15_00_4355_3_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_16() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_16",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_16_00_4085_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_16_00_4085_2_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_16_00_4095_4_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_16_00_4095_4_V1",
		&returns,
	)
	return returns
}

func SqlServerEngineVersion_VER_16_00_4105_2_V1() SqlServerEngineVersion {
	_init_.Initialize()
	var returns SqlServerEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		"VER_16_00_4105_2_V1",
		&returns,
	)
	return returns
}

