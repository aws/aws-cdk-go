package previewawsdmsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdms/previewawsdmsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provides information that defines a migration project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMigrationProjectPropsMixin := awscdkmixinspreview.Mixins.NewCfnMigrationProjectPropsMixin(&CfnMigrationProjectMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html
//
type CfnMigrationProjectPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMigrationProjectMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMigrationProjectPropsMixin
type jsiiProxy_CfnMigrationProjectPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnMigrationProjectPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DMS::MigrationProject`.
func NewCfnMigrationProjectPropsMixin(props *CfnMigrationProjectMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMigrationProjectPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMigrationProjectPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMigrationProjectPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DMS::MigrationProject`.
func NewCfnMigrationProjectPropsMixin_Override(c CfnMigrationProjectPropsMixin, props *CfnMigrationProjectMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMigrationProjectPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMigrationProjectPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnMigrationProjectPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMigrationProjectPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

