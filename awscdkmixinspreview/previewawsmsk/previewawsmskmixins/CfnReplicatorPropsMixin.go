package previewawsmskmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmsk/previewawsmskmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates the replicator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReplicatorPropsMixin := awscdkmixinspreview.Mixins.NewCfnReplicatorPropsMixin(&CfnReplicatorMixinProps{
//   	Description: jsii.String("description"),
//   	KafkaClusters: []interface{}{
//   		&KafkaClusterProperty{
//   			AmazonMskCluster: &AmazonMskClusterProperty{
//   				MskClusterArn: jsii.String("mskClusterArn"),
//   			},
//   			VpcConfig: &KafkaClusterClientVpcConfigProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   	},
//   	ReplicationInfoList: []interface{}{
//   		&ReplicationInfoProperty{
//   			ConsumerGroupReplication: &ConsumerGroupReplicationProperty{
//   				ConsumerGroupsToExclude: []*string{
//   					jsii.String("consumerGroupsToExclude"),
//   				},
//   				ConsumerGroupsToReplicate: []*string{
//   					jsii.String("consumerGroupsToReplicate"),
//   				},
//   				DetectAndCopyNewConsumerGroups: jsii.Boolean(false),
//   				SynchroniseConsumerGroupOffsets: jsii.Boolean(false),
//   			},
//   			SourceKafkaClusterArn: jsii.String("sourceKafkaClusterArn"),
//   			TargetCompressionType: jsii.String("targetCompressionType"),
//   			TargetKafkaClusterArn: jsii.String("targetKafkaClusterArn"),
//   			TopicReplication: &TopicReplicationProperty{
//   				CopyAccessControlListsForTopics: jsii.Boolean(false),
//   				CopyTopicConfigurations: jsii.Boolean(false),
//   				DetectAndCopyNewTopics: jsii.Boolean(false),
//   				StartingPosition: &ReplicationStartingPositionProperty{
//   					Type: jsii.String("type"),
//   				},
//   				TopicNameConfiguration: &ReplicationTopicNameConfigurationProperty{
//   					Type: jsii.String("type"),
//   				},
//   				TopicsToExclude: []*string{
//   					jsii.String("topicsToExclude"),
//   				},
//   				TopicsToReplicate: []*string{
//   					jsii.String("topicsToReplicate"),
//   				},
//   			},
//   		},
//   	},
//   	ReplicatorName: jsii.String("replicatorName"),
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html
//
type CfnReplicatorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnReplicatorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReplicatorPropsMixin
type jsiiProxy_CfnReplicatorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnReplicatorPropsMixin) Props() *CfnReplicatorMixinProps {
	var returns *CfnReplicatorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicatorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MSK::Replicator`.
func NewCfnReplicatorPropsMixin(props *CfnReplicatorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnReplicatorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReplicatorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicatorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MSK::Replicator`.
func NewCfnReplicatorPropsMixin_Override(c CfnReplicatorPropsMixin, props *CfnReplicatorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnReplicatorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicatorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicatorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicatorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnReplicatorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

