package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of authentication protocol or the VPC components for your event source's SourceAccessConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAccessConfigurationType := awscdk.Aws_lambda.SourceAccessConfigurationType_BASIC_AUTH()
//
// See: https://docs.aws.amazon.com/lambda/latest/dg/API_SourceAccessConfiguration.html#SSS-Type-SourceAccessConfiguration-Type
//
type SourceAccessConfigurationType interface {
	// The key to use in `SourceAccessConfigurationProperty.Type` property in CloudFormation.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html#cfn-lambda-eventsourcemapping-sourceaccessconfiguration-type
	//
	Type() *string
}

// The jsii proxy struct for SourceAccessConfigurationType
type jsiiProxy_SourceAccessConfigurationType struct {
	_ byte // padding
}

func (j *jsiiProxy_SourceAccessConfigurationType) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// A custom source access configuration property.
func SourceAccessConfigurationType_Of(name *string) SourceAccessConfigurationType {
	_init_.Initialize()

	if err := validateSourceAccessConfigurationType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns SourceAccessConfigurationType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func SourceAccessConfigurationType_BASIC_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"BASIC_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_CLIENT_CERTIFICATE_TLS_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"CLIENT_CERTIFICATE_TLS_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_SASL_SCRAM_256_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"SASL_SCRAM_256_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_SASL_SCRAM_512_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"SASL_SCRAM_512_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_SERVER_ROOT_CA_CERTIFICATE() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"SERVER_ROOT_CA_CERTIFICATE",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_VPC_SECURITY_GROUP() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"VPC_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_VPC_SUBNET() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"VPC_SUBNET",
		&returns,
	)
	return returns
}

