package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the PostgreSQL instance engines (those returned by {@link DatabaseInstanceEngine.postgres}).
//
// Example:
//   var vpc vpc
//
//   engine := rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   	version: rds.postgresEngineVersion_VER_12_3(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstance(this, jsii.String("InstanceWithCustomizedSecret"), &databaseInstanceProps{
//   	engine: engine,
//   	vpc: vpc,
//   	credentials: rds.credentials.fromGeneratedSecret(jsii.String("postgres"), &credentialsBaseOptions{
//   		secretName: jsii.String("my-cool-name"),
//   		encryptionKey: myKey,
//   		excludeCharacters: jsii.String("!&*^#@()"),
//   		replicaRegions: []replicaRegion{
//   			&replicaRegion{
//   				region: jsii.String("eu-west-1"),
//   			},
//   			&replicaRegion{
//   				region: jsii.String("eu-west-2"),
//   			},
//   		},
//   	}),
//   })
//
type PostgresEngineVersion interface {
	// The full version string, for example, "13.11".
	PostgresFullVersion() *string
	// The major version of the engine, for example, "13".
	PostgresMajorVersion() *string
}

// The jsii proxy struct for PostgresEngineVersion
type jsiiProxy_PostgresEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_PostgresEngineVersion) PostgresFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEngineVersion) PostgresMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresMajorVersion",
		&returns,
	)
	return returns
}


// Create a new PostgresEngineVersion with an arbitrary version.
func PostgresEngineVersion_Of(postgresFullVersion *string, postgresMajorVersion *string, postgresFeatures *PostgresEngineFeatures) PostgresEngineVersion {
	_init_.Initialize()

	if err := validatePostgresEngineVersion_OfParameters(postgresFullVersion, postgresMajorVersion, postgresFeatures); err != nil {
		panic(err)
	}
	var returns PostgresEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"of",
		[]interface{}{postgresFullVersion, postgresMajorVersion, postgresFeatures},
		&returns,
	)

	return returns
}

func PostgresEngineVersion_VER_10() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_1() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_1",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_10() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_10",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_11() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_11",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_12() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_12",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_13() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_13",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_14() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_14",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_15() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_15",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_16() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_16",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_17() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_17",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_18() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_18",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_19() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_19",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_20() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_20",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_3() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_3",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_4() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_4",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_5() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_5",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_6() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_6",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_7() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_7",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_10_9() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_10_9",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_1() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_1",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_10() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_10",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_11() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_11",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_12() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_12",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_13() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_13",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_14() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_14",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_15() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_15",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_16() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_16",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_2() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_2",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_4() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_4",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_5() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_5",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_6() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_6",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_7() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_7",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_8() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_8",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_11_9() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_11_9",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_10() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_10",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_11() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_11",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_2() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_2",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_3() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_3",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_4() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_4",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_5() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_5",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_6() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_6",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_7() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_7",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_8() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_8",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_12_9() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_12_9",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13_1() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13_1",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13_2() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13_2",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13_3() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13_3",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13_4() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13_4",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13_5() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13_5",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13_6() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13_6",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_13_7() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_13_7",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_14() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_14",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_14_1() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_14_1",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_14_2() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_14_2",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_14_3() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_14_3",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_14_4() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_14_4",
		&returns,
	)
	return returns
}

func PostgresEngineVersion_VER_9_6_24() PostgresEngineVersion {
	_init_.Initialize()
	var returns PostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		"VER_9_6_24",
		&returns,
	)
	return returns
}

