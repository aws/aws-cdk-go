package previewawssecretsmanagermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecretsmanager/previewawssecretsmanagermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configure the rotation schedule and Lambda rotation function for a secret. For more information, see [How rotation works](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html) .
//
// For database credentials, refer to the following resources:
//
// - Amazon RDS master user credentials: [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html)
// - Amazon Redshift admin user credentials: [AWS::Redshift::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html)
//
// Choose one of the following options for the rotation function:
//
// - Create a new rotation function using `HostedRotationLambda` based on a [Secrets Manager rotation function template](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) .
// - Use an existing rotation function by specifying its ARN with `RotationLambdaARN` .
//
// > For database secrets defined in the same CloudFormation template as the database or service:
// >
// > - Use the [AWS::SecretsManager::SecretTargetAttachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html) resource to populate the secret with connection details.
// > - Add a `DependsOn` attribute to the `RotationSchedule` resource that uses a `SecretTargetAttachment` . This ensures the rotation is configured after the secret is populated with connection details. > You can define only one rotation schedule per secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRotationSchedulePropsMixin := awscdkmixinspreview.Mixins.NewCfnRotationSchedulePropsMixin(&CfnRotationScheduleMixinProps{
//   	ExternalSecretRotationMetadata: []interface{}{
//   		&ExternalSecretRotationMetadataItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ExternalSecretRotationRoleArn: jsii.String("externalSecretRotationRoleArn"),
//   	HostedRotationLambda: &HostedRotationLambdaProperty{
//   		ExcludeCharacters: jsii.String("excludeCharacters"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		MasterSecretArn: jsii.String("masterSecretArn"),
//   		MasterSecretKmsKeyArn: jsii.String("masterSecretKmsKeyArn"),
//   		RotationLambdaName: jsii.String("rotationLambdaName"),
//   		RotationType: jsii.String("rotationType"),
//   		Runtime: jsii.String("runtime"),
//   		SuperuserSecretArn: jsii.String("superuserSecretArn"),
//   		SuperuserSecretKmsKeyArn: jsii.String("superuserSecretKmsKeyArn"),
//   		VpcSecurityGroupIds: jsii.String("vpcSecurityGroupIds"),
//   		VpcSubnetIds: jsii.String("vpcSubnetIds"),
//   	},
//   	RotateImmediatelyOnUpdate: jsii.Boolean(false),
//   	RotationLambdaArn: jsii.String("rotationLambdaArn"),
//   	RotationRules: &RotationRulesProperty{
//   		AutomaticallyAfterDays: jsii.Number(123),
//   		Duration: jsii.String("duration"),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	SecretId: jsii.String("secretId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html
//
type CfnRotationSchedulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRotationScheduleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRotationSchedulePropsMixin
type jsiiProxy_CfnRotationSchedulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRotationSchedulePropsMixin) Props() *CfnRotationScheduleMixinProps {
	var returns *CfnRotationScheduleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecretsManager::RotationSchedule`.
func NewCfnRotationSchedulePropsMixin(props *CfnRotationScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRotationSchedulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRotationSchedulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRotationSchedulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_secretsmanager.mixins.CfnRotationSchedulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecretsManager::RotationSchedule`.
func NewCfnRotationSchedulePropsMixin_Override(c CfnRotationSchedulePropsMixin, props *CfnRotationScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_secretsmanager.mixins.CfnRotationSchedulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRotationSchedulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRotationSchedulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_secretsmanager.mixins.CfnRotationSchedulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRotationSchedulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_secretsmanager.mixins.CfnRotationSchedulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRotationSchedulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRotationSchedulePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

