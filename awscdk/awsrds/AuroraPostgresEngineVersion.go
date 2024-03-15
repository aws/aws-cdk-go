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
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
//   		Version: rds.AuroraPostgresEngineVersion_VER_15_2(),
//   	}),
//   	Credentials: rds.Credentials_FromUsername(jsii.String("adminuser"), &CredentialsFromUsernameOptions{
//   		Password: awscdk.SecretValue_UnsafePlainText(jsii.String("7959866cacc02c2d243ecfe177464fe6")),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		PubliclyAccessible: jsii.Boolean(false),
//   	}),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_*Provisioned(jsii.String("reader")),
//   	},
//   	StorageType: rds.DBClusterStorageType_AURORA_IOPT1,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	Vpc: Vpc,
//   })
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

func AuroraPostgresEngineVersion_VER_11_17() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_17",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_18() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_18",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_19() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_19",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_20() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_20",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_11_21() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_11_21",
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

func AuroraPostgresEngineVersion_VER_12_12() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_12",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_13() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_13",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_14() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_14",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_15() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_15",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_16() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_16",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_12_17() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_12_17",
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

func AuroraPostgresEngineVersion_VER_13_10() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_10",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_11() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_11",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_12() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_12",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_13() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_13",
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

func AuroraPostgresEngineVersion_VER_13_8() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_8",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_13_9() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_13_9",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_14_10() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_10",
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

func AuroraPostgresEngineVersion_VER_14_6() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_6",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_14_7() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_7",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_14_8() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_8",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_14_9() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_14_9",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_15_2() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_15_2",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_15_3() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_15_3",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_15_4() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_15_4",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_15_5() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_15_5",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_16_0() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_16_0",
		&returns,
	)
	return returns
}

func AuroraPostgresEngineVersion_VER_16_1() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_16_1",
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

func AuroraPostgresEngineVersion_VER_9_6_22() AuroraPostgresEngineVersion {
	_init_.Initialize()
	var returns AuroraPostgresEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		"VER_9_6_22",
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

