package previewawsathenamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsathena/previewawsathenamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::Athena::WorkGroup resource specifies an Amazon Athena workgroup, which contains a name, description, creation time, state, and other configuration, listed under `WorkGroupConfiguration` .
//
// Each workgroup enables you to isolate queries for you or your group from other queries in the same account. For more information, see [CreateWorkGroup](https://docs.aws.amazon.com/athena/latest/APIReference/API_CreateWorkGroup.html) in the *Amazon Athena API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkGroupPropsMixin(&CfnWorkGroupMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RecursiveDeleteOption: jsii.Boolean(false),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkGroupConfiguration: &WorkGroupConfigurationProperty{
//   		AdditionalConfiguration: jsii.String("additionalConfiguration"),
//   		BytesScannedCutoffPerQuery: jsii.Number(123),
//   		CustomerContentEncryptionConfiguration: &CustomerContentEncryptionConfigurationProperty{
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   		EnforceWorkGroupConfiguration: jsii.Boolean(false),
//   		EngineConfiguration: &EngineConfigurationProperty{
//   			AdditionalConfigs: map[string]*string{
//   				"additionalConfigsKey": jsii.String("additionalConfigs"),
//   			},
//   			Classifications: []interface{}{
//   				&ClassificationProperty{
//   					Name: jsii.String("name"),
//   					Properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			CoordinatorDpuSize: jsii.Number(123),
//   			DefaultExecutorDpuSize: jsii.Number(123),
//   			MaxConcurrentDpus: jsii.Number(123),
//   			SparkProperties: map[string]*string{
//   				"sparkPropertiesKey": jsii.String("sparkProperties"),
//   			},
//   		},
//   		EngineVersion: &EngineVersionProperty{
//   			EffectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   			SelectedEngineVersion: jsii.String("selectedEngineVersion"),
//   		},
//   		ExecutionRole: jsii.String("executionRole"),
//   		ManagedQueryResultsConfiguration: &ManagedQueryResultsConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionConfiguration: &ManagedStorageEncryptionConfigurationProperty{
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   		},
//   		MonitoringConfiguration: &MonitoringConfigurationProperty{
//   			CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   				LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   				LogTypes: map[string][]*string{
//   					"logTypesKey": []*string{
//   						jsii.String("logTypes"),
//   					},
//   				},
//   			},
//   			ManagedLoggingConfiguration: &ManagedLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			S3LoggingConfiguration: &S3LoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   				LogLocation: jsii.String("logLocation"),
//   			},
//   		},
//   		PublishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   		RequesterPaysEnabled: jsii.Boolean(false),
//   		ResultConfiguration: &ResultConfigurationProperty{
//   			AclConfiguration: &AclConfigurationProperty{
//   				S3AclOption: jsii.String("s3AclOption"),
//   			},
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				EncryptionOption: jsii.String("encryptionOption"),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputLocation: jsii.String("outputLocation"),
//   		},
//   	},
//   	WorkGroupConfigurationUpdates: &WorkGroupConfigurationUpdatesProperty{
//   		AdditionalConfiguration: jsii.String("additionalConfiguration"),
//   		BytesScannedCutoffPerQuery: jsii.Number(123),
//   		CustomerContentEncryptionConfiguration: &CustomerContentEncryptionConfigurationProperty{
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   		EnforceWorkGroupConfiguration: jsii.Boolean(false),
//   		EngineConfiguration: &EngineConfigurationProperty{
//   			AdditionalConfigs: map[string]*string{
//   				"additionalConfigsKey": jsii.String("additionalConfigs"),
//   			},
//   			Classifications: []interface{}{
//   				&ClassificationProperty{
//   					Name: jsii.String("name"),
//   					Properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			CoordinatorDpuSize: jsii.Number(123),
//   			DefaultExecutorDpuSize: jsii.Number(123),
//   			MaxConcurrentDpus: jsii.Number(123),
//   			SparkProperties: map[string]*string{
//   				"sparkPropertiesKey": jsii.String("sparkProperties"),
//   			},
//   		},
//   		EngineVersion: &EngineVersionProperty{
//   			EffectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   			SelectedEngineVersion: jsii.String("selectedEngineVersion"),
//   		},
//   		ExecutionRole: jsii.String("executionRole"),
//   		ManagedQueryResultsConfiguration: &ManagedQueryResultsConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionConfiguration: &ManagedStorageEncryptionConfigurationProperty{
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   		},
//   		MonitoringConfiguration: &MonitoringConfigurationProperty{
//   			CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   				LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   				LogTypes: map[string][]*string{
//   					"logTypesKey": []*string{
//   						jsii.String("logTypes"),
//   					},
//   				},
//   			},
//   			ManagedLoggingConfiguration: &ManagedLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			S3LoggingConfiguration: &S3LoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   				LogLocation: jsii.String("logLocation"),
//   			},
//   		},
//   		PublishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   		RemoveBytesScannedCutoffPerQuery: jsii.Boolean(false),
//   		RemoveCustomerContentEncryptionConfiguration: jsii.Boolean(false),
//   		RequesterPaysEnabled: jsii.Boolean(false),
//   		ResultConfigurationUpdates: &ResultConfigurationUpdatesProperty{
//   			AclConfiguration: &AclConfigurationProperty{
//   				S3AclOption: jsii.String("s3AclOption"),
//   			},
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				EncryptionOption: jsii.String("encryptionOption"),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputLocation: jsii.String("outputLocation"),
//   			RemoveAclConfiguration: jsii.Boolean(false),
//   			RemoveEncryptionConfiguration: jsii.Boolean(false),
//   			RemoveExpectedBucketOwner: jsii.Boolean(false),
//   			RemoveOutputLocation: jsii.Boolean(false),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html
//
type CfnWorkGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkGroupPropsMixin
type jsiiProxy_CfnWorkGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkGroupPropsMixin) Props() *CfnWorkGroupMixinProps {
	var returns *CfnWorkGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Athena::WorkGroup`.
func NewCfnWorkGroupPropsMixin(props *CfnWorkGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnWorkGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Athena::WorkGroup`.
func NewCfnWorkGroupPropsMixin_Override(c CfnWorkGroupPropsMixin, props *CfnWorkGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnWorkGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnWorkGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnWorkGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

