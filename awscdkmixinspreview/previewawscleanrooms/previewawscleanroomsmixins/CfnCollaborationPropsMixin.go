package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanrooms/previewawscleanroomsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new collaboration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCollaborationPropsMixin := awscdkmixinspreview.Mixins.NewCfnCollaborationPropsMixin(&CfnCollaborationMixinProps{
//   	AllowedResultRegions: []*string{
//   		jsii.String("allowedResultRegions"),
//   	},
//   	AnalyticsEngine: jsii.String("analyticsEngine"),
//   	AutoApprovedChangeTypes: []*string{
//   		jsii.String("autoApprovedChangeTypes"),
//   	},
//   	CreatorDisplayName: jsii.String("creatorDisplayName"),
//   	CreatorMemberAbilities: []*string{
//   		jsii.String("creatorMemberAbilities"),
//   	},
//   	CreatorMlMemberAbilities: &MLMemberAbilitiesProperty{
//   		CustomMlMemberAbilities: []*string{
//   			jsii.String("customMlMemberAbilities"),
//   		},
//   	},
//   	CreatorPaymentConfiguration: &PaymentConfigurationProperty{
//   		JobCompute: &JobComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   		MachineLearning: &MLPaymentConfigProperty{
//   			ModelInference: &ModelInferencePaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   			ModelTraining: &ModelTrainingPaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   		},
//   		QueryCompute: &QueryComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   	},
//   	DataEncryptionMetadata: &DataEncryptionMetadataProperty{
//   		AllowCleartext: jsii.Boolean(false),
//   		AllowDuplicates: jsii.Boolean(false),
//   		AllowJoinsOnColumnsWithDifferentNames: jsii.Boolean(false),
//   		PreserveNulls: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	JobLogStatus: jsii.String("jobLogStatus"),
//   	Members: []interface{}{
//   		&MemberSpecificationProperty{
//   			AccountId: jsii.String("accountId"),
//   			DisplayName: jsii.String("displayName"),
//   			MemberAbilities: []*string{
//   				jsii.String("memberAbilities"),
//   			},
//   			MlMemberAbilities: &MLMemberAbilitiesProperty{
//   				CustomMlMemberAbilities: []*string{
//   					jsii.String("customMlMemberAbilities"),
//   				},
//   			},
//   			PaymentConfiguration: &PaymentConfigurationProperty{
//   				JobCompute: &JobComputePaymentConfigProperty{
//   					IsResponsible: jsii.Boolean(false),
//   				},
//   				MachineLearning: &MLPaymentConfigProperty{
//   					ModelInference: &ModelInferencePaymentConfigProperty{
//   						IsResponsible: jsii.Boolean(false),
//   					},
//   					ModelTraining: &ModelTrainingPaymentConfigProperty{
//   						IsResponsible: jsii.Boolean(false),
//   					},
//   				},
//   				QueryCompute: &QueryComputePaymentConfigProperty{
//   					IsResponsible: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html
//
type CfnCollaborationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCollaborationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCollaborationPropsMixin
type jsiiProxy_CfnCollaborationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCollaborationPropsMixin) Props() *CfnCollaborationMixinProps {
	var returns *CfnCollaborationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCollaborationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::Collaboration`.
func NewCfnCollaborationPropsMixin(props *CfnCollaborationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCollaborationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCollaborationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCollaborationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::Collaboration`.
func NewCfnCollaborationPropsMixin_Override(c CfnCollaborationPropsMixin, props *CfnCollaborationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCollaborationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCollaborationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCollaborationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCollaborationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCollaborationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

