package awscdkneptunealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible Instances Types to use in Neptune cluster used for defining `DatabaseInstanceProps.instanceType`.
//
// Example:
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("ServerlessDatabase"), &DatabaseClusterProps{
//   	Vpc: Vpc,
//   	InstanceType: neptune.InstanceType_SERVERLESS(),
//   	ServerlessScalingConfiguration: &ServerlessScalingConfiguration{
//   		MinCapacity: jsii.Number(1),
//   		MaxCapacity: jsii.Number(5),
//   	},
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

	if err := validateInstanceType_OfParameters(instanceType); err != nil {
		panic(err)
	}
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

func InstanceType_R5_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_16XLARGE",
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

func InstanceType_R5D_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5D_XLARGE",
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

func InstanceType_R6I_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_32XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_32XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6I_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6I_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_SERVERLESS() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"SERVERLESS",
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

func InstanceType_X2G_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2G_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2G_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2G_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2G_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2G_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2G_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2G_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2G_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2G_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2G_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2G_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2G_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2G_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2IEDN_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2IEDN_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2IEDN_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2IEDN_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2IEDN_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2IEDN_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2IEDN_32XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2IEDN_32XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2IEDN_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2IEDN_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2IEDN_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2IEDN_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_X2IEDN_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"X2IEDN_XLARGE",
		&returns,
	)
	return returns
}

