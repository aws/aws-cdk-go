// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible Instances Types to use in Neptune cluster used for defining {@link DatabaseInstanceProps.instanceType}.
//
// Example:
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	vpc: vpc,
//   	instanceType: neptune.instanceType_R5_LARGE(),
//   	instances: jsii.Number(2),
//   })
//
// Experimental.
type InstanceType interface {
}

// The jsii proxy struct for InstanceType
type jsiiProxy_InstanceType struct {
	_ byte // padding
}

// Build an InstanceType from given string or token, such as CfnParameter.
// Experimental.
func InstanceType_Of(instanceType *string) InstanceType {
	_init_.Initialize()

	var returns InstanceType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"of",
		[]interface{}{instanceType},
		&returns,
	)

	return returns
}

func InstanceType_R4_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_T3_MEDIUM() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"T3_MEDIUM",
		&returns,
	)
	return returns
}

func InstanceType_T4G_MEDIUM() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"T4G_MEDIUM",
		&returns,
	)
	return returns
}

