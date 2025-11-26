package previewawscassandramixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscassandra/previewawscassandramixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// You can use the `AWS::Cassandra::Table` resource to create a new table in Amazon Keyspaces (for Apache Cassandra).
//
// For more information, see [Create a table](https://docs.aws.amazon.com/keyspaces/latest/devguide/getting-started.tables.html) in the *Amazon Keyspaces Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTablePropsMixin := awscdkmixinspreview.Mixins.NewCfnTablePropsMixin(&CfnTableMixinProps{
//   	AutoScalingSpecifications: &AutoScalingSpecificationProperty{
//   		ReadCapacityAutoScaling: &AutoScalingSettingProperty{
//   			AutoScalingDisabled: jsii.Boolean(false),
//   			MaximumUnits: jsii.Number(123),
//   			MinimumUnits: jsii.Number(123),
//   			ScalingPolicy: &ScalingPolicyProperty{
//   				TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   					DisableScaleIn: jsii.Boolean(false),
//   					ScaleInCooldown: jsii.Number(123),
//   					ScaleOutCooldown: jsii.Number(123),
//   					TargetValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   		WriteCapacityAutoScaling: &AutoScalingSettingProperty{
//   			AutoScalingDisabled: jsii.Boolean(false),
//   			MaximumUnits: jsii.Number(123),
//   			MinimumUnits: jsii.Number(123),
//   			ScalingPolicy: &ScalingPolicyProperty{
//   				TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   					DisableScaleIn: jsii.Boolean(false),
//   					ScaleInCooldown: jsii.Number(123),
//   					ScaleOutCooldown: jsii.Number(123),
//   					TargetValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	BillingMode: &BillingModeProperty{
//   		Mode: jsii.String("mode"),
//   		ProvisionedThroughput: &ProvisionedThroughputProperty{
//   			ReadCapacityUnits: jsii.Number(123),
//   			WriteCapacityUnits: jsii.Number(123),
//   		},
//   	},
//   	CdcSpecification: &CdcSpecificationProperty{
//   		Status: jsii.String("status"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ViewType: jsii.String("viewType"),
//   	},
//   	ClientSideTimestampsEnabled: jsii.Boolean(false),
//   	ClusteringKeyColumns: []interface{}{
//   		&ClusteringKeyColumnProperty{
//   			Column: &ColumnProperty{
//   				ColumnName: jsii.String("columnName"),
//   				ColumnType: jsii.String("columnType"),
//   			},
//   			OrderBy: jsii.String("orderBy"),
//   		},
//   	},
//   	DefaultTimeToLive: jsii.Number(123),
//   	EncryptionSpecification: &EncryptionSpecificationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	},
//   	KeyspaceName: jsii.String("keyspaceName"),
//   	PartitionKeyColumns: []interface{}{
//   		&ColumnProperty{
//   			ColumnName: jsii.String("columnName"),
//   			ColumnType: jsii.String("columnType"),
//   		},
//   	},
//   	PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	RegularColumns: []interface{}{
//   		&ColumnProperty{
//   			ColumnName: jsii.String("columnName"),
//   			ColumnType: jsii.String("columnType"),
//   		},
//   	},
//   	ReplicaSpecifications: []interface{}{
//   		&ReplicaSpecificationProperty{
//   			ReadCapacityAutoScaling: &AutoScalingSettingProperty{
//   				AutoScalingDisabled: jsii.Boolean(false),
//   				MaximumUnits: jsii.Number(123),
//   				MinimumUnits: jsii.Number(123),
//   				ScalingPolicy: &ScalingPolicyProperty{
//   					TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   						DisableScaleIn: jsii.Boolean(false),
//   						ScaleInCooldown: jsii.Number(123),
//   						ScaleOutCooldown: jsii.Number(123),
//   						TargetValue: jsii.Number(123),
//   					},
//   				},
//   			},
//   			ReadCapacityUnits: jsii.Number(123),
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	TableName: jsii.String("tableName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WarmThroughput: &WarmThroughputProperty{
//   		ReadUnitsPerSecond: jsii.Number(123),
//   		WriteUnitsPerSecond: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html
//
type CfnTablePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTableMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTablePropsMixin
type jsiiProxy_CfnTablePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTablePropsMixin) Props() *CfnTableMixinProps {
	var returns *CfnTableMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTablePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cassandra::Table`.
func NewCfnTablePropsMixin(props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cassandra::Table`.
func NewCfnTablePropsMixin_Override(c CfnTablePropsMixin, props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTablePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTablePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTablePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cassandra.mixins.CfnTablePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTablePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTablePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

