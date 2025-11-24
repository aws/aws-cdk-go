package mixinsawsiotfleetwise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiotfleetwise/mixinsawsiotfleetwise/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an orchestration of data collection rules.
//
// The AWS IoT FleetWise Edge Agent software running in vehicles uses campaigns to decide how to collect and transfer data to the cloud. You create campaigns in the cloud. After you or your team approve campaigns, AWS IoT FleetWise automatically deploys them to vehicles.
//
// For more information, see [Campaigns](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/campaigns.html) in the *AWS IoT FleetWise Developer Guide* .
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCampaignPropsMixin := awscdkmixinspreview.Mixins.NewCfnCampaignPropsMixin(&CfnCampaignMixinProps{
//   	Action: jsii.String("action"),
//   	CollectionScheme: &CollectionSchemeProperty{
//   		ConditionBasedCollectionScheme: &ConditionBasedCollectionSchemeProperty{
//   			ConditionLanguageVersion: jsii.Number(123),
//   			Expression: jsii.String("expression"),
//   			MinimumTriggerIntervalMs: jsii.Number(123),
//   			TriggerMode: jsii.String("triggerMode"),
//   		},
//   		TimeBasedCollectionScheme: &TimeBasedCollectionSchemeProperty{
//   			PeriodMs: jsii.Number(123),
//   		},
//   	},
//   	Compression: jsii.String("compression"),
//   	DataDestinationConfigs: []interface{}{
//   		&DataDestinationConfigProperty{
//   			MqttTopicConfig: &MqttTopicConfigProperty{
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				MqttTopicArn: jsii.String("mqttTopicArn"),
//   			},
//   			S3Config: &S3ConfigProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   				DataFormat: jsii.String("dataFormat"),
//   				Prefix: jsii.String("prefix"),
//   				StorageCompressionFormat: jsii.String("storageCompressionFormat"),
//   			},
//   			TimestreamConfig: &TimestreamConfigProperty{
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				TimestreamTableArn: jsii.String("timestreamTableArn"),
//   			},
//   		},
//   	},
//   	DataExtraDimensions: []*string{
//   		jsii.String("dataExtraDimensions"),
//   	},
//   	DataPartitions: []interface{}{
//   		&DataPartitionProperty{
//   			Id: jsii.String("id"),
//   			StorageOptions: &DataPartitionStorageOptionsProperty{
//   				MaximumSize: &StorageMaximumSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				MinimumTimeToLive: &StorageMinimumTimeToLiveProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				StorageLocation: jsii.String("storageLocation"),
//   			},
//   			UploadOptions: &DataPartitionUploadOptionsProperty{
//   				ConditionLanguageVersion: jsii.Number(123),
//   				Expression: jsii.String("expression"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DiagnosticsMode: jsii.String("diagnosticsMode"),
//   	ExpiryTime: jsii.String("expiryTime"),
//   	Name: jsii.String("name"),
//   	PostTriggerCollectionDuration: jsii.Number(123),
//   	Priority: jsii.Number(123),
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//   	SignalsToCollect: []interface{}{
//   		&SignalInformationProperty{
//   			DataPartitionId: jsii.String("dataPartitionId"),
//   			MaxSampleCount: jsii.Number(123),
//   			MinimumSamplingIntervalMs: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	SignalsToFetch: []interface{}{
//   		&SignalFetchInformationProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			ConditionLanguageVersion: jsii.Number(123),
//   			FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			SignalFetchConfig: &SignalFetchConfigProperty{
//   				ConditionBased: &ConditionBasedSignalFetchConfigProperty{
//   					ConditionExpression: jsii.String("conditionExpression"),
//   					TriggerMode: jsii.String("triggerMode"),
//   				},
//   				TimeBased: &TimeBasedSignalFetchConfigProperty{
//   					ExecutionFrequencyMs: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	SpoolingMode: jsii.String("spoolingMode"),
//   	StartTime: jsii.String("startTime"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetArn: jsii.String("targetArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html
//
type CfnCampaignPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCampaignMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCampaignPropsMixin
type jsiiProxy_CfnCampaignPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Props() *CfnCampaignMixinProps {
	var returns *CfnCampaignMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTFleetWise::Campaign`.
func NewCfnCampaignPropsMixin(props *CfnCampaignMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCampaignPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCampaignPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaignPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTFleetWise::Campaign`.
func NewCfnCampaignPropsMixin_Override(c CfnCampaignPropsMixin, props *CfnCampaignMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCampaignPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaignPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCampaignPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCampaignPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCampaignPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

