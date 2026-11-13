package previewawsdsqlmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdsql/previewawsdsqlmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DSQL::Cluster` resource specifies an cluster. You can use this resource to create, modify, and manage clusters.
//
// This resource supports both single-Region clusters and multi-Region clusters through the `MultiRegionProperties` parameter.
//
// > Creating multi-Region clusters requires additional IAM permissions beyond those needed for single-Region clusters. > - The witness Region specified in `multiRegionProperties.witnessRegion` cannot be the same as the cluster's Region.
//
// *Required permissions*
//
// - **dsql:CreateCluster** - Required to create a cluster.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
// - **dsql:TagResource** - Permission to add tags to a resource.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
// - **dsql:PutMultiRegionProperties** - Permission to configure multi-Region properties for a cluster.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
// - **dsql:AddPeerCluster** - When specifying `multiRegionProperties.clusters` , permission to add peer clusters.
//
// Resources:
//
// - Local cluster: `arn:aws:dsql:region:account-id:cluster/*`
// - Each peer cluster: exact ARN of each specified peer cluster
// - **dsql:PutWitnessRegion** - When specifying `multiRegionProperties.witnessRegion` , permission to set a witness Region. This permission is checked both in the cluster Region and in the witness Region.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
//
// Condition Keys: `dsql:WitnessRegion` (matching the specified witness region).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnClusterLogsMixin := awscdkmixinspreview.Mixins.NewCfnClusterLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html
//
type CfnClusterLogsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterLogsMixin
type jsiiProxy_CfnClusterLogsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnClusterLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::DSQL::Cluster`.
func NewCfnClusterLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnClusterLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::DSQL::Cluster`.
func NewCfnClusterLogsMixin_Override(c CfnClusterLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnClusterLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterLogsMixin_AUDIT_LOGS() CfnClusterAuditLogs {
	_init_.Initialize()
	var returns CfnClusterAuditLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterLogsMixin",
		"AUDIT_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClusterLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

