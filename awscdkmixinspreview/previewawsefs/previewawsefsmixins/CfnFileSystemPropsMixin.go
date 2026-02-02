package previewawsefsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsefs/previewawsefsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::EFS::FileSystem` resource creates a new, empty file system in Amazon Elastic File System ( Amazon EFS ).
//
// You must create a mount target ( [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) ) to mount your EFS file system on an Amazon EC2 or other AWS cloud compute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var fileSystemPolicy interface{}
//
//   cfnFileSystemPropsMixin := awscdkmixinspreview.Mixins.NewCfnFileSystemPropsMixin(&CfnFileSystemMixinProps{
//   	AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   	BackupPolicy: &BackupPolicyProperty{
//   		Status: jsii.String("status"),
//   	},
//   	BypassPolicyLockoutSafetyCheck: jsii.Boolean(false),
//   	Encrypted: jsii.Boolean(false),
//   	FileSystemPolicy: fileSystemPolicy,
//   	FileSystemProtection: &FileSystemProtectionProperty{
//   		ReplicationOverwriteProtection: jsii.String("replicationOverwriteProtection"),
//   	},
//   	FileSystemTags: []ElasticFileSystemTagProperty{
//   		&ElasticFileSystemTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LifecyclePolicies: []interface{}{
//   		&LifecyclePolicyProperty{
//   			TransitionToArchive: jsii.String("transitionToArchive"),
//   			TransitionToIa: jsii.String("transitionToIa"),
//   			TransitionToPrimaryStorageClass: jsii.String("transitionToPrimaryStorageClass"),
//   		},
//   	},
//   	PerformanceMode: jsii.String("performanceMode"),
//   	ProvisionedThroughputInMibps: jsii.Number(123),
//   	ReplicationConfiguration: &ReplicationConfigurationProperty{
//   		Destinations: []interface{}{
//   			&ReplicationDestinationProperty{
//   				AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   				FileSystemId: jsii.String("fileSystemId"),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				Region: jsii.String("region"),
//   				RoleArn: jsii.String("roleArn"),
//   				Status: jsii.String("status"),
//   				StatusMessage: jsii.String("statusMessage"),
//   			},
//   		},
//   	},
//   	ThroughputMode: jsii.String("throughputMode"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
//
type CfnFileSystemPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFileSystemMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFileSystemPropsMixin
type jsiiProxy_CfnFileSystemPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFileSystemPropsMixin) Props() *CfnFileSystemMixinProps {
	var returns *CfnFileSystemMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystemPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EFS::FileSystem`.
func NewCfnFileSystemPropsMixin(props *CfnFileSystemMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFileSystemPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFileSystemPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFileSystemPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EFS::FileSystem`.
func NewCfnFileSystemPropsMixin_Override(c CfnFileSystemPropsMixin, props *CfnFileSystemMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFileSystemPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileSystemPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFileSystemPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFileSystemPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFileSystemPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

