// The CDK Construct Library for AWS::MSK
package awscdkmskalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmskalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Kafka cluster version.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V2_8_1(),
//   	Vpc: Vpc,
//   	EncryptionInTransit: &EncryptionInTransitConfig{
//   		ClientBroker: msk.ClientBrokerEncryption_TLS,
//   	},
//   	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
//   		Scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type KafkaVersion interface {
	// cluster version number.
	// Experimental.
	Version() *string
}

// The jsii proxy struct for KafkaVersion
type jsiiProxy_KafkaVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_KafkaVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom cluster version.
// Experimental.
func KafkaVersion_Of(version *string) KafkaVersion {
	_init_.Initialize()

	if err := validateKafkaVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns KafkaVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KafkaVersion_V1_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V1_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_2_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_2_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_3_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_3_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_4_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_4_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_5_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_5_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_2",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_3() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_3",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_7_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_7_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_7_2",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_8_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_8_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_8_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_8_1",
		&returns,
	)
	return returns
}

