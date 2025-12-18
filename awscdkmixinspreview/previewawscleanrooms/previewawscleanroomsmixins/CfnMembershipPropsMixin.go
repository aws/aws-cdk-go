package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanrooms/previewawscleanroomsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a membership for a specific collaboration identifier and joins the collaboration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMembershipPropsMixin := awscdkmixinspreview.Mixins.NewCfnMembershipPropsMixin(&CfnMembershipMixinProps{
//   	CollaborationIdentifier: jsii.String("collaborationIdentifier"),
//   	DefaultJobResultConfiguration: &MembershipProtectedJobResultConfigurationProperty{
//   		OutputConfiguration: &MembershipProtectedJobOutputConfigurationProperty{
//   			S3: &ProtectedJobS3OutputConfigurationInputProperty{
//   				Bucket: jsii.String("bucket"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	DefaultResultConfiguration: &MembershipProtectedQueryResultConfigurationProperty{
//   		OutputConfiguration: &MembershipProtectedQueryOutputConfigurationProperty{
//   			S3: &ProtectedQueryS3OutputConfigurationProperty{
//   				Bucket: jsii.String("bucket"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   				ResultFormat: jsii.String("resultFormat"),
//   				SingleFileOutput: jsii.Boolean(false),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	JobLogStatus: jsii.String("jobLogStatus"),
//   	PaymentConfiguration: &MembershipPaymentConfigurationProperty{
//   		JobCompute: &MembershipJobComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   		MachineLearning: &MembershipMLPaymentConfigProperty{
//   			ModelInference: &MembershipModelInferencePaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   			ModelTraining: &MembershipModelTrainingPaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   			SyntheticDataGeneration: &MembershipSyntheticDataGenerationPaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   		},
//   		QueryCompute: &MembershipQueryComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   	},
//   	QueryLogStatus: jsii.String("queryLogStatus"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-membership.html
//
type CfnMembershipPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMembershipMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMembershipPropsMixin
type jsiiProxy_CfnMembershipPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMembershipPropsMixin) Props() *CfnMembershipMixinProps {
	var returns *CfnMembershipMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMembershipPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::Membership`.
func NewCfnMembershipPropsMixin(props *CfnMembershipMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMembershipPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMembershipPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMembershipPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::Membership`.
func NewCfnMembershipPropsMixin_Override(c CfnMembershipPropsMixin, props *CfnMembershipMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMembershipPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMembershipPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMembershipPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMembershipPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMembershipPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

