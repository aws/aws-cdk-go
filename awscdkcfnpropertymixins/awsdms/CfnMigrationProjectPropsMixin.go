package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provides information that defines a migration project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMigrationProjectPropsMixin := awscdkcfnpropertymixins.Aws_dms.NewCfnMigrationProjectPropsMixin(&CfnMigrationProjectMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceProfileArn: jsii.String("instanceProfileArn"),
//   	InstanceProfileIdentifier: jsii.String("instanceProfileIdentifier"),
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   	MigrationProjectCreationTime: jsii.String("migrationProjectCreationTime"),
//   	MigrationProjectIdentifier: jsii.String("migrationProjectIdentifier"),
//   	MigrationProjectName: jsii.String("migrationProjectName"),
//   	SchemaConversionApplicationAttributes: &SchemaConversionApplicationAttributesProperty{
//   		S3BucketPath: jsii.String("s3BucketPath"),
//   		S3BucketRoleArn: jsii.String("s3BucketRoleArn"),
//   	},
//   	SourceDataProviderDescriptors: []interface{}{
//   		&DataProviderDescriptorProperty{
//   			DataProviderArn: jsii.String("dataProviderArn"),
//   			DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   			DataProviderName: jsii.String("dataProviderName"),
//   			SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   			SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetDataProviderDescriptors: []interface{}{
//   		&DataProviderDescriptorProperty{
//   			DataProviderArn: jsii.String("dataProviderArn"),
//   			DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   			DataProviderName: jsii.String("dataProviderName"),
//   			SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   			SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		},
//   	},
//   	TransformationRules: jsii.String("transformationRules"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html
//
type CfnMigrationProjectPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMigrationProjectMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMigrationProjectPropsMixin
type jsiiProxy_CfnMigrationProjectPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMigrationProjectPropsMixin) Props() *CfnMigrationProjectMixinProps {
	var returns *CfnMigrationProjectMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMigrationProjectPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DMS::MigrationProject`.
func NewCfnMigrationProjectPropsMixin(props *CfnMigrationProjectMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMigrationProjectPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMigrationProjectPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMigrationProjectPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnMigrationProjectPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DMS::MigrationProject`.
func NewCfnMigrationProjectPropsMixin_Override(c CfnMigrationProjectPropsMixin, props *CfnMigrationProjectMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnMigrationProjectPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMigrationProjectPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMigrationProjectPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnMigrationProjectPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMigrationProjectPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnMigrationProjectPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMigrationProjectPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMigrationProjectPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

