package mixinsawsredshiftserverless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsredshiftserverless/mixinsawsredshiftserverless/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// A collection of database objects and users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var namespaceResourcePolicy interface{}
//
//   cfnNamespacePropsMixin := awscdkmixinspreview.Mixins.NewCfnNamespacePropsMixin(&CfnNamespaceMixinProps{
//   	AdminPasswordSecretKmsKeyId: jsii.String("adminPasswordSecretKmsKeyId"),
//   	AdminUsername: jsii.String("adminUsername"),
//   	AdminUserPassword: jsii.String("adminUserPassword"),
//   	DbName: jsii.String("dbName"),
//   	DefaultIamRoleArn: jsii.String("defaultIamRoleArn"),
//   	FinalSnapshotName: jsii.String("finalSnapshotName"),
//   	FinalSnapshotRetentionPeriod: jsii.Number(123),
//   	IamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogExports: []*string{
//   		jsii.String("logExports"),
//   	},
//   	ManageAdminPassword: jsii.Boolean(false),
//   	NamespaceName: jsii.String("namespaceName"),
//   	NamespaceResourcePolicy: namespaceResourcePolicy,
//   	RedshiftIdcApplicationArn: jsii.String("redshiftIdcApplicationArn"),
//   	SnapshotCopyConfigurations: []interface{}{
//   		&SnapshotCopyConfigurationProperty{
//   			DestinationKmsKeyId: jsii.String("destinationKmsKeyId"),
//   			DestinationRegion: jsii.String("destinationRegion"),
//   			SnapshotRetentionPeriod: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-namespace.html
//
type CfnNamespacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNamespaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNamespacePropsMixin
type jsiiProxy_CfnNamespacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNamespacePropsMixin) Props() *CfnNamespaceMixinProps {
	var returns *CfnNamespaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RedshiftServerless::Namespace`.
func NewCfnNamespacePropsMixin(props *CfnNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNamespacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNamespacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNamespacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RedshiftServerless::Namespace`.
func NewCfnNamespacePropsMixin_Override(c CfnNamespacePropsMixin, props *CfnNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNamespacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNamespacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNamespacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNamespacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNamespacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

