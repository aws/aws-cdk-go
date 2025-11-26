package previewawsm2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsm2/previewawsm2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a runtime environment for a given runtime engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentPropsMixin(&CfnEnvironmentMixinProps{
//   	Description: jsii.String("description"),
//   	EngineType: jsii.String("engineType"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	HighAvailabilityConfig: &HighAvailabilityConfigProperty{
//   		DesiredCapacity: jsii.Number(123),
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	NetworkType: jsii.String("networkType"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	StorageConfigurations: []interface{}{
//   		&StorageConfigurationProperty{
//   			Efs: &EfsStorageConfigurationProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   				MountPoint: jsii.String("mountPoint"),
//   			},
//   			Fsx: &FsxStorageConfigurationProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   				MountPoint: jsii.String("mountPoint"),
//   			},
//   		},
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-environment.html
//
type CfnEnvironmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentPropsMixin
type jsiiProxy_CfnEnvironmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Props() *CfnEnvironmentMixinProps {
	var returns *CfnEnvironmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::M2::Environment`.
func NewCfnEnvironmentPropsMixin(props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::M2::Environment`.
func NewCfnEnvironmentPropsMixin_Override(c CfnEnvironmentPropsMixin, props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEnvironmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

