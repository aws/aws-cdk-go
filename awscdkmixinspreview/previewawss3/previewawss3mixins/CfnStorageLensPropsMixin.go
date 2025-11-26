package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::S3::StorageLens resource creates an Amazon S3 Storage Lens configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sses3 interface{}
//
//   cfnStorageLensPropsMixin := awscdkmixinspreview.Mixins.NewCfnStorageLensPropsMixin(&CfnStorageLensMixinProps{
//   	StorageLensConfiguration: &StorageLensConfigurationProperty{
//   		AccountLevel: &AccountLevelProperty{
//   			ActivityMetrics: &ActivityMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			BucketLevel: &BucketLevelProperty{
//   				ActivityMetrics: &ActivityMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				PrefixLevel: &PrefixLevelProperty{
//   					StorageMetrics: &PrefixLevelStorageMetricsProperty{
//   						IsEnabled: jsii.Boolean(false),
//   						SelectionCriteria: &SelectionCriteriaProperty{
//   							Delimiter: jsii.String("delimiter"),
//   							MaxDepth: jsii.Number(123),
//   							MinStorageBytesPercentage: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   			DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			StorageLensGroupLevel: &StorageLensGroupLevelProperty{
//   				StorageLensGroupSelectionCriteria: &StorageLensGroupSelectionCriteriaProperty{
//   					Exclude: []*string{
//   						jsii.String("exclude"),
//   					},
//   					Include: []*string{
//   						jsii.String("include"),
//   					},
//   				},
//   			},
//   		},
//   		AwsOrg: &AwsOrgProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		DataExport: &DataExportProperty{
//   			CloudWatchMetrics: &CloudWatchMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			S3BucketDestination: &S3BucketDestinationProperty{
//   				AccountId: jsii.String("accountId"),
//   				Arn: jsii.String("arn"),
//   				Encryption: &EncryptionProperty{
//   					Ssekms: &SSEKMSProperty{
//   						KeyId: jsii.String("keyId"),
//   					},
//   					Sses3: sses3,
//   				},
//   				Format: jsii.String("format"),
//   				OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   		Exclude: &BucketsAndRegionsProperty{
//   			Buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		Id: jsii.String("id"),
//   		Include: &BucketsAndRegionsProperty{
//   			Buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		IsEnabled: jsii.Boolean(false),
//   		StorageLensArn: jsii.String("storageLensArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html
//
type CfnStorageLensPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStorageLensMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStorageLensPropsMixin
type jsiiProxy_CfnStorageLensPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStorageLensPropsMixin) Props() *CfnStorageLensMixinProps {
	var returns *CfnStorageLensMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLensPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3::StorageLens`.
func NewCfnStorageLensPropsMixin(props *CfnStorageLensMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStorageLensPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStorageLensPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStorageLensPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3::StorageLens`.
func NewCfnStorageLensPropsMixin_Override(c CfnStorageLensPropsMixin, props *CfnStorageLensMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStorageLensPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStorageLensPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStorageLensPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStorageLensPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStorageLensPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

