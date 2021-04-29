package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

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
// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.architecture
//
// Experimental.
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
// Experimental.
type AuroraClusterEngineProps struct {
	// The version of the Aurora cluster engine.
	// Experimental.
	Version AuroraEngineVersion `json:"version"`
}

// The versions for the Aurora cluster engine (those returned by {@link DatabaseClusterEngine.aurora}).
// Experimental.
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
// Experimental.
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
// Experimental.
type AuroraMysqlClusterEngineProps struct {
	// The version of the Aurora MySQL cluster engine.
	// Experimental.
	Version AuroraMysqlEngineVersion `json:"version"`
}

// The versions for the Aurora MySQL cluster engine (those returned by {@link DatabaseClusterEngine.auroraMysql}).
// Experimental.
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
// Experimental.
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
// Experimental.
type AuroraPostgresClusterEngineProps struct {
	// The version of the Aurora PostgreSQL cluster engine.
	// Experimental.
	Version AuroraPostgresEngineVersion `json:"version"`
}

// Features supported by this version of the Aurora Postgres cluster engine.
// Experimental.
type AuroraPostgresEngineFeatures struct {
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data import feature.
	// Experimental.
	S3Export *bool `json:"s3Export"`
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data import feature.
	// Experimental.
	S3Import *bool `json:"s3Import"`
}

// The versions for the Aurora PostgreSQL cluster engine (those returned by {@link DatabaseClusterEngine.auroraPostgres}).
// Experimental.
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
// Experimental.
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
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
//
// Experimental.
type BackupProps struct {
	// How many days to retain the backup.
	// Experimental.
	Retention awscdk.Duration `json:"retention"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'
	// Experimental.
	PreferredWindow *string `json:"preferredWindow"`
}

// A CloudFormation `AWS::RDS::DBCluster`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDBCluster_DBClusterRoleProperty struct {
	// `CfnDBCluster.DBClusterRoleProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnDBCluster.DBClusterRoleProperty.FeatureName`.
	FeatureName *string `json:"featureName"`
}

type CfnDBCluster_ScalingConfigurationProperty struct {
	// `CfnDBCluster.ScalingConfigurationProperty.AutoPause`.
	AutoPause interface{} `json:"autoPause"`
	// `CfnDBCluster.ScalingConfigurationProperty.MaxCapacity`.
	MaxCapacity *float64 `json:"maxCapacity"`
	// `CfnDBCluster.ScalingConfigurationProperty.MinCapacity`.
	MinCapacity *float64 `json:"minCapacity"`
	// `CfnDBCluster.ScalingConfigurationProperty.SecondsUntilAutoPause`.
	SecondsUntilAutoPause *float64 `json:"secondsUntilAutoPause"`
}

// A CloudFormation `AWS::RDS::DBClusterParameterGroup`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBClusterParameterGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBClusterParameterGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::RDS::DBClusterParameterGroup`.
type CfnDBClusterParameterGroupProps struct {
	// `AWS::RDS::DBClusterParameterGroup.Description`.
	Description *string `json:"description"`
	// `AWS::RDS::DBClusterParameterGroup.Family`.
	Family *string `json:"family"`
	// `AWS::RDS::DBClusterParameterGroup.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `AWS::RDS::DBClusterParameterGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// Properties for defining a `AWS::RDS::DBCluster`.
type CfnDBClusterProps struct {
	// `AWS::RDS::DBCluster.Engine`.
	Engine *string `json:"engine"`
	// `AWS::RDS::DBCluster.AssociatedRoles`.
	AssociatedRoles interface{} `json:"associatedRoles"`
	// `AWS::RDS::DBCluster.AvailabilityZones`.
	AvailabilityZones *[]*string `json:"availabilityZones"`
	// `AWS::RDS::DBCluster.BacktrackWindow`.
	BacktrackWindow *float64 `json:"backtrackWindow"`
	// `AWS::RDS::DBCluster.BackupRetentionPeriod`.
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod"`
	// `AWS::RDS::DBCluster.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `AWS::RDS::DBCluster.DBClusterIdentifier`.
	DbClusterIdentifier *string `json:"dbClusterIdentifier"`
	// `AWS::RDS::DBCluster.DBClusterParameterGroupName`.
	DbClusterParameterGroupName *string `json:"dbClusterParameterGroupName"`
	// `AWS::RDS::DBCluster.DBSubnetGroupName`.
	DbSubnetGroupName *string `json:"dbSubnetGroupName"`
	// `AWS::RDS::DBCluster.DeletionProtection`.
	DeletionProtection interface{} `json:"deletionProtection"`
	// `AWS::RDS::DBCluster.EnableCloudwatchLogsExports`.
	EnableCloudwatchLogsExports *[]*string `json:"enableCloudwatchLogsExports"`
	// `AWS::RDS::DBCluster.EnableHttpEndpoint`.
	EnableHttpEndpoint interface{} `json:"enableHttpEndpoint"`
	// `AWS::RDS::DBCluster.EnableIAMDatabaseAuthentication`.
	EnableIamDatabaseAuthentication interface{} `json:"enableIamDatabaseAuthentication"`
	// `AWS::RDS::DBCluster.EngineMode`.
	EngineMode *string `json:"engineMode"`
	// `AWS::RDS::DBCluster.EngineVersion`.
	EngineVersion *string `json:"engineVersion"`
	// `AWS::RDS::DBCluster.GlobalClusterIdentifier`.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier"`
	// `AWS::RDS::DBCluster.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::RDS::DBCluster.MasterUsername`.
	MasterUsername *string `json:"masterUsername"`
	// `AWS::RDS::DBCluster.MasterUserPassword`.
	MasterUserPassword *string `json:"masterUserPassword"`
	// `AWS::RDS::DBCluster.Port`.
	Port *float64 `json:"port"`
	// `AWS::RDS::DBCluster.PreferredBackupWindow`.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// `AWS::RDS::DBCluster.PreferredMaintenanceWindow`.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// `AWS::RDS::DBCluster.ReplicationSourceIdentifier`.
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier"`
	// `AWS::RDS::DBCluster.RestoreType`.
	RestoreType *string `json:"restoreType"`
	// `AWS::RDS::DBCluster.ScalingConfiguration`.
	ScalingConfiguration interface{} `json:"scalingConfiguration"`
	// `AWS::RDS::DBCluster.SnapshotIdentifier`.
	SnapshotIdentifier *string `json:"snapshotIdentifier"`
	// `AWS::RDS::DBCluster.SourceDBClusterIdentifier`.
	SourceDbClusterIdentifier *string `json:"sourceDbClusterIdentifier"`
	// `AWS::RDS::DBCluster.SourceRegion`.
	SourceRegion *string `json:"sourceRegion"`
	// `AWS::RDS::DBCluster.StorageEncrypted`.
	StorageEncrypted interface{} `json:"storageEncrypted"`
	// `AWS::RDS::DBCluster.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::RDS::DBCluster.UseLatestRestorableTime`.
	UseLatestRestorableTime interface{} `json:"useLatestRestorableTime"`
	// `AWS::RDS::DBCluster.VpcSecurityGroupIds`.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::RDS::DBInstance`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDBInstance_DBInstanceRoleProperty struct {
	// `CfnDBInstance.DBInstanceRoleProperty.FeatureName`.
	FeatureName *string `json:"featureName"`
	// `CfnDBInstance.DBInstanceRoleProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
}

type CfnDBInstance_ProcessorFeatureProperty struct {
	// `CfnDBInstance.ProcessorFeatureProperty.Name`.
	Name *string `json:"name"`
	// `CfnDBInstance.ProcessorFeatureProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::RDS::DBInstance`.
type CfnDBInstanceProps struct {
	// `AWS::RDS::DBInstance.DBInstanceClass`.
	DbInstanceClass *string `json:"dbInstanceClass"`
	// `AWS::RDS::DBInstance.AllocatedStorage`.
	AllocatedStorage *string `json:"allocatedStorage"`
	// `AWS::RDS::DBInstance.AllowMajorVersionUpgrade`.
	AllowMajorVersionUpgrade interface{} `json:"allowMajorVersionUpgrade"`
	// `AWS::RDS::DBInstance.AssociatedRoles`.
	AssociatedRoles interface{} `json:"associatedRoles"`
	// `AWS::RDS::DBInstance.AutoMinorVersionUpgrade`.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade"`
	// `AWS::RDS::DBInstance.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
	// `AWS::RDS::DBInstance.BackupRetentionPeriod`.
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod"`
	// `AWS::RDS::DBInstance.CACertificateIdentifier`.
	CaCertificateIdentifier *string `json:"caCertificateIdentifier"`
	// `AWS::RDS::DBInstance.CharacterSetName`.
	CharacterSetName *string `json:"characterSetName"`
	// `AWS::RDS::DBInstance.CopyTagsToSnapshot`.
	CopyTagsToSnapshot interface{} `json:"copyTagsToSnapshot"`
	// `AWS::RDS::DBInstance.DBClusterIdentifier`.
	DbClusterIdentifier *string `json:"dbClusterIdentifier"`
	// `AWS::RDS::DBInstance.DBInstanceIdentifier`.
	DbInstanceIdentifier *string `json:"dbInstanceIdentifier"`
	// `AWS::RDS::DBInstance.DBName`.
	DbName *string `json:"dbName"`
	// `AWS::RDS::DBInstance.DBParameterGroupName`.
	DbParameterGroupName *string `json:"dbParameterGroupName"`
	// `AWS::RDS::DBInstance.DBSecurityGroups`.
	DbSecurityGroups *[]*string `json:"dbSecurityGroups"`
	// `AWS::RDS::DBInstance.DBSnapshotIdentifier`.
	DbSnapshotIdentifier *string `json:"dbSnapshotIdentifier"`
	// `AWS::RDS::DBInstance.DBSubnetGroupName`.
	DbSubnetGroupName *string `json:"dbSubnetGroupName"`
	// `AWS::RDS::DBInstance.DeleteAutomatedBackups`.
	DeleteAutomatedBackups interface{} `json:"deleteAutomatedBackups"`
	// `AWS::RDS::DBInstance.DeletionProtection`.
	DeletionProtection interface{} `json:"deletionProtection"`
	// `AWS::RDS::DBInstance.Domain`.
	Domain *string `json:"domain"`
	// `AWS::RDS::DBInstance.DomainIAMRoleName`.
	DomainIamRoleName *string `json:"domainIamRoleName"`
	// `AWS::RDS::DBInstance.EnableCloudwatchLogsExports`.
	EnableCloudwatchLogsExports *[]*string `json:"enableCloudwatchLogsExports"`
	// `AWS::RDS::DBInstance.EnableIAMDatabaseAuthentication`.
	EnableIamDatabaseAuthentication interface{} `json:"enableIamDatabaseAuthentication"`
	// `AWS::RDS::DBInstance.EnablePerformanceInsights`.
	EnablePerformanceInsights interface{} `json:"enablePerformanceInsights"`
	// `AWS::RDS::DBInstance.Engine`.
	Engine *string `json:"engine"`
	// `AWS::RDS::DBInstance.EngineVersion`.
	EngineVersion *string `json:"engineVersion"`
	// `AWS::RDS::DBInstance.Iops`.
	Iops *float64 `json:"iops"`
	// `AWS::RDS::DBInstance.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::RDS::DBInstance.LicenseModel`.
	LicenseModel *string `json:"licenseModel"`
	// `AWS::RDS::DBInstance.MasterUsername`.
	MasterUsername *string `json:"masterUsername"`
	// `AWS::RDS::DBInstance.MasterUserPassword`.
	MasterUserPassword *string `json:"masterUserPassword"`
	// `AWS::RDS::DBInstance.MaxAllocatedStorage`.
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage"`
	// `AWS::RDS::DBInstance.MonitoringInterval`.
	MonitoringInterval *float64 `json:"monitoringInterval"`
	// `AWS::RDS::DBInstance.MonitoringRoleArn`.
	MonitoringRoleArn *string `json:"monitoringRoleArn"`
	// `AWS::RDS::DBInstance.MultiAZ`.
	MultiAz interface{} `json:"multiAz"`
	// `AWS::RDS::DBInstance.OptionGroupName`.
	OptionGroupName *string `json:"optionGroupName"`
	// `AWS::RDS::DBInstance.PerformanceInsightsKMSKeyId`.
	PerformanceInsightsKmsKeyId *string `json:"performanceInsightsKmsKeyId"`
	// `AWS::RDS::DBInstance.PerformanceInsightsRetentionPeriod`.
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod"`
	// `AWS::RDS::DBInstance.Port`.
	Port *string `json:"port"`
	// `AWS::RDS::DBInstance.PreferredBackupWindow`.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// `AWS::RDS::DBInstance.PreferredMaintenanceWindow`.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// `AWS::RDS::DBInstance.ProcessorFeatures`.
	ProcessorFeatures interface{} `json:"processorFeatures"`
	// `AWS::RDS::DBInstance.PromotionTier`.
	PromotionTier *float64 `json:"promotionTier"`
	// `AWS::RDS::DBInstance.PubliclyAccessible`.
	PubliclyAccessible interface{} `json:"publiclyAccessible"`
	// `AWS::RDS::DBInstance.SourceDBInstanceIdentifier`.
	SourceDbInstanceIdentifier *string `json:"sourceDbInstanceIdentifier"`
	// `AWS::RDS::DBInstance.SourceRegion`.
	SourceRegion *string `json:"sourceRegion"`
	// `AWS::RDS::DBInstance.StorageEncrypted`.
	StorageEncrypted interface{} `json:"storageEncrypted"`
	// `AWS::RDS::DBInstance.StorageType`.
	StorageType *string `json:"storageType"`
	// `AWS::RDS::DBInstance.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::RDS::DBInstance.Timezone`.
	Timezone *string `json:"timezone"`
	// `AWS::RDS::DBInstance.UseDefaultProcessorFeatures`.
	UseDefaultProcessorFeatures interface{} `json:"useDefaultProcessorFeatures"`
	// `AWS::RDS::DBInstance.VPCSecurityGroups`.
	VpcSecurityGroups *[]*string `json:"vpcSecurityGroups"`
}

// A CloudFormation `AWS::RDS::DBParameterGroup`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBParameterGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBParameterGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBParameterGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::RDS::DBParameterGroup`.
type CfnDBParameterGroupProps struct {
	// `AWS::RDS::DBParameterGroup.Description`.
	Description *string `json:"description"`
	// `AWS::RDS::DBParameterGroup.Family`.
	Family *string `json:"family"`
	// `AWS::RDS::DBParameterGroup.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `AWS::RDS::DBParameterGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::RDS::DBProxy`.
type CfnDBProxy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDbProxyArn() *string
	AttrEndpoint() *string
	AttrVpcId() *string
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

func (j *jsiiProxy_CfnDBProxy) AttrVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcId",
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBProxy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBProxy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBProxy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDBProxy_AuthFormatProperty struct {
	// `CfnDBProxy.AuthFormatProperty.AuthScheme`.
	AuthScheme *string `json:"authScheme"`
	// `CfnDBProxy.AuthFormatProperty.Description`.
	Description *string `json:"description"`
	// `CfnDBProxy.AuthFormatProperty.IAMAuth`.
	IamAuth *string `json:"iamAuth"`
	// `CfnDBProxy.AuthFormatProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDBProxy.AuthFormatProperty.UserName`.
	UserName *string `json:"userName"`
}

type CfnDBProxy_TagFormatProperty struct {
	// `CfnDBProxy.TagFormatProperty.Key`.
	Key *string `json:"key"`
	// `CfnDBProxy.TagFormatProperty.Value`.
	Value *string `json:"value"`
}

// A CloudFormation `AWS::RDS::DBProxyEndpoint`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBProxyEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBProxyEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBProxyEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDBProxyEndpoint_TagFormatProperty struct {
	// `CfnDBProxyEndpoint.TagFormatProperty.Key`.
	Key *string `json:"key"`
	// `CfnDBProxyEndpoint.TagFormatProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::RDS::DBProxyEndpoint`.
type CfnDBProxyEndpointProps struct {
	// `AWS::RDS::DBProxyEndpoint.DBProxyEndpointName`.
	DbProxyEndpointName *string `json:"dbProxyEndpointName"`
	// `AWS::RDS::DBProxyEndpoint.DBProxyName`.
	DbProxyName *string `json:"dbProxyName"`
	// `AWS::RDS::DBProxyEndpoint.VpcSubnetIds`.
	VpcSubnetIds *[]*string `json:"vpcSubnetIds"`
	// `AWS::RDS::DBProxyEndpoint.Tags`.
	Tags *[]*CfnDBProxyEndpoint_TagFormatProperty `json:"tags"`
	// `AWS::RDS::DBProxyEndpoint.TargetRole`.
	TargetRole *string `json:"targetRole"`
	// `AWS::RDS::DBProxyEndpoint.VpcSecurityGroupIds`.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds"`
}

// Properties for defining a `AWS::RDS::DBProxy`.
type CfnDBProxyProps struct {
	// `AWS::RDS::DBProxy.Auth`.
	Auth interface{} `json:"auth"`
	// `AWS::RDS::DBProxy.DBProxyName`.
	DbProxyName *string `json:"dbProxyName"`
	// `AWS::RDS::DBProxy.EngineFamily`.
	EngineFamily *string `json:"engineFamily"`
	// `AWS::RDS::DBProxy.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::RDS::DBProxy.VpcSubnetIds`.
	VpcSubnetIds *[]*string `json:"vpcSubnetIds"`
	// `AWS::RDS::DBProxy.DebugLogging`.
	DebugLogging interface{} `json:"debugLogging"`
	// `AWS::RDS::DBProxy.IdleClientTimeout`.
	IdleClientTimeout *float64 `json:"idleClientTimeout"`
	// `AWS::RDS::DBProxy.RequireTLS`.
	RequireTls interface{} `json:"requireTls"`
	// `AWS::RDS::DBProxy.Tags`.
	Tags *[]*CfnDBProxy_TagFormatProperty `json:"tags"`
	// `AWS::RDS::DBProxy.VpcSecurityGroupIds`.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::RDS::DBProxyTargetGroup`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBProxyTargetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBProxyTargetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDBProxyTargetGroup_ConnectionPoolConfigurationInfoFormatProperty struct {
	// `CfnDBProxyTargetGroup.ConnectionPoolConfigurationInfoFormatProperty.ConnectionBorrowTimeout`.
	ConnectionBorrowTimeout *float64 `json:"connectionBorrowTimeout"`
	// `CfnDBProxyTargetGroup.ConnectionPoolConfigurationInfoFormatProperty.InitQuery`.
	InitQuery *string `json:"initQuery"`
	// `CfnDBProxyTargetGroup.ConnectionPoolConfigurationInfoFormatProperty.MaxConnectionsPercent`.
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent"`
	// `CfnDBProxyTargetGroup.ConnectionPoolConfigurationInfoFormatProperty.MaxIdleConnectionsPercent`.
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent"`
	// `CfnDBProxyTargetGroup.ConnectionPoolConfigurationInfoFormatProperty.SessionPinningFilters`.
	SessionPinningFilters *[]*string `json:"sessionPinningFilters"`
}

// Properties for defining a `AWS::RDS::DBProxyTargetGroup`.
type CfnDBProxyTargetGroupProps struct {
	// `AWS::RDS::DBProxyTargetGroup.DBProxyName`.
	DbProxyName *string `json:"dbProxyName"`
	// `AWS::RDS::DBProxyTargetGroup.TargetGroupName`.
	TargetGroupName *string `json:"targetGroupName"`
	// `AWS::RDS::DBProxyTargetGroup.ConnectionPoolConfigurationInfo`.
	ConnectionPoolConfigurationInfo interface{} `json:"connectionPoolConfigurationInfo"`
	// `AWS::RDS::DBProxyTargetGroup.DBClusterIdentifiers`.
	DbClusterIdentifiers *[]*string `json:"dbClusterIdentifiers"`
	// `AWS::RDS::DBProxyTargetGroup.DBInstanceIdentifiers`.
	DbInstanceIdentifiers *[]*string `json:"dbInstanceIdentifiers"`
}

// A CloudFormation `AWS::RDS::DBSecurityGroup`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBSecurityGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBSecurityGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBSecurityGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDBSecurityGroup_IngressProperty struct {
	// `CfnDBSecurityGroup.IngressProperty.CIDRIP`.
	Cidrip *string `json:"cidrip"`
	// `CfnDBSecurityGroup.IngressProperty.EC2SecurityGroupId`.
	Ec2SecurityGroupId *string `json:"ec2SecurityGroupId"`
	// `CfnDBSecurityGroup.IngressProperty.EC2SecurityGroupName`.
	Ec2SecurityGroupName *string `json:"ec2SecurityGroupName"`
	// `CfnDBSecurityGroup.IngressProperty.EC2SecurityGroupOwnerId`.
	Ec2SecurityGroupOwnerId *string `json:"ec2SecurityGroupOwnerId"`
}

// A CloudFormation `AWS::RDS::DBSecurityGroupIngress`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBSecurityGroupIngress) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::RDS::DBSecurityGroupIngress`.
type CfnDBSecurityGroupIngressProps struct {
	// `AWS::RDS::DBSecurityGroupIngress.DBSecurityGroupName`.
	DbSecurityGroupName *string `json:"dbSecurityGroupName"`
	// `AWS::RDS::DBSecurityGroupIngress.CIDRIP`.
	Cidrip *string `json:"cidrip"`
	// `AWS::RDS::DBSecurityGroupIngress.EC2SecurityGroupId`.
	Ec2SecurityGroupId *string `json:"ec2SecurityGroupId"`
	// `AWS::RDS::DBSecurityGroupIngress.EC2SecurityGroupName`.
	Ec2SecurityGroupName *string `json:"ec2SecurityGroupName"`
	// `AWS::RDS::DBSecurityGroupIngress.EC2SecurityGroupOwnerId`.
	Ec2SecurityGroupOwnerId *string `json:"ec2SecurityGroupOwnerId"`
}

// Properties for defining a `AWS::RDS::DBSecurityGroup`.
type CfnDBSecurityGroupProps struct {
	// `AWS::RDS::DBSecurityGroup.DBSecurityGroupIngress`.
	DbSecurityGroupIngress interface{} `json:"dbSecurityGroupIngress"`
	// `AWS::RDS::DBSecurityGroup.GroupDescription`.
	GroupDescription *string `json:"groupDescription"`
	// `AWS::RDS::DBSecurityGroup.EC2VpcId`.
	Ec2VpcId *string `json:"ec2VpcId"`
	// `AWS::RDS::DBSecurityGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::RDS::DBSubnetGroup`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDBSubnetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDBSubnetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::RDS::DBSubnetGroup`.
type CfnDBSubnetGroupProps struct {
	// `AWS::RDS::DBSubnetGroup.DBSubnetGroupDescription`.
	DbSubnetGroupDescription *string `json:"dbSubnetGroupDescription"`
	// `AWS::RDS::DBSubnetGroup.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `AWS::RDS::DBSubnetGroup.DBSubnetGroupName`.
	DbSubnetGroupName *string `json:"dbSubnetGroupName"`
	// `AWS::RDS::DBSubnetGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::RDS::EventSubscription`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnEventSubscription) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnEventSubscription) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::RDS::EventSubscription`.
type CfnEventSubscriptionProps struct {
	// `AWS::RDS::EventSubscription.SnsTopicArn`.
	SnsTopicArn *string `json:"snsTopicArn"`
	// `AWS::RDS::EventSubscription.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `AWS::RDS::EventSubscription.EventCategories`.
	EventCategories *[]*string `json:"eventCategories"`
	// `AWS::RDS::EventSubscription.SourceIds`.
	SourceIds *[]*string `json:"sourceIds"`
	// `AWS::RDS::EventSubscription.SourceType`.
	SourceType *string `json:"sourceType"`
}

// A CloudFormation `AWS::RDS::GlobalCluster`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnGlobalCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnGlobalCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnGlobalCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::RDS::GlobalCluster`.
type CfnGlobalClusterProps struct {
	// `AWS::RDS::GlobalCluster.DeletionProtection`.
	DeletionProtection interface{} `json:"deletionProtection"`
	// `AWS::RDS::GlobalCluster.Engine`.
	Engine *string `json:"engine"`
	// `AWS::RDS::GlobalCluster.EngineVersion`.
	EngineVersion *string `json:"engineVersion"`
	// `AWS::RDS::GlobalCluster.GlobalClusterIdentifier`.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier"`
	// `AWS::RDS::GlobalCluster.SourceDBClusterIdentifier`.
	SourceDbClusterIdentifier *string `json:"sourceDbClusterIdentifier"`
	// `AWS::RDS::GlobalCluster.StorageEncrypted`.
	StorageEncrypted interface{} `json:"storageEncrypted"`
}

// A CloudFormation `AWS::RDS::OptionGroup`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnOptionGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnOptionGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnOptionGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnOptionGroup_OptionConfigurationProperty struct {
	// `CfnOptionGroup.OptionConfigurationProperty.OptionName`.
	OptionName *string `json:"optionName"`
	// `CfnOptionGroup.OptionConfigurationProperty.DBSecurityGroupMemberships`.
	DbSecurityGroupMemberships *[]*string `json:"dbSecurityGroupMemberships"`
	// `CfnOptionGroup.OptionConfigurationProperty.OptionSettings`.
	OptionSettings interface{} `json:"optionSettings"`
	// `CfnOptionGroup.OptionConfigurationProperty.OptionVersion`.
	OptionVersion *string `json:"optionVersion"`
	// `CfnOptionGroup.OptionConfigurationProperty.Port`.
	Port *float64 `json:"port"`
	// `CfnOptionGroup.OptionConfigurationProperty.VpcSecurityGroupMemberships`.
	VpcSecurityGroupMemberships *[]*string `json:"vpcSecurityGroupMemberships"`
}

type CfnOptionGroup_OptionSettingProperty struct {
	// `CfnOptionGroup.OptionSettingProperty.Name`.
	Name *string `json:"name"`
	// `CfnOptionGroup.OptionSettingProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::RDS::OptionGroup`.
type CfnOptionGroupProps struct {
	// `AWS::RDS::OptionGroup.EngineName`.
	EngineName *string `json:"engineName"`
	// `AWS::RDS::OptionGroup.MajorEngineVersion`.
	MajorEngineVersion *string `json:"majorEngineVersion"`
	// `AWS::RDS::OptionGroup.OptionConfigurations`.
	OptionConfigurations interface{} `json:"optionConfigurations"`
	// `AWS::RDS::OptionGroup.OptionGroupDescription`.
	OptionGroupDescription *string `json:"optionGroupDescription"`
	// `AWS::RDS::OptionGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// The extra options passed to the {@link IClusterEngine.bindToCluster} method.
// Experimental.
type ClusterEngineBindOptions struct {
	// The customer-provided ParameterGroup.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// The role used for S3 exporting.
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// The role used for S3 importing.
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
}

// The type returned from the {@link IClusterEngine.bindToCluster} method.
// Experimental.
type ClusterEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	// Experimental.
	Features *ClusterEngineFeatures `json:"features"`
	// The ParameterGroup to use for the cluster.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// The port to use for this cluster, unless the customer specified the port directly.
	// Experimental.
	Port *float64 `json:"port"`
}

// Represents Database Engine features.
// Experimental.
type ClusterEngineFeatures struct {
	// Feature name for the DB instance that the IAM role to export to S3 bucket is to be associated with.
	// Experimental.
	S3Export *string `json:"s3Export"`
	// Feature name for the DB instance that the IAM role to access the S3 bucket for import is to be associated with.
	// Experimental.
	S3Import *string `json:"s3Import"`
}

// Username and password combination.
// Experimental.
type Credentials interface {
	EncryptionKey() awskms.IKey
	ExcludeCharacters() *string
	Password() awscdk.SecretValue
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


// Experimental.
func NewCredentials_Override(c Credentials) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Credentials",
		nil, // no parameters
		c,
	)
}

// Creates Credentials with a password generated and stored in Secrets Manager.
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Base options for creating Credentials.
// Experimental.
type CredentialsBaseOptions struct {
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The characters to exclude from the generated password.
	//
	// Has no effect if {@link password} has been provided.
	// Experimental.
	ExcludeCharacters *string `json:"excludeCharacters"`
	// The name of the secret.
	// Experimental.
	SecretName *string `json:"secretName"`
}

// Create a clustered database with a given number of instances.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewDatabaseCluster_Override(d DatabaseCluster, scope constructs.Construct, id *string, props *DatabaseClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseCluster",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing DatabaseCluster from properties.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type DatabaseClusterAttributes struct {
	// Identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `json:"clusterIdentifier"`
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `json:"clusterEndpointAddress"`
	// The engine of the existing Cluster.
	// Experimental.
	Engine IClusterEngine `json:"engine"`
	// Endpoint addresses of individual instances.
	// Experimental.
	InstanceEndpointAddresses *[]*string `json:"instanceEndpointAddresses"`
	// Identifier for the instances.
	// Experimental.
	InstanceIdentifiers *[]*string `json:"instanceIdentifiers"`
	// The database port.
	// Experimental.
	Port *float64 `json:"port"`
	// Reader endpoint address.
	// Experimental.
	ReaderEndpointAddress *string `json:"readerEndpointAddress"`
	// The security groups of the database cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
}

// A new or imported clustered database.
// Experimental.
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


// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseClusterBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type DatabaseClusterEngine interface {
}

// The jsii proxy struct for DatabaseClusterEngine
type jsiiProxy_DatabaseClusterEngine struct {
	_ byte // padding
}

// Experimental.
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

// Experimental.
func NewDatabaseClusterEngine_Override(d DatabaseClusterEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		nil, // no parameters
		d,
	)
}

// Creates a new plain Aurora database cluster engine.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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


// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseClusterFromSnapshot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type DatabaseClusterFromSnapshotProps struct {
	// What kind of database to start.
	// Experimental.
	Engine IClusterEngine `json:"engine"`
	// Settings for the individual instances that are launched.
	// Experimental.
	InstanceProps *InstanceProps `json:"instanceProps"`
	// The identifier for the DB instance snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot.
	// However, you can use only the ARN to specify a DB instance snapshot.
	// Experimental.
	SnapshotIdentifier *string `json:"snapshotIdentifier"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
	//
	// Experimental.
	Backup *BackupProps `json:"backup"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `json:"clusterIdentifier"`
	// Name of a database which is automatically created inside the cluster.
	// Experimental.
	DefaultDatabaseName *string `json:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Experimental.
	InstanceIdentifierBase *string `json:"instanceIdentifierBase"`
	// How many replicas/instances to create.
	//
	// Has to be at least 1.
	// Experimental.
	Instances *float64 `json:"instances"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instances.
	// Experimental.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval"`
	// Role that will be used to manage DB instances monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `json:"monitoringRole"`
	// Additional parameters to pass to the database engine.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// What port to listen on.
	// Experimental.
	Port *float64 `json:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// S3 buckets that you want to load data into. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 export.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// S3 buckets that you want to load data from. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 import.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// Existing subnet group for the cluster.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
}

// Properties for a new database cluster.
// Experimental.
type DatabaseClusterProps struct {
	// What kind of database to start.
	// Experimental.
	Engine IClusterEngine `json:"engine"`
	// Settings for the individual instances that are launched.
	// Experimental.
	InstanceProps *InstanceProps `json:"instanceProps"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
	//
	// Experimental.
	Backup *BackupProps `json:"backup"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `json:"clusterIdentifier"`
	// Credentials for the administrative user.
	// Experimental.
	Credentials Credentials `json:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	// Experimental.
	DefaultDatabaseName *string `json:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Experimental.
	InstanceIdentifierBase *string `json:"instanceIdentifierBase"`
	// How many replicas/instances to create.
	//
	// Has to be at least 1.
	// Experimental.
	Instances *float64 `json:"instances"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instances.
	// Experimental.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval"`
	// Role that will be used to manage DB instances monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `json:"monitoringRole"`
	// Additional parameters to pass to the database engine.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// What port to listen on.
	// Experimental.
	Port *float64 `json:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// S3 buckets that you want to load data into. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 export.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// S3 buckets that you want to load data from. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 import.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For MySQL:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// Whether to enable storage encryption.
	// Experimental.
	StorageEncrypted *bool `json:"storageEncrypted"`
	// The KMS key for storage encryption.
	//
	// If specified, {@link storageEncrypted} will be set to `true`.
	// Experimental.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey"`
	// Existing subnet group for the cluster.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
}

// A database instance.
// Experimental.
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


// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (d *jsiiProxy_DatabaseInstance) SetLogRetention() {
	_jsii_.InvokeVoid(
		d,
		"setLogRetention",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
// Experimental.
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
// Experimental.
type DatabaseInstanceAttributes struct {
	// The endpoint address.
	// Experimental.
	InstanceEndpointAddress *string `json:"instanceEndpointAddress"`
	// The instance identifier.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier"`
	// The database port.
	// Experimental.
	Port *float64 `json:"port"`
	// The security groups of the instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The engine of the existing database Instance.
	// Experimental.
	Engine IInstanceEngine `json:"engine"`
}

// A new or imported database instance.
// Experimental.
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


// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseInstanceBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type DatabaseInstanceEngine interface {
}

// The jsii proxy struct for DatabaseInstanceEngine
type jsiiProxy_DatabaseInstanceEngine struct {
	_ byte // padding
}

// Experimental.
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

// Experimental.
func NewDatabaseInstanceEngine_Override(d DatabaseInstanceEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		nil, // no parameters
		d,
	)
}

// Creates a new MariaDB instance engine.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// A database instance restored from a snapshot.
// Experimental.
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


// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (d *jsiiProxy_DatabaseInstanceFromSnapshot) SetLogRetention() {
	_jsii_.InvokeVoid(
		d,
		"setLogRetention",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
// Experimental.
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
// Experimental.
type DatabaseInstanceFromSnapshotProps struct {
	// The VPC network where the DB subnet group should be created.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	// Experimental.
	BackupRetention awscdk.Duration `json:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	// Experimental.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	// Experimental.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	// Experimental.
	Domain *string `json:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	// Experimental.
	DomainRole awsiam.IRole `json:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	// Experimental.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `json:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	// Experimental.
	Iops *float64 `json:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	// Experimental.
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	// Experimental.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `json:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	// Experimental.
	MultiAz *bool `json:"multiAz"`
	// The option group to associate with the instance.
	// Experimental.
	OptionGroup IOptionGroup `json:"optionGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Experimental.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Experimental.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention"`
	// The port for the instance.
	// Experimental.
	Port *float64 `json:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	// Experimental.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	// Experimental.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	// Experimental.
	PubliclyAccessible *bool `json:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	// Experimental.
	StorageType StorageType `json:"storageType"`
	// Existing subnet group for the instance.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The database engine.
	// Experimental.
	Engine IInstanceEngine `json:"engine"`
	// The allocated storage size, specified in gigabytes (GB).
	// Experimental.
	AllocatedStorage *float64 `json:"allocatedStorage"`
	// Whether to allow major version upgrades.
	// Experimental.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade"`
	// The name of the database.
	// Experimental.
	DatabaseName *string `json:"databaseName"`
	// The name of the compute and memory capacity for the instance.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The license model.
	// Experimental.
	LicenseModel LicenseModel `json:"licenseModel"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// The time zone of the instance.
	//
	// This is currently supported only by Microsoft Sql Server.
	// Experimental.
	Timezone *string `json:"timezone"`
	// The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance.
	//
	// If you're restoring from a shared manual DB
	// snapshot, you must specify the ARN of the snapshot.
	// Experimental.
	SnapshotIdentifier *string `json:"snapshotIdentifier"`
	// Master user credentials.
	//
	// Note - It is not possible to change the master username for a snapshot;
	// however, it is possible to provide (or generate) a new password.
	// Experimental.
	Credentials SnapshotCredentials `json:"credentials"`
}

// Construction properties for a DatabaseInstanceNew.
// Experimental.
type DatabaseInstanceNewProps struct {
	// The VPC network where the DB subnet group should be created.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	// Experimental.
	BackupRetention awscdk.Duration `json:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	// Experimental.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	// Experimental.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	// Experimental.
	Domain *string `json:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	// Experimental.
	DomainRole awsiam.IRole `json:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	// Experimental.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `json:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	// Experimental.
	Iops *float64 `json:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	// Experimental.
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	// Experimental.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `json:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	// Experimental.
	MultiAz *bool `json:"multiAz"`
	// The option group to associate with the instance.
	// Experimental.
	OptionGroup IOptionGroup `json:"optionGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Experimental.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Experimental.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention"`
	// The port for the instance.
	// Experimental.
	Port *float64 `json:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	// Experimental.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	// Experimental.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	// Experimental.
	PubliclyAccessible *bool `json:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	// Experimental.
	StorageType StorageType `json:"storageType"`
	// Existing subnet group for the instance.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// Construction properties for a DatabaseInstance.
// Experimental.
type DatabaseInstanceProps struct {
	// The VPC network where the DB subnet group should be created.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	// Experimental.
	BackupRetention awscdk.Duration `json:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	// Experimental.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	// Experimental.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	// Experimental.
	Domain *string `json:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	// Experimental.
	DomainRole awsiam.IRole `json:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	// Experimental.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `json:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	// Experimental.
	Iops *float64 `json:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	// Experimental.
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	// Experimental.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `json:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	// Experimental.
	MultiAz *bool `json:"multiAz"`
	// The option group to associate with the instance.
	// Experimental.
	OptionGroup IOptionGroup `json:"optionGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Experimental.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Experimental.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention"`
	// The port for the instance.
	// Experimental.
	Port *float64 `json:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	// Experimental.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	// Experimental.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	// Experimental.
	PubliclyAccessible *bool `json:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	// Experimental.
	StorageType StorageType `json:"storageType"`
	// Existing subnet group for the instance.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The database engine.
	// Experimental.
	Engine IInstanceEngine `json:"engine"`
	// The allocated storage size, specified in gigabytes (GB).
	// Experimental.
	AllocatedStorage *float64 `json:"allocatedStorage"`
	// Whether to allow major version upgrades.
	// Experimental.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade"`
	// The name of the database.
	// Experimental.
	DatabaseName *string `json:"databaseName"`
	// The name of the compute and memory capacity for the instance.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The license model.
	// Experimental.
	LicenseModel LicenseModel `json:"licenseModel"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// The time zone of the instance.
	//
	// This is currently supported only by Microsoft Sql Server.
	// Experimental.
	Timezone *string `json:"timezone"`
	// For supported engines, specifies the character set to associate with the DB instance.
	// Experimental.
	CharacterSetName *string `json:"characterSetName"`
	// Credentials for the administrative user.
	// Experimental.
	Credentials Credentials `json:"credentials"`
	// Indicates whether the DB instance is encrypted.
	// Experimental.
	StorageEncrypted *bool `json:"storageEncrypted"`
	// The KMS key that's used to encrypt the DB instance.
	// Experimental.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey"`
}

// A read replica database instance.
// Experimental.
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


// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseInstanceReadReplica) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (d *jsiiProxy_DatabaseInstanceReadReplica) SetLogRetention() {
	_jsii_.InvokeVoid(
		d,
		"setLogRetention",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
// Experimental.
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
// Experimental.
type DatabaseInstanceReadReplicaProps struct {
	// The VPC network where the DB subnet group should be created.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	// Experimental.
	BackupRetention awscdk.Duration `json:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	// Experimental.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	// Experimental.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	// Experimental.
	Domain *string `json:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	// Experimental.
	DomainRole awsiam.IRole `json:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	// Experimental.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `json:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	// Experimental.
	Iops *float64 `json:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	// Experimental.
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	// Experimental.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `json:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	// Experimental.
	MultiAz *bool `json:"multiAz"`
	// The option group to associate with the instance.
	// Experimental.
	OptionGroup IOptionGroup `json:"optionGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Experimental.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Experimental.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention"`
	// The port for the instance.
	// Experimental.
	Port *float64 `json:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	// Experimental.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	// Experimental.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	// Experimental.
	PubliclyAccessible *bool `json:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	// Experimental.
	StorageType StorageType `json:"storageType"`
	// Existing subnet group for the instance.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The name of the compute and memory capacity classes.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The source database instance.
	//
	// Each DB instance can have a limited number of read replicas. For more
	// information, see https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html.
	// Experimental.
	SourceDatabaseInstance IDatabaseInstance `json:"sourceDatabaseInstance"`
	// Indicates whether the DB instance is encrypted.
	// Experimental.
	StorageEncrypted *bool `json:"storageEncrypted"`
	// The KMS key that's used to encrypt the DB instance.
	// Experimental.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey"`
}

// Construction properties for a DatabaseInstanceSource.
// Experimental.
type DatabaseInstanceSourceProps struct {
	// The VPC network where the DB subnet group should be created.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	// Experimental.
	BackupRetention awscdk.Duration `json:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `json:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `json:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `json:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	// Experimental.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	// Experimental.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	// Experimental.
	Domain *string `json:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	// Experimental.
	DomainRole awsiam.IRole `json:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	// Experimental.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `json:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	// Experimental.
	Iops *float64 `json:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	// Experimental.
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	// Experimental.
	MonitoringInterval awscdk.Duration `json:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `json:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	// Experimental.
	MultiAz *bool `json:"multiAz"`
	// The option group to associate with the instance.
	// Experimental.
	OptionGroup IOptionGroup `json:"optionGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Experimental.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Experimental.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention"`
	// The port for the instance.
	// Experimental.
	Port *float64 `json:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	// Experimental.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	// Experimental.
	ProcessorFeatures *ProcessorFeatures `json:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	// Experimental.
	PubliclyAccessible *bool `json:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `json:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This feature is only supported by the Microsoft SQL Server and Oracle engines.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `json:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	// Experimental.
	StorageType StorageType `json:"storageType"`
	// Existing subnet group for the instance.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The database engine.
	// Experimental.
	Engine IInstanceEngine `json:"engine"`
	// The allocated storage size, specified in gigabytes (GB).
	// Experimental.
	AllocatedStorage *float64 `json:"allocatedStorage"`
	// Whether to allow major version upgrades.
	// Experimental.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade"`
	// The name of the database.
	// Experimental.
	DatabaseName *string `json:"databaseName"`
	// The name of the compute and memory capacity for the instance.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The license model.
	// Experimental.
	LicenseModel LicenseModel `json:"licenseModel"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// The time zone of the instance.
	//
	// This is currently supported only by Microsoft Sql Server.
	// Experimental.
	Timezone *string `json:"timezone"`
}

// RDS Database Proxy.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewDatabaseProxy_Override(d DatabaseProxy, scope constructs.Construct, id *string, props *DatabaseProxyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseProxy",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing database proxy.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseProxy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type DatabaseProxyAttributes struct {
	// DB Proxy ARN.
	// Experimental.
	DbProxyArn *string `json:"dbProxyArn"`
	// DB Proxy Name.
	// Experimental.
	DbProxyName *string `json:"dbProxyName"`
	// Endpoint.
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// The security groups of the instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
}

// Options for a new DatabaseProxy.
// Experimental.
type DatabaseProxyOptions struct {
	// The secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	// One or more secrets are required.
	// Experimental.
	Secrets *[]awssecretsmanager.ISecret `json:"secrets"`
	// The VPC to associate with the new proxy.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The duration for a proxy to wait for a connection to become available in the connection pool.
	//
	// Only applies when the proxy has opened its maximum number of connections and all connections are busy with client
	// sessions.
	//
	// Value must be between 1 second and 1 hour, or `Duration.seconds(0)` to represent unlimited.
	// Experimental.
	BorrowTimeout awscdk.Duration `json:"borrowTimeout"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	// An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens;
	// it can't end with a hyphen or contain two consecutive hyphens.
	// Experimental.
	DbProxyName *string `json:"dbProxyName"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections.
	// The debug information includes the text of SQL statements that you submit through the proxy.
	// Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive
	// information that appears in the logs.
	// Experimental.
	DebugLogging *bool `json:"debugLogging"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	// Experimental.
	IamAuth *bool `json:"iamAuth"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	// Experimental.
	IdleClientTimeout awscdk.Duration `json:"idleClientTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Typically used with SET statements to make sure that each connection has identical settings such as time zone
	// and character set.
	// For multiple statements, use semicolons as the separator.
	// You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	//
	// not currently supported for PostgreSQL.
	// Experimental.
	InitQuery *string `json:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB
	// cluster used by the target group.
	//
	// 1-100
	// Experimental.
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// A high value enables the proxy to leave a high percentage of idle connections open.
	// A low value causes the proxy to close idle client connections and return the underlying database connections
	// to the connection pool.
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance
	// or Aurora DB cluster used by the target group.
	//
	// between 0 and MaxConnectionsPercent
	// Experimental.
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	// Experimental.
	RequireTLS *bool `json:"requireTLS"`
	// IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// One or more VPC security groups to associate with the new proxy.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Including an item in the list exempts that class of SQL operations from the pinning behavior.
	// Experimental.
	SessionPinningFilters *[]SessionPinningFilter `json:"sessionPinningFilters"`
	// The subnets used by the proxy.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// Construction properties for a DatabaseProxy.
// Experimental.
type DatabaseProxyProps struct {
	// The secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	// One or more secrets are required.
	// Experimental.
	Secrets *[]awssecretsmanager.ISecret `json:"secrets"`
	// The VPC to associate with the new proxy.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The duration for a proxy to wait for a connection to become available in the connection pool.
	//
	// Only applies when the proxy has opened its maximum number of connections and all connections are busy with client
	// sessions.
	//
	// Value must be between 1 second and 1 hour, or `Duration.seconds(0)` to represent unlimited.
	// Experimental.
	BorrowTimeout awscdk.Duration `json:"borrowTimeout"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	// An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens;
	// it can't end with a hyphen or contain two consecutive hyphens.
	// Experimental.
	DbProxyName *string `json:"dbProxyName"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections.
	// The debug information includes the text of SQL statements that you submit through the proxy.
	// Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive
	// information that appears in the logs.
	// Experimental.
	DebugLogging *bool `json:"debugLogging"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	// Experimental.
	IamAuth *bool `json:"iamAuth"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	// Experimental.
	IdleClientTimeout awscdk.Duration `json:"idleClientTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Typically used with SET statements to make sure that each connection has identical settings such as time zone
	// and character set.
	// For multiple statements, use semicolons as the separator.
	// You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	//
	// not currently supported for PostgreSQL.
	// Experimental.
	InitQuery *string `json:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB
	// cluster used by the target group.
	//
	// 1-100
	// Experimental.
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// A high value enables the proxy to leave a high percentage of idle connections open.
	// A low value causes the proxy to close idle client connections and return the underlying database connections
	// to the connection pool.
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance
	// or Aurora DB cluster used by the target group.
	//
	// between 0 and MaxConnectionsPercent
	// Experimental.
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	// Experimental.
	RequireTLS *bool `json:"requireTLS"`
	// IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// One or more VPC security groups to associate with the new proxy.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Including an item in the list exempts that class of SQL operations from the pinning behavior.
	// Experimental.
	SessionPinningFilters *[]SessionPinningFilter `json:"sessionPinningFilters"`
	// The subnets used by the proxy.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// DB proxy target: Instance or Cluster.
	// Experimental.
	ProxyTarget ProxyTarget `json:"proxyTarget"`
}

// A database secret.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewDatabaseSecret_Override(d DatabaseSecret, scope constructs.Construct, id *string, props *DatabaseSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing secret into the Stack.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (d *jsiiProxy_DatabaseSecret) AddReplicaRegion(region *string, encryptionKey awskms.IKey) {
	_jsii_.InvokeVoid(
		d,
		"addReplicaRegion",
		[]interface{}{region, encryptionKey},
	)
}

// Adds a rotation schedule to the secret.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
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
// Experimental.
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
// Experimental.
func (d *jsiiProxy_DatabaseSecret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		d,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type DatabaseSecretProps struct {
	// The username.
	// Experimental.
	Username *string `json:"username"`
	// The KMS key to use to encrypt the secret.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// Characters to not include in the generated password.
	// Experimental.
	ExcludeCharacters *string `json:"excludeCharacters"`
	// The master secret which will be used to rotate this secret.
	// Experimental.
	MasterSecret awssecretsmanager.ISecret `json:"masterSecret"`
	// Whether to replace this secret when the criteria for the password change.
	//
	// This is achieved by overriding the logical id of the AWS::SecretsManager::Secret
	// with a hash of the options that influence the password generation. This
	// way a new secret will be created when the password is regenerated and the
	// cluster or instance consuming this secret will have its credentials updated.
	// Experimental.
	ReplaceOnPasswordCriteriaChanges *bool `json:"replaceOnPasswordCriteriaChanges"`
	// A name for the secret.
	// Experimental.
	SecretName *string `json:"secretName"`
}

// Connection endpoint of a database cluster or instance.
//
// Consists of a combination of hostname and port.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewEndpoint_Override(e Endpoint, address *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Endpoint",
		[]interface{}{address, port},
		e,
	)
}

// A version of an engine - for either a cluster, or instance.
// Experimental.
type EngineVersion struct {
	// The major version of the engine, for example, "5.6". Used in specifying the ParameterGroup family and OptionGroup version for this engine.
	// Experimental.
	MajorVersion *string `json:"majorVersion"`
	// The full version string of the engine, for example, "5.6.mysql_aurora.1.22.1". It can be undefined, which means RDS should use whatever version it deems appropriate for the given engine type.
	// Experimental.
	FullVersion *string `json:"fullVersion"`
}

// The interface representing a database cluster (as opposed to instance) engine.
// Experimental.
type IClusterEngine interface {
	IEngine
	// Method called when the engine is used to create a new cluster.
	// Experimental.
	BindToCluster(scope constructs.Construct, options *ClusterEngineBindOptions) *ClusterEngineConfig
	// The application used by this engine to perform rotation for a multi-user scenario.
	// Experimental.
	MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The application used by this engine to perform rotation for a single-user scenario.
	// Experimental.
	SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The log types that are available with this engine type.
	// Experimental.
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
// Experimental.
type IDatabaseCluster interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Add a new db proxy to this cluster.
	// Experimental.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	// Return the given named metric for this DBCluster.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes
	// Experimental.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes
	// Experimental.
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of deadlocks in the database per second.
	//
	// Average over 5 minutes
	// Experimental.
	MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of time that the instance has been running, in seconds.
	//
	// Average over 5 minutes
	// Experimental.
	MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory, in bytes.
	//
	// Average over 5 minutes
	// Experimental.
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of local storage available, in bytes.
	//
	// Average over 5 minutes
	// Experimental.
	MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput received from clients by each instance, in bytes per second.
	//
	// Average over 5 minutes
	// Experimental.
	MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput both received from and transmitted to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes
	// Experimental.
	MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput sent to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes
	// Experimental.
	MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes consumed by all Aurora snapshots outside its backup retention window.
	//
	// Average over 5 minutes
	// Experimental.
	MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes for which you are billed.
	//
	// Average over 5 minutes
	// Experimental.
	MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of storage used by your Aurora DB instance, in bytes.
	//
	// Average over 5 minutes
	// Experimental.
	MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of billed read I/O operations from a cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes
	// Experimental.
	MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of write disk I/O operations to the cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes
	// Experimental.
	MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	// Experimental.
	ClusterReadEndpoint() Endpoint
	// The engine of this Cluster.
	//
	// May be not known for imported Clusters if it wasn't provided explicitly.
	// Experimental.
	Engine() IClusterEngine
	// Endpoints which address each individual replica.
	// Experimental.
	InstanceEndpoints() *[]Endpoint
	// Identifiers of the replicas.
	// Experimental.
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
// Experimental.
type IDatabaseInstance interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Add a new db proxy to this instance.
	// Experimental.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	// Grant the given identity connection access to the database.
	// Experimental.
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this DBInstance.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes
	// Experimental.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes
	// Experimental.
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory.
	//
	// Average over 5 minutes
	// Experimental.
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available storage space.
	//
	// Average over 5 minutes
	// Experimental.
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk write I/O operations per second.
	//
	// Average over 5 minutes
	// Experimental.
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk read I/O operations per second.
	//
	// Average over 5 minutes
	// Experimental.
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CloudWatch event rule which triggers for instance events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The instance endpoint address.
	// Experimental.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	// Experimental.
	DbInstanceEndpointPort() *string
	// The engine of this database Instance.
	//
	// May be not known for imported Instances if it wasn't provided explicitly,
	// or for read replicas.
	// Experimental.
	Engine() IInstanceEngine
	// The instance arn.
	// Experimental.
	InstanceArn() *string
	// The instance endpoint.
	// Experimental.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	// Experimental.
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
// Experimental.
type IDatabaseProxy interface {
	awscdk.IResource
	// Grant the given identity connection access to the proxy.
	// Experimental.
	GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant
	// DB Proxy ARN.
	// Experimental.
	DbProxyArn() *string
	// DB Proxy Name.
	// Experimental.
	DbProxyName() *string
	// Endpoint.
	// Experimental.
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
// Experimental.
type IEngine interface {
	// The default name of the master database user if one was not provided explicitly.
	//
	// The global default of 'admin' will be used if this is `undefined`.
	// Note that 'admin' is a reserved word in PostgreSQL and cannot be used.
	// Experimental.
	DefaultUsername() *string
	// The family this engine belongs to, like "MYSQL", or "POSTGRESQL".
	//
	// This property is used when creating a Database Proxy.
	// Most engines don't belong to any family
	// (and because of that, you can't create Database Proxies for their Clusters or Instances).
	// Experimental.
	EngineFamily() *string
	// The type of the engine, for example "mysql".
	// Experimental.
	EngineType() *string
	// The exact version of the engine that is used, for example "5.1.42".
	// Experimental.
	EngineVersion() *EngineVersion
	// The family to use for ParameterGroups using this engine.
	//
	// This is usually equal to "<engineType><engineMajorVersion>",
	// but can sometimes be a variation of that.
	// You can pass this property when creating new ParameterGroup.
	// Experimental.
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
// Experimental.
type IInstanceEngine interface {
	IEngine
	// Method called when the engine is used to create a new instance.
	// Experimental.
	BindToInstance(scope constructs.Construct, options *InstanceEngineBindOptions) *InstanceEngineConfig
	// The application used by this engine to perform rotation for a multi-user scenario.
	// Experimental.
	MultiUserRotationApplication() awssecretsmanager.SecretRotationApplication
	// The application used by this engine to perform rotation for a single-user scenario.
	// Experimental.
	SingleUserRotationApplication() awssecretsmanager.SecretRotationApplication
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

// An option group.
// Experimental.
type IOptionGroup interface {
	awscdk.IResource
	// Adds a configuration to this OptionGroup.
	//
	// This method is a no-op for an imported OptionGroup.
	//
	// Returns: true if the OptionConfiguration was successfully added.
	// Experimental.
	AddConfiguration(configuration *OptionConfiguration) *bool
	// The name of the option group.
	// Experimental.
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
// Experimental.
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
	// Experimental.
	AddParameter(key *string, value *string) *bool
	// Method called when this Parameter Group is used when defining a database cluster.
	// Experimental.
	BindToCluster(options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig
	// Method called when this Parameter Group is used when defining a database instance.
	// Experimental.
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
// Experimental.
type IServerlessCluster interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Grant the given identity to access to the Data API.
	// Experimental.
	GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the cluster.
	// Experimental.
	ClusterArn() *string
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	// Experimental.
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
// Experimental.
type ISubnetGroup interface {
	awscdk.IResource
	// The name of the subnet group.
	// Experimental.
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
// Experimental.
type InstanceEngineBindOptions struct {
	// The Active Directory directory ID to create the DB instance in.
	// Experimental.
	Domain *string `json:"domain"`
	// The option group of the database.
	// Experimental.
	OptionGroup IOptionGroup `json:"optionGroup"`
	// The role used for S3 exporting.
	// Experimental.
	S3ExportRole awsiam.IRole `json:"s3ExportRole"`
	// The role used for S3 importing.
	// Experimental.
	S3ImportRole awsiam.IRole `json:"s3ImportRole"`
	// The timezone of the database, set by the customer.
	// Experimental.
	Timezone *string `json:"timezone"`
}

// The type returned from the {@link IInstanceEngine.bind} method.
// Experimental.
type InstanceEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	// Experimental.
	Features *InstanceEngineFeatures `json:"features"`
	// Option group of the database.
	// Experimental.
	OptionGroup IOptionGroup `json:"optionGroup"`
}

// Represents Database Engine features.
// Experimental.
type InstanceEngineFeatures struct {
	// Feature name for the DB instance that the IAM role to export to S3 bucket is to be associated with.
	// Experimental.
	S3Export *string `json:"s3Export"`
	// Feature name for the DB instance that the IAM role to access the S3 bucket for import is to be associated with.
	// Experimental.
	S3Import *string `json:"s3Import"`
}

// Instance properties for database instances.
// Experimental.
type InstanceProps struct {
	// What subnets to run the RDS instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Whether to allow upgrade of major version for the DB instance.
	// Experimental.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade"`
	// Whether to remove automated backups immediately after the DB instance is deleted for the DB instance.
	// Experimental.
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups"`
	// Whether to enable Performance Insights for the DB instance.
	// Experimental.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights"`
	// What type of instance to start for the replicas.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Experimental.
	PerformanceInsightEncryptionKey awskms.IKey `json:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Experimental.
	PerformanceInsightRetention PerformanceInsightRetention `json:"performanceInsightRetention"`
	// Indicates whether the DB instance is an internet-facing instance.
	// Experimental.
	PubliclyAccessible *bool `json:"publiclyAccessible"`
	// Security group.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Where to place the instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// The license model.
// Experimental.
type LicenseModel string

const (
	LicenseModel_LICENSE_INCLUDED LicenseModel = "LICENSE_INCLUDED"
	LicenseModel_BRING_YOUR_OWN_LICENSE LicenseModel = "BRING_YOUR_OWN_LICENSE"
	LicenseModel_GENERAL_PUBLIC_LICENSE LicenseModel = "GENERAL_PUBLIC_LICENSE"
)

// The versions for the MariaDB instance engines (those returned by {@link DatabaseInstanceEngine.mariaDb}).
// Experimental.
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
// Experimental.
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

func MariaDbEngineVersion_VER_10_0() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_17() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0_17",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_24() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0_24",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_28() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0_28",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_31() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0_31",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_32() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0_32",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_34() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0_34",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_0_35() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_0_35",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_1",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_14() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_1_14",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_19() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_1_19",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_23() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_1_23",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_26() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_1_26",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_31() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_1_31",
		&returns,
	)
	return returns
}

func MariaDbEngineVersion_VER_10_1_34() MariaDbEngineVersion {
	_init_.Initialize()
	var returns MariaDbEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		"VER_10_1_34",
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

// Properties for MariaDB instance engines.
//
// Used in {@link DatabaseInstanceEngine.mariaDb}.
// Experimental.
type MariaDbInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version MariaDbEngineVersion `json:"version"`
}

// Properties for MySQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.mysql}.
// Experimental.
type MySqlInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version MysqlEngineVersion `json:"version"`
}

// The versions for the MySQL instance engines (those returned by {@link DatabaseInstanceEngine.mysql}).
// Experimental.
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
// Experimental.
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

func MysqlEngineVersion_VER_5_5() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_5",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_46() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_5_46",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_53() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_5_53",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_57() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_5_57",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_59() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_5_59",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_61() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_5_61",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_34() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_34",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_35() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_35",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_37() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_37",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_39() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_39",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_40() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_40",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_41() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_41",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_43() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_43",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_44() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_44",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_46() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_46",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_48() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		"VER_5_6_48",
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

// Configuration properties for an option.
// Experimental.
type OptionConfiguration struct {
	// The name of the option.
	// Experimental.
	Name *string `json:"name"`
	// The port number that this option uses.
	//
	// If `port` is specified then `vpc`
	// must also be specified.
	// Experimental.
	Port *float64 `json:"port"`
	// Optional list of security groups to use for this option, if `vpc` is specified.
	//
	// If no groups are provided, a default one will be created.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The settings for the option.
	// Experimental.
	Settings *map[string]*string `json:"settings"`
	// The version for the option.
	// Experimental.
	Version *string `json:"version"`
	// The VPC where a security group should be created for this option.
	//
	// If `vpc`
	// is specified then `port` must also be specified.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// An option group.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewOptionGroup_Override(o OptionGroup, scope constructs.Construct, id *string, props *OptionGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.OptionGroup",
		[]interface{}{scope, id, props},
		o,
	)
}

// Import an existing option group.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (o *jsiiProxy_OptionGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type OptionGroupProps struct {
	// The configurations for this option group.
	// Experimental.
	Configurations *[]*OptionConfiguration `json:"configurations"`
	// The database engine that this option group is associated with.
	// Experimental.
	Engine IInstanceEngine `json:"engine"`
	// A description of the option group.
	// Experimental.
	Description *string `json:"description"`
}

// Properties for Oracle Enterprise Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleEe}.
// Experimental.
type OracleEeInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version OracleEngineVersion `json:"version"`
}

// The versions for the Oracle instance engines (those returned by {@link DatabaseInstanceEngine.oracleSe2} and {@link DatabaseInstanceEngine.oracleEe}).
// Experimental.
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
// Experimental.
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

// Properties for Oracle Standard Edition 2 instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleSe2}.
// Experimental.
type OracleSe2InstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version OracleEngineVersion `json:"version"`
}

// A parameter group.
//
// Represents both a cluster parameter group,
// and an instance parameter group.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewParameterGroup_Override(p ParameterGroup, scope constructs.Construct, id *string, props *ParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a parameter group.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (p *jsiiProxy_ParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Method called when this Parameter Group is used when defining a database cluster.
// Experimental.
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
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type ParameterGroupClusterBindOptions struct {
}

// The type returned from {@link IParameterGroup.bindToCluster}.
// Experimental.
type ParameterGroupClusterConfig struct {
	// The name of this parameter group.
	// Experimental.
	ParameterGroupName *string `json:"parameterGroupName"`
}

// Options for {@link IParameterGroup.bindToInstance}. Empty for now, but can be extended later.
// Experimental.
type ParameterGroupInstanceBindOptions struct {
}

// The type returned from {@link IParameterGroup.bindToInstance}.
// Experimental.
type ParameterGroupInstanceConfig struct {
	// The name of this parameter group.
	// Experimental.
	ParameterGroupName *string `json:"parameterGroupName"`
}

// Properties for a parameter group.
// Experimental.
type ParameterGroupProps struct {
	// The database engine for this parameter group.
	// Experimental.
	Engine IEngine `json:"engine"`
	// Description for this parameter group.
	// Experimental.
	Description *string `json:"description"`
	// The parameters in this parameter group.
	// Experimental.
	Parameters *map[string]*string `json:"parameters"`
}

// The retention period for Performance Insight.
// Experimental.
type PerformanceInsightRetention string

const (
	PerformanceInsightRetention_DEFAULT PerformanceInsightRetention = "DEFAULT"
	PerformanceInsightRetention_LONG_TERM PerformanceInsightRetention = "LONG_TERM"
)

// Features supported by the Postgres database engine.
// Experimental.
type PostgresEngineFeatures struct {
	// Whether this version of the Postgres engine supports the S3 data import feature.
	// Experimental.
	S3Import *bool `json:"s3Import"`
}

// The versions for the PostgreSQL instance engines (those returned by {@link DatabaseInstanceEngine.postgres}).
// Experimental.
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
// Experimental.
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

// Properties for PostgreSQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.postgres}.
// Experimental.
type PostgresInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version PostgresEngineVersion `json:"version"`
}

// The processor features.
// Experimental.
type ProcessorFeatures struct {
	// The number of CPU core.
	// Experimental.
	CoreCount *float64 `json:"coreCount"`
	// The number of threads per core.
	// Experimental.
	ThreadsPerCore *float64 `json:"threadsPerCore"`
}

// Proxy target: Instance or Cluster.
//
// A target group is a collection of databases that the proxy can connect to.
// Currently, you can specify only one RDS DB instance or Aurora DB cluster.
// Experimental.
type ProxyTarget interface {
	Bind(proxy DatabaseProxy) *ProxyTargetConfig
}

// The jsii proxy struct for ProxyTarget
type jsiiProxy_ProxyTarget struct {
	_ byte // padding
}

// From cluster.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type ProxyTargetConfig struct {
	// The engine family of the database instance or cluster this proxy connects with.
	// Experimental.
	EngineFamily *string `json:"engineFamily"`
	// The database clusters to which this proxy connects.
	//
	// Either this or `dbInstances` will be set and the other `undefined`.
	// Experimental.
	DbClusters *[]IDatabaseCluster `json:"dbClusters"`
	// The database instances to which this proxy connects.
	//
	// Either this or `dbClusters` will be set and the other `undefined`.
	// Experimental.
	DbInstances *[]IDatabaseInstance `json:"dbInstances"`
}

// Options to add the multi user rotation.
// Experimental.
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
	// Experimental.
	Secret awssecretsmanager.ISecret `json:"secret"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Experimental.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter"`
	// Specifies characters to not include in generated passwords.
	// Experimental.
	ExcludeCharacters *string `json:"excludeCharacters"`
}

// Options to add the multi user rotation.
// Experimental.
type RotationSingleUserOptions struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Experimental.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter"`
	// Specifies characters to not include in generated passwords.
	// Experimental.
	ExcludeCharacters *string `json:"excludeCharacters"`
}

// Create an Aurora Serverless Cluster.
// Experimental.
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
	Node() constructs.Node
	PhysicalName() *string
	Secret() awssecretsmanager.ISecret
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

func (j *jsiiProxy_ServerlessCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (s *jsiiProxy_ServerlessCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Renders the secret attachment target specifications.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type ServerlessClusterAttributes struct {
	// Identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `json:"clusterIdentifier"`
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `json:"clusterEndpointAddress"`
	// The database port.
	// Experimental.
	Port *float64 `json:"port"`
	// Reader endpoint address.
	// Experimental.
	ReaderEndpointAddress *string `json:"readerEndpointAddress"`
	// The secret attached to the database cluster.
	// Experimental.
	Secret awssecretsmanager.ISecret `json:"secret"`
	// The security groups of the database cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
}

// Properties to configure an Aurora Serverless Cluster.
// Experimental.
type ServerlessClusterProps struct {
	// What kind of database to start.
	// Experimental.
	Engine IClusterEngine `json:"engine"`
	// The VPC that this Aurora Serverless cluster has been created in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Automatic backup retention cannot be disabled on serverless clusters.
	// Must be a value from 1 day to 35 days.
	// Experimental.
	BackupRetention awscdk.Duration `json:"backupRetention"`
	// An optional identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `json:"clusterIdentifier"`
	// Credentials for the administrative user.
	// Experimental.
	Credentials Credentials `json:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	// Experimental.
	DefaultDatabaseName *string `json:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection"`
	// Whether to enable the Data API.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
	//
	// Experimental.
	EnableDataApi *bool `json:"enableDataApi"`
	// Additional parameters to pass to the database engine.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Scaling configuration of an Aurora Serverless database cluster.
	// Experimental.
	Scaling *ServerlessScalingOptions `json:"scaling"`
	// Security group.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The KMS key for storage encryption.
	// Experimental.
	StorageEncryptionKey awskms.IKey `json:"storageEncryptionKey"`
	// Existing subnet group for the cluster.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup"`
	// Where to place the instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// Options for configuring scaling on an Aurora Serverless cluster.
// Experimental.
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
	// Experimental.
	AutoPause awscdk.Duration `json:"autoPause"`
	// The maximum capacity for an Aurora Serverless database cluster.
	// Experimental.
	MaxCapacity AuroraCapacityUnit `json:"maxCapacity"`
	// The minimum capacity for an Aurora Serverless database cluster.
	// Experimental.
	MinCapacity AuroraCapacityUnit `json:"minCapacity"`
}

// SessionPinningFilter.
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html#rds-proxy-pinning
//
// Experimental.
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
// Experimental.
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
// Experimental.
type SnapshotCredentials interface {
	EncryptionKey() awskms.IKey
	ExcludeCharacters() *string
	GeneratePassword() *bool
	Password() awscdk.SecretValue
	ReplaceOnPasswordCriteriaChanges() *bool
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


// Experimental.
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
// The new credentials are stored in Secrets Manager.
//
// Note - The username must match the existing master username of the snapshot.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type SnapshotCredentialsFromGeneratedPasswordOptions struct {
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The characters to exclude from the generated password.
	// Experimental.
	ExcludeCharacters *string `json:"excludeCharacters"`
}

// Properties for SQL Server Enterprise Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerEe}.
// Experimental.
type SqlServerEeInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version SqlServerEngineVersion `json:"version"`
}

// The versions for the SQL Server instance engines (those returned by {@link DatabaseInstanceEngine.sqlServerSe}, {@link DatabaseInstanceEngine.sqlServerEx}, {@link DatabaseInstanceEngine.sqlServerWeb} and {@link DatabaseInstanceEngine.sqlServerEe}).
// Experimental.
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
// Experimental.
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

// Properties for SQL Server Express Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerEx}.
// Experimental.
type SqlServerExInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version SqlServerEngineVersion `json:"version"`
}

// Properties for SQL Server Standard Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerSe}.
// Experimental.
type SqlServerSeInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version SqlServerEngineVersion `json:"version"`
}

// Properties for SQL Server Web Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerWeb}.
// Experimental.
type SqlServerWebInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version SqlServerEngineVersion `json:"version"`
}

// The type of storage.
// Experimental.
type StorageType string

const (
	StorageType_STANDARD StorageType = "STANDARD"
	StorageType_GP2 StorageType = "GP2"
	StorageType_IO1 StorageType = "IO1"
)

// Class for creating a RDS DB subnet group.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewSubnetGroup_Override(s SubnetGroup, scope constructs.Construct, id *string, props *SubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.SubnetGroup",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports an existing subnet group by name.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (s *jsiiProxy_SubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type SubnetGroupProps struct {
	// Description of the subnet group.
	// Experimental.
	Description *string `json:"description"`
	// The VPC to place the subnet group in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The removal policy to apply when the subnet group are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// The name of the subnet group.
	// Experimental.
	SubnetGroupName *string `json:"subnetGroupName"`
	// Which subnets within the VPC to associate with this group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

