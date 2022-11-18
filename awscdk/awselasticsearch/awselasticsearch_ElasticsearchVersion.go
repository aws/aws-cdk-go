package awselasticsearch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Elasticsearch version.
//
// Example:
//   domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: es.elasticsearchVersion_V7_4(),
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
type ElasticsearchVersion interface {
	// Elasticsearch version number.
	Version() *string
}

// The jsii proxy struct for ElasticsearchVersion
type jsiiProxy_ElasticsearchVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_ElasticsearchVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom Elasticsearch version.
func ElasticsearchVersion_Of(version *string) ElasticsearchVersion {
	_init_.Initialize()

	if err := validateElasticsearchVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns ElasticsearchVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func ElasticsearchVersion_V1_5() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V1_5",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V2_3() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V2_3",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_1() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_1",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_3() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_3",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_5() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_5",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_6() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_6",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_0() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_0",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_2() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_2",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_3() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_3",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_4() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_4",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_5() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_5",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_7() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_7",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_8() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_8",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_1() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_1",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_10() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_10",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_4() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_4",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_7() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_7",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_8() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_8",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_9() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_9",
		&returns,
	)
	return returns
}

