package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the Aurora cluster engine (those returned by `DatabaseClusterEngine.aurora`).
//
// Example:
//   var vpc vpc
//
//   rds.NewDatabaseClusterFromSnapshot(this, jsii.String("Database"), &DatabaseClusterFromSnapshotProps{
//   	Engine: rds.DatabaseClusterEngine_Aurora(&AuroraClusterEngineProps{
//   		Version: rds.AuroraEngineVersion_VER_1_22_2(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
//   	Vpc: Vpc,
//   	SnapshotIdentifier: jsii.String("mySnapshot"),
//   })
//
type AuroraEngineVersion interface {
	// The full version string, for example, "5.6.mysql_aurora.1.78.3.6".
	AuroraFullVersion() *string
	// The major version of the engine.
	//
	// Currently, it's always "5.6".
	AuroraMajorVersion() *string
}

// The jsii proxy struct for AuroraEngineVersion
type jsiiProxy_AuroraEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AuroraEngineVersion) AuroraFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraEngineVersion) AuroraMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraMajorVersion",
		&returns,
	)
	return returns
}


// Create a new AuroraEngineVersion with an arbitrary version.
func AuroraEngineVersion_Of(auroraFullVersion *string, auroraMajorVersion *string) AuroraEngineVersion {
	_init_.Initialize()

	if err := validateAuroraEngineVersion_OfParameters(auroraFullVersion); err != nil {
		panic(err)
	}
	var returns AuroraEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"of",
		[]interface{}{auroraFullVersion, auroraMajorVersion},
		&returns,
	)

	return returns
}

func AuroraEngineVersion_VER_1_17_9() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_17_9",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_19_0() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_19_0",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_19_1() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_19_1",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_19_2() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_19_2",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_19_5() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_19_5",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_19_6() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_19_6",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_20_0() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_20_0",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_20_1() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_20_1",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_21_0() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_21_0",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_22_0() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_22_0",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_22_1() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_22_1",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_22_1_3() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_22_1_3",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_22_2() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_22_2",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_22_3() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_22_3",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_22_4() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_22_4",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_22_5() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_22_5",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_23_0() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_23_0",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_23_1() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_23_1",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_23_2() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_23_2",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_23_3() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_23_3",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_1_23_4() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_1_23_4",
		&returns,
	)
	return returns
}

func AuroraEngineVersion_VER_10A() AuroraEngineVersion {
	_init_.Initialize()
	var returns AuroraEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		"VER_10A",
		&returns,
	)
	return returns
}

