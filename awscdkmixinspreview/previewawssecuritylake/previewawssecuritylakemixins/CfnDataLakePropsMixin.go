package previewawssecuritylakemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecuritylake/previewawssecuritylakemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Initializes an Amazon Security Lake instance with the provided (or default) configuration.
//
// You can enable Security Lake in AWS Regions with customized settings before enabling log collection in Regions. To specify particular Regions, configure these Regions using the `configurations` parameter. If you have already enabled Security Lake in a Region when you call this command, the command will update the Region if you provide new configuration parameters. If you have not already enabled Security Lake in the Region when you call this API, it will set up the data lake in the Region with the specified configurations.
//
// When you enable Security Lake , it starts ingesting security data after the `CreateAwsLogSource` call. This includes ingesting security data from sources, storing data, and making data accessible to subscribers. Security Lake also enables all the existing settings and resources that it stores or maintains for your AWS account in the current Region, including security log and event data. For more information, see the [Amazon Security Lake User Guide](https://docs.aws.amazon.com//security-lake/latest/userguide/what-is-security-lake.html) .
//
// > If you use this template to create multiple data lakes in different AWS Regions , and more than one of your data lakes include an [AWS::SecurityLake::AwsLogSource](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html) resource, then you must deploy these data lakes sequentially. This is required because data lakes operate globally, and `AwsLogSource` resources must be deployed one at a time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataLakePropsMixin := awscdkmixinspreview.Mixins.NewCfnDataLakePropsMixin(&CfnDataLakeMixinProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		Expiration: &ExpirationProperty{
//   			Days: jsii.Number(123),
//   		},
//   		Transitions: []interface{}{
//   			&TransitionsProperty{
//   				Days: jsii.Number(123),
//   				StorageClass: jsii.String("storageClass"),
//   			},
//   		},
//   	},
//   	MetaStoreManagerRoleArn: jsii.String("metaStoreManagerRoleArn"),
//   	ReplicationConfiguration: &ReplicationConfigurationProperty{
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-datalake.html
//
type CfnDataLakePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataLakeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataLakePropsMixin
type jsiiProxy_CfnDataLakePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataLakePropsMixin) Props() *CfnDataLakeMixinProps {
	var returns *CfnDataLakeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityLake::DataLake`.
func NewCfnDataLakePropsMixin(props *CfnDataLakeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataLakePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataLakePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataLakePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityLake::DataLake`.
func NewCfnDataLakePropsMixin_Override(c CfnDataLakePropsMixin, props *CfnDataLakeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataLakePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataLakePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataLakePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataLakePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDataLakePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

