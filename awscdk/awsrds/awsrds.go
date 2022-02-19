package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Aurora capacity units (ACUs).
//
// Each ACU is a combination of processing and memory capacity.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.architecture
//
type AuroraCapacityUnit string

const (
	AuroraCapacityUnit_ACU_1 AuroraCapacityUnit = "ACU_1"
	AuroraCapacityUnit_ACU_2 AuroraCapacityUnit = "ACU_2"
	AuroraCapacityUnit_ACU_4 AuroraCapacityUnit = "ACU_4"
	AuroraCapacityUnit_ACU_8 AuroraCapacityUnit = "ACU_8"
	AuroraCapacityUnit_ACU_16 AuroraCapacityUnit = "ACU_16"
	AuroraCapacityUnit_ACU_32 AuroraCapacityUnit = "ACU_32"
	AuroraCapacityUnit_ACU_64 AuroraCapacityUnit = "ACU_64"
	AuroraCapacityUnit_ACU_128 AuroraCapacityUnit = "ACU_128"
	AuroraCapacityUnit_ACU_192 AuroraCapacityUnit = "ACU_192"
	AuroraCapacityUnit_ACU_256 AuroraCapacityUnit = "ACU_256"
	AuroraCapacityUnit_ACU_384 AuroraCapacityUnit = "ACU_384"
)

// Creation properties of the plain Aurora database cluster engine.
//
// Used in {@link DatabaseClusterEngine.aurora}.
//
// TODO: EXAMPLE
//
type AuroraClusterEngineProps struct {
	// The version of the Aurora cluster engine.
	Version AuroraEngineVersion `json:"version" yaml:"version"`
}

// The versions for the Aurora cluster engine (those returned by {@link DatabaseClusterEngine.aurora}).
//
// TODO: EXAMPLE
//
type AuroraEngineVersion interface {
	AuroraFullVersion() *string
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

// Creation properties of the Aurora MySQL database cluster engine.
//
// Used in {@link DatabaseClusterEngine.auroraMysql}.
//
// TODO: EXAMPLE
//
type AuroraMysqlClusterEngineProps struct {
	// The version of the Aurora MySQL cluster engine.
	Version AuroraMysqlEngineVersion `json:"version" yaml:"version"`
}

// The versions for the Aurora MySQL cluster engine (those returned by {@link DatabaseClusterEngine.auroraMysql}).
//
// TODO: EXAMPLE
//
type AuroraMysqlEngineVersion interface {
	AuroraMysqlFullVersion() *string
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

	var returns AuroraMysqlEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		"of",
		[]interface{}{auroraMysqlFullVersion, auroraMysqlMajorVersion},
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

// Creation properties of the Aurora PostgreSQL database cluster engine.
//
// Used in {@link DatabaseClusterEngine.auroraPostgres}.
//
// TODO: EXAMPLE
//
type AuroraPostgresClusterEngineProps struct {
	// The version of the Aurora PostgreSQL cluster engine.
	Version AuroraPostgresEngineVersion `json:"version" yaml:"version"`
}

// Features supported by this version of the Aurora Postgres cluster engine.
//
// TODO: EXAMPLE
//
type AuroraPostgresEngineFeatures struct {
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data export feature.
	S3Export *bool `json:"s3Export" yaml:"s3Export"`
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data import feature.
	S3Import *bool `json:"s3Import" yaml:"s3Import"`
}

// The versions for the Aurora PostgreSQL cluster engine (those returned by {@link DatabaseClusterEngine.auroraPostgres}).
//
// TODO: EXAMPLE
//
type AuroraPostgresEngineVersion interface {
	AuroraPostgresFullVersion() *string
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

// Backup configuration for RDS databases.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
//
type BackupProps struct {
	// How many days to retain the backup.
	Retention awscdk.Duration `json:"retention" yaml:"retention"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'
	PreferredWindow *string `json:"preferredWindow" yaml:"preferredWindow"`
}

// A CloudFormation `AWS::RDS::DBCluster`.
//
// The `AWS::RDS::DBCluster` resource creates an Amazon Aurora DB cluster. For more information, see [Managing an Amazon Aurora DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Aurora.html) in the *Amazon Aurora User Guide* .
//
// > You can only create this resource in AWS Regions where Amazon Aurora is supported.
//
// This topic covers the resource for Amazon Aurora DB clusters. For the documentation on the resource for Amazon RDS DB instances, see [AWS::RDS::DBInstance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html) .
//
// *Updating DB clusters*
//
// When properties labeled " *Update requires:* [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) " are updated, AWS CloudFormation first creates a replacement DB cluster, then changes references from other dependent resources to point to the replacement DB cluster, and finally deletes the old DB cluster.
//
// > We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB cluster. To preserve your data, perform the following procedure:
// >
// > - Deactivate any applications that are using the DB cluster so that there's no activity on the DB instance.
// > - Create a snapshot of the DB cluster. For more information about creating DB snapshots, see [Creating a DB Cluster Snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CreateSnapshotCluster.html) .
// > - If you want to restore your DB cluster using a DB cluster snapshot, modify the updated template with your DB cluster changes and add the `SnapshotIdentifier` property with the ID of the DB cluster snapshot that you want to use.
// >
// > After you restore a DB cluster with a `SnapshotIdentifier` property, you must specify the same `SnapshotIdentifier` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the DB cluster snapshot again, and the data in the database is not changed. However, if you don't specify the `SnapshotIdentifier` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified `SnapshotIdentifier` property, and the original DB cluster is deleted.
// > - Update the stack.
//
// Currently, when you are updating the stack for an Aurora Serverless DB cluster, you can't include changes to any other properties when you specify one of the following properties: `PreferredBackupWindow` , `PreferredMaintenanceWindow` , and `Port` . This limitation doesn't apply to provisioned DB clusters.
//
// For more information about updating other properties of this resource, see `[ModifyDBCluster](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_ModifyDBCluster.html)` . For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// *Deleting DB clusters*
//
// The default `DeletionPolicy` for `AWS::RDS::DBCluster` resources is `Snapshot` . For more information about how AWS CloudFormation deletes resources, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// TODO: EXAMPLE
//
type CfnDBCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AssociatedRoles() interface{}
	SetAssociatedRoles(val interface{})
	AttrEndpointAddress() *string
	AttrEndpointPort() *string
	AttrReadEndpointAddress() *string
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	BacktrackWindow() *float64
	SetBacktrackWindow(val *float64)
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbClusterParameterGroupName() *string
	SetDbClusterParameterGroupName(val *string)
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	EnableHttpEndpoint() interface{}
	SetEnableHttpEndpoint(val interface{})
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	Engine() *string
	SetEngine(val *string)
	EngineMode() *string
	SetEngineMode(val *string)
	EngineVersion() *string
	SetEngineVersion(val *string)
	GlobalClusterIdentifier() *string
	SetGlobalClusterIdentifier(val *string)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	Ref() *string
	ReplicationSourceIdentifier() *string
	SetReplicationSourceIdentifier(val *string)
	RestoreType() *string
	SetRestoreType(val *string)
	ScalingConfiguration() interface{}
	SetScalingConfiguration(val interface{})
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	SourceRegion() *string
	SetSourceRegion(val *string)
	Stack() awscdk.Stack
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBCluster
type jsiiProxy_CfnDBCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBCluster) AssociatedRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrReadEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReadEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) BacktrackWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EnableHttpEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EnableIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EngineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) ReplicationSourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) RestoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) ScalingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBCluster`.
func NewCfnDBCluster(scope constructs.Construct, id *string, props *CfnDBClusterProps) CfnDBCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnDBCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBCluster`.
func NewCfnDBCluster_Override(c CfnDBCluster, scope constructs.Construct, id *string, props *CfnDBClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetAssociatedRoles(val interface{}) {
	_jsii_.Set(
		j,
		"associatedRoles",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetBacktrackWindow(val *float64) {
	_jsii_.Set(
		j,
		"backtrackWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetBackupRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetCopyTagsToSnapshot(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDbClusterParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEnableCloudwatchLogsExports(val *[]*string) {
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEnableHttpEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"enableHttpEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEnableIamDatabaseAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"enableIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEngineMode(val *string) {
	_jsii_.Set(
		j,
		"engineMode",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetGlobalClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetReplicationSourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"replicationSourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetRestoreType(val *string) {
	_jsii_.Set(
		j,
		"restoreType",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetScalingConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"scalingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetSourceDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetSourceRegion(val *string) {
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetStorageEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetUseLatestRestorableTime(val interface{}) {
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetVpcSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBCluster) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBCluster) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBCluster) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
//
// TODO: EXAMPLE
//
type CfnDBCluster_DBClusterRoleProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf. For the list of supported feature names, see the `SupportedFeatureNames` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference* .
	FeatureName *string `json:"featureName" yaml:"featureName"`
}

// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
//
// For more information, see [Using Amazon Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html) in the *Amazon Aurora User Guide* .
//
// TODO: EXAMPLE
//
type CfnDBCluster_ScalingConfigurationProperty struct {
	// A value that indicates whether to allow or disallow automatic pause for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// A DB cluster can be paused only when it's idle (it has no connections).
	//
	// > If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it.
	AutoPause interface{} `json:"autoPause" yaml:"autoPause"`
	// The maximum capacity for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are `1` , `2` , `4` , `8` , `16` , `32` , `64` , `128` , and `256` .
	//
	// For Aurora PostgreSQL, valid capacity values are `2` , `4` , `8` , `16` , `32` , `64` , `192` , and `384` .
	//
	// The maximum capacity must be greater than or equal to the minimum capacity.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are `1` , `2` , `4` , `8` , `16` , `32` , `64` , `128` , and `256` .
	//
	// For Aurora PostgreSQL, valid capacity values are `2` , `4` , `8` , `16` , `32` , `64` , `192` , and `384` .
	//
	// The minimum capacity must be less than or equal to the maximum capacity.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// The time, in seconds, before an Aurora DB cluster in `serverless` mode is paused.
	//
	// Specify a value between 300 and 86,400 seconds.
	SecondsUntilAutoPause *float64 `json:"secondsUntilAutoPause" yaml:"secondsUntilAutoPause"`
}

// A CloudFormation `AWS::RDS::DBClusterParameterGroup`.
//
// The `AWS::RDS::DBClusterParameterGroup` resource creates a new Amazon RDS DB cluster parameter group.
//
// For information about configuring parameters for Amazon Aurora DB instances, see [Working with DB parameter groups and DB cluster parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide* .
//
// > If you apply a parameter group to a DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
// >
// > If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.
//
// TODO: EXAMPLE
//
type CfnDBClusterParameterGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	Family() *string
	SetFamily(val *string)
	LogicalId() *string
	Node() constructs.Node
	Parameters() interface{}
	SetParameters(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBClusterParameterGroup
type jsiiProxy_CfnDBClusterParameterGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBClusterParameterGroup`.
func NewCfnDBClusterParameterGroup(scope constructs.Construct, id *string, props *CfnDBClusterParameterGroupProps) CfnDBClusterParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDBClusterParameterGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBClusterParameterGroup`.
func NewCfnDBClusterParameterGroup_Override(c CfnDBClusterParameterGroup, scope constructs.Construct, id *string, props *CfnDBClusterParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) SetFamily(val *string) {
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBClusterParameterGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBClusterParameterGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBClusterParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBClusterParameterGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBClusterParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBClusterParameterGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBClusterParameterGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBClusterParameterGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBClusterParameterGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBClusterParameterGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBClusterParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDBClusterParameterGroup`.
//
// TODO: EXAMPLE
//
type CfnDBClusterParameterGroupProps struct {
	// A friendly description for this DB cluster parameter group.
	Description *string `json:"description" yaml:"description"`
	// The DB cluster parameter group family name.
	//
	// A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
	//
	// > The DB cluster parameter group family can't be changed when updating a DB cluster parameter group.
	//
	// To list all of the available parameter group families, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"`
	//
	// The output contains duplicates.
	//
	// For more information, see `[CreateDBClusterParameterGroup](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_CreateDBClusterParameterGroup.html)` .
	Family *string `json:"family" yaml:"family"`
	// Provides a list of parameters for the DB cluster parameter group.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// Tags to assign to the DB cluster parameter group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnDBCluster`.
//
// TODO: EXAMPLE
//
type CfnDBClusterProps struct {
	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values: `aurora` (for MySQL 5.6-compatible Aurora), `aurora-mysql` (for MySQL 5.7-compatible Aurora), and `aurora-postgresql`
	Engine *string `json:"engine" yaml:"engine"`
	// Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf.
	AssociatedRoles interface{} `json:"associatedRoles" yaml:"associatedRoles"`
	// A list of Availability Zones (AZs) where instances in the DB cluster can be created.
	//
	// For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide* .
	AvailabilityZones *[]*string `json:"availabilityZones" yaml:"availabilityZones"`
	// The target backtrack window, in seconds. To disable backtracking, set this value to 0.
	//
	// > Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	//
	// Default: 0
	//
	// Constraints:
	//
	// - If specified, this value must be set to a number from 0 to 259,200 (72 hours).
	BacktrackWindow *float64 `json:"backtrackWindow" yaml:"backtrackWindow"`
	// The number of days for which automated backups are retained.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
	//
	// The default is not to copy them.
	CopyTagsToSnapshot interface{} `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The name of your database.
	//
	// If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide* .
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `my-cluster1`
	DbClusterIdentifier *string `json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The name of the DB cluster parameter group to associate with this DB cluster.
	//
	// > If you apply a parameter group to an existing DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
	// >
	// > If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.
	//
	// To list all of the available DB cluster parameter group names, use the following command:
	//
	// `aws rds describe-db-cluster-parameter-groups --query "DBClusterParameterGroups[].DBClusterParameterGroupName" --output text`
	DbClusterParameterGroupName *string `json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// A DB subnet group that you want to associate with this DB cluster.
	//
	// If you are restoring a DB cluster to a point in time with `RestoreType` set to `copy-on-write` , and don't specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.
	DbSubnetGroupName *string `json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// A value that indicates whether the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection interface{} `json:"deletionProtection" yaml:"deletionProtection"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Aurora User Guide* .
	//
	// *Aurora MySQL*
	//
	// Valid values: `audit` , `error` , `general` , `slowquery`
	//
	// *Aurora PostgreSQL*
	//
	// Valid values: `postgresql`
	EnableCloudwatchLogsExports *[]*string `json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// A value that indicates whether to enable the HTTP endpoint for an Aurora Serverless DB cluster.
	//
	// By default, the HTTP endpoint is disabled.
	//
	// When enabled, the HTTP endpoint provides a connectionless web service API for running SQL queries on the Aurora Serverless DB cluster. You can also query your database from inside the RDS console with the query editor.
	//
	// For more information, see [Using the Data API for Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the *Amazon Aurora User Guide* .
	EnableHttpEndpoint interface{} `json:"enableHttpEndpoint" yaml:"enableHttpEndpoint"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// For more information, see [IAM Database Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon Aurora User Guide.*
	EnableIamDatabaseAuthentication interface{} `json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// The DB engine mode of the DB cluster, either `provisioned` , `serverless` , `parallelquery` , `global` , or `multimaster` .
	//
	// The `parallelquery` engine mode isn't required for Aurora MySQL version 1.23 and higher 1.x versions, and version 2.09 and higher 2.x versions.
	//
	// The `global` engine mode isn't required for Aurora MySQL version 1.22 and higher 1.x versions, and `global` engine mode isn't required for any 2.x versions.
	//
	// The `multimaster` engine mode only applies for DB clusters created with Aurora MySQL version 5.6.10a.
	//
	// For Aurora PostgreSQL, the `global` engine mode isn't required, and both the `parallelquery` and the `multimaster` engine modes currently aren't supported.
	//
	// Limitations and requirements apply to some DB engine modes. For more information, see the following sections in the *Amazon Aurora User Guide* :
	//
	// - [Limitations of Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)
	// - [Limitations of Parallel Query](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)
	// - [Limitations of Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)
	// - [Limitations of Multi-Master Clusters](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-multi-master.html#aurora-multi-master-limitations)
	EngineMode *string `json:"engineMode" yaml:"engineMode"`
	// The version number of the database engine to use.
	//
	// To list all of the available engine versions for `aurora` (for MySQL 5.6-compatible Aurora), use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for `aurora-mysql` (for MySQL 5.7-compatible Aurora), use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-mysql --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for `aurora-postgresql` , use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-postgresql --query "DBEngineVersions[].EngineVersion"`
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster.
	//
	// To define the primary database cluster of the global cluster, use the [AWS::RDS::GlobalCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html) resource.
	//
	// If you aren't configuring a global database cluster, don't specify this property.
	//
	// > To remove the DB cluster from a global database cluster, specify an empty value for the `GlobalClusterIdentifier` property.
	//
	// For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) in the *Amazon Aurora User Guide* .
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	//
	// If you enable the `StorageEncrypted` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the `StorageEncrypted` property to `true` .
	//
	// If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property value is inherited from the snapshot, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the master user for the DB cluster.
	//
	// > If you specify the `SourceDBClusterIdentifier` or `SnapshotIdentifier` property, don't specify this property. The value is inherited from the source DB instance or snapshot.
	MasterUsername *string `json:"masterUsername" yaml:"masterUsername"`
	// The master password for the DB instance.
	//
	// > If you specify the `SourceDBClusterIdentifier` or `SnapshotIdentifier` property, don't specify this property. The value is inherited from the source DB instance or snapshot.
	MasterUserPassword *string `json:"masterUserPassword" yaml:"masterUserPassword"`
	// The port number on which the DB instances in the DB cluster accept connections.
	//
	// Default:
	//
	// - When `EngineMode` is `provisioned` , `3306` (for both Aurora MySQL and Aurora PostgreSQL)
	// - When `EngineMode` is `serverless` :
	//
	// - `3306` when `Engine` is `aurora` or `aurora-mysql`
	// - `5432` when `Engine` is `aurora-postgresql`
	//
	// > The `No interruption` on update behavior only applies to DB clusters. If you are updating a DB instance, see [Port](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-port) for the AWS::RDS::DBInstance resource.
	Port *float64 `json:"port" yaml:"port"`
	// The daily time range during which automated backups are created.
	//
	// For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the *Amazon Aurora User Guide.*
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Adjusting the Preferred DB Cluster Maintenance Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora) in the *Amazon Aurora User Guide.*
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// The type of restore to be performed. You can specify one of the following values:.
	//
	// - `full-copy` - The new DB cluster is restored as a full copy of the source DB cluster.
	// - `copy-on-write` - The new DB cluster is restored as a clone of the source DB cluster.
	//
	// Constraints: You can't specify `copy-on-write` if the engine version of the source DB cluster is earlier than 1.11.
	//
	// If you don't specify a `RestoreType` value, then the new DB cluster is restored as a full copy of the source DB cluster.
	RestoreType *string `json:"restoreType" yaml:"restoreType"`
	// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
	ScalingConfiguration interface{} `json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// The identifier for the DB snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
	//
	// After you restore a DB cluster with a `SnapshotIdentifier` property, you must specify the same `SnapshotIdentifier` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed. However, if you don't specify the `SnapshotIdentifier` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified `SnapshotIdentifier` property, and the original DB cluster is deleted.
	//
	// If you specify the `SnapshotIdentifier` property to restore a DB cluster (as opposed to specifying it for DB cluster updates), then don't specify the following properties:
	//
	// - `GlobalClusterIdentifier`
	// - `MasterUsername`
	// - `ReplicationSourceIdentifier`
	// - `RestoreType`
	// - `SourceDBClusterIdentifier`
	// - `SourceRegion`
	// - `StorageEncrypted`
	// - `UseLatestRestorableTime`
	//
	// Constraints:
	//
	// - Must match the identifier of an existing Snapshot.
	SnapshotIdentifier *string `json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing DBCluster.
	SourceDbClusterIdentifier *string `json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The AWS Region which contains the source DB cluster when replicating a DB cluster.
	//
	// For example, `us-east-1` .
	SourceRegion *string `json:"sourceRegion" yaml:"sourceRegion"`
	// Indicates whether the DB cluster is encrypted.
	//
	// If you specify the `KmsKeyId` property, then you must enable encryption.
	//
	// If you specify the `SnapshotIdentifier` or `SourceDBClusterIdentifier` property, don't specify this property. The value is inherited from the snapshot or source DB cluster, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	StorageEncrypted interface{} `json:"storageEncrypted" yaml:"storageEncrypted"`
	// Tags to assign to the DB cluster.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// A value that indicates whether to restore the DB cluster to the latest restorable backup time.
	//
	// By default, the DB cluster is not restored to the latest restorable backup time.
	UseLatestRestorableTime interface{} `json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of EC2 VPC security groups to associate with this DB cluster.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::RDS::DBInstance`.
//
// The `AWS::RDS::DBInstance` resource creates an Amazon RDS DB instance.
//
// If you import an existing DB instance, and the template configuration doesn't match the actual configuration of the DB instance, AWS CloudFormation applies the changes in the template during the import operation.
//
// > If a DB instance is deleted or replaced during an update, AWS CloudFormation deletes all automated snapshots. However, it retains manual DB snapshots. During an update that requires replacement, you can apply a stack policy to prevent DB instances from being replaced. For more information, see [Prevent Updates to Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html) .
//
// This topic covers the resource for Amazon RDS DB instances. For the documentation on the resource for Amazon Aurora DB clusters, see [AWS::RDS::DBCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html) .
//
// *Updating DB instances*
//
// When properties labeled " *Update requires:* [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) " are updated, AWS CloudFormation first creates a replacement DB instance, then changes references from other dependent resources to point to the replacement DB instance, and finally deletes the old DB instance.
//
// > We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB instance. To preserve your data, perform the following procedure:
// >
// > - Deactivate any applications that are using the DB instance so that there's no activity on the DB instance.
// > - Create a snapshot of the DB instance. For more information about creating DB snapshots, see [Creating a DB Snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html) .
// > - If you want to restore your instance using a DB snapshot, modify the updated template with your DB instance changes and add the `DBSnapshotIdentifier` property with the ID of the DB snapshot that you want to use.
// >
// > After you restore a DB instance with a `DBSnapshotIdentifier` property, you must specify the same `DBSnapshotIdentifier` property for any future updates to the DB instance. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the `DBSnapshotIdentifier` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified `DBSnapshotIdentifier` property, and the original DB instance is deleted.
// > - Update the stack.
//
// For more information about updating other properties of this resource, see `[ModifyDBInstance](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_ModifyDBInstance.html)` . For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// *Deleting DB instances*
//
// For DB instances that are part of an Aurora DB cluster, you can set a deletion policy for your DB instance to control how AWS CloudFormation handles the DB instance when the stack is deleted. For Amazon RDS DB instances, you can choose to *retain* the DB instance, to *delete* the DB instance, or to *create a snapshot* of the DB instance. The default AWS CloudFormation behavior depends on the `DBClusterIdentifier` property:
//
// - For `AWS::RDS::DBInstance` resources that don't specify the `DBClusterIdentifier` property, AWS CloudFormation saves a snapshot of the DB instance.
// - For `AWS::RDS::DBInstance` resources that do specify the `DBClusterIdentifier` property, AWS CloudFormation deletes the DB instance.
//
// For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// TODO: EXAMPLE
//
type CfnDBInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AllocatedStorage() *string
	SetAllocatedStorage(val *string)
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	AssociatedRoles() interface{}
	SetAssociatedRoles(val interface{})
	AttrEndpointAddress() *string
	AttrEndpointPort() *string
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CharacterSetName() *string
	SetCharacterSetName(val *string)
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	CreationStack() *[]*string
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbInstanceClass() *string
	SetDbInstanceClass(val *string)
	DbInstanceIdentifier() *string
	SetDbInstanceIdentifier(val *string)
	DbName() *string
	SetDbName(val *string)
	DbParameterGroupName() *string
	SetDbParameterGroupName(val *string)
	DbSecurityGroups() *[]*string
	SetDbSecurityGroups(val *[]*string)
	DbSnapshotIdentifier() *string
	SetDbSnapshotIdentifier(val *string)
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DeleteAutomatedBackups() interface{}
	SetDeleteAutomatedBackups(val interface{})
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	Domain() *string
	SetDomain(val *string)
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	EnablePerformanceInsights() interface{}
	SetEnablePerformanceInsights(val interface{})
	Engine() *string
	SetEngine(val *string)
	EngineVersion() *string
	SetEngineVersion(val *string)
	Iops() *float64
	SetIops(val *float64)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LicenseModel() *string
	SetLicenseModel(val *string)
	LogicalId() *string
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MaxAllocatedStorage() *float64
	SetMaxAllocatedStorage(val *float64)
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	MultiAz() interface{}
	SetMultiAz(val interface{})
	Node() constructs.Node
	OptionGroupName() *string
	SetOptionGroupName(val *string)
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	Port() *string
	SetPort(val *string)
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	ProcessorFeatures() interface{}
	SetProcessorFeatures(val interface{})
	PromotionTier() *float64
	SetPromotionTier(val *float64)
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	Ref() *string
	SourceDbInstanceIdentifier() *string
	SetSourceDbInstanceIdentifier(val *string)
	SourceRegion() *string
	SetSourceRegion(val *string)
	Stack() awscdk.Stack
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	StorageType() *string
	SetStorageType(val *string)
	Tags() awscdk.TagManager
	Timezone() *string
	SetTimezone(val *string)
	UpdatedProperites() *map[string]interface{}
	UseDefaultProcessorFeatures() interface{}
	SetUseDefaultProcessorFeatures(val interface{})
	VpcSecurityGroups() *[]*string
	SetVpcSecurityGroups(val *[]*string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBInstance
type jsiiProxy_CfnDBInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBInstance) AllocatedStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AssociatedRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EnableIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EnablePerformanceInsights() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) ProcessorFeatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PromotionTier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) SourceDbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) UseDefaultProcessorFeatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultProcessorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) VpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroups",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBInstance`.
func NewCfnDBInstance(scope constructs.Construct, id *string, props *CfnDBInstanceProps) CfnDBInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnDBInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBInstance`.
func NewCfnDBInstance_Override(c CfnDBInstance, scope constructs.Construct, id *string, props *CfnDBInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetAllocatedStorage(val *string) {
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetAllowMajorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetAssociatedRoles(val interface{}) {
	_jsii_.Set(
		j,
		"associatedRoles",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetAutoMinorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetBackupRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetCaCertificateIdentifier(val *string) {
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetCharacterSetName(val *string) {
	_jsii_.Set(
		j,
		"characterSetName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetCopyTagsToSnapshot(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbInstanceClass(val *string) {
	_jsii_.Set(
		j,
		"dbInstanceClass",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbInstanceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbName(val *string) {
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"dbSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDeleteAutomatedBackups(val interface{}) {
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDomainIamRoleName(val *string) {
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetEnableCloudwatchLogsExports(val *[]*string) {
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetEnableIamDatabaseAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"enableIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetEnablePerformanceInsights(val interface{}) {
	_jsii_.Set(
		j,
		"enablePerformanceInsights",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetIops(val *float64) {
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetLicenseModel(val *string) {
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetMaxAllocatedStorage(val *float64) {
	_jsii_.Set(
		j,
		"maxAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetMonitoringInterval(val *float64) {
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetMonitoringRoleArn(val *string) {
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetMultiAz(val interface{}) {
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetOptionGroupName(val *string) {
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPerformanceInsightsKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPerformanceInsightsRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPort(val *string) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetProcessorFeatures(val interface{}) {
	_jsii_.Set(
		j,
		"processorFeatures",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPromotionTier(val *float64) {
	_jsii_.Set(
		j,
		"promotionTier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetSourceDbInstanceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceDbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetSourceRegion(val *string) {
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetStorageEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetUseDefaultProcessorFeatures(val interface{}) {
	_jsii_.Set(
		j,
		"useDefaultProcessorFeatures",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetVpcSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroups",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBInstance) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBInstance) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBInstance) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBInstance) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBInstance) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB instance.
//
// TODO: EXAMPLE
//
type CfnDBInstance_DBInstanceRoleProperty struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// IAM roles that are associated with a DB instance grant permission for the DB instance to access other AWS services on your behalf. For the list of supported feature names, see the `SupportedFeatureNames` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference* .
	FeatureName *string `json:"featureName" yaml:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// The `ProcessorFeature` property type specifies the processor features of a DB instance class status.
//
// TODO: EXAMPLE
//
type CfnDBInstance_ProcessorFeatureProperty struct {
	// The name of the processor feature.
	//
	// Valid names are `coreCount` and `threadsPerCore` .
	Name *string `json:"name" yaml:"name"`
	// The value of a processor feature name.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnDBInstance`.
//
// TODO: EXAMPLE
//
type CfnDBInstanceProps struct {
	// The compute and memory capacity of the DB instance, for example, `db.m4.large` . Not all DB instance classes are available in all AWS Regions, or for all database engines.
	//
	// For the full list of DB instance classes, and availability for your engine, see [DB Instance Class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide.* For more information about DB instance class pricing and AWS Region support for DB instance classes, see [Amazon RDS Pricing](https://docs.aws.amazon.com/rds/pricing/) .
	DbInstanceClass *string `json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// The amount of storage (in gigabytes) to be initially allocated for the database instance.
	//
	// > If any value is set in the `Iops` parameter, `AllocatedStorage` must be at least 100 GiB, which corresponds to the minimum Iops value of 1,000. If you increase the `Iops` value (in 1,000 IOPS increments), then you must also increase the `AllocatedStorage` value (in 100-GiB increments).
	//
	// *Amazon Aurora*
	//
	// Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.
	//
	// *MySQL*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// *MariaDB*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// *PostgreSQL*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// *Oracle*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 10 to 3072.
	//
	// *SQL Server*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2):
	//
	// - Enterprise and Standard editions: Must be an integer from 20 to 16384.
	// - Web and Express editions: Must be an integer from 20 to 16384.
	// - Provisioned IOPS storage (io1):
	//
	// - Enterprise and Standard editions: Must be an integer from 20 to 16384.
	// - Web and Express editions: Must be an integer from 20 to 16384.
	// - Magnetic storage (standard):
	//
	// - Enterprise and Standard editions: Must be an integer from 20 to 1024.
	// - Web and Express editions: Must be an integer from 20 to 1024.
	AllocatedStorage *string `json:"allocatedStorage" yaml:"allocatedStorage"`
	// A value that indicates whether major version upgrades are allowed.
	//
	// Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.
	//
	// Constraints: Major version upgrades must be allowed when specifying a value for the `EngineVersion` parameter that is a different major version than the DB instance's current version.
	AllowMajorVersionUpgrade interface{} `json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// The AWS Identity and Access Management (IAM) roles associated with the DB instance.
	AssociatedRoles interface{} `json:"associatedRoles" yaml:"associatedRoles"`
	// A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Availability Zone (AZ) where the database will be created.
	//
	// For information on AWS Regions and Availability Zones, see [Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html) .
	//
	// *Amazon Aurora*
	//
	// Not applicable. Availability Zones are managed by the DB cluster.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS Region.
	//
	// Example: `us-east-1d`
	//
	// Constraint: The `AvailabilityZone` parameter can't be specified if the DB instance is a Multi-AZ deployment. The specified Availability Zone must be in the same AWS Region as the current endpoint.
	//
	// > If you're creating a DB instance in an RDS on VMware environment, specify the identifier of the custom Availability Zone to create the DB instance in.
	// >
	// > For more information about RDS on VMware, see the [RDS on VMware User Guide.](https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days for which automated backups are retained.
	//
	// Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The retention period for automated backups is managed by the DB cluster.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 0 to 35
	// - Can't be set to 0 if the DB instance is a source to read replicas
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// The identifier of the CA certificate for this DB instance.
	//
	// > Specifying or updating this property triggers a reboot.
	//
	// For more information about CA certificate identifiers for RDS DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon RDS User Guide* .
	//
	// For more information about CA certificate identifiers for Aurora DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon Aurora User Guide* .
	CaCertificateIdentifier *string `json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// For supported engines, indicates that the DB instance should be associated with the specified character set.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The character set is managed by the DB cluster. For more information, see [AWS::RDS::DBCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html) .
	CharacterSetName *string `json:"characterSetName" yaml:"characterSetName"`
	// A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance.
	//
	// By default, tags are not copied.
	//
	// *Amazon Aurora*
	//
	// Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.
	CopyTagsToSnapshot interface{} `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The identifier of the DB cluster that the instance will belong to.
	DbClusterIdentifier *string `json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DbInstanceIdentifier *string `json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The meaning of this parameter differs according to the database engine you use.
	//
	// > If you specify the `[DBSnapshotIdentifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsnapshotidentifier)` property, AWS CloudFormation ignores this property.
	// >
	// > If you restore DB instances from snapshots, this property doesn't apply to the MySQL, PostgreSQL, or MariaDB engines.
	//
	// *MySQL*
	//
	// The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.
	//
	// Constraints:
	//
	// - Must contain 1 to 64 letters or numbers.
	// - Can't be a word reserved by the specified database engine
	//
	// *MariaDB*
	//
	// The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.
	//
	// Constraints:
	//
	// - Must contain 1 to 64 letters or numbers.
	// - Can't be a word reserved by the specified database engine
	//
	// *PostgreSQL*
	//
	// The name of the database to create when the DB instance is created. If this parameter is not specified, the default `postgres` database is created in the DB instance.
	//
	// Constraints:
	//
	// - Must contain 1 to 63 letters, numbers, or underscores.
	// - Must begin with a letter or an underscore. Subsequent characters can be letters, underscores, or digits (0-9).
	// - Can't be a word reserved by the specified database engine
	//
	// *Oracle*
	//
	// The Oracle System ID (SID) of the created DB instance. If you specify `null` , the default value `ORCL` is used. You can't specify the string NULL, or any other reserved word, for `DBName` .
	//
	// Default: `ORCL`
	//
	// Constraints:
	//
	// - Can't be longer than 8 characters
	//
	// *SQL Server*
	//
	// Not applicable. Must be null.
	//
	// *Amazon Aurora MySQL*
	//
	// The name of the database to create when the primary DB instance of the Aurora MySQL DB cluster is created. If this parameter isn't specified for an Aurora MySQL DB cluster, no database is created in the DB cluster.
	//
	// Constraints:
	//
	// - It must contain 1 to 64 alphanumeric characters.
	// - It can't be a word reserved by the database engine.
	//
	// *Amazon Aurora PostgreSQL*
	//
	// The name of the database to create when the primary DB instance of the Aurora PostgreSQL DB cluster is created. If this parameter isn't specified for an Aurora PostgreSQL DB cluster, a database named `postgres` is created in the DB cluster.
	//
	// Constraints:
	//
	// - It must contain 1 to 63 alphanumeric characters.
	// - It must begin with a letter or an underscore. Subsequent characters can be letters, underscores, or digits (0 to 9).
	// - It can't be a word reserved by the database engine.
	DbName *string `json:"dbName" yaml:"dbName"`
	// The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html) resource created in the template.
	//
	// To list all of the available DB parameter group names, use the following command:
	//
	// `aws rds describe-db-parameter-groups --query "DBParameterGroups[].DBParameterGroupName" --output text`
	//
	// > If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.
	//
	// If you don't specify a value for the `DBParameterGroupName` property, the default DB parameter group for the specified engine and engine version is used.
	DbParameterGroupName *string `json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// A list of the DB security groups to assign to the DB instance.
	//
	// The list can include both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup resources created in the template.
	//
	// If you set DBSecurityGroups, you must not set VPCSecurityGroups, and vice versa. Also, note that the DBSecurityGroups property exists only for backwards compatibility with older regions and is no longer recommended for providing security information to an RDS DB instance. Instead, use VPCSecurityGroups.
	//
	// > If you specify this property, AWS CloudFormation sends only the following properties (if specified) to Amazon RDS during create operations:
	// >
	// > - `AllocatedStorage`
	// > - `AutoMinorVersionUpgrade`
	// > - `AvailabilityZone`
	// > - `BackupRetentionPeriod`
	// > - `CharacterSetName`
	// > - `DBInstanceClass`
	// > - `DBName`
	// > - `DBParameterGroupName`
	// > - `DBSecurityGroups`
	// > - `DBSubnetGroupName`
	// > - `Engine`
	// > - `EngineVersion`
	// > - `Iops`
	// > - `LicenseModel`
	// > - `MasterUsername`
	// > - `MasterUserPassword`
	// > - `MultiAZ`
	// > - `OptionGroupName`
	// > - `PreferredBackupWindow`
	// > - `PreferredMaintenanceWindow`
	// >
	// > All other properties are ignored. Specify a virtual private cloud (VPC) security group if you want to submit other properties, such as `StorageType` , `StorageEncrypted` , or `KmsKeyId` . If you're already using the `DBSecurityGroups` property, you can't use these other properties by updating your DB instance to use a VPC security group. You must recreate the DB instance.
	DbSecurityGroups *[]*string `json:"dbSecurityGroups" yaml:"dbSecurityGroups"`
	// The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance.
	//
	// If you're restoring from a shared manual DB snapshot, you must specify the ARN of the snapshot.
	//
	// By specifying this property, you can create a DB instance from the specified DB snapshot. If the `DBSnapshotIdentifier` property is an empty string or the `AWS::RDS::DBInstance` declaration has no `DBSnapshotIdentifier` property, AWS CloudFormation creates a new database. If the property contains a value (other than an empty string), AWS CloudFormation creates a database from the specified snapshot. If a snapshot with the specified name doesn't exist, AWS CloudFormation can't create the database and it rolls back the stack.
	//
	// Some DB instance properties aren't valid when you restore from a snapshot, such as the `MasterUsername` and `MasterUserPassword` properties. For information about the properties that you can specify, see the `RestoreDBInstanceFromDBSnapshot` action in the *Amazon RDS API Reference* .
	//
	// After you restore a DB instance with a `DBSnapshotIdentifier` property, you must specify the same `DBSnapshotIdentifier` property for any future updates to the DB instance. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the `DBSnapshotIdentifier` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified `DBSnapshotIdentifier` property, and the original DB instance is deleted.
	//
	// If you specify the `DBSnapshotIdentifier` property to restore a DB instance (as opposed to specifying it for DB instance updates), then don't specify the following properties:
	//
	// - `CharacterSetName`
	// - `DBClusterIdentifier`
	// - `DBName`
	// - `DeleteAutomatedBackups`
	// - `EnablePerformanceInsights`
	// - `KmsKeyId`
	// - `MasterUsername`
	// - `MonitoringInterval`
	// - `MonitoringRoleArn`
	// - `PerformanceInsightsKMSKeyId`
	// - `PerformanceInsightsRetentionPeriod`
	// - `PromotionTier`
	// - `SourceDBInstanceIdentifier`
	// - `SourceRegion`
	// - `StorageEncrypted`
	// - `Timezone`
	DbSnapshotIdentifier *string `json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// A DB subnet group to associate with the DB instance.
	//
	// If you update this value, the new subnet group must be a subnet group in a new VPC.
	//
	// If there's no DB subnet group, then the DB instance isn't a VPC DB instance.
	//
	// For more information about using Amazon RDS in a VPC, see [Using Amazon RDS with Amazon Virtual Private Cloud (VPC)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon Relational Database Service Developer Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. The DB subnet group is managed by the DB cluster. If specified, the setting must match the DB cluster setting.
	DbSubnetGroupName *string `json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// A value that indicates whether to remove automated backups immediately after the DB instance is deleted.
	//
	// This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.
	DeleteAutomatedBackups interface{} `json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// A value that indicates whether the DB instance has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled. For more information, see [Deleting a DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html) .
	//
	// *Amazon Aurora*
	//
	// Not applicable. You can enable or disable deletion protection for the DB cluster. For more information, see `CreateDBCluster` . DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.
	DeletionProtection interface{} `json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	//
	// Currently, only Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.
	//
	// For more information, see [Kerberos Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the *Amazon RDS User Guide* .
	Domain *string `json:"domain" yaml:"domain"`
	// Specify the name of the IAM role to be used when making API calls to the Directory Service.
	//
	// This setting doesn't apply to RDS Custom.
	DomainIamRoleName *string `json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Relational Database Service User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. CloudWatch Logs exports are managed by the DB cluster.
	//
	// *MariaDB*
	//
	// Valid values: `audit` , `error` , `general` , `slowquery`
	//
	// *Microsoft SQL Server*
	//
	// Valid values: `agent` , `error`
	//
	// *MySQL*
	//
	// Valid values: `audit` , `error` , `general` , `slowquery`
	//
	// *Oracle*
	//
	// Valid values: `alert` , `audit` , `listener` , `trace`
	//
	// *PostgreSQL*
	//
	// Valid values: `postgresql` , `upgrade`
	EnableCloudwatchLogsExports *[]*string `json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// For more information, see [IAM Database Authentication for MySQL and PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon RDS User Guide.*
	//
	// *Amazon Aurora*
	//
	// Not applicable. Mapping AWS IAM accounts to database accounts is managed by the DB cluster.
	EnableIamDatabaseAuthentication interface{} `json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// A value that indicates whether to enable Performance Insights for the DB instance.
	//
	// For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon Relational Database Service User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	EnablePerformanceInsights interface{} `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The name of the database engine that you want to use for this DB instance.
	//
	// > When you are creating a DB instance, the `Engine` property is required.
	//
	// Valid Values:
	//
	// - `aurora` (for MySQL 5.6-compatible Aurora)
	// - `aurora-mysql` (for MySQL 5.7-compatible Aurora)
	// - `aurora-postgresql`
	// - `mariadb`
	// - `mysql`
	// - `oracle-ee`
	// - `oracle-se2`
	// - `oracle-se1`
	// - `oracle-se`
	// - `postgres`
	// - `sqlserver-ee`
	// - `sqlserver-se`
	// - `sqlserver-ex`
	// - `sqlserver-web`
	Engine *string `json:"engine" yaml:"engine"`
	// The version number of the database engine to use.
	//
	// For a list of valid engine versions, use the `DescribeDBEngineVersions` action.
	//
	// The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every AWS Region.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.
	//
	// *MariaDB*
	//
	// See [MariaDB on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the *Amazon RDS User Guide.*
	//
	// *Microsoft SQL Server*
	//
	// See [Microsoft SQL Server Versions on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the *Amazon RDS User Guide.*
	//
	// *MySQL*
	//
	// See [MySQL on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the *Amazon RDS User Guide.*
	//
	// *Oracle*
	//
	// See [Oracle Database Engine Release Notes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the *Amazon RDS User Guide.*
	//
	// *PostgreSQL*
	//
	// See [Supported PostgreSQL Database Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts.General.DBVersions) in the *Amazon RDS User Guide.*
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	//
	// If you specify this property, you must follow the range of allowed ratios of your requested IOPS rate to the amount of storage that you allocate (IOPS to allocated storage). For example, you can provision an Oracle database instance with 1000 IOPS and 200 GiB of storage (a ratio of 5:1), or specify 2000 IOPS with 200 GiB of storage (a ratio of 10:1). For more information, see [Amazon RDS Provisioned IOPS Storage to Improve Performance](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide* .
	//
	// > If you specify `io1` for the `StorageType` property, then you must also specify the `Iops` property.
	Iops *float64 `json:"iops" yaml:"iops"`
	// The ARN of the AWS KMS key that's used to encrypt the DB instance, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	//
	// If you enable the StorageEncrypted property but don't specify this property, AWS CloudFormation uses the default KMS key. If you specify this property, you must set the StorageEncrypted property to true.
	//
	// If you specify the `SourceDBInstanceIdentifier` property, the value is inherited from the source DB instance if the read replica is created in the same region.
	//
	// If you create an encrypted read replica in a different AWS Region, then you must specify a KMS key for the destination AWS Region. KMS encryption keys are specific to the region that they're created in, and you can't use encryption keys from one region in another region.
	//
	// If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property value is inherited from the snapshot, and if the DB instance is encrypted, the specified `KmsKeyId` property is used.
	//
	// If you specify `DBSecurityGroups` , AWS CloudFormation ignores this property. To specify both a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. The KMS key identifier is managed by the DB cluster.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// License model information for this DB instance.
	//
	// Valid values:
	//
	// - Aurora MySQL - `general-public-license`
	// - Aurora PostgreSQL - `postgresql-license`
	// - MariaDB - `general-public-license`
	// - Microsoft SQL Server - `license-included`
	// - MySQL - `general-public-license`
	// - Oracle - `bring-your-own-license` or `license-included`
	// - PostgreSQL - `postgresql-license`
	//
	// > If you've specified `DBSecurityGroups` and then you update the license model, AWS CloudFormation replaces the underlying DB instance. This will incur some interruptions to database availability.
	LicenseModel *string `json:"licenseModel" yaml:"licenseModel"`
	// The master user name for the DB instance.
	//
	// > If you specify the `SourceDBInstanceIdentifier` or `DBSnapshotIdentifier` property, don't specify this property. The value is inherited from the source DB instance or snapshot.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The name for the master user is managed by the DB cluster.
	//
	// *MariaDB*
	//
	// Constraints:
	//
	// - Required for MariaDB.
	// - Must be 1 to 16 letters or numbers.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *Microsoft SQL Server*
	//
	// Constraints:
	//
	// - Required for SQL Server.
	// - Must be 1 to 128 letters or numbers.
	// - The first character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *MySQL*
	//
	// Constraints:
	//
	// - Required for MySQL.
	// - Must be 1 to 16 letters or numbers.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *Oracle*
	//
	// Constraints:
	//
	// - Required for Oracle.
	// - Must be 1 to 30 letters or numbers.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *PostgreSQL*
	//
	// Constraints:
	//
	// - Required for PostgreSQL.
	// - Must be 1 to 63 letters or numbers.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	MasterUsername *string `json:"masterUsername" yaml:"masterUsername"`
	// The password for the master user. The password can include any printable ASCII character except "/", """, or "@".
	//
	// *Amazon Aurora*
	//
	// Not applicable. The password for the master user is managed by the DB cluster.
	//
	// *MariaDB*
	//
	// Constraints: Must contain from 8 to 41 characters.
	//
	// *Microsoft SQL Server*
	//
	// Constraints: Must contain from 8 to 128 characters.
	//
	// *MySQL*
	//
	// Constraints: Must contain from 8 to 41 characters.
	//
	// *Oracle*
	//
	// Constraints: Must contain from 8 to 30 characters.
	//
	// *PostgreSQL*
	//
	// Constraints: Must contain from 8 to 128 characters.
	MasterUserPassword *string `json:"masterUserPassword" yaml:"masterUserPassword"`
	// The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.
	//
	// For more information about this setting, including limitations that apply to it, see [Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the *Amazon RDS User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	//
	// To disable collection of Enhanced Monitoring metrics, specify 0. The default is 0.
	//
	// If `MonitoringRoleArn` is specified, then you must set `MonitoringInterval` to a value other than 0.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Valid Values: `0, 1, 5, 10, 15, 30, 60`
	MonitoringInterval *float64 `json:"monitoringInterval" yaml:"monitoringInterval"`
	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.
	//
	// For example, `arn:aws:iam:123456789012:role/emaccess` . For information on creating a monitoring role, see [Setting Up and Enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide* .
	//
	// If `MonitoringInterval` is set to a value other than 0, then you must supply a `MonitoringRoleArn` value.
	//
	// This setting doesn't apply to RDS Custom.
	MonitoringRoleArn *string `json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// Specifies whether the database instance is a multiple Availability Zone deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the `MultiAZ` parameter is set to true.
	//
	// *Amazon Aurora*
	//
	// Not applicable. Amazon Aurora storage is replicated across all of the Availability Zones and doesn't require the `MultiAZ` option to be set.
	MultiAz interface{} `json:"multiAz" yaml:"multiAz"`
	// Indicates that the DB instance should be associated with the specified option group.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance once it is associated with a DB instance.
	OptionGroupName *string `json:"optionGroupName" yaml:"optionGroupName"`
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	//
	// If you do not specify a value for `PerformanceInsightsKMSKeyId` , then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account. Your AWS account has a different default KMS key for each AWS Region.
	//
	// For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights) .
	PerformanceInsightsKmsKeyId *string `json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// The amount of time, in days, to retain Performance Insights data. Valid values are 7 or 731 (2 years).
	//
	// For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights) .
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// The port number on which the database accepts connections.
	Port *string `json:"port" yaml:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled, using the `BackupRetentionPeriod` parameter.
	//
	// For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the *Amazon RDS User Guide.*
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The daily time range for creating automated backups is managed by the DB cluster.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Adjusting the Preferred DB Instance Maintenance Window](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow) in the *Amazon RDS User Guide.*
	//
	// > This property applies when AWS CloudFormation initially creates the DB instance. If you use AWS CloudFormation to update the DB instance, those updates are applied immediately.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
	//
	// This setting doesn't apply to RDS Custom.
	ProcessorFeatures interface{} `json:"processorFeatures" yaml:"processorFeatures"`
	// A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.
	//
	// For more information, see [Fault Tolerance for an Aurora DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the *Amazon Aurora User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Default: 1
	//
	// Valid Values: 0 - 15
	PromotionTier *float64 `json:"promotionTier" yaml:"promotionTier"`
	// Indicates whether the DB instance is an internet-facing instance.
	//
	// If you specify `true` , AWS CloudFormation creates an instance with a publicly resolvable DNS name, which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address.
	//
	// The default behavior value depends on your VPC setup and the database subnet group. For more information, see the `PubliclyAccessible` parameter in [`CreateDBInstance`](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) in the *Amazon RDS API Reference* .
	//
	// If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the *DependsOn* attribute to declare a dependency on the VPC-gateway attachment. For more information, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	//
	// > If you specify DBSecurityGroups, AWS CloudFormation ignores this property. To specify a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon RDS User Guide* .
	PubliclyAccessible interface{} `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// If you want to create a read replica DB instance, specify the ID of the source DB instance.
	//
	// Each DB instance can have a limited number of read replicas. For more information, see [Working with Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html) in the *Amazon Relational Database Service Developer Guide* .
	//
	// The `SourceDBInstanceIdentifier` property determines whether a DB instance is a read replica. If you remove the `SourceDBInstanceIdentifier` property from your template and then update your stack, AWS CloudFormation deletes the Read Replica and creates a new DB instance (not a read replica).
	//
	// > - If you specify a source DB instance that uses VPC security groups, we recommend that you specify the `VPCSecurityGroups` property. If you don't specify the property, the read replica inherits the value of the `VPCSecurityGroups` property from the source DB when you create the replica. However, if you update the stack, AWS CloudFormation reverts the replica's `VPCSecurityGroups` property to the default value because it's not defined in the stack's template. This change might cause unexpected issues.
	// > - Read replicas don't support deletion policies. AWS CloudFormation ignores any deletion policy that's associated with a read replica.
	// > - If you specify `SourceDBInstanceIdentifier` , don't specify the `DBSnapshotIdentifier` property. You can't create a read replica from a snapshot.
	// > - Don't set the `BackupRetentionPeriod` , `DBName` , `MasterUsername` , `MasterUserPassword` , and `PreferredBackupWindow` properties. The database attributes are inherited from the source DB instance, and backups are disabled for read replicas.
	// > - If the source DB instance is in a different region than the read replica, specify the source region in `SourceRegion` , and specify an ARN for a valid DB instance in `SourceDBInstanceIdentifier` . For more information, see [Constructing a Amazon RDS Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.ARN) in the *Amazon RDS User Guide* .
	// > - For DB instances in Amazon Aurora clusters, don't specify this property. Amazon RDS automatically assigns writer and reader DB instances.
	SourceDbInstanceIdentifier *string `json:"sourceDbInstanceIdentifier" yaml:"sourceDbInstanceIdentifier"`
	// The ID of the region that contains the source DB instance for the read replica.
	SourceRegion *string `json:"sourceRegion" yaml:"sourceRegion"`
	// A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.
	//
	// If you specify the `KmsKeyId` property, then you must enable encryption.
	//
	// If you specify the `SnapshotIdentifier` or `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the snapshot or source DB instance, and if the DB instance is encrypted, the specified `KmsKeyId` property is used.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The encryption for DB instances is managed by the DB cluster.
	StorageEncrypted interface{} `json:"storageEncrypted" yaml:"storageEncrypted"`
	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: `standard | gp2 | io1`
	//
	// The `standard` value is also known as magnetic.
	//
	// If you specify `io1` , you must also include a value for the `Iops` parameter.
	//
	// Default: `io1` if the `Iops` parameter is specified, otherwise `standard`
	//
	// For more information, see [Amazon RDS DB Instance Storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. Aurora data is stored in the cluster volume, which is a single, virtual volume that uses solid state drives (SSDs).
	StorageType *string `json:"storageType" yaml:"storageType"`
	// Tags to assign to the DB instance.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The time zone of the DB instance.
	//
	// The time zone parameter is currently supported only by [Microsoft SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone) .
	Timezone *string `json:"timezone" yaml:"timezone"`
	// A value that indicates whether the DB instance class of the DB instance uses its default processor features.
	//
	// This setting doesn't apply to RDS Custom.
	UseDefaultProcessorFeatures interface{} `json:"useDefaultProcessorFeatures" yaml:"useDefaultProcessorFeatures"`
	// A list of the VPC security group IDs to assign to the DB instance.
	//
	// The list can include both the physical IDs of existing VPC security groups and references to [AWS::EC2::SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	//
	// If you set `VPCSecurityGroups` , you must not set [`DBSecurityGroups`](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups) , and vice versa.
	//
	// > You can migrate a DB instance in your stack from an RDS DB security group to a VPC security group, but keep the following in mind:
	// >
	// > - You can't revert to using an RDS security group after you establish a VPC security group membership.
	// > - When you migrate your DB instance to VPC security groups, if your stack update rolls back because the DB instance update fails or because an update fails in another AWS CloudFormation resource, the rollback fails because it can't revert to an RDS security group.
	// > - To use the properties that are available when you use a VPC security group, you must recreate the DB instance. If you don't, AWS CloudFormation submits only the property values that are listed in the [`DBSecurityGroups`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups) property.
	//
	// To avoid this situation, migrate your DB instance to using VPC security groups only when that is the only change in your stack template.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster. If specified, the setting must match the DB cluster setting.
	VpcSecurityGroups *[]*string `json:"vpcSecurityGroups" yaml:"vpcSecurityGroups"`
}

// A CloudFormation `AWS::RDS::DBParameterGroup`.
//
// The `AWS::RDS::DBParameterGroup` resource creates a custom parameter group for an RDS database family.
//
// This type can be declared in a template and referenced in the `DBParameterGroupName` property of an `[AWS::RDS::DBInstance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html)` resource.
//
// For information about configuring parameters for Amazon RDS DB instances, see [Working with DB parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide* .
//
// For information about configuring parameters for Amazon Aurora DB instances, see [Working with DB parameter groups and DB cluster parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide* .
//
// > Applying a parameter group to a DB instance may require the DB instance to reboot, resulting in a database outage for the duration of the reboot.
//
// TODO: EXAMPLE
//
type CfnDBParameterGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	Family() *string
	SetFamily(val *string)
	LogicalId() *string
	Node() constructs.Node
	Parameters() interface{}
	SetParameters(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBParameterGroup
type jsiiProxy_CfnDBParameterGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBParameterGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBParameterGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBParameterGroup`.
func NewCfnDBParameterGroup(scope constructs.Construct, id *string, props *CfnDBParameterGroupProps) CfnDBParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDBParameterGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBParameterGroup`.
func NewCfnDBParameterGroup_Override(c CfnDBParameterGroup, scope constructs.Construct, id *string, props *CfnDBParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBParameterGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDBParameterGroup) SetFamily(val *string) {
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_CfnDBParameterGroup) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBParameterGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBParameterGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBParameterGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBParameterGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBParameterGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBParameterGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBParameterGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBParameterGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBParameterGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBParameterGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBParameterGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBParameterGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBParameterGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBParameterGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBParameterGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBParameterGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDBParameterGroup`.
//
// TODO: EXAMPLE
//
type CfnDBParameterGroupProps struct {
	// Provides the customer-specified description for this DB parameter group.
	Description *string `json:"description" yaml:"description"`
	// The DB parameter group family name.
	//
	// A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a DB engine and engine version compatible with that DB parameter group family.
	//
	// > The DB parameter group family can't be changed when updating a DB parameter group.
	//
	// To list all of the available parameter group families, use the following command:
	//
	// `aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"`
	//
	// The output contains duplicates.
	//
	// For more information, see `[CreateDBParameterGroup](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_CreateDBParameterGroup.html)` .
	Family *string `json:"family" yaml:"family"`
	// An array of parameter names and values for the parameter update.
	//
	// At least one parameter name and value must be supplied. Subsequent arguments are optional.
	//
	// For more information about DB parameters and DB parameter groups for Amazon RDS DB engines, see [Working with DB Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide* .
	//
	// For more information about DB cluster and DB instance parameters and parameter groups for Amazon Aurora DB engines, see [Working with DB Parameter Groups and DB Cluster Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide* .
	//
	// > AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// Tags to assign to the DB parameter group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::RDS::DBProxy`.
//
// The `AWS::RDS::DBProxy` resource creates or updates a DB proxy.
//
// For information about RDS Proxy for Amazon RDS, see [Managing Connections with Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html) in the *Amazon RDS User Guide* .
//
// For information about RDS Proxy for Amazon Aurora, see [Managing Connections with Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html) in the *Amazon Aurora User Guide* .
//
// > Limitations apply to RDS Proxy, including DB engine version limitations and AWS Region limitations.
// >
// > For information about limitations that apply to RDS Proxy for Amazon RDS, see [Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html#rds-proxy.limitations) in the *Amazon RDS User Guide* .
// >
// > For information about that apply to RDS Proxy for Amazon Aurora, see [Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html#rds-proxy.limitations) in the *Amazon Aurora User Guide* .
//
// TODO: EXAMPLE
//
type CfnDBProxy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDbProxyArn() *string
	AttrEndpoint() *string
	Auth() interface{}
	SetAuth(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DbProxyName() *string
	SetDbProxyName(val *string)
	DebugLogging() interface{}
	SetDebugLogging(val interface{})
	EngineFamily() *string
	SetEngineFamily(val *string)
	IdleClientTimeout() *float64
	SetIdleClientTimeout(val *float64)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RequireTls() interface{}
	SetRequireTls(val interface{})
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	Tags() *[]*CfnDBProxy_TagFormatProperty
	SetTags(val *[]*CfnDBProxy_TagFormatProperty)
	UpdatedProperites() *map[string]interface{}
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSubnetIds() *[]*string
	SetVpcSubnetIds(val *[]*string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBProxy
type jsiiProxy_CfnDBProxy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBProxy) AttrDbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) Auth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) DbProxyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) DebugLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) EngineFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) IdleClientTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleClientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) RequireTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) Tags() *[]*CfnDBProxy_TagFormatProperty {
	var returns *[]*CfnDBProxy_TagFormatProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxy) VpcSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnetIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBProxy`.
func NewCfnDBProxy(scope constructs.Construct, id *string, props *CfnDBProxyProps) CfnDBProxy {
	_init_.Initialize()

	j := jsiiProxy_CfnDBProxy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBProxy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBProxy`.
func NewCfnDBProxy_Override(c CfnDBProxy, scope constructs.Construct, id *string, props *CfnDBProxyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBProxy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetAuth(val interface{}) {
	_jsii_.Set(
		j,
		"auth",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetDbProxyName(val *string) {
	_jsii_.Set(
		j,
		"dbProxyName",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetDebugLogging(val interface{}) {
	_jsii_.Set(
		j,
		"debugLogging",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetEngineFamily(val *string) {
	_jsii_.Set(
		j,
		"engineFamily",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetIdleClientTimeout(val *float64) {
	_jsii_.Set(
		j,
		"idleClientTimeout",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetRequireTls(val interface{}) {
	_jsii_.Set(
		j,
		"requireTls",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetTags(val *[]*CfnDBProxy_TagFormatProperty) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetVpcSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxy) SetVpcSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSubnetIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBProxy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBProxy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBProxy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBProxy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBProxy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBProxy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBProxy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBProxy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBProxy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBProxy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBProxy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBProxy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBProxy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBProxy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBProxy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBProxy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBProxy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBProxy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBProxy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBProxy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the details of authentication used by a proxy to log in as a specific database user.
//
// TODO: EXAMPLE
//
type CfnDBProxy_AuthFormatProperty struct {
	// The type of authentication that the proxy uses for connections from the proxy to the underlying database.
	//
	// Valid Values: `SECRETS`
	AuthScheme *string `json:"authScheme" yaml:"authScheme"`
	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	Description *string `json:"description" yaml:"description"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	//
	// Valid Values: `DISABLED | REQUIRED`
	IamAuth *string `json:"iamAuth" yaml:"iamAuth"`
	// The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The name of the database user to which the proxy connects.
	UserName *string `json:"userName" yaml:"userName"`
}

// Metadata assigned to a DB proxy consisting of a key-value pair.
//
// TODO: EXAMPLE
//
type CfnDBProxy_TagFormatProperty struct {
	// A key is the required name of the tag.
	//
	// The string value can be 1-128 Unicode characters in length and can't be prefixed with `aws:` . The string can contain only the set of Unicode letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Key *string `json:"key" yaml:"key"`
	// A value is the optional value of the tag.
	//
	// The string value can be 1-256 Unicode characters in length and can't be prefixed with `aws:` . The string can contain only the set of Unicode letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Value *string `json:"value" yaml:"value"`
}

// A CloudFormation `AWS::RDS::DBProxyEndpoint`.
//
// The `AWS::RDS::DBProxyEndpoint` resource creates or updates a DB proxy endpoint. You can use custom proxy endpoints to access a proxy through a different VPC than the proxy's default VPC.
//
// For more information about RDS Proxy, see [AWS::RDS::DBProxy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html) .
//
// TODO: EXAMPLE
//
type CfnDBProxyEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDbProxyEndpointArn() *string
	AttrEndpoint() *string
	AttrIsDefault() awscdk.IResolvable
	AttrVpcId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DbProxyEndpointName() *string
	SetDbProxyEndpointName(val *string)
	DbProxyName() *string
	SetDbProxyName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() *[]*CfnDBProxyEndpoint_TagFormatProperty
	SetTags(val *[]*CfnDBProxyEndpoint_TagFormatProperty)
	TargetRole() *string
	SetTargetRole(val *string)
	UpdatedProperites() *map[string]interface{}
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSubnetIds() *[]*string
	SetVpcSubnetIds(val *[]*string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBProxyEndpoint
type jsiiProxy_CfnDBProxyEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBProxyEndpoint) AttrDbProxyEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbProxyEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) AttrIsDefault() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) AttrVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) DbProxyEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) DbProxyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) Tags() *[]*CfnDBProxyEndpoint_TagFormatProperty {
	var returns *[]*CfnDBProxyEndpoint_TagFormatProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) TargetRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpoint) VpcSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnetIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBProxyEndpoint`.
func NewCfnDBProxyEndpoint(scope constructs.Construct, id *string, props *CfnDBProxyEndpointProps) CfnDBProxyEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnDBProxyEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBProxyEndpoint`.
func NewCfnDBProxyEndpoint_Override(c CfnDBProxyEndpoint, scope constructs.Construct, id *string, props *CfnDBProxyEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBProxyEndpoint) SetDbProxyEndpointName(val *string) {
	_jsii_.Set(
		j,
		"dbProxyEndpointName",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyEndpoint) SetDbProxyName(val *string) {
	_jsii_.Set(
		j,
		"dbProxyName",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyEndpoint) SetTags(val *[]*CfnDBProxyEndpoint_TagFormatProperty) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyEndpoint) SetTargetRole(val *string) {
	_jsii_.Set(
		j,
		"targetRole",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyEndpoint) SetVpcSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyEndpoint) SetVpcSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSubnetIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBProxyEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBProxyEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBProxyEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBProxyEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBProxyEndpoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBProxyEndpoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBProxyEndpoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBProxyEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBProxyEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBProxyEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBProxyEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBProxyEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBProxyEndpoint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBProxyEndpoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBProxyEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBProxyEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBProxyEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBProxyEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBProxyEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Metadata assigned to a DB proxy endpoint consisting of a key-value pair.
//
// TODO: EXAMPLE
//
type CfnDBProxyEndpoint_TagFormatProperty struct {
	// A value is the optional value of the tag.
	//
	// The string value can be 1-256 Unicode characters in length and can't be prefixed with `aws:` . The string can contain only the set of Unicode letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Key *string `json:"key" yaml:"key"`
	// Metadata assigned to a DB instance consisting of a key-value pair.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnDBProxyEndpoint`.
//
// TODO: EXAMPLE
//
type CfnDBProxyEndpointProps struct {
	// The name of the DB proxy endpoint to create.
	DbProxyEndpointName *string `json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName *string `json:"dbProxyName" yaml:"dbProxyName"`
	// The VPC subnet IDs for the DB proxy endpoint that you create.
	//
	// You can specify a different set of subnet IDs than for the original DB proxy.
	VpcSubnetIds *[]*string `json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	Tags *[]*CfnDBProxyEndpoint_TagFormatProperty `json:"tags" yaml:"tags"`
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	//
	// Valid Values: `READ_WRITE | READ_ONLY`
	TargetRole *string `json:"targetRole" yaml:"targetRole"`
	// The VPC security group IDs for the DB proxy endpoint that you create.
	//
	// You can specify a different set of security group IDs than for the original DB proxy. The default is the default security group for the VPC.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

// Properties for defining a `CfnDBProxy`.
//
// TODO: EXAMPLE
//
type CfnDBProxyProps struct {
	// The authorization mechanism that the proxy uses.
	Auth interface{} `json:"auth" yaml:"auth"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyName *string `json:"dbProxyName" yaml:"dbProxyName"`
	// The kinds of databases that the proxy can connect to.
	//
	// This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora.
	//
	// *Valid values* : `MYSQL` | `POSTGRESQL`
	EngineFamily *string `json:"engineFamily" yaml:"engineFamily"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds *[]*string `json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging interface{} `json:"debugLogging" yaml:"debugLogging"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *float64 `json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls interface{} `json:"requireTls" yaml:"requireTls"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	Tags *[]*CfnDBProxy_TagFormatProperty `json:"tags" yaml:"tags"`
	// One or more VPC security group IDs to associate with the new proxy.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::RDS::DBProxyTargetGroup`.
//
// The `AWS::RDS::DBProxyTargetGroup` resource represents a set of RDS DB instances, Aurora DB clusters, or both that a proxy can connect to. Currently, each target group is associated with exactly one RDS DB instance or Aurora DB cluster.
//
// This data type is used as a response element in the `DescribeDBProxyTargetGroups` action.
//
// For information about RDS Proxy for Amazon RDS, see [Managing Connections with Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html) in the *Amazon RDS User Guide* .
//
// For information about RDS Proxy for Amazon Aurora, see [Managing Connections with Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html) in the *Amazon Aurora User Guide* .
//
// For a sample template that creates a DB proxy and registers a DB instance, see [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#aws-resource-rds-dbproxy--examples) in AWS::RDS::DBProxy.
//
// > Limitations apply to RDS Proxy, including DB engine version limitations and AWS Region limitations.
// >
// > For information about limitations that apply to RDS Proxy for Amazon RDS, see [Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html#rds-proxy.limitations) in the *Amazon RDS User Guide* .
// >
// > For information about that apply to RDS Proxy for Amazon Aurora, see [Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html#rds-proxy.limitations) in the *Amazon Aurora User Guide* .
//
// TODO: EXAMPLE
//
type CfnDBProxyTargetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrTargetGroupArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectionPoolConfigurationInfo() interface{}
	SetConnectionPoolConfigurationInfo(val interface{})
	CreationStack() *[]*string
	DbClusterIdentifiers() *[]*string
	SetDbClusterIdentifiers(val *[]*string)
	DbInstanceIdentifiers() *[]*string
	SetDbInstanceIdentifiers(val *[]*string)
	DbProxyName() *string
	SetDbProxyName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	TargetGroupName() *string
	SetTargetGroupName(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBProxyTargetGroup
type jsiiProxy_CfnDBProxyTargetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) AttrTargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTargetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) ConnectionPoolConfigurationInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionPoolConfigurationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) DbClusterIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbClusterIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) DbInstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbInstanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) DbProxyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBProxyTargetGroup`.
func NewCfnDBProxyTargetGroup(scope constructs.Construct, id *string, props *CfnDBProxyTargetGroupProps) CfnDBProxyTargetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDBProxyTargetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBProxyTargetGroup`.
func NewCfnDBProxyTargetGroup_Override(c CfnDBProxyTargetGroup, scope constructs.Construct, id *string, props *CfnDBProxyTargetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) SetConnectionPoolConfigurationInfo(val interface{}) {
	_jsii_.Set(
		j,
		"connectionPoolConfigurationInfo",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) SetDbClusterIdentifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifiers",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) SetDbInstanceIdentifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"dbInstanceIdentifiers",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) SetDbProxyName(val *string) {
	_jsii_.Set(
		j,
		"dbProxyName",
		val,
	)
}

func (j *jsiiProxy_CfnDBProxyTargetGroup) SetTargetGroupName(val *string) {
	_jsii_.Set(
		j,
		"targetGroupName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBProxyTargetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBProxyTargetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBProxyTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBProxyTargetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBProxyTargetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBProxyTargetGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBProxyTargetGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBProxyTargetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBProxyTargetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBProxyTargetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBProxyTargetGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBProxyTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBProxyTargetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the settings that control the size and behavior of the connection pool associated with a `DBProxyTargetGroup` .
//
// TODO: EXAMPLE
//
type CfnDBProxyTargetGroup_ConnectionPoolConfigurationInfoFormatProperty struct {
	// The number of seconds for a proxy to wait for a connection to become available in the connection pool.
	//
	// Only applies when the proxy has opened its maximum number of connections and all connections are busy with client sessions.
	//
	// Default: 120
	//
	// Constraints: between 1 and 3600, or 0 representing unlimited
	ConnectionBorrowTimeout *float64 `json:"connectionBorrowTimeout" yaml:"connectionBorrowTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Typically used with `SET` statements to make sure that each connection has identical settings such as time zone and character set. For multiple statements, use semicolons as the separator. You can also include multiple variables in a single `SET` statement, such as `SET x=1, y=2` .
	//
	// Default: no initialization query
	InitQuery *string `json:"initQuery" yaml:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// The value is expressed as a percentage of the `max_connections` setting for the RDS DB instance or Aurora DB cluster used by the target group.
	//
	// Default: 100
	//
	// Constraints: between 1 and 100
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// The value is expressed as a percentage of the `max_connections` setting for the RDS DB instance or Aurora DB cluster used by the target group. With a high value, the proxy leaves a high percentage of idle database connections open. A low value causes the proxy to close more idle connections and return them to the database.
	//
	// Default: 50
	//
	// Constraints: between 0 and `MaxConnectionsPercent`
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Including an item in the list exempts that class of SQL operations from the pinning behavior.
	//
	// Default: no session pinning filters
	SessionPinningFilters *[]*string `json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
}

// Properties for defining a `CfnDBProxyTargetGroup`.
//
// TODO: EXAMPLE
//
type CfnDBProxyTargetGroupProps struct {
	// The identifier of the `DBProxy` that is associated with the `DBProxyTargetGroup` .
	DbProxyName *string `json:"dbProxyName" yaml:"dbProxyName"`
	// The identifier for the target group.
	//
	// > Currently, this property must be set to `default` .
	TargetGroupName *string `json:"targetGroupName" yaml:"targetGroupName"`
	// Settings that control the size and behavior of the connection pool associated with a `DBProxyTargetGroup` .
	ConnectionPoolConfigurationInfo interface{} `json:"connectionPoolConfigurationInfo" yaml:"connectionPoolConfigurationInfo"`
	// One or more DB cluster identifiers.
	DbClusterIdentifiers *[]*string `json:"dbClusterIdentifiers" yaml:"dbClusterIdentifiers"`
	// One or more DB instance identifiers.
	DbInstanceIdentifiers *[]*string `json:"dbInstanceIdentifiers" yaml:"dbInstanceIdentifiers"`
}

// A CloudFormation `AWS::RDS::DBSecurityGroup`.
//
// The `AWS::RDS::DBSecurityGroup` resource creates or updates an Amazon RDS DB security group.
//
// > DB security groups are a part of the EC2-Classic Platform and as such are not supported in all regions. It is advised to use the `AWS::EC2::SecurityGroup` resource in those regions instead. To determine which platform you are on, see [Determining Whether You Are Using the EC2-VPC or EC2-Classic Platform](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.FindDefaultVPC.html) . For more information on the `AWS::EC2::SecurityGroup` , see the documentation for [EC2 security groups](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) .
//
// TODO: EXAMPLE
//
type CfnDBSecurityGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DbSecurityGroupIngress() interface{}
	SetDbSecurityGroupIngress(val interface{})
	Ec2VpcId() *string
	SetEc2VpcId(val *string)
	GroupDescription() *string
	SetGroupDescription(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBSecurityGroup
type jsiiProxy_CfnDBSecurityGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBSecurityGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) DbSecurityGroupIngress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbSecurityGroupIngress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) Ec2VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2VpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) GroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBSecurityGroup`.
func NewCfnDBSecurityGroup(scope constructs.Construct, id *string, props *CfnDBSecurityGroupProps) CfnDBSecurityGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDBSecurityGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBSecurityGroup`.
func NewCfnDBSecurityGroup_Override(c CfnDBSecurityGroup, scope constructs.Construct, id *string, props *CfnDBSecurityGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroup) SetDbSecurityGroupIngress(val interface{}) {
	_jsii_.Set(
		j,
		"dbSecurityGroupIngress",
		val,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroup) SetEc2VpcId(val *string) {
	_jsii_.Set(
		j,
		"ec2VpcId",
		val,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroup) SetGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"groupDescription",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBSecurityGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBSecurityGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBSecurityGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBSecurityGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBSecurityGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBSecurityGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBSecurityGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBSecurityGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBSecurityGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBSecurityGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBSecurityGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBSecurityGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBSecurityGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBSecurityGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBSecurityGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBSecurityGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBSecurityGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBSecurityGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSecurityGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `Ingress` property type specifies an individual ingress rule within an `AWS::RDS::DBSecurityGroup` resource.
//
// TODO: EXAMPLE
//
type CfnDBSecurityGroup_IngressProperty struct {
	// The IP range to authorize.
	Cidrip *string `json:"cidrip" yaml:"cidrip"`
	// Id of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupId *string `json:"ec2SecurityGroupId" yaml:"ec2SecurityGroupId"`
	// Name of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupName *string `json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// AWS account number of the owner of the EC2 security group specified in the `EC2SecurityGroupName` parameter.
	//
	// The AWS access key ID isn't an acceptable value. For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupOwnerId *string `json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

// A CloudFormation `AWS::RDS::DBSecurityGroupIngress`.
//
// The `AWS::RDS::DBSecurityGroupIngress` resource enables ingress to a DB security group using one of two forms of authorization. First, you can add EC2 or VPC security groups to the DB security group if the application using the database is running on EC2 or VPC instances. Second, IP ranges are available if the application accessing your database is running on the Internet.
//
// This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// For details about the settings for DB security group ingress, see [AuthorizeDBSecurityGroupIngress](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_AuthorizeDBSecurityGroupIngress.html) .
//
// TODO: EXAMPLE
//
type CfnDBSecurityGroupIngress interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Cidrip() *string
	SetCidrip(val *string)
	CreationStack() *[]*string
	DbSecurityGroupName() *string
	SetDbSecurityGroupName(val *string)
	Ec2SecurityGroupId() *string
	SetEc2SecurityGroupId(val *string)
	Ec2SecurityGroupName() *string
	SetEc2SecurityGroupName(val *string)
	Ec2SecurityGroupOwnerId() *string
	SetEc2SecurityGroupOwnerId(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBSecurityGroupIngress
type jsiiProxy_CfnDBSecurityGroupIngress struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) Cidrip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) DbSecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) Ec2SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2SecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) Ec2SecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2SecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) Ec2SecurityGroupOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2SecurityGroupOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBSecurityGroupIngress`.
func NewCfnDBSecurityGroupIngress(scope constructs.Construct, id *string, props *CfnDBSecurityGroupIngressProps) CfnDBSecurityGroupIngress {
	_init_.Initialize()

	j := jsiiProxy_CfnDBSecurityGroupIngress{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngress",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBSecurityGroupIngress`.
func NewCfnDBSecurityGroupIngress_Override(c CfnDBSecurityGroupIngress, scope constructs.Construct, id *string, props *CfnDBSecurityGroupIngressProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngress",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) SetCidrip(val *string) {
	_jsii_.Set(
		j,
		"cidrip",
		val,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) SetDbSecurityGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) SetEc2SecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"ec2SecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) SetEc2SecurityGroupName(val *string) {
	_jsii_.Set(
		j,
		"ec2SecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBSecurityGroupIngress) SetEc2SecurityGroupOwnerId(val *string) {
	_jsii_.Set(
		j,
		"ec2SecurityGroupOwnerId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBSecurityGroupIngress_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngress",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBSecurityGroupIngress_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngress",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBSecurityGroupIngress_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngress",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBSecurityGroupIngress_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngress",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBSecurityGroupIngress) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBSecurityGroupIngress) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBSecurityGroupIngress) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBSecurityGroupIngress) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSecurityGroupIngress) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDBSecurityGroupIngress`.
//
// TODO: EXAMPLE
//
type CfnDBSecurityGroupIngressProps struct {
	// The name of the DB security group to add authorization to.
	DbSecurityGroupName *string `json:"dbSecurityGroupName" yaml:"dbSecurityGroupName"`
	// The IP range to authorize.
	Cidrip *string `json:"cidrip" yaml:"cidrip"`
	// Id of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupId *string `json:"ec2SecurityGroupId" yaml:"ec2SecurityGroupId"`
	// Name of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupName *string `json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// AWS account number of the owner of the EC2 security group specified in the `EC2SecurityGroupName` parameter.
	//
	// The AWS access key ID isn't an acceptable value. For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupOwnerId *string `json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

// Properties for defining a `CfnDBSecurityGroup`.
//
// TODO: EXAMPLE
//
type CfnDBSecurityGroupProps struct {
	// Ingress rules to be applied to the DB security group.
	DbSecurityGroupIngress interface{} `json:"dbSecurityGroupIngress" yaml:"dbSecurityGroupIngress"`
	// Provides the description of the DB security group.
	GroupDescription *string `json:"groupDescription" yaml:"groupDescription"`
	// The identifier of an Amazon VPC. This property indicates the VPC that this DB security group belongs to.
	//
	// > The `EC2VpcId` property is for backward compatibility with older regions, and is no longer recommended for providing security information to an RDS DB instance.
	Ec2VpcId *string `json:"ec2VpcId" yaml:"ec2VpcId"`
	// Tags to assign to the DB security group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::RDS::DBSubnetGroup`.
//
// The `AWS::RDS::DBSubnetGroup` resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.
//
// For more information, see [Working with DB subnet groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html#USER_VPC.Subnets) in the *Amazon RDS User Guide* .
//
// TODO: EXAMPLE
//
type CfnDBSubnetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DbSubnetGroupDescription() *string
	SetDbSubnetGroupDescription(val *string)
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBSubnetGroup
type jsiiProxy_CfnDBSubnetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBSubnetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) DbSubnetGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBSubnetGroup`.
func NewCfnDBSubnetGroup(scope constructs.Construct, id *string, props *CfnDBSubnetGroupProps) CfnDBSubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDBSubnetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBSubnetGroup`.
func NewCfnDBSubnetGroup_Override(c CfnDBSubnetGroup, scope constructs.Construct, id *string, props *CfnDBSubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBSubnetGroup) SetDbSubnetGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDBSubnetGroup) SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBSubnetGroup) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDBSubnetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDBSubnetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDBSubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBSubnetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDBSubnetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDBSubnetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBSubnetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDBSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDBSubnetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDBSubnetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDBSubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDBSubnetGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDBSubnetGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDBSubnetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDBSubnetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDBSubnetGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDBSubnetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDBSubnetGroup`.
//
// TODO: EXAMPLE
//
type CfnDBSubnetGroupProps struct {
	// The description for the DB subnet group.
	DbSubnetGroupDescription *string `json:"dbSubnetGroupDescription" yaml:"dbSubnetGroupDescription"`
	// The EC2 Subnet IDs for the DB subnet group.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// The name for the DB subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 lowercase alphanumeric characters or hyphens. Must not be "Default".
	//
	// Example: `mysubnetgroup`
	DbSubnetGroupName *string `json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Tags to assign to the DB subnet group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::RDS::EventSubscription`.
//
// The `AWS::RDS::EventSubscription` resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see [Using Amazon RDS Event Notification](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html) in the *Amazon RDS User Guide* .
//
// TODO: EXAMPLE
//
type CfnEventSubscription interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EventCategories() *[]*string
	SetEventCategories(val *[]*string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	SourceIds() *[]*string
	SetSourceIds(val *[]*string)
	SourceType() *string
	SetSourceType(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnEventSubscription
type jsiiProxy_CfnEventSubscription struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventSubscription) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) EventCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::EventSubscription`.
func NewCfnEventSubscription(scope constructs.Construct, id *string, props *CfnEventSubscriptionProps) CfnEventSubscription {
	_init_.Initialize()

	j := jsiiProxy_CfnEventSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnEventSubscription",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::EventSubscription`.
func NewCfnEventSubscription_Override(c CfnEventSubscription, scope constructs.Construct, id *string, props *CfnEventSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnEventSubscription",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetEventCategories(val *[]*string) {
	_jsii_.Set(
		j,
		"eventCategories",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"snsTopicArn",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSourceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceIds",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSourceType(val *string) {
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventSubscription_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnEventSubscription",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventSubscription_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnEventSubscription",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnEventSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventSubscription_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnEventSubscription",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEventSubscription) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnEventSubscription) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnEventSubscription) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEventSubscription) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEventSubscription) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnEventSubscription) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnEventSubscription) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnEventSubscription) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnEventSubscription) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEventSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventSubscription) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnEventSubscription) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnEventSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSubscription) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEventSubscription`.
//
// TODO: EXAMPLE
//
type CfnEventSubscriptionProps struct {
	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	//
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn *string `json:"snsTopicArn" yaml:"snsTopicArn"`
	// A value that indicates whether to activate the subscription.
	//
	// If the event notification subscription isn't activated, the subscription is created but not active.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// A list of event categories for a particular source type ( `SourceType` ) that you want to subscribe to.
	//
	// You can see a list of the categories for a given source type in the "Amazon RDS event categories and event messages" section of the [*Amazon RDS User Guide*](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.Messages.html) or the [*Amazon Aurora User Guide*](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Messages.html) . You can also see this list by using the `DescribeEventCategories` operation.
	EventCategories *[]*string `json:"eventCategories" yaml:"eventCategories"`
	// The list of identifiers of the event sources for which events are returned.
	//
	// If not specified, then all sources are included in the response. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens. It can't end with a hyphen or contain two consecutive hyphens.
	//
	// Constraints:
	//
	// - If a `SourceIds` value is supplied, `SourceType` must also be provided.
	// - If the source type is a DB instance, a `DBInstanceIdentifier` value must be supplied.
	// - If the source type is a DB cluster, a `DBClusterIdentifier` value must be supplied.
	// - If the source type is a DB parameter group, a `DBParameterGroupName` value must be supplied.
	// - If the source type is a DB security group, a `DBSecurityGroupName` value must be supplied.
	// - If the source type is a DB snapshot, a `DBSnapshotIdentifier` value must be supplied.
	// - If the source type is a DB cluster snapshot, a `DBClusterSnapshotIdentifier` value must be supplied.
	SourceIds *[]*string `json:"sourceIds" yaml:"sourceIds"`
	// The type of source that is generating the events.
	//
	// For example, if you want to be notified of events generated by a DB instance, set this parameter to `db-instance` . If this value isn't specified, all events are returned.
	//
	// Valid values: `db-instance` | `db-cluster` | `db-parameter-group` | `db-security-group` | `db-snapshot` | `db-cluster-snapshot`
	SourceType *string `json:"sourceType" yaml:"sourceType"`
}

// A CloudFormation `AWS::RDS::GlobalCluster`.
//
// The `AWS::RDS::GlobalCluster` resource creates or updates an Amazon Aurora global database spread across multiple AWS Regions.
//
// The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
//
// You can create a global database that is initially empty, and then add a primary cluster and a secondary cluster to it.
//
// For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) in the *Amazon Aurora User Guide* .
//
// TODO: EXAMPLE
//
type CfnGlobalCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	Engine() *string
	SetEngine(val *string)
	EngineVersion() *string
	SetEngineVersion(val *string)
	GlobalClusterIdentifier() *string
	SetGlobalClusterIdentifier(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	Stack() awscdk.Stack
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGlobalCluster
type jsiiProxy_CfnGlobalCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGlobalCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::GlobalCluster`.
func NewCfnGlobalCluster(scope constructs.Construct, id *string, props *CfnGlobalClusterProps) CfnGlobalCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnGlobalCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnGlobalCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::GlobalCluster`.
func NewCfnGlobalCluster_Override(c CfnGlobalCluster, scope constructs.Construct, id *string, props *CfnGlobalClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnGlobalCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGlobalCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalCluster) SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalCluster) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalCluster) SetGlobalClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalCluster) SetSourceDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalCluster) SetStorageEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGlobalCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnGlobalCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGlobalCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnGlobalCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnGlobalCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnGlobalCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnGlobalCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnGlobalCluster) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnGlobalCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnGlobalCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnGlobalCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnGlobalCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnGlobalCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnGlobalCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnGlobalCluster) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnGlobalCluster) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnGlobalCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnGlobalCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGlobalCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnGlobalCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnGlobalCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnGlobalCluster`.
//
// TODO: EXAMPLE
//
type CfnGlobalClusterProps struct {
	// The deletion protection setting for the new global database.
	//
	// The global database can't be deleted when deletion protection is enabled.
	DeletionProtection interface{} `json:"deletionProtection" yaml:"deletionProtection"`
	// The name of the database engine to be used for this DB cluster.
	//
	// If this property isn't specified, the database engine is derived from the source DB cluster specified by the `SourceDBClusterIdentifier` property.
	//
	// > If the `SourceDBClusterIdentifier` property isn't specified, this property is required. If the `SourceDBClusterIdentifier` property is specified, make sure this property isn't specified.
	Engine *string `json:"engine" yaml:"engine"`
	// The engine version of the Aurora global database.
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// The cluster identifier of the global database cluster.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The DB cluster identifier or Amazon Resource Name (ARN) to use as the primary cluster of the global database.
	//
	// > If the `Engine` property isn't specified, this property is required. If the `Engine` property is specified, make sure this property isn't specified.
	SourceDbClusterIdentifier *string `json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The storage encryption setting for the global database cluster.
	StorageEncrypted interface{} `json:"storageEncrypted" yaml:"storageEncrypted"`
}

// A CloudFormation `AWS::RDS::OptionGroup`.
//
// The `AWS::RDS::OptionGroup` resource creates or updates an option group, to enable and configure features that are specific to a particular DB engine.
//
// TODO: EXAMPLE
//
type CfnOptionGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EngineName() *string
	SetEngineName(val *string)
	LogicalId() *string
	MajorEngineVersion() *string
	SetMajorEngineVersion(val *string)
	Node() constructs.Node
	OptionConfigurations() interface{}
	SetOptionConfigurations(val interface{})
	OptionGroupDescription() *string
	SetOptionGroupDescription(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnOptionGroup
type jsiiProxy_CfnOptionGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnOptionGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) MajorEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"majorEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) OptionConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) OptionGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOptionGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::OptionGroup`.
func NewCfnOptionGroup(scope constructs.Construct, id *string, props *CfnOptionGroupProps) CfnOptionGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnOptionGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnOptionGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::OptionGroup`.
func NewCfnOptionGroup_Override(c CfnOptionGroup, scope constructs.Construct, id *string, props *CfnOptionGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnOptionGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOptionGroup) SetEngineName(val *string) {
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_CfnOptionGroup) SetMajorEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"majorEngineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnOptionGroup) SetOptionConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"optionConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnOptionGroup) SetOptionGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"optionGroupDescription",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnOptionGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnOptionGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnOptionGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnOptionGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnOptionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnOptionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOptionGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CfnOptionGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnOptionGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnOptionGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnOptionGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnOptionGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnOptionGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnOptionGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnOptionGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnOptionGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnOptionGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnOptionGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnOptionGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOptionGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnOptionGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnOptionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOptionGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `OptionConfiguration` property type specifies an individual option, and its settings, within an `AWS::RDS::OptionGroup` resource.
//
// TODO: EXAMPLE
//
type CfnOptionGroup_OptionConfigurationProperty struct {
	// The configuration of options to include in a group.
	OptionName *string `json:"optionName" yaml:"optionName"`
	// A list of DBSecurityGroupMembership name strings used for this option.
	DbSecurityGroupMemberships *[]*string `json:"dbSecurityGroupMemberships" yaml:"dbSecurityGroupMemberships"`
	// The option settings to include in an option group.
	OptionSettings interface{} `json:"optionSettings" yaml:"optionSettings"`
	// The version for the option.
	OptionVersion *string `json:"optionVersion" yaml:"optionVersion"`
	// The optional port for the option.
	Port *float64 `json:"port" yaml:"port"`
	// A list of VpcSecurityGroupMembership name strings used for this option.
	VpcSecurityGroupMemberships *[]*string `json:"vpcSecurityGroupMemberships" yaml:"vpcSecurityGroupMemberships"`
}

// The `OptionSetting` property type specifies the value for an option within an `OptionSetting` property.
//
// TODO: EXAMPLE
//
type CfnOptionGroup_OptionSettingProperty struct {
	// The name of the option that has settings that you can set.
	Name *string `json:"name" yaml:"name"`
	// The current value of the option setting.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnOptionGroup`.
//
// TODO: EXAMPLE
//
type CfnOptionGroupProps struct {
	// Specifies the name of the engine that this option group should be associated with.
	//
	// Valid Values:
	//
	// - `mariadb`
	// - `mysql`
	// - `oracle-ee`
	// - `oracle-se2`
	// - `oracle-se1`
	// - `oracle-se`
	// - `postgres`
	// - `sqlserver-ee`
	// - `sqlserver-se`
	// - `sqlserver-ex`
	// - `sqlserver-web`
	EngineName *string `json:"engineName" yaml:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion *string `json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// A list of options and the settings for each option.
	OptionConfigurations interface{} `json:"optionConfigurations" yaml:"optionConfigurations"`
	// The description of the option group.
	OptionGroupDescription *string `json:"optionGroupDescription" yaml:"optionGroupDescription"`
	// Tags to assign to the option group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// The extra options passed to the {@link IClusterEngine.bindToCluster} method.
//
// TODO: EXAMPLE
//
type ClusterEngineBindOptions struct {
	// The customer-provided ParameterGroup.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The role used for S3 exporting.
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// The role used for S3 importing.
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
}

// The type returned from the {@link IClusterEngine.bindToCluster} method.
//
// TODO: EXAMPLE
//
type ClusterEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	Features *ClusterEngineFeatures `json:"features" yaml:"features"`
	// The ParameterGroup to use for the cluster.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The port to use for this cluster, unless the customer specified the port directly.
	Port *float64 `json:"port" yaml:"port"`
}

// Represents Database Engine features.
//
// TODO: EXAMPLE
//
type ClusterEngineFeatures struct {
	// Feature name for the DB instance that the IAM role to export to S3 bucket is to be associated with.
	S3Export *string `json:"s3Export" yaml:"s3Export"`
	// Feature name for the DB instance that the IAM role to access the S3 bucket for import is to be associated with.
	S3Import *string `json:"s3Import" yaml:"s3Import"`
}

// Username and password combination.
//
// TODO: EXAMPLE
//
type Credentials interface {
	EncryptionKey() awskms.IKey
	ExcludeCharacters() *string
	Password() awscdk.SecretValue
	ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion
	Secret() awssecretsmanager.ISecret
	SecretName() *string
	Username() *string
	UsernameAsString() *bool
}

// The jsii proxy struct for Credentials
type jsiiProxy_Credentials struct {
	_ byte // padding
}

func (j *jsiiProxy_Credentials) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion {
	var returns *[]*awssecretsmanager.ReplicaRegion
	_jsii_.Get(
		j,
		"replicaRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) UsernameAsString() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"usernameAsString",
		&returns,
	)
	return returns
}


func NewCredentials_Override(c Credentials) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Credentials",
		nil, // no parameters
		c,
	)
}

// Creates Credentials with a password generated and stored in Secrets Manager.
func Credentials_FromGeneratedSecret(username *string, options *CredentialsBaseOptions) Credentials {
	_init_.Initialize()

	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromGeneratedSecret",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Creates Credentials from a password.
//
// Do not put passwords in your CDK code directly.
func Credentials_FromPassword(username *string, password awscdk.SecretValue) Credentials {
	_init_.Initialize()

	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromPassword",
		[]interface{}{username, password},
		&returns,
	)

	return returns
}

// Creates Credentials from an existing Secrets Manager ``Secret`` (or ``DatabaseSecret``).
//
// The Secret must be a JSON string with a ``username`` and ``password`` field:
// ```
// {
//    ...
//    "username": <required: username>,
//    "password": <required: password>,
// }
// ```
func Credentials_FromSecret(secret awssecretsmanager.ISecret, username *string) Credentials {
	_init_.Initialize()

	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromSecret",
		[]interface{}{secret, username},
		&returns,
	)

	return returns
}

// Creates Credentials for the given username, and optional password and key.
//
// If no password is provided, one will be generated and stored in Secrets Manager.
func Credentials_FromUsername(username *string, options *CredentialsFromUsernameOptions) Credentials {
	_init_.Initialize()

	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromUsername",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Base options for creating Credentials.
//
// TODO: EXAMPLE
//
type CredentialsBaseOptions struct {
	// KMS encryption key to encrypt the generated secret.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// The characters to exclude from the generated password.
	//
	// Has no effect if {@link password} has been provided.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `json:"replicaRegions" yaml:"replicaRegions"`
	// The name of the secret.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// Options for creating Credentials from a username.
//
// TODO: EXAMPLE
//
type CredentialsFromUsernameOptions struct {
	// KMS encryption key to encrypt the generated secret.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// The characters to exclude from the generated password.
	//
	// Has no effect if {@link password} has been provided.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `json:"replicaRegions" yaml:"replicaRegions"`
	// The name of the secret.
	SecretName *string `json:"secretName" yaml:"secretName"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	Password awscdk.SecretValue `json:"password" yaml:"password"`
}

// Create a clustered database with a given number of instances.
//
// TODO: EXAMPLE
//
type DatabaseCluster interface {
	DatabaseClusterBase
	ClusterEndpoint() Endpoint
	ClusterIdentifier() *string
	ClusterReadEndpoint() Endpoint
	Connections() awsec2.Connections
	Engine() IClusterEngine
	Env() *awscdk.ResourceEnvironment
	InstanceEndpoints() *[]Endpoint
	InstanceIdentifiers() *[]*string
	NewCfnProps() *CfnDBClusterProps
	Node() constructs.Node
	PhysicalName() *string
	Secret() awssecretsmanager.ISecret
	SecurityGroups() *[]awsec2.ISecurityGroup
	Stack() awscdk.Stack
	SubnetGroup() ISubnetGroup
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation
	AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for DatabaseCluster
type jsiiProxy_DatabaseCluster struct {
	jsiiProxy_DatabaseClusterBase
}

func (j *jsiiProxy_DatabaseCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Engine() IClusterEngine {
	var returns IClusterEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) NewCfnProps() *CfnDBClusterProps {
	var returns *CfnDBClusterProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) SubnetGroup() ISubnetGroup {
	var returns ISubnetGroup
	_jsii_.Get(
		j,
		"subnetGroup",
		&returns,
	)
	return returns
}


func NewDatabaseCluster(scope constructs.Construct, id *string, props *DatabaseClusterProps) DatabaseCluster {
	_init_.Initialize()

	j := jsiiProxy_DatabaseCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseCluster_Override(d DatabaseCluster, scope constructs.Construct, id *string, props *DatabaseClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseCluster",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing DatabaseCluster from properties.
func DatabaseCluster_FromDatabaseClusterAttributes(scope constructs.Construct, id *string, attrs *DatabaseClusterAttributes) IDatabaseCluster {
	_init_.Initialize()

	var returns IDatabaseCluster

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseCluster",
		"fromDatabaseClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseCluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseCluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new db proxy to this cluster.
func (d *jsiiProxy_DatabaseCluster) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds the multi user rotation to this cluster.
func (d *jsiiProxy_DatabaseCluster) AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationMultiUser",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds the single user rotation of the master password to this cluster.
func (d *jsiiProxy_DatabaseCluster) AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseCluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseCluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the given named metric for this DBCluster.
func (d *jsiiProxy_DatabaseCluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The percentage of CPU utilization.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of database connections in use.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of deadlocks in the database per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDeadlocks",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of time that the instance has been running, in seconds.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricEngineUptime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available random access memory, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of local storage available, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeLocalStorage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput received from clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkReceiveThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput both received from and transmitted to clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput sent to clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkTransmitThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total amount of backup storage in bytes consumed by all Aurora snapshots outside its backup retention window.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSnapshotStorageUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total amount of backup storage in bytes for which you are billed.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTotalBackupStorageBilled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of storage used by your Aurora DB instance, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeBytesUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of billed read I/O operations from a cluster volume, reported at 5-minute intervals.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeReadIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of write disk I/O operations to the cluster volume, reported at 5-minute intervals.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseCluster) MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeWriteIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing cluster instance.
//
// TODO: EXAMPLE
//
type DatabaseClusterAttributes struct {
	// Identifier for the cluster.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Cluster endpoint address.
	ClusterEndpointAddress *string `json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// The engine of the existing Cluster.
	Engine IClusterEngine `json:"engine" yaml:"engine"`
	// Endpoint addresses of individual instances.
	InstanceEndpointAddresses *[]*string `json:"instanceEndpointAddresses" yaml:"instanceEndpointAddresses"`
	// Identifier for the instances.
	InstanceIdentifiers *[]*string `json:"instanceIdentifiers" yaml:"instanceIdentifiers"`
	// The database port.
	Port *float64 `json:"port" yaml:"port"`
	// Reader endpoint address.
	ReaderEndpointAddress *string `json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The security groups of the database cluster.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
}

// A new or imported clustered database.
type DatabaseClusterBase interface {
	awscdk.Resource
	IDatabaseCluster
	ClusterEndpoint() Endpoint
	ClusterIdentifier() *string
	ClusterReadEndpoint() Endpoint
	Connections() awsec2.Connections
	Engine() IClusterEngine
	Env() *awscdk.ResourceEnvironment
	InstanceEndpoints() *[]Endpoint
	InstanceIdentifiers() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for DatabaseClusterBase
type jsiiProxy_DatabaseClusterBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabaseCluster
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Engine() IClusterEngine {
	var returns IClusterEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDatabaseClusterBase_Override(d DatabaseClusterBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterBase",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseClusterBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseClusterBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new db proxy to this cluster.
func (d *jsiiProxy_DatabaseClusterBase) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseClusterBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseClusterBase) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseClusterBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseClusterBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the given named metric for this DBCluster.
func (d *jsiiProxy_DatabaseClusterBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The percentage of CPU utilization.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of database connections in use.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of deadlocks in the database per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDeadlocks",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of time that the instance has been running, in seconds.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricEngineUptime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available random access memory, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of local storage available, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeLocalStorage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput received from clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkReceiveThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput both received from and transmitted to clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput sent to clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkTransmitThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total amount of backup storage in bytes consumed by all Aurora snapshots outside its backup retention window.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSnapshotStorageUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total amount of backup storage in bytes for which you are billed.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTotalBackupStorageBilled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of storage used by your Aurora DB instance, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeBytesUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of billed read I/O operations from a cluster volume, reported at 5-minute intervals.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeReadIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of write disk I/O operations to the cluster volume, reported at 5-minute intervals.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterBase) MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeWriteIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseClusterBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A database cluster engine.
//
// Provides mapping to the serverless application
// used for secret rotation.
//
// TODO: EXAMPLE
//
type DatabaseClusterEngine interface {
}

// The jsii proxy struct for DatabaseClusterEngine
type jsiiProxy_DatabaseClusterEngine struct {
	_ byte // padding
}

func NewDatabaseClusterEngine() DatabaseClusterEngine {
	_init_.Initialize()

	j := jsiiProxy_DatabaseClusterEngine{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDatabaseClusterEngine_Override(d DatabaseClusterEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		nil, // no parameters
		d,
	)
}

// Creates a new plain Aurora database cluster engine.
func DatabaseClusterEngine_Aurora(props *AuroraClusterEngineProps) IClusterEngine {
	_init_.Initialize()

	var returns IClusterEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"aurora",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Aurora MySQL database cluster engine.
func DatabaseClusterEngine_AuroraMysql(props *AuroraMysqlClusterEngineProps) IClusterEngine {
	_init_.Initialize()

	var returns IClusterEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"auroraMysql",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Aurora PostgreSQL database cluster engine.
func DatabaseClusterEngine_AuroraPostgres(props *AuroraPostgresClusterEngineProps) IClusterEngine {
	_init_.Initialize()

	var returns IClusterEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"auroraPostgres",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func DatabaseClusterEngine_AURORA() IClusterEngine {
	_init_.Initialize()
	var returns IClusterEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"AURORA",
		&returns,
	)
	return returns
}

func DatabaseClusterEngine_AURORA_MYSQL() IClusterEngine {
	_init_.Initialize()
	var returns IClusterEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"AURORA_MYSQL",
		&returns,
	)
	return returns
}

func DatabaseClusterEngine_AURORA_POSTGRESQL() IClusterEngine {
	_init_.Initialize()
	var returns IClusterEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"AURORA_POSTGRESQL",
		&returns,
	)
	return returns
}

// A database cluster restored from a snapshot.
//
// TODO: EXAMPLE
//
type DatabaseClusterFromSnapshot interface {
	DatabaseClusterBase
	ClusterEndpoint() Endpoint
	ClusterIdentifier() *string
	ClusterReadEndpoint() Endpoint
	Connections() awsec2.Connections
	Engine() IClusterEngine
	Env() *awscdk.ResourceEnvironment
	InstanceEndpoints() *[]Endpoint
	InstanceIdentifiers() *[]*string
	NewCfnProps() *CfnDBClusterProps
	Node() constructs.Node
	PhysicalName() *string
	SecurityGroups() *[]awsec2.ISecurityGroup
	Stack() awscdk.Stack
	SubnetGroup() ISubnetGroup
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for DatabaseClusterFromSnapshot
type jsiiProxy_DatabaseClusterFromSnapshot struct {
	jsiiProxy_DatabaseClusterBase
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) Engine() IClusterEngine {
	var returns IClusterEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) NewCfnProps() *CfnDBClusterProps {
	var returns *CfnDBClusterProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterFromSnapshot) SubnetGroup() ISubnetGroup {
	var returns ISubnetGroup
	_jsii_.Get(
		j,
		"subnetGroup",
		&returns,
	)
	return returns
}


func NewDatabaseClusterFromSnapshot(scope constructs.Construct, id *string, props *DatabaseClusterFromSnapshotProps) DatabaseClusterFromSnapshot {
	_init_.Initialize()

	j := jsiiProxy_DatabaseClusterFromSnapshot{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterFromSnapshot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseClusterFromSnapshot_Override(d DatabaseClusterFromSnapshot, scope constructs.Construct, id *string, props *DatabaseClusterFromSnapshotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterFromSnapshot",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseClusterFromSnapshot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterFromSnapshot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseClusterFromSnapshot_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterFromSnapshot",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new db proxy to this cluster.
func (d *jsiiProxy_DatabaseClusterFromSnapshot) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseClusterFromSnapshot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseClusterFromSnapshot) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseClusterFromSnapshot) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseClusterFromSnapshot) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return the given named metric for this DBCluster.
func (d *jsiiProxy_DatabaseClusterFromSnapshot) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The percentage of CPU utilization.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of database connections in use.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of deadlocks in the database per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDeadlocks",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of time that the instance has been running, in seconds.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricEngineUptime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available random access memory, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of local storage available, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeLocalStorage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput received from clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkReceiveThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput both received from and transmitted to clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of network throughput sent to clients by each instance, in bytes per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkTransmitThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total amount of backup storage in bytes consumed by all Aurora snapshots outside its backup retention window.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSnapshotStorageUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total amount of backup storage in bytes for which you are billed.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTotalBackupStorageBilled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of storage used by your Aurora DB instance, in bytes.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeBytesUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of billed read I/O operations from a cluster volume, reported at 5-minute intervals.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeReadIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of write disk I/O operations to the cluster volume, reported at 5-minute intervals.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseClusterFromSnapshot) MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeWriteIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseClusterFromSnapshot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for ``DatabaseClusterFromSnapshot``.
//
// TODO: EXAMPLE
//
type DatabaseClusterFromSnapshotProps struct {
	// What kind of database to start.
	Engine IClusterEngine `json:"engine" yaml:"engine"`
	// Settings for the individual instances that are launched.
	InstanceProps *InstanceProps `json:"instanceProps" yaml:"instanceProps"`
	// The identifier for the DB instance snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot.
	// However, you can use only the ARN to specify a DB instance snapshot.
	SnapshotIdentifier *string `json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The number of seconds to set a cluster's target backtrack window to.
	//
	// This feature is only supported by the Aurora MySQL database engine and
	// cannot be enabled on existing clusters.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html
	//
	BacktrackWindow awscdk.Duration `json:"backtrackWindow" yaml:"backtrackWindow"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
	//
	Backup *BackupProps `json:"backup" yaml:"backup"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Name of a database which is automatically created inside the cluster.
	DefaultDatabaseName *string `json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	InstanceIdentifierBase *string `json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// How many replicas/instances to create.
	//
	// Has to be at least 1.
	Instances *float64 `json:"instances" yaml:"instances"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instances.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instances monitoring.
	MonitoringRole awsiam.IRole `json:"monitoringRole" yaml:"monitoringRole"`
	// Additional parameters to pass to the database engine.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBClusterParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBClusterParameterGroup.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// What port to listen on.
	Port *float64 `json:"port" yaml:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 export.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 import.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// Existing subnet group for the cluster.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
}

// Properties for a new database cluster.
//
// TODO: EXAMPLE
//
type DatabaseClusterProps struct {
	// What kind of database to start.
	Engine IClusterEngine `json:"engine" yaml:"engine"`
	// Settings for the individual instances that are launched.
	InstanceProps *InstanceProps `json:"instanceProps" yaml:"instanceProps"`
	// The number of seconds to set a cluster's target backtrack window to.
	//
	// This feature is only supported by the Aurora MySQL database engine and
	// cannot be enabled on existing clusters.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html
	//
	BacktrackWindow awscdk.Duration `json:"backtrackWindow" yaml:"backtrackWindow"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
	//
	Backup *BackupProps `json:"backup" yaml:"backup"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Whether to copy tags to the snapshot when a snapshot is created.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Credentials for the administrative user.
	Credentials Credentials `json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	DefaultDatabaseName *string `json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	InstanceIdentifierBase *string `json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// How many replicas/instances to create.
	//
	// Has to be at least 1.
	Instances *float64 `json:"instances" yaml:"instances"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instances.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instances monitoring.
	MonitoringRole awsiam.IRole `json:"monitoringRole" yaml:"monitoringRole"`
	// Additional parameters to pass to the database engine.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBClusterParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBClusterParameterGroup.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// What port to listen on.
	Port *float64 `json:"port" yaml:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 export.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 import.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// Whether to enable storage encryption.
	StorageEncrypted *bool `json:"storageEncrypted" yaml:"storageEncrypted"`
	// The KMS key for storage encryption.
	//
	// If specified, {@link storageEncrypted} will be set to `true`.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
	// Existing subnet group for the cluster.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
}

// A database instance.
//
// TODO: EXAMPLE
//
type DatabaseInstance interface {
	DatabaseInstanceBase
	IDatabaseInstance
	Connections() awsec2.Connections
	DbInstanceEndpointAddress() *string
	DbInstanceEndpointPort() *string
	EnableIamAuthentication() *bool
	SetEnableIamAuthentication(val *bool)
	Engine() IInstanceEngine
	Env() *awscdk.ResourceEnvironment
	InstanceArn() *string
	InstanceEndpoint() Endpoint
	InstanceIdentifier() *string
	InstanceType() awsec2.InstanceType
	NewCfnProps() *CfnDBInstanceProps
	Node() constructs.Node
	PhysicalName() *string
	Secret() awssecretsmanager.ISecret
	SourceCfnProps() *CfnDBInstanceProps
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	VpcPlacement() *awsec2.SubnetSelection
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation
	AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	SetLogRetention()
	ToString() *string
}

// The jsii proxy struct for DatabaseInstance
type jsiiProxy_DatabaseInstance struct {
	jsiiProxy_DatabaseInstanceBase
	jsiiProxy_IDatabaseInstance
}

func (j *jsiiProxy_DatabaseInstance) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EnableIamAuthentication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIamAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Engine() IInstanceEngine {
	var returns IInstanceEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceType() awsec2.InstanceType {
	var returns awsec2.InstanceType
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) NewCfnProps() *CfnDBInstanceProps {
	var returns *CfnDBInstanceProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) SourceCfnProps() *CfnDBInstanceProps {
	var returns *CfnDBInstanceProps
	_jsii_.Get(
		j,
		"sourceCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) VpcPlacement() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcPlacement",
		&returns,
	)
	return returns
}


func NewDatabaseInstance(scope constructs.Construct, id *string, props *DatabaseInstanceProps) DatabaseInstance {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseInstance_Override(d DatabaseInstance, scope constructs.Construct, id *string, props *DatabaseInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatabaseInstance) SetEnableIamAuthentication(val *bool) {
	_jsii_.Set(
		j,
		"enableIamAuthentication",
		val,
	)
}

// Import an existing database instance.
func DatabaseInstance_FromDatabaseInstanceAttributes(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) IDatabaseInstance {
	_init_.Initialize()

	var returns IDatabaseInstance

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		"fromDatabaseInstanceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new db proxy to this instance.
func (d *jsiiProxy_DatabaseInstance) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds the multi user rotation to this instance.
func (d *jsiiProxy_DatabaseInstance) AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationMultiUser",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds the single user rotation of the master password to this instance.
func (d *jsiiProxy_DatabaseInstance) AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseInstance) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseInstance) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity connection access to the database.
func (d *jsiiProxy_DatabaseInstance) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this DBInstance.
func (d *jsiiProxy_DatabaseInstance) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The percentage of CPU utilization.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstance) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of database connections in use.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstance) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available random access memory.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstance) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available storage space.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstance) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk write I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstance) MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricReadIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk read I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstance) MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricWriteIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule which triggers for instance events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
func (d *jsiiProxy_DatabaseInstance) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		d,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) SetLogRetention() {
	_jsii_.InvokeVoid(
		d,
		"setLogRetention",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing instance.
//
// TODO: EXAMPLE
//
type DatabaseInstanceAttributes struct {
	// The endpoint address.
	InstanceEndpointAddress *string `json:"instanceEndpointAddress" yaml:"instanceEndpointAddress"`
	// The instance identifier.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The database port.
	Port *float64 `json:"port" yaml:"port"`
	// The security groups of the instance.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The engine of the existing database Instance.
	Engine IInstanceEngine `json:"engine" yaml:"engine"`
}

// A new or imported database instance.
//
// TODO: EXAMPLE
//
type DatabaseInstanceBase interface {
	awscdk.Resource
	IDatabaseInstance
	Connections() awsec2.Connections
	DbInstanceEndpointAddress() *string
	DbInstanceEndpointPort() *string
	EnableIamAuthentication() *bool
	SetEnableIamAuthentication(val *bool)
	Engine() IInstanceEngine
	Env() *awscdk.ResourceEnvironment
	InstanceArn() *string
	InstanceEndpoint() Endpoint
	InstanceIdentifier() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	ToString() *string
}

// The jsii proxy struct for DatabaseInstanceBase
type jsiiProxy_DatabaseInstanceBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabaseInstance
}

func (j *jsiiProxy_DatabaseInstanceBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) EnableIamAuthentication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIamAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) Engine() IInstanceEngine {
	var returns IInstanceEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDatabaseInstanceBase_Override(d DatabaseInstanceBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceBase",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatabaseInstanceBase) SetEnableIamAuthentication(val *bool) {
	_jsii_.Set(
		j,
		"enableIamAuthentication",
		val,
	)
}

// Import an existing database instance.
func DatabaseInstanceBase_FromDatabaseInstanceAttributes(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) IDatabaseInstance {
	_init_.Initialize()

	var returns IDatabaseInstance

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceBase",
		"fromDatabaseInstanceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseInstanceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseInstanceBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new db proxy to this instance.
func (d *jsiiProxy_DatabaseInstanceBase) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseInstanceBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseInstanceBase) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseInstanceBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseInstanceBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity connection access to the database.
func (d *jsiiProxy_DatabaseInstanceBase) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this DBInstance.
func (d *jsiiProxy_DatabaseInstanceBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The percentage of CPU utilization.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceBase) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of database connections in use.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceBase) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available random access memory.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceBase) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available storage space.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceBase) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk write I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceBase) MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricReadIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk read I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceBase) MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricWriteIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule which triggers for instance events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
func (d *jsiiProxy_DatabaseInstanceBase) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		d,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseInstanceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A database instance engine.
//
// Provides mapping to DatabaseEngine used for
// secret rotation.
//
// TODO: EXAMPLE
//
type DatabaseInstanceEngine interface {
}

// The jsii proxy struct for DatabaseInstanceEngine
type jsiiProxy_DatabaseInstanceEngine struct {
	_ byte // padding
}

func NewDatabaseInstanceEngine() DatabaseInstanceEngine {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstanceEngine{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDatabaseInstanceEngine_Override(d DatabaseInstanceEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		nil, // no parameters
		d,
	)
}

// Creates a new MariaDB instance engine.
func DatabaseInstanceEngine_MariaDb(props *MariaDbInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"mariaDb",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new MySQL instance engine.
func DatabaseInstanceEngine_Mysql(props *MySqlInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"mysql",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Enterprise Edition instance engine.
func DatabaseInstanceEngine_OracleEe(props *OracleEeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"oracleEe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Standard Edition 1 instance engine.
func DatabaseInstanceEngine_OracleSe2(props *OracleSe2InstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"oracleSe2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new PostgreSQL instance engine.
func DatabaseInstanceEngine_Postgres(props *PostgresInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"postgres",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Enterprise Edition instance engine.
func DatabaseInstanceEngine_SqlServerEe(props *SqlServerEeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerEe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Express Edition instance engine.
func DatabaseInstanceEngine_SqlServerEx(props *SqlServerExInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerEx",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Standard Edition instance engine.
func DatabaseInstanceEngine_SqlServerSe(props *SqlServerSeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerSe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Web Edition instance engine.
func DatabaseInstanceEngine_SqlServerWeb(props *SqlServerWebInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerWeb",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func DatabaseInstanceEngine_MARIADB() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"MARIADB",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_MYSQL() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"MYSQL",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_EE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"ORACLE_EE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_SE2() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"ORACLE_SE2",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_POSTGRES() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"POSTGRES",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_EE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_EE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_EX() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_EX",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_SE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_SE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_WEB() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_WEB",
		&returns,
	)
	return returns
}

// A database instance restored from a snapshot.
//
// TODO: EXAMPLE
//
type DatabaseInstanceFromSnapshot interface {
	DatabaseInstanceBase
	IDatabaseInstance
	Connections() awsec2.Connections
	DbInstanceEndpointAddress() *string
	DbInstanceEndpointPort() *string
	EnableIamAuthentication() *bool
	SetEnableIamAuthentication(val *bool)
	Engine() IInstanceEngine
	Env() *awscdk.ResourceEnvironment
	InstanceArn() *string
	InstanceEndpoint() Endpoint
	InstanceIdentifier() *string
	InstanceType() awsec2.InstanceType
	NewCfnProps() *CfnDBInstanceProps
	Node() constructs.Node
	PhysicalName() *string
	Secret() awssecretsmanager.ISecret
	SourceCfnProps() *CfnDBInstanceProps
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	VpcPlacement() *awsec2.SubnetSelection
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation
	AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	SetLogRetention()
	ToString() *string
}

// The jsii proxy struct for DatabaseInstanceFromSnapshot
type jsiiProxy_DatabaseInstanceFromSnapshot struct {
	jsiiProxy_DatabaseInstanceBase
	jsiiProxy_IDatabaseInstance
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) EnableIamAuthentication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIamAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) Engine() IInstanceEngine {
	var returns IInstanceEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) InstanceType() awsec2.InstanceType {
	var returns awsec2.InstanceType
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) NewCfnProps() *CfnDBInstanceProps {
	var returns *CfnDBInstanceProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) SourceCfnProps() *CfnDBInstanceProps {
	var returns *CfnDBInstanceProps
	_jsii_.Get(
		j,
		"sourceCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) VpcPlacement() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcPlacement",
		&returns,
	)
	return returns
}


func NewDatabaseInstanceFromSnapshot(scope constructs.Construct, id *string, props *DatabaseInstanceFromSnapshotProps) DatabaseInstanceFromSnapshot {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstanceFromSnapshot{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceFromSnapshot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseInstanceFromSnapshot_Override(d DatabaseInstanceFromSnapshot, scope constructs.Construct, id *string, props *DatabaseInstanceFromSnapshotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceFromSnapshot",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatabaseInstanceFromSnapshot) SetEnableIamAuthentication(val *bool) {
	_jsii_.Set(
		j,
		"enableIamAuthentication",
		val,
	)
}

// Import an existing database instance.
func DatabaseInstanceFromSnapshot_FromDatabaseInstanceAttributes(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) IDatabaseInstance {
	_init_.Initialize()

	var returns IDatabaseInstance

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceFromSnapshot",
		"fromDatabaseInstanceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseInstanceFromSnapshot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceFromSnapshot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseInstanceFromSnapshot_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceFromSnapshot",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new db proxy to this instance.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds the multi user rotation to this instance.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationMultiUser",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds the single user rotation of the master password to this instance.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity connection access to the database.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this DBInstance.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The percentage of CPU utilization.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of database connections in use.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available random access memory.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available storage space.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk write I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricReadIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk read I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricWriteIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule which triggers for instance events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		d,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) SetLogRetention() {
	_jsii_.InvokeVoid(
		d,
		"setLogRetention",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a DatabaseInstanceFromSnapshot.
//
// TODO: EXAMPLE
//
type DatabaseInstanceFromSnapshotProps struct {
	// The VPC network where the DB subnet group should be created.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	Domain *string `json:"domain" yaml:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	DomainRole awsiam.IRole `json:"domainRole" yaml:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	MonitoringRole awsiam.IRole `json:"monitoringRole" yaml:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	MultiAz *bool `json:"multiAz" yaml:"multiAz"`
	// The option group to associate with the instance.
	OptionGroup IOptionGroup `json:"optionGroup" yaml:"optionGroup"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// The port for the instance.
	Port *float64 `json:"port" yaml:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures" yaml:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType StorageType `json:"storageType" yaml:"storageType"`
	// Existing subnet group for the instance.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// The database engine.
	Engine IInstanceEngine `json:"engine" yaml:"engine"`
	// The allocated storage size, specified in gigabytes (GB).
	AllocatedStorage *float64 `json:"allocatedStorage" yaml:"allocatedStorage"`
	// Whether to allow major version upgrades.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// The name of the database.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The name of the compute and memory capacity for the instance.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The license model.
	LicenseModel LicenseModel `json:"licenseModel" yaml:"licenseModel"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The time zone of the instance.
	//
	// This is currently supported only by Microsoft Sql Server.
	Timezone *string `json:"timezone" yaml:"timezone"`
	// The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance.
	//
	// If you're restoring from a shared manual DB
	// snapshot, you must specify the ARN of the snapshot.
	SnapshotIdentifier *string `json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Master user credentials.
	//
	// Note - It is not possible to change the master username for a snapshot;
	// however, it is possible to provide (or generate) a new password.
	Credentials SnapshotCredentials `json:"credentials" yaml:"credentials"`
}

// Construction properties for a DatabaseInstanceNew.
//
// TODO: EXAMPLE
//
type DatabaseInstanceNewProps struct {
	// The VPC network where the DB subnet group should be created.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	Domain *string `json:"domain" yaml:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	DomainRole awsiam.IRole `json:"domainRole" yaml:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	MonitoringRole awsiam.IRole `json:"monitoringRole" yaml:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	MultiAz *bool `json:"multiAz" yaml:"multiAz"`
	// The option group to associate with the instance.
	OptionGroup IOptionGroup `json:"optionGroup" yaml:"optionGroup"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// The port for the instance.
	Port *float64 `json:"port" yaml:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures" yaml:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType StorageType `json:"storageType" yaml:"storageType"`
	// Existing subnet group for the instance.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Construction properties for a DatabaseInstance.
//
// TODO: EXAMPLE
//
type DatabaseInstanceProps struct {
	// The VPC network where the DB subnet group should be created.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	Domain *string `json:"domain" yaml:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	DomainRole awsiam.IRole `json:"domainRole" yaml:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	MonitoringRole awsiam.IRole `json:"monitoringRole" yaml:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	MultiAz *bool `json:"multiAz" yaml:"multiAz"`
	// The option group to associate with the instance.
	OptionGroup IOptionGroup `json:"optionGroup" yaml:"optionGroup"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// The port for the instance.
	Port *float64 `json:"port" yaml:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures" yaml:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType StorageType `json:"storageType" yaml:"storageType"`
	// Existing subnet group for the instance.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// The database engine.
	Engine IInstanceEngine `json:"engine" yaml:"engine"`
	// The allocated storage size, specified in gigabytes (GB).
	AllocatedStorage *float64 `json:"allocatedStorage" yaml:"allocatedStorage"`
	// Whether to allow major version upgrades.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// The name of the database.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The name of the compute and memory capacity for the instance.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The license model.
	LicenseModel LicenseModel `json:"licenseModel" yaml:"licenseModel"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The time zone of the instance.
	//
	// This is currently supported only by Microsoft Sql Server.
	Timezone *string `json:"timezone" yaml:"timezone"`
	// For supported engines, specifies the character set to associate with the DB instance.
	CharacterSetName *string `json:"characterSetName" yaml:"characterSetName"`
	// Credentials for the administrative user.
	Credentials Credentials `json:"credentials" yaml:"credentials"`
	// Indicates whether the DB instance is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted" yaml:"storageEncrypted"`
	// The KMS key that's used to encrypt the DB instance.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
}

// A read replica database instance.
//
// TODO: EXAMPLE
//
type DatabaseInstanceReadReplica interface {
	DatabaseInstanceBase
	IDatabaseInstance
	Connections() awsec2.Connections
	DbInstanceEndpointAddress() *string
	DbInstanceEndpointPort() *string
	EnableIamAuthentication() *bool
	SetEnableIamAuthentication(val *bool)
	Engine() IInstanceEngine
	Env() *awscdk.ResourceEnvironment
	InstanceArn() *string
	InstanceEndpoint() Endpoint
	InstanceIdentifier() *string
	InstanceType() awsec2.InstanceType
	NewCfnProps() *CfnDBInstanceProps
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	VpcPlacement() *awsec2.SubnetSelection
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	SetLogRetention()
	ToString() *string
}

// The jsii proxy struct for DatabaseInstanceReadReplica
type jsiiProxy_DatabaseInstanceReadReplica struct {
	jsiiProxy_DatabaseInstanceBase
	jsiiProxy_IDatabaseInstance
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) EnableIamAuthentication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIamAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) Engine() IInstanceEngine {
	var returns IInstanceEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) InstanceType() awsec2.InstanceType {
	var returns awsec2.InstanceType
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) NewCfnProps() *CfnDBInstanceProps {
	var returns *CfnDBInstanceProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) VpcPlacement() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcPlacement",
		&returns,
	)
	return returns
}


func NewDatabaseInstanceReadReplica(scope constructs.Construct, id *string, props *DatabaseInstanceReadReplicaProps) DatabaseInstanceReadReplica {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstanceReadReplica{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceReadReplica",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseInstanceReadReplica_Override(d DatabaseInstanceReadReplica, scope constructs.Construct, id *string, props *DatabaseInstanceReadReplicaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceReadReplica",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatabaseInstanceReadReplica) SetEnableIamAuthentication(val *bool) {
	_jsii_.Set(
		j,
		"enableIamAuthentication",
		val,
	)
}

// Import an existing database instance.
func DatabaseInstanceReadReplica_FromDatabaseInstanceAttributes(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) IDatabaseInstance {
	_init_.Initialize()

	var returns IDatabaseInstance

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceReadReplica",
		"fromDatabaseInstanceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseInstanceReadReplica_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceReadReplica",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseInstanceReadReplica_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceReadReplica",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new db proxy to this instance.
func (d *jsiiProxy_DatabaseInstanceReadReplica) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseInstanceReadReplica) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseInstanceReadReplica) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseInstanceReadReplica) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseInstanceReadReplica) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity connection access to the database.
func (d *jsiiProxy_DatabaseInstanceReadReplica) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this DBInstance.
func (d *jsiiProxy_DatabaseInstanceReadReplica) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The percentage of CPU utilization.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceReadReplica) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of database connections in use.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceReadReplica) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available random access memory.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceReadReplica) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The amount of available storage space.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceReadReplica) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk write I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceReadReplica) MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricReadIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of disk read I/O operations per second.
//
// Average over 5 minutes
func (d *jsiiProxy_DatabaseInstanceReadReplica) MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricWriteIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule which triggers for instance events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
func (d *jsiiProxy_DatabaseInstanceReadReplica) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		d,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) SetLogRetention() {
	_jsii_.InvokeVoid(
		d,
		"setLogRetention",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseInstanceReadReplica) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a DatabaseInstanceReadReplica.
//
// TODO: EXAMPLE
//
type DatabaseInstanceReadReplicaProps struct {
	// The VPC network where the DB subnet group should be created.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	Domain *string `json:"domain" yaml:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	DomainRole awsiam.IRole `json:"domainRole" yaml:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	MonitoringRole awsiam.IRole `json:"monitoringRole" yaml:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	MultiAz *bool `json:"multiAz" yaml:"multiAz"`
	// The option group to associate with the instance.
	OptionGroup IOptionGroup `json:"optionGroup" yaml:"optionGroup"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// The port for the instance.
	Port *float64 `json:"port" yaml:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures" yaml:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType StorageType `json:"storageType" yaml:"storageType"`
	// Existing subnet group for the instance.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// The name of the compute and memory capacity classes.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The source database instance.
	//
	// Each DB instance can have a limited number of read replicas. For more
	// information, see https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html.
	SourceDatabaseInstance IDatabaseInstance `json:"sourceDatabaseInstance" yaml:"sourceDatabaseInstance"`
	// Indicates whether the DB instance is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted" yaml:"storageEncrypted"`
	// The KMS key that's used to encrypt the DB instance.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
}

// Construction properties for a DatabaseInstanceSource.
//
// TODO: EXAMPLE
//
type DatabaseInstanceSourceProps struct {
	// The VPC network where the DB subnet group should be created.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	Domain *string `json:"domain" yaml:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	DomainRole awsiam.IRole `json:"domainRole" yaml:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	MonitoringRole awsiam.IRole `json:"monitoringRole" yaml:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	MultiAz *bool `json:"multiAz" yaml:"multiAz"`
	// The option group to associate with the instance.
	OptionGroup IOptionGroup `json:"optionGroup" yaml:"optionGroup"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// The port for the instance.
	Port *float64 `json:"port" yaml:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures" yaml:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType StorageType `json:"storageType" yaml:"storageType"`
	// Existing subnet group for the instance.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// The database engine.
	Engine IInstanceEngine `json:"engine" yaml:"engine"`
	// The allocated storage size, specified in gigabytes (GB).
	AllocatedStorage *float64 `json:"allocatedStorage" yaml:"allocatedStorage"`
	// Whether to allow major version upgrades.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// The name of the database.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The name of the compute and memory capacity for the instance.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The license model.
	LicenseModel LicenseModel `json:"licenseModel" yaml:"licenseModel"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The time zone of the instance.
	//
	// This is currently supported only by Microsoft Sql Server.
	Timezone *string `json:"timezone" yaml:"timezone"`
}

// RDS Database Proxy.
//
// TODO: EXAMPLE
//
type DatabaseProxy interface {
	awscdk.Resource
	awsec2.IConnectable
	IDatabaseProxy
	awssecretsmanager.ISecretAttachmentTarget
	Connections() awsec2.Connections
	DbProxyArn() *string
	DbProxyName() *string
	Endpoint() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for DatabaseProxy
type jsiiProxy_DatabaseProxy struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IConnectable
	jsiiProxy_IDatabaseProxy
	internal.Type__awssecretsmanagerISecretAttachmentTarget
}

func (j *jsiiProxy_DatabaseProxy) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxy) DbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxy) DbProxyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxy) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseProxy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDatabaseProxy(scope constructs.Construct, id *string, props *DatabaseProxyProps) DatabaseProxy {
	_init_.Initialize()

	j := jsiiProxy_DatabaseProxy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseProxy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseProxy_Override(d DatabaseProxy, scope constructs.Construct, id *string, props *DatabaseProxyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseProxy",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing database proxy.
func DatabaseProxy_FromDatabaseProxyAttributes(scope constructs.Construct, id *string, attrs *DatabaseProxyAttributes) IDatabaseProxy {
	_init_.Initialize()

	var returns IDatabaseProxy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseProxy",
		"fromDatabaseProxyAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseProxy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseProxy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseProxy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseProxy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseProxy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (d *jsiiProxy_DatabaseProxy) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseProxy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseProxy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseProxy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity connection access to the proxy.
func (d *jsiiProxy_DatabaseProxy) GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee, dbUser},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseProxy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing DB Proxy.
//
// TODO: EXAMPLE
//
type DatabaseProxyAttributes struct {
	// DB Proxy ARN.
	DbProxyArn *string `json:"dbProxyArn" yaml:"dbProxyArn"`
	// DB Proxy Name.
	DbProxyName *string `json:"dbProxyName" yaml:"dbProxyName"`
	// Endpoint.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// The security groups of the instance.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
}

// Options for a new DatabaseProxy.
//
// TODO: EXAMPLE
//
type DatabaseProxyOptions struct {
	// The secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	// One or more secrets are required.
	Secrets *[]awssecretsmanager.ISecret `json:"secrets" yaml:"secrets"`
	// The VPC to associate with the new proxy.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The duration for a proxy to wait for a connection to become available in the connection pool.
	//
	// Only applies when the proxy has opened its maximum number of connections and all connections are busy with client
	// sessions.
	//
	// Value must be between 1 second and 1 hour, or `Duration.seconds(0)` to represent unlimited.
	BorrowTimeout awscdk.Duration `json:"borrowTimeout" yaml:"borrowTimeout"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	// An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens;
	// it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyName *string `json:"dbProxyName" yaml:"dbProxyName"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections.
	// The debug information includes the text of SQL statements that you submit through the proxy.
	// Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive
	// information that appears in the logs.
	DebugLogging *bool `json:"debugLogging" yaml:"debugLogging"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	IamAuth *bool `json:"iamAuth" yaml:"iamAuth"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout awscdk.Duration `json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Typically used with SET statements to make sure that each connection has identical settings such as time zone
	// and character set.
	// For multiple statements, use semicolons as the separator.
	// You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	//
	// not currently supported for PostgreSQL.
	InitQuery *string `json:"initQuery" yaml:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB
	// cluster used by the target group.
	//
	// 1-100
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// A high value enables the proxy to leave a high percentage of idle connections open.
	// A low value causes the proxy to close idle client connections and return the underlying database connections
	// to the connection pool.
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance
	// or Aurora DB cluster used by the target group.
	//
	// between 0 and MaxConnectionsPercent
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTLS *bool `json:"requireTLS" yaml:"requireTLS"`
	// IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// One or more VPC security groups to associate with the new proxy.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Including an item in the list exempts that class of SQL operations from the pinning behavior.
	SessionPinningFilters *[]SessionPinningFilter `json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
	// The subnets used by the proxy.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Construction properties for a DatabaseProxy.
//
// TODO: EXAMPLE
//
type DatabaseProxyProps struct {
	// The secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	// One or more secrets are required.
	Secrets *[]awssecretsmanager.ISecret `json:"secrets" yaml:"secrets"`
	// The VPC to associate with the new proxy.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The duration for a proxy to wait for a connection to become available in the connection pool.
	//
	// Only applies when the proxy has opened its maximum number of connections and all connections are busy with client
	// sessions.
	//
	// Value must be between 1 second and 1 hour, or `Duration.seconds(0)` to represent unlimited.
	BorrowTimeout awscdk.Duration `json:"borrowTimeout" yaml:"borrowTimeout"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	// An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens;
	// it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyName *string `json:"dbProxyName" yaml:"dbProxyName"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections.
	// The debug information includes the text of SQL statements that you submit through the proxy.
	// Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive
	// information that appears in the logs.
	DebugLogging *bool `json:"debugLogging" yaml:"debugLogging"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	IamAuth *bool `json:"iamAuth" yaml:"iamAuth"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout awscdk.Duration `json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Typically used with SET statements to make sure that each connection has identical settings such as time zone
	// and character set.
	// For multiple statements, use semicolons as the separator.
	// You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	//
	// not currently supported for PostgreSQL.
	InitQuery *string `json:"initQuery" yaml:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB
	// cluster used by the target group.
	//
	// 1-100
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// A high value enables the proxy to leave a high percentage of idle connections open.
	// A low value causes the proxy to close idle client connections and return the underlying database connections
	// to the connection pool.
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance
	// or Aurora DB cluster used by the target group.
	//
	// between 0 and MaxConnectionsPercent
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTLS *bool `json:"requireTLS" yaml:"requireTLS"`
	// IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// One or more VPC security groups to associate with the new proxy.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Including an item in the list exempts that class of SQL operations from the pinning behavior.
	SessionPinningFilters *[]SessionPinningFilter `json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
	// The subnets used by the proxy.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// DB proxy target: Instance or Cluster.
	ProxyTarget ProxyTarget `json:"proxyTarget" yaml:"proxyTarget"`
}

// A database secret.
//
// TODO: EXAMPLE
//
type DatabaseSecret interface {
	awssecretsmanager.Secret
	ArnForPolicies() *string
	AutoCreatePolicy() *bool
	EncryptionKey() awskms.IKey
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	SecretArn() *string
	SecretFullArn() *string
	SecretName() *string
	SecretValue() awscdk.SecretValue
	Stack() awscdk.Stack
	AddReplicaRegion(region *string, encryptionKey awskms.IKey)
	AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret
	DenyAccountRootDelete()
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	SecretValueFromJson(jsonField *string) awscdk.SecretValue
	ToString() *string
}

// The jsii proxy struct for DatabaseSecret
type jsiiProxy_DatabaseSecret struct {
	internal.Type__awssecretsmanagerSecret
}

func (j *jsiiProxy_DatabaseSecret) ArnForPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDatabaseSecret(scope constructs.Construct, id *string, props *DatabaseSecretProps) DatabaseSecret {
	_init_.Initialize()

	j := jsiiProxy_DatabaseSecret{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseSecret_Override(d DatabaseSecret, scope constructs.Construct, id *string, props *DatabaseSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing secret into the Stack.
func DatabaseSecret_FromSecretAttributes(scope constructs.Construct, id *string, attrs *awssecretsmanager.SecretAttributes) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a secret by complete ARN.
//
// The complete ARN is the ARN with the Secrets Manager-supplied suffix.
func DatabaseSecret_FromSecretCompleteArn(scope constructs.Construct, id *string, secretCompleteArn *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretCompleteArn",
		[]interface{}{scope, id, secretCompleteArn},
		&returns,
	)

	return returns
}

// Imports a secret by secret name.
//
// A secret with this name must exist in the same account & region.
// Replaces the deprecated `fromSecretName`.
func DatabaseSecret_FromSecretNameV2(scope constructs.Construct, id *string, secretName *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretNameV2",
		[]interface{}{scope, id, secretName},
		&returns,
	)

	return returns
}

// Imports a secret by partial ARN.
//
// The partial ARN is the ARN without the Secrets Manager-supplied suffix.
func DatabaseSecret_FromSecretPartialArn(scope constructs.Construct, id *string, secretPartialArn *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"fromSecretPartialArn",
		[]interface{}{scope, id, secretPartialArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseSecret_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a replica region for the secret.
func (d *jsiiProxy_DatabaseSecret) AddReplicaRegion(region *string, encryptionKey awskms.IKey) {
	_jsii_.InvokeVoid(
		d,
		"addReplicaRegion",
		[]interface{}{region, encryptionKey},
	)
}

// Adds a rotation schedule to the secret.
func (d *jsiiProxy_DatabaseSecret) AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule {
	var returns awssecretsmanager.RotationSchedule

	_jsii_.Invoke(
		d,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a statement to the IAM resource policy associated with this secret.
//
// If this secret was created in this stack, a resource policy will be
// automatically created upon the first call to `addToResourcePolicy`. If
// the secret is imported, then this is a no-op.
func (d *jsiiProxy_DatabaseSecret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		d,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_DatabaseSecret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Attach a target to this secret.
//
// Returns: An attached secret
func (d *jsiiProxy_DatabaseSecret) Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret

	_jsii_.Invoke(
		d,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

// Denies the `DeleteSecret` action to all principals within the current account.
func (d *jsiiProxy_DatabaseSecret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		d,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecret) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_DatabaseSecret) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_DatabaseSecret) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grants reading the secret value to some role.
func (d *jsiiProxy_DatabaseSecret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

// Grants writing and updating the secret value to some role.
func (d *jsiiProxy_DatabaseSecret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
func (d *jsiiProxy_DatabaseSecret) SecretValueFromJson(jsonField *string) awscdk.SecretValue {
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		d,
		"secretValueFromJson",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DatabaseSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a DatabaseSecret.
//
// TODO: EXAMPLE
//
type DatabaseSecretProps struct {
	// The username.
	Username *string `json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Characters to not include in the generated password.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret which will be used to rotate this secret.
	MasterSecret awssecretsmanager.ISecret `json:"masterSecret" yaml:"masterSecret"`
	// Whether to replace this secret when the criteria for the password change.
	//
	// This is achieved by overriding the logical id of the AWS::SecretsManager::Secret
	// with a hash of the options that influence the password generation. This
	// way a new secret will be created when the password is regenerated and the
	// cluster or instance consuming this secret will have its credentials updated.
	ReplaceOnPasswordCriteriaChanges *bool `json:"replaceOnPasswordCriteriaChanges" yaml:"replaceOnPasswordCriteriaChanges"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `json:"replicaRegions" yaml:"replicaRegions"`
	// A name for the secret.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// Connection endpoint of a database cluster or instance.
//
// Consists of a combination of hostname and port.
//
// TODO: EXAMPLE
//
type Endpoint interface {
	Hostname() *string
	Port() *float64
	SocketAddress() *string
}

// The jsii proxy struct for Endpoint
type jsiiProxy_Endpoint struct {
	_ byte // padding
}

func (j *jsiiProxy_Endpoint) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) SocketAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketAddress",
		&returns,
	)
	return returns
}


func NewEndpoint(address *string, port *float64) Endpoint {
	_init_.Initialize()

	j := jsiiProxy_Endpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Endpoint",
		[]interface{}{address, port},
		&j,
	)

	return &j
}

func NewEndpoint_Override(e Endpoint, address *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Endpoint",
		[]interface{}{address, port},
		e,
	)
}

// A version of an engine - for either a cluster, or instance.
//
// TODO: EXAMPLE
//
type EngineVersion struct {
	// The major version of the engine, for example, "5.6". Used in specifying the ParameterGroup family and OptionGroup version for this engine.
	MajorVersion *string `json:"majorVersion" yaml:"majorVersion"`
	// The full version string of the engine, for example, "5.6.mysql_aurora.1.22.1". It can be undefined, which means RDS should use whatever version it deems appropriate for the given engine type.
	FullVersion *string `json:"fullVersion" yaml:"fullVersion"`
}

// The interface representing a database cluster (as opposed to instance) engine.
type IClusterEngine interface {
	IEngine
	// Method called when the engine is used to create a new cluster.
	BindToCluster(scope constructs.Construct, options *ClusterEngineBindOptions) *ClusterEngineConfig
	// The application used by this engine to perform rotation for a multi-user scenario.
	MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The application used by this engine to perform rotation for a single-user scenario.
	SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The log types that are available with this engine type.
	SupportedLogTypes() *[]*string
}

// The jsii proxy for IClusterEngine
type jsiiProxy_IClusterEngine struct {
	jsiiProxy_IEngine
}

func (i *jsiiProxy_IClusterEngine) BindToCluster(scope constructs.Construct, options *ClusterEngineBindOptions) *ClusterEngineConfig {
	var returns *ClusterEngineConfig

	_jsii_.Invoke(
		i,
		"bindToCluster",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IClusterEngine) MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"multiUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterEngine) SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"singleUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterEngine) SupportedLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLogTypes",
		&returns,
	)
	return returns
}

// Create a clustered database with a given number of instances.
type IDatabaseCluster interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Add a new db proxy to this cluster.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	// Return the given named metric for this DBCluster.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of deadlocks in the database per second.
	//
	// Average over 5 minutes
	MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of time that the instance has been running, in seconds.
	//
	// Average over 5 minutes
	MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory, in bytes.
	//
	// Average over 5 minutes
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of local storage available, in bytes.
	//
	// Average over 5 minutes
	MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput received from clients by each instance, in bytes per second.
	//
	// Average over 5 minutes
	MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput both received from and transmitted to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes
	MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput sent to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes
	MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes consumed by all Aurora snapshots outside its backup retention window.
	//
	// Average over 5 minutes
	MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes for which you are billed.
	//
	// Average over 5 minutes
	MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of storage used by your Aurora DB instance, in bytes.
	//
	// Average over 5 minutes
	MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of billed read I/O operations from a cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes
	MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of write disk I/O operations to the cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes
	MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The endpoint to use for read/write operations.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	ClusterReadEndpoint() Endpoint
	// The engine of this Cluster.
	//
	// May be not known for imported Clusters if it wasn't provided explicitly.
	Engine() IClusterEngine
	// Endpoints which address each individual replica.
	InstanceEndpoints() *[]Endpoint
	// Identifiers of the replicas.
	InstanceIdentifiers() *[]*string
}

// The jsii proxy for IDatabaseCluster
type jsiiProxy_IDatabaseCluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
	internal.Type__awssecretsmanagerISecretAttachmentTarget
}

func (i *jsiiProxy_IDatabaseCluster) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		i,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDeadlocks",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEngineUptime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeLocalStorage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkReceiveThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkTransmitThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSnapshotStorageUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalBackupStorageBilled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeBytesUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeReadIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeWriteIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IDatabaseCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Engine() IClusterEngine {
	var returns IClusterEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// A database instance.
type IDatabaseInstance interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Add a new db proxy to this instance.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	// Grant the given identity connection access to the database.
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this DBInstance.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory.
	//
	// Average over 5 minutes
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available storage space.
	//
	// Average over 5 minutes
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk write I/O operations per second.
	//
	// Average over 5 minutes
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk read I/O operations per second.
	//
	// Average over 5 minutes
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CloudWatch event rule which triggers for instance events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The instance endpoint address.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	DbInstanceEndpointPort() *string
	// The engine of this database Instance.
	//
	// May be not known for imported Instances if it wasn't provided explicitly,
	// or for read replicas.
	Engine() IInstanceEngine
	// The instance arn.
	InstanceArn() *string
	// The instance endpoint.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	InstanceIdentifier() *string
}

// The jsii proxy for IDatabaseInstance
type jsiiProxy_IDatabaseInstance struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
	internal.Type__awssecretsmanagerISecretAttachmentTarget
}

func (i *jsiiProxy_IDatabaseInstance) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	var returns DatabaseProxy

	_jsii_.Invoke(
		i,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricReadIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricWriteIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IDatabaseInstance) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Engine() IInstanceEngine {
	var returns IInstanceEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// DB Proxy.
type IDatabaseProxy interface {
	awscdk.IResource
	// Grant the given identity connection access to the proxy.
	GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant
	// DB Proxy ARN.
	DbProxyArn() *string
	// DB Proxy Name.
	DbProxyName() *string
	// Endpoint.
	Endpoint() *string
}

// The jsii proxy for IDatabaseProxy
type jsiiProxy_IDatabaseProxy struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDatabaseProxy) GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee, dbUser},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDatabaseProxy) DbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxy) DbProxyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxy) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

// A common interface for database engines.
//
// Don't implement this interface directly,
// instead implement one of the known sub-interfaces,
// like IClusterEngine and IInstanceEngine.
type IEngine interface {
	// The default name of the master database user if one was not provided explicitly.
	//
	// The global default of 'admin' will be used if this is `undefined`.
	// Note that 'admin' is a reserved word in PostgreSQL and cannot be used.
	DefaultUsername() *string
	// The family this engine belongs to, like "MYSQL", or "POSTGRESQL".
	//
	// This property is used when creating a Database Proxy.
	// Most engines don't belong to any family
	// (and because of that, you can't create Database Proxies for their Clusters or Instances).
	EngineFamily() *string
	// The type of the engine, for example "mysql".
	EngineType() *string
	// The exact version of the engine that is used, for example "5.1.42".
	EngineVersion() *EngineVersion
	// The family to use for ParameterGroups using this engine.
	//
	// This is usually equal to "<engineType><engineMajorVersion>",
	// but can sometimes be a variation of that.
	// You can pass this property when creating new ParameterGroup.
	ParameterGroupFamily() *string
}

// The jsii proxy for IEngine
type jsiiProxy_IEngine struct {
	_ byte // padding
}

func (j *jsiiProxy_IEngine) DefaultUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) EngineFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) EngineVersion() *EngineVersion {
	var returns *EngineVersion
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEngine) ParameterGroupFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamily",
		&returns,
	)
	return returns
}

// Interface representing a database instance (as opposed to cluster) engine.
type IInstanceEngine interface {
	IEngine
	// Method called when the engine is used to create a new instance.
	BindToInstance(scope constructs.Construct, options *InstanceEngineBindOptions) *InstanceEngineConfig
	// The application used by this engine to perform rotation for a multi-user scenario.
	MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The application used by this engine to perform rotation for a single-user scenario.
	SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// Whether this engine supports automatic backups of a read replica instance.
	SupportsReadReplicaBackups() *bool
}

// The jsii proxy for IInstanceEngine
type jsiiProxy_IInstanceEngine struct {
	jsiiProxy_IEngine
}

func (i *jsiiProxy_IInstanceEngine) BindToInstance(scope constructs.Construct, options *InstanceEngineBindOptions) *InstanceEngineConfig {
	var returns *InstanceEngineConfig

	_jsii_.Invoke(
		i,
		"bindToInstance",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IInstanceEngine) MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"multiUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceEngine) SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication {
	var returns awssecretsmanager.SecretRotationApplication
	_jsii_.Get(
		j,
		"singleUserRotationApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceEngine) SupportsReadReplicaBackups() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsReadReplicaBackups",
		&returns,
	)
	return returns
}

// An option group.
type IOptionGroup interface {
	awscdk.IResource
	// Adds a configuration to this OptionGroup.
	//
	// This method is a no-op for an imported OptionGroup.
	//
	// Returns: true if the OptionConfiguration was successfully added.
	AddConfiguration(configuration *OptionConfiguration) *bool
	// The name of the option group.
	OptionGroupName() *string
}

// The jsii proxy for IOptionGroup
type jsiiProxy_IOptionGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOptionGroup) AddConfiguration(configuration *OptionConfiguration) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"addConfiguration",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IOptionGroup) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

// A parameter group.
//
// Represents both a cluster parameter group,
// and an instance parameter group.
type IParameterGroup interface {
	awscdk.IResource
	// Adds a parameter to this group.
	//
	// If this is an imported parameter group,
	// this method does nothing.
	//
	// Returns: true if the parameter was actually added
	// (i.e., this ParameterGroup is not imported),
	// false otherwise
	AddParameter(key *string, value *string) *bool
	// Method called when this Parameter Group is used when defining a database cluster.
	BindToCluster(options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig
	// Method called when this Parameter Group is used when defining a database instance.
	BindToInstance(options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig
}

// The jsii proxy for IParameterGroup
type jsiiProxy_IParameterGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IParameterGroup) AddParameter(key *string, value *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"addParameter",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IParameterGroup) BindToCluster(options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig {
	var returns *ParameterGroupClusterConfig

	_jsii_.Invoke(
		i,
		"bindToCluster",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IParameterGroup) BindToInstance(options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig {
	var returns *ParameterGroupInstanceConfig

	_jsii_.Invoke(
		i,
		"bindToInstance",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Interface representing a serverless database cluster.
type IServerlessCluster interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Grant the given identity to access to the Data API.
	GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the cluster.
	ClusterArn() *string
	// The endpoint to use for read/write operations.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	ClusterReadEndpoint() Endpoint
}

// The jsii proxy for IServerlessCluster
type jsiiProxy_IServerlessCluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
	internal.Type__awssecretsmanagerISecretAttachmentTarget
}

func (i *jsiiProxy_IServerlessCluster) GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDataApiAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IServerlessCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IServerlessCluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Interface for a subnet group.
type ISubnetGroup interface {
	awscdk.IResource
	// The name of the subnet group.
	SubnetGroupName() *string
}

// The jsii proxy for ISubnetGroup
type jsiiProxy_ISubnetGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISubnetGroup) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

// The options passed to {@link IInstanceEngine.bind}.
//
// TODO: EXAMPLE
//
type InstanceEngineBindOptions struct {
	// The Active Directory directory ID to create the DB instance in.
	Domain *string `json:"domain" yaml:"domain"`
	// The option group of the database.
	OptionGroup IOptionGroup `json:"optionGroup" yaml:"optionGroup"`
	// The role used for S3 exporting.
	S3ExportRole awsiam.IRole `json:"s3ExportRole" yaml:"s3ExportRole"`
	// The role used for S3 importing.
	S3ImportRole awsiam.IRole `json:"s3ImportRole" yaml:"s3ImportRole"`
	// The timezone of the database, set by the customer.
	Timezone *string `json:"timezone" yaml:"timezone"`
}

// The type returned from the {@link IInstanceEngine.bind} method.
//
// TODO: EXAMPLE
//
type InstanceEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	Features *InstanceEngineFeatures `json:"features" yaml:"features"`
	// Option group of the database.
	OptionGroup IOptionGroup `json:"optionGroup" yaml:"optionGroup"`
}

// Represents Database Engine features.
//
// TODO: EXAMPLE
//
type InstanceEngineFeatures struct {
	// Feature name for the DB instance that the IAM role to export to S3 bucket is to be associated with.
	S3Export *string `json:"s3Export" yaml:"s3Export"`
	// Feature name for the DB instance that the IAM role to access the S3 bucket for import is to be associated with.
	S3Import *string `json:"s3Import" yaml:"s3Import"`
}

// Instance properties for database instances.
//
// TODO: EXAMPLE
//
type InstanceProps struct {
	// What subnets to run the RDS instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Whether to allow upgrade of major version for the DB instance.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Whether to remove automated backups immediately after the DB instance is deleted for the DB instance.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// What type of instance to start for the replicas.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Security group.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Where to place the instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// The license model.
//
// TODO: EXAMPLE
//
type LicenseModel string

const (
	LicenseModel_LICENSE_INCLUDED LicenseModel = "LICENSE_INCLUDED"
	LicenseModel_BRING_YOUR_OWN_LICENSE LicenseModel = "BRING_YOUR_OWN_LICENSE"
	LicenseModel_GENERAL_PUBLIC_LICENSE LicenseModel = "GENERAL_PUBLIC_LICENSE"
)

// The versions for the MariaDB instance engines (those returned by {@link DatabaseInstanceEngine.mariaDb}).
//
// TODO: EXAMPLE
//
type MariaDbEngineVersion interface {
	MariaDbFullVersion() *string
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

	var returns MariaDbEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"of",
		[]interface{}{mariaDbFullVersion, mariaDbMajorVersion},
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

// Properties for MariaDB instance engines.
//
// Used in {@link DatabaseInstanceEngine.mariaDb}.
//
// TODO: EXAMPLE
//
type MariaDbInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version MariaDbEngineVersion `json:"version" yaml:"version"`
}

// Properties for MySQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.mysql}.
//
// TODO: EXAMPLE
//
type MySqlInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version MysqlEngineVersion `json:"version" yaml:"version"`
}

// The versions for the MySQL instance engines (those returned by {@link DatabaseInstanceEngine.mysql}).
//
// TODO: EXAMPLE
//
type MysqlEngineVersion interface {
	MysqlFullVersion() *string
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

	var returns MysqlEngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"of",
		[]interface{}{mysqlFullVersion, mysqlMajorVersion},
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

// Configuration properties for an option.
//
// TODO: EXAMPLE
//
type OptionConfiguration struct {
	// The name of the option.
	Name *string `json:"name" yaml:"name"`
	// The port number that this option uses.
	//
	// If `port` is specified then `vpc`
	// must also be specified.
	Port *float64 `json:"port" yaml:"port"`
	// Optional list of security groups to use for this option, if `vpc` is specified.
	//
	// If no groups are provided, a default one will be created.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The settings for the option.
	Settings *map[string]*string `json:"settings" yaml:"settings"`
	// The version for the option.
	Version *string `json:"version" yaml:"version"`
	// The VPC where a security group should be created for this option.
	//
	// If `vpc`
	// is specified then `port` must also be specified.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// An option group.
//
// TODO: EXAMPLE
//
type OptionGroup interface {
	awscdk.Resource
	IOptionGroup
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	OptionConnections() *map[string]awsec2.Connections
	OptionGroupName() *string
	PhysicalName() *string
	Stack() awscdk.Stack
	AddConfiguration(configuration *OptionConfiguration) *bool
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for OptionGroup
type jsiiProxy_OptionGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IOptionGroup
}

func (j *jsiiProxy_OptionGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) OptionConnections() *map[string]awsec2.Connections {
	var returns *map[string]awsec2.Connections
	_jsii_.Get(
		j,
		"optionConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewOptionGroup(scope constructs.Construct, id *string, props *OptionGroupProps) OptionGroup {
	_init_.Initialize()

	j := jsiiProxy_OptionGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.OptionGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewOptionGroup_Override(o OptionGroup, scope constructs.Construct, id *string, props *OptionGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.OptionGroup",
		[]interface{}{scope, id, props},
		o,
	)
}

// Import an existing option group.
func OptionGroup_FromOptionGroupName(scope constructs.Construct, id *string, optionGroupName *string) IOptionGroup {
	_init_.Initialize()

	var returns IOptionGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OptionGroup",
		"fromOptionGroupName",
		[]interface{}{scope, id, optionGroupName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OptionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OptionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func OptionGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OptionGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a configuration to this OptionGroup.
//
// This method is a no-op for an imported OptionGroup.
func (o *jsiiProxy_OptionGroup) AddConfiguration(configuration *OptionConfiguration) *bool {
	var returns *bool

	_jsii_.Invoke(
		o,
		"addConfiguration",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (o *jsiiProxy_OptionGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OptionGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (o *jsiiProxy_OptionGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (o *jsiiProxy_OptionGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OptionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for an OptionGroup.
//
// TODO: EXAMPLE
//
type OptionGroupProps struct {
	// The configurations for this option group.
	Configurations *[]*OptionConfiguration `json:"configurations" yaml:"configurations"`
	// The database engine that this option group is associated with.
	Engine IInstanceEngine `json:"engine" yaml:"engine"`
	// A description of the option group.
	Description *string `json:"description" yaml:"description"`
}

// Properties for Oracle Enterprise Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleEe}.
//
// TODO: EXAMPLE
//
type OracleEeInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version OracleEngineVersion `json:"version" yaml:"version"`
}

// The versions for the Oracle instance engines (those returned by {@link DatabaseInstanceEngine.oracleSe2} and {@link DatabaseInstanceEngine.oracleEe}).
//
// TODO: EXAMPLE
//
type OracleEngineVersion interface {
	OracleFullVersion() *string
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

// Properties for Oracle Standard Edition 2 instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleSe2}.
//
// TODO: EXAMPLE
//
type OracleSe2InstanceEngineProps struct {
	// The exact version of the engine to use.
	Version OracleEngineVersion `json:"version" yaml:"version"`
}

// A parameter group.
//
// Represents both a cluster parameter group,
// and an instance parameter group.
//
// TODO: EXAMPLE
//
type ParameterGroup interface {
	awscdk.Resource
	IParameterGroup
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddParameter(key *string, value *string) *bool
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	BindToCluster(_options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig
	BindToInstance(_options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ParameterGroup
type jsiiProxy_ParameterGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IParameterGroup
}

func (j *jsiiProxy_ParameterGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParameterGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewParameterGroup(scope constructs.Construct, id *string, props *ParameterGroupProps) ParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_ParameterGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewParameterGroup_Override(p ParameterGroup, scope constructs.Construct, id *string, props *ParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a parameter group.
func ParameterGroup_FromParameterGroupName(scope constructs.Construct, id *string, parameterGroupName *string) IParameterGroup {
	_init_.Initialize()

	var returns IParameterGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"fromParameterGroupName",
		[]interface{}{scope, id, parameterGroupName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ParameterGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a parameter to this parameter group.
func (p *jsiiProxy_ParameterGroup) AddParameter(key *string, value *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		p,
		"addParameter",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (p *jsiiProxy_ParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Method called when this Parameter Group is used when defining a database cluster.
func (p *jsiiProxy_ParameterGroup) BindToCluster(_options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig {
	var returns *ParameterGroupClusterConfig

	_jsii_.Invoke(
		p,
		"bindToCluster",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

// Method called when this Parameter Group is used when defining a database instance.
func (p *jsiiProxy_ParameterGroup) BindToInstance(_options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig {
	var returns *ParameterGroupInstanceConfig

	_jsii_.Invoke(
		p,
		"bindToInstance",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (p *jsiiProxy_ParameterGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (p *jsiiProxy_ParameterGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_ParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for {@link IParameterGroup.bindToCluster}. Empty for now, but can be extended later.
//
// TODO: EXAMPLE
//
type ParameterGroupClusterBindOptions struct {
}

// The type returned from {@link IParameterGroup.bindToCluster}.
//
// TODO: EXAMPLE
//
type ParameterGroupClusterConfig struct {
	// The name of this parameter group.
	ParameterGroupName *string `json:"parameterGroupName" yaml:"parameterGroupName"`
}

// Options for {@link IParameterGroup.bindToInstance}. Empty for now, but can be extended later.
//
// TODO: EXAMPLE
//
type ParameterGroupInstanceBindOptions struct {
}

// The type returned from {@link IParameterGroup.bindToInstance}.
//
// TODO: EXAMPLE
//
type ParameterGroupInstanceConfig struct {
	// The name of this parameter group.
	ParameterGroupName *string `json:"parameterGroupName" yaml:"parameterGroupName"`
}

// Properties for a parameter group.
//
// TODO: EXAMPLE
//
type ParameterGroupProps struct {
	// The database engine for this parameter group.
	Engine IEngine `json:"engine" yaml:"engine"`
	// Description for this parameter group.
	Description *string `json:"description" yaml:"description"`
	// The parameters in this parameter group.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
}

// The retention period for Performance Insight.
type PerformanceInsightRetention string

const (
	PerformanceInsightRetention_DEFAULT PerformanceInsightRetention = "DEFAULT"
	PerformanceInsightRetention_LONG_TERM PerformanceInsightRetention = "LONG_TERM"
)

// Features supported by the Postgres database engine.
//
// TODO: EXAMPLE
//
type PostgresEngineFeatures struct {
	// Whether this version of the Postgres engine supports the S3 data export feature.
	S3Export *bool `json:"s3Export" yaml:"s3Export"`
	// Whether this version of the Postgres engine supports the S3 data import feature.
	S3Import *bool `json:"s3Import" yaml:"s3Import"`
}

// The versions for the PostgreSQL instance engines (those returned by {@link DatabaseInstanceEngine.postgres}).
//
// TODO: EXAMPLE
//
type PostgresEngineVersion interface {
	PostgresFullVersion() *string
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

// Properties for PostgreSQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.postgres}.
//
// TODO: EXAMPLE
//
type PostgresInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version PostgresEngineVersion `json:"version" yaml:"version"`
}

// The processor features.
//
// TODO: EXAMPLE
//
type ProcessorFeatures struct {
	// The number of CPU core.
	CoreCount *float64 `json:"coreCount" yaml:"coreCount"`
	// The number of threads per core.
	ThreadsPerCore *float64 `json:"threadsPerCore" yaml:"threadsPerCore"`
}

// Proxy target: Instance or Cluster.
//
// A target group is a collection of databases that the proxy can connect to.
// Currently, you can specify only one RDS DB instance or Aurora DB cluster.
//
// TODO: EXAMPLE
//
type ProxyTarget interface {
	Bind(proxy DatabaseProxy) *ProxyTargetConfig
}

// The jsii proxy struct for ProxyTarget
type jsiiProxy_ProxyTarget struct {
	_ byte // padding
}

// From cluster.
func ProxyTarget_FromCluster(cluster IDatabaseCluster) ProxyTarget {
	_init_.Initialize()

	var returns ProxyTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ProxyTarget",
		"fromCluster",
		[]interface{}{cluster},
		&returns,
	)

	return returns
}

// From instance.
func ProxyTarget_FromInstance(instance IDatabaseInstance) ProxyTarget {
	_init_.Initialize()

	var returns ProxyTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ProxyTarget",
		"fromInstance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// Bind this target to the specified database proxy.
func (p *jsiiProxy_ProxyTarget) Bind(proxy DatabaseProxy) *ProxyTargetConfig {
	var returns *ProxyTargetConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{proxy},
		&returns,
	)

	return returns
}

// The result of binding a `ProxyTarget` to a `DatabaseProxy`.
//
// TODO: EXAMPLE
//
type ProxyTargetConfig struct {
	// The engine family of the database instance or cluster this proxy connects with.
	EngineFamily *string `json:"engineFamily" yaml:"engineFamily"`
	// The database clusters to which this proxy connects.
	//
	// Either this or `dbInstances` will be set and the other `undefined`.
	DbClusters *[]IDatabaseCluster `json:"dbClusters" yaml:"dbClusters"`
	// The database instances to which this proxy connects.
	//
	// Either this or `dbClusters` will be set and the other `undefined`.
	DbInstances *[]IDatabaseInstance `json:"dbInstances" yaml:"dbInstances"`
}

// Options to add the multi user rotation.
//
// TODO: EXAMPLE
//
type RotationMultiUserOptions struct {
	// The secret to rotate.
	//
	// It must be a JSON string with the following format:
	// ```
	// {
	//    "engine": <required: database engine>,
	//    "host": <required: instance host name>,
	//    "username": <required: username>,
	//    "password": <required: password>,
	//    "dbname": <optional: database name>,
	//    "port": <optional: if not specified, default port will be used>,
	//    "masterarn": <required: the arn of the master secret which will be used to create users/change passwords>
	// }
	// ```
	Secret awssecretsmanager.ISecret `json:"secret" yaml:"secret"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// The VPC interface endpoint to use for the Secrets Manager API.
	//
	// If you enable private DNS hostnames for your VPC private endpoint (the default), you don't
	// need to specify an endpoint. The standard Secrets Manager DNS hostname the Secrets Manager
	// CLI and SDKs use by default (https://secretsmanager.<region>.amazonaws.com) automatically
	// resolves to your VPC endpoint.
	Endpoint awsec2.IInterfaceVpcEndpoint `json:"endpoint" yaml:"endpoint"`
	// Specifies characters to not include in generated passwords.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// Where to place the rotation Lambda function.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Options to add the multi user rotation.
//
// TODO: EXAMPLE
//
type RotationSingleUserOptions struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// The VPC interface endpoint to use for the Secrets Manager API.
	//
	// If you enable private DNS hostnames for your VPC private endpoint (the default), you don't
	// need to specify an endpoint. The standard Secrets Manager DNS hostname the Secrets Manager
	// CLI and SDKs use by default (https://secretsmanager.<region>.amazonaws.com) automatically
	// resolves to your VPC endpoint.
	Endpoint awsec2.IInterfaceVpcEndpoint `json:"endpoint" yaml:"endpoint"`
	// Specifies characters to not include in generated passwords.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// Where to place the rotation Lambda function.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Create an Aurora Serverless Cluster.
//
// TODO: EXAMPLE
//
type ServerlessCluster interface {
	awscdk.Resource
	IServerlessCluster
	ClusterArn() *string
	ClusterEndpoint() Endpoint
	ClusterIdentifier() *string
	ClusterReadEndpoint() Endpoint
	Connections() awsec2.Connections
	EnableDataApi() *bool
	SetEnableDataApi(val *bool)
	Env() *awscdk.ResourceEnvironment
	NewCfnProps() *CfnDBClusterProps
	Node() constructs.Node
	PhysicalName() *string
	Secret() awssecretsmanager.ISecret
	SecurityGroups() *[]awsec2.ISecurityGroup
	Stack() awscdk.Stack
	AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation
	AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for ServerlessCluster
type jsiiProxy_ServerlessCluster struct {
	internal.Type__awscdkResource
	jsiiProxy_IServerlessCluster
}

func (j *jsiiProxy_ServerlessCluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) EnableDataApi() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDataApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) NewCfnProps() *CfnDBClusterProps {
	var returns *CfnDBClusterProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewServerlessCluster(scope constructs.Construct, id *string, props *ServerlessClusterProps) ServerlessCluster {
	_init_.Initialize()

	j := jsiiProxy_ServerlessCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ServerlessCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewServerlessCluster_Override(s ServerlessCluster, scope constructs.Construct, id *string, props *ServerlessClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ServerlessCluster",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_ServerlessCluster) SetEnableDataApi(val *bool) {
	_jsii_.Set(
		j,
		"enableDataApi",
		val,
	)
}

// Import an existing DatabaseCluster from properties.
func ServerlessCluster_FromServerlessClusterAttributes(scope constructs.Construct, id *string, attrs *ServerlessClusterAttributes) IServerlessCluster {
	_init_.Initialize()

	var returns IServerlessCluster

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ServerlessCluster",
		"fromServerlessClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServerlessCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ServerlessCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ServerlessCluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ServerlessCluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds the multi user rotation to this cluster.
func (s *jsiiProxy_ServerlessCluster) AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		s,
		"addRotationMultiUser",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds the single user rotation of the master password to this cluster.
func (s *jsiiProxy_ServerlessCluster) AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		s,
		"addRotationSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (s *jsiiProxy_ServerlessCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (s *jsiiProxy_ServerlessCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		s,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (s *jsiiProxy_ServerlessCluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (s *jsiiProxy_ServerlessCluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity to access to the Data API, including read access to the secret attached to the cluster if present.
func (s *jsiiProxy_ServerlessCluster) GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantDataApiAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServerlessCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing cluster instance.
//
// TODO: EXAMPLE
//
type ServerlessClusterAttributes struct {
	// Identifier for the cluster.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Cluster endpoint address.
	ClusterEndpointAddress *string `json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// The database port.
	Port *float64 `json:"port" yaml:"port"`
	// Reader endpoint address.
	ReaderEndpointAddress *string `json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The secret attached to the database cluster.
	Secret awssecretsmanager.ISecret `json:"secret" yaml:"secret"`
	// The security groups of the database cluster.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
}

// A Aurora Serverless Cluster restored from a snapshot.
//
// TODO: EXAMPLE
//
type ServerlessClusterFromSnapshot interface {
	awscdk.Resource
	IServerlessCluster
	ClusterArn() *string
	ClusterEndpoint() Endpoint
	ClusterIdentifier() *string
	ClusterReadEndpoint() Endpoint
	Connections() awsec2.Connections
	EnableDataApi() *bool
	SetEnableDataApi(val *bool)
	Env() *awscdk.ResourceEnvironment
	NewCfnProps() *CfnDBClusterProps
	Node() constructs.Node
	PhysicalName() *string
	Secret() awssecretsmanager.ISecret
	SecurityGroups() *[]awsec2.ISecurityGroup
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for ServerlessClusterFromSnapshot
type jsiiProxy_ServerlessClusterFromSnapshot struct {
	internal.Type__awscdkResource
	jsiiProxy_IServerlessCluster
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) EnableDataApi() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDataApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) NewCfnProps() *CfnDBClusterProps {
	var returns *CfnDBClusterProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewServerlessClusterFromSnapshot(scope constructs.Construct, id *string, props *ServerlessClusterFromSnapshotProps) ServerlessClusterFromSnapshot {
	_init_.Initialize()

	j := jsiiProxy_ServerlessClusterFromSnapshot{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ServerlessClusterFromSnapshot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewServerlessClusterFromSnapshot_Override(s ServerlessClusterFromSnapshot, scope constructs.Construct, id *string, props *ServerlessClusterFromSnapshotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ServerlessClusterFromSnapshot",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_ServerlessClusterFromSnapshot) SetEnableDataApi(val *bool) {
	_jsii_.Set(
		j,
		"enableDataApi",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServerlessClusterFromSnapshot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ServerlessClusterFromSnapshot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ServerlessClusterFromSnapshot_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ServerlessClusterFromSnapshot",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (s *jsiiProxy_ServerlessClusterFromSnapshot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
func (s *jsiiProxy_ServerlessClusterFromSnapshot) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		s,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessClusterFromSnapshot) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (s *jsiiProxy_ServerlessClusterFromSnapshot) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (s *jsiiProxy_ServerlessClusterFromSnapshot) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity to access to the Data API, including read access to the secret attached to the cluster if present.
func (s *jsiiProxy_ServerlessClusterFromSnapshot) GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantDataApiAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServerlessClusterFromSnapshot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for ``ServerlessClusterFromSnapshot``.
//
// TODO: EXAMPLE
//
type ServerlessClusterFromSnapshotProps struct {
	// What kind of database to start.
	Engine IClusterEngine `json:"engine" yaml:"engine"`
	// The identifier for the DB instance snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot.
	// However, you can use only the ARN to specify a DB instance snapshot.
	SnapshotIdentifier *string `json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The VPC that this Aurora Serverless cluster has been created in.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Automatic backup retention cannot be disabled on serverless clusters.
	// Must be a value from 1 day to 35 days.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// An optional identifier for the cluster.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Master user credentials.
	//
	// Note - It is not possible to change the master username for a snapshot;
	// however, it is possible to provide (or generate) a new password.
	Credentials SnapshotCredentials `json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	DefaultDatabaseName *string `json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable the Data API.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
	//
	EnableDataApi *bool `json:"enableDataApi" yaml:"enableDataApi"`
	// Additional parameters to pass to the database engine.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Scaling configuration of an Aurora Serverless database cluster.
	Scaling *ServerlessScalingOptions `json:"scaling" yaml:"scaling"`
	// Security group.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Existing subnet group for the cluster.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// Where to place the instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Properties for a new Aurora Serverless Cluster.
//
// TODO: EXAMPLE
//
type ServerlessClusterProps struct {
	// What kind of database to start.
	Engine IClusterEngine `json:"engine" yaml:"engine"`
	// The VPC that this Aurora Serverless cluster has been created in.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Automatic backup retention cannot be disabled on serverless clusters.
	// Must be a value from 1 day to 35 days.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// An optional identifier for the cluster.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Credentials for the administrative user.
	Credentials Credentials `json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	DefaultDatabaseName *string `json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable the Data API.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
	//
	EnableDataApi *bool `json:"enableDataApi" yaml:"enableDataApi"`
	// Additional parameters to pass to the database engine.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Scaling configuration of an Aurora Serverless database cluster.
	Scaling *ServerlessScalingOptions `json:"scaling" yaml:"scaling"`
	// Security group.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The KMS key for storage encryption.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
	// Existing subnet group for the cluster.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// Where to place the instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Options for configuring scaling on an Aurora Serverless cluster.
//
// TODO: EXAMPLE
//
type ServerlessScalingOptions struct {
	// The time before an Aurora Serverless database cluster is paused.
	//
	// A database cluster can be paused only when it is idle (it has no connections).
	// Auto pause time must be between 5 minutes and 1 day.
	//
	// If a DB cluster is paused for more than seven days, the DB cluster might be
	// backed up with a snapshot. In this case, the DB cluster is restored when there
	// is a request to connect to it.
	//
	// Set to 0 to disable
	AutoPause awscdk.Duration `json:"autoPause" yaml:"autoPause"`
	// The maximum capacity for an Aurora Serverless database cluster.
	MaxCapacity AuroraCapacityUnit `json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity for an Aurora Serverless database cluster.
	MinCapacity AuroraCapacityUnit `json:"minCapacity" yaml:"minCapacity"`
}

// SessionPinningFilter.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html#rds-proxy-pinning
//
type SessionPinningFilter interface {
	FilterName() *string
}

// The jsii proxy struct for SessionPinningFilter
type jsiiProxy_SessionPinningFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_SessionPinningFilter) FilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterName",
		&returns,
	)
	return returns
}


// custom filter.
func SessionPinningFilter_Of(filterName *string) SessionPinningFilter {
	_init_.Initialize()

	var returns SessionPinningFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SessionPinningFilter",
		"of",
		[]interface{}{filterName},
		&returns,
	)

	return returns
}

func SessionPinningFilter_EXCLUDE_VARIABLE_SETS() SessionPinningFilter {
	_init_.Initialize()
	var returns SessionPinningFilter
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.SessionPinningFilter",
		"EXCLUDE_VARIABLE_SETS",
		&returns,
	)
	return returns
}

// Credentials to update the password for a ``DatabaseInstanceFromSnapshot``.
//
// TODO: EXAMPLE
//
type SnapshotCredentials interface {
	EncryptionKey() awskms.IKey
	ExcludeCharacters() *string
	GeneratePassword() *bool
	Password() awscdk.SecretValue
	ReplaceOnPasswordCriteriaChanges() *bool
	ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion
	Secret() awssecretsmanager.Secret
	Username() *string
}

// The jsii proxy struct for SnapshotCredentials
type jsiiProxy_SnapshotCredentials struct {
	_ byte // padding
}

func (j *jsiiProxy_SnapshotCredentials) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) GeneratePassword() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"generatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) ReplaceOnPasswordCriteriaChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"replaceOnPasswordCriteriaChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion {
	var returns *[]*awssecretsmanager.ReplicaRegion
	_jsii_.Get(
		j,
		"replicaRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) Secret() awssecretsmanager.Secret {
	var returns awssecretsmanager.Secret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


func NewSnapshotCredentials_Override(s SnapshotCredentials) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		nil, // no parameters
		s,
	)
}

// Generate a new password for the snapshot, using the existing username and an optional encryption key.
//
// Note - The username must match the existing master username of the snapshot.
//
// NOTE: use `fromGeneratedSecret()` for new Clusters and Instances.
func SnapshotCredentials_FromGeneratedPassword(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) SnapshotCredentials {
	_init_.Initialize()

	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromGeneratedPassword",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Generate a new password for the snapshot, using the existing username and an optional encryption key.
//
// The new credentials are stored in Secrets Manager.
//
// Note - The username must match the existing master username of the snapshot.
func SnapshotCredentials_FromGeneratedSecret(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) SnapshotCredentials {
	_init_.Initialize()

	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromGeneratedSecret",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Update the snapshot login with an existing password.
func SnapshotCredentials_FromPassword(password awscdk.SecretValue) SnapshotCredentials {
	_init_.Initialize()

	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromPassword",
		[]interface{}{password},
		&returns,
	)

	return returns
}

// Update the snapshot login with an existing password from a Secret.
//
// The Secret must be a JSON string with a ``password`` field:
// ```
// {
//    ...
//    "password": <required: password>,
// }
// ```
func SnapshotCredentials_FromSecret(secret awssecretsmanager.Secret) SnapshotCredentials {
	_init_.Initialize()

	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromSecret",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

// Options used in the {@link SnapshotCredentials.fromGeneratedPassword} method.
//
// TODO: EXAMPLE
//
type SnapshotCredentialsFromGeneratedPasswordOptions struct {
	// KMS encryption key to encrypt the generated secret.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// The characters to exclude from the generated password.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `json:"replicaRegions" yaml:"replicaRegions"`
}

// Properties for SQL Server Enterprise Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerEe}.
//
// TODO: EXAMPLE
//
type SqlServerEeInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `json:"version" yaml:"version"`
}

// The versions for the SQL Server instance engines (those returned by {@link DatabaseInstanceEngine.sqlServerSe}, {@link DatabaseInstanceEngine.sqlServerEx}, {@link DatabaseInstanceEngine.sqlServerWeb} and {@link DatabaseInstanceEngine.sqlServerEe}).
//
// TODO: EXAMPLE
//
type SqlServerEngineVersion interface {
	SqlServerFullVersion() *string
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

// Properties for SQL Server Express Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerEx}.
//
// TODO: EXAMPLE
//
type SqlServerExInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `json:"version" yaml:"version"`
}

// Properties for SQL Server Standard Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerSe}.
//
// TODO: EXAMPLE
//
type SqlServerSeInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `json:"version" yaml:"version"`
}

// Properties for SQL Server Web Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerWeb}.
//
// TODO: EXAMPLE
//
type SqlServerWebInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `json:"version" yaml:"version"`
}

// The type of storage.
//
// TODO: EXAMPLE
//
type StorageType string

const (
	StorageType_STANDARD StorageType = "STANDARD"
	StorageType_GP2 StorageType = "GP2"
	StorageType_IO1 StorageType = "IO1"
)

// Class for creating a RDS DB subnet group.
//
// TODO: EXAMPLE
//
type SubnetGroup interface {
	awscdk.Resource
	ISubnetGroup
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	SubnetGroupName() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for SubnetGroup
type jsiiProxy_SubnetGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_ISubnetGroup
}

func (j *jsiiProxy_SubnetGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}


func NewSubnetGroup(scope constructs.Construct, id *string, props *SubnetGroupProps) SubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_SubnetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.SubnetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSubnetGroup_Override(s SubnetGroup, scope constructs.Construct, id *string, props *SubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.SubnetGroup",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports an existing subnet group by name.
func SubnetGroup_FromSubnetGroupName(scope constructs.Construct, id *string, subnetGroupName *string) ISubnetGroup {
	_init_.Initialize()

	var returns ISubnetGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SubnetGroup",
		"fromSubnetGroupName",
		[]interface{}{scope, id, subnetGroupName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SubnetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SubnetGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SubnetGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (s *jsiiProxy_SubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SubnetGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (s *jsiiProxy_SubnetGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (s *jsiiProxy_SubnetGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SubnetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a SubnetGroup.
//
// TODO: EXAMPLE
//
type SubnetGroupProps struct {
	// Description of the subnet group.
	Description *string `json:"description" yaml:"description"`
	// The VPC to place the subnet group in.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The removal policy to apply when the subnet group are removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// The name of the subnet group.
	SubnetGroupName *string `json:"subnetGroupName" yaml:"subnetGroupName"`
	// Which subnets within the VPC to associate with this group.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

