package awsopensearchservice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// OpenSearch version.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	ebs: &ebsOptions{
//   		volumeSize: jsii.Number(100),
//   		volumeType: ec2.ebsDeviceVolumeType_GENERAL_PURPOSE_SSD,
//   	},
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   })
//
type EngineVersion interface {
	// engine version identifier.
	Version() *string
}

// The jsii proxy struct for EngineVersion
type jsiiProxy_EngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_EngineVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom ElasticSearch version.
func EngineVersion_Elasticsearch(version *string) EngineVersion {
	_init_.Initialize()

	var returns EngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"elasticsearch",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// Custom OpenSearch version.
func EngineVersion_OpenSearch(version *string) EngineVersion {
	_init_.Initialize()

	var returns EngineVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"openSearch",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func EngineVersion_ELASTICSEARCH_1_5() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_1_5",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_2_3() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_2_3",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_5_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_5_1",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_5_3() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_5_3",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_5_5() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_5_5",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_5_6() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_5_6",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_6_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_6_0",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_6_2() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_6_2",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_6_3() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_6_3",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_6_4() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_6_4",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_6_5() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_6_5",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_6_7() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_6_7",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_6_8() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_6_8",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_7_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_7_1",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_7_10() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_7_10",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_7_4() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_7_4",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_7_7() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_7_7",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_7_8() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_7_8",
		&returns,
	)
	return returns
}

func EngineVersion_ELASTICSEARCH_7_9() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"ELASTICSEARCH_7_9",
		&returns,
	)
	return returns
}

func EngineVersion_OPENSEARCH_1_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"OPENSEARCH_1_0",
		&returns,
	)
	return returns
}

func EngineVersion_OPENSEARCH_1_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"OPENSEARCH_1_1",
		&returns,
	)
	return returns
}

func EngineVersion_OPENSEARCH_1_2() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"OPENSEARCH_1_2",
		&returns,
	)
	return returns
}

func EngineVersion_OPENSEARCH_1_3() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		"OPENSEARCH_1_3",
		&returns,
	)
	return returns
}

