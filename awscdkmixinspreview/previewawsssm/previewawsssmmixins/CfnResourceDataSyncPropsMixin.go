package previewawsssmmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssm/previewawsssmmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSM::ResourceDataSync` resource creates, updates, or deletes a resource data sync for AWS Systems Manager .
//
// A resource data sync helps you view data from multiple sources in a single location. Systems Manager offers two types of resource data sync: `SyncToDestination` and `SyncFromSource` .
//
// You can configure Systems Manager Inventory to use the `SyncToDestination` type to synchronize Inventory data from multiple AWS Regions to a single Amazon S3 bucket.
//
// You can configure Systems Manager Explorer to use the `SyncFromSource` type to synchronize operational work items (OpsItems) and operational data (OpsData) from multiple AWS Regions . This type can synchronize OpsItems and OpsData from multiple AWS accounts and Regions or from an `EntireOrganization` by using AWS Organizations .
//
// A resource data sync is an asynchronous operation that returns immediately. After a successful initial sync is completed, the system continuously syncs data.
//
// By default, data is not encrypted in Amazon S3 . We strongly recommend that you enable encryption in Amazon S3 to ensure secure data storage. We also recommend that you secure access to the Amazon S3 bucket by creating a restrictive bucket policy.
//
// For more information, see [Configuring Inventory Collection](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-inventory-configuring.html#sysman-inventory-datasync) and [Setting Up Systems Manager Explorer to Display Data from Multiple Accounts and Regions](https://docs.aws.amazon.com/systems-manager/latest/userguide/Explorer-resource-data-sync.html) in the *AWS Systems Manager User Guide* .
//
// > The following *Syntax* section shows all fields that are supported for a resource data sync. The *Examples* section below shows the recommended way to specify configurations for each sync type. Refer to the *Examples* section when you create your resource data sync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceDataSyncPropsMixin := awscdkmixinspreview.Mixins.NewCfnResourceDataSyncPropsMixin(&CfnResourceDataSyncMixinProps{
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	S3Destination: &S3DestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		BucketRegion: jsii.String("bucketRegion"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SyncFormat: jsii.String("syncFormat"),
//   	},
//   	SyncFormat: jsii.String("syncFormat"),
//   	SyncName: jsii.String("syncName"),
//   	SyncSource: &SyncSourceProperty{
//   		AwsOrganizationsSource: &AwsOrganizationsSourceProperty{
//   			OrganizationalUnits: []*string{
//   				jsii.String("organizationalUnits"),
//   			},
//   			OrganizationSourceType: jsii.String("organizationSourceType"),
//   		},
//   		IncludeFutureRegions: jsii.Boolean(false),
//   		SourceRegions: []*string{
//   			jsii.String("sourceRegions"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//   	},
//   	SyncType: jsii.String("syncType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html
//
type CfnResourceDataSyncPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResourceDataSyncMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourceDataSyncPropsMixin
type jsiiProxy_CfnResourceDataSyncPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResourceDataSyncPropsMixin) Props() *CfnResourceDataSyncMixinProps {
	var returns *CfnResourceDataSyncMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDataSyncPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::ResourceDataSync`.
func NewCfnResourceDataSyncPropsMixin(props *CfnResourceDataSyncMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResourceDataSyncPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourceDataSyncPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourceDataSyncPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::ResourceDataSync`.
func NewCfnResourceDataSyncPropsMixin_Override(c CfnResourceDataSyncPropsMixin, props *CfnResourceDataSyncMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResourceDataSyncPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourceDataSyncPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceDataSyncPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceDataSyncPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResourceDataSyncPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

