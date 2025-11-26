package previewawssimspaceweavermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssimspaceweaver/previewawssimspaceweavermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::SimSpaceWeaver::Simulation` resource to specify a simulation that CloudFormation starts in the AWS Cloud , in your AWS account .
//
// In the resource properties section of your template, provide the name of an existing IAM role configured with the proper permissions, and the name of an existing Amazon S3 bucket. Your account must have permissions to read the Amazon S3 bucket. The Amazon S3 bucket must contain a valid schema. The schema must refer to simulation assets that are already uploaded to the AWS Cloud . For more information, see the [detailed tutorial](https://docs.aws.amazon.com/simspaceweaver/latest/userguide/getting-started_detailed.html) in the *AWS SimSpace Weaver User Guide* .
//
// Specify a `SnapshotS3Location` to start a simulation from a snapshot instead of from a schema. When you start a simulation from a snapshot, SimSpace Weaver initializes the entity data in the State Fabric with data saved in the snapshot, starts the spatial and service apps that were running when the snapshot was created, and restores the clock to the appropriate tick. Your app zip files must be in the same location in Amazon S3 as they were in for the original simulation. You must start any custom apps separately. For more information about snapshots, see [Snapshots](https://docs.aws.amazon.com/simspaceweaver/latest/userguide/working-with_snapshots.html) in the *AWS SimSpace Weaver User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSimulationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSimulationPropsMixin(&CfnSimulationMixinProps{
//   	MaximumDuration: jsii.String("maximumDuration"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	SchemaS3Location: &S3LocationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//   	},
//   	SnapshotS3Location: &S3LocationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-simspaceweaver-simulation.html
//
type CfnSimulationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSimulationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSimulationPropsMixin
type jsiiProxy_CfnSimulationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSimulationPropsMixin) Props() *CfnSimulationMixinProps {
	var returns *CfnSimulationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimulationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SimSpaceWeaver::Simulation`.
func NewCfnSimulationPropsMixin(props *CfnSimulationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSimulationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSimulationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSimulationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_simspaceweaver.mixins.CfnSimulationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SimSpaceWeaver::Simulation`.
func NewCfnSimulationPropsMixin_Override(c CfnSimulationPropsMixin, props *CfnSimulationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_simspaceweaver.mixins.CfnSimulationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSimulationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSimulationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_simspaceweaver.mixins.CfnSimulationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSimulationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_simspaceweaver.mixins.CfnSimulationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSimulationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSimulationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

