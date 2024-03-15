package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the Aurora MySQL cluster engine (those returned by `DatabaseClusterEngine.auroraMysql`).
//
// https://docs.aws.amazon.com/AmazonRDS/latest/AuroraMySQLReleaseNotes/Welcome.html
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
//   	}),
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
//   	 // Optional - will default to 'admin' username and generated password
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		PubliclyAccessible: jsii.Boolean(false),
//   	}),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_*Provisioned(jsii.String("reader1"), &ProvisionedClusterInstanceProps{
//   			PromotionTier: jsii.Number(1),
//   		}),
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader2")),
//   	},
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	Vpc: Vpc,
//   })
//
type AuroraMysqlEngineVersion interface {
	// The full version string, for example, "5.7.mysql_aurora.1.78.3.6".
	AuroraMysqlFullVersion() *string
	// The major version of the engine.
	//
	// Currently, it's either "5.7", or "8.0".
	AuroraMysqlMajorVersion() *string
}

// The jsii proxy struct for AuroraMysqlEngineVersion
type jsiiProxy_AuroraMysqlEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AuroraMysqlEngineVersion) AuroraMysqlFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraMysqlFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraMysqlEngineVersion) AuroraMysqlMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraMysqlMajorVersion",
		&returns,
	)
	return returns
}


// Create a new AuroraMysqlEngineVersion with an arbitrary version.
func AuroraMysqlEngineVersion_Of(auroraMysqlFullVersion *string, auroraMysqlMajorVersion *string) AuroraMysqlEngineVersion {
	_init_.Initialize()

	if err := validateAuroraMysqlEngineVersion_OfParameters(auroraMysqlFullVersion); err != nil {
		panic(err)
	}
	var returns AuroraMysqlEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"of",
		[]interface{}{auroraMysqlFullVersion, auroraMysqlMajorVersion},
		&returns,
	)

	return returns
}

func AuroraMysqlEngineVersion_VER_2_02_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_02_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_03_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_03_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_03_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_03_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_03_4() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_03_4",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_4() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_4",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_5() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_5",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_6() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_6",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_7() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_7",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_8() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_8",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_9() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_9",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_05_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_05_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_05_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_05_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_06_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_06_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_10() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_10",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_4() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_4",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_5() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_5",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_6() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_6",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_7() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_7",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_8() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_8",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_9() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_9",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_4() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_4",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_10_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_10_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_10_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_10_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_10_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_10_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_10_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_10_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_11_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_11_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_11_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_11_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_11_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_11_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_11_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_11_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_11_4() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_11_4",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_12_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_12_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_12_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_12_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_01_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_01_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_01_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_01_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_02_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_02_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_02_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_02_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_02_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_02_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_02_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_02_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_03_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_03_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_03_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_03_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_03_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_03_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_03_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_03_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_04_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_04_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_04_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_04_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_05_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_05_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_05_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_05_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_05_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_05_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_06_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_06_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_5_7_12() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"VER_5_7_12",
		&returns,
	)
	return returns
}

