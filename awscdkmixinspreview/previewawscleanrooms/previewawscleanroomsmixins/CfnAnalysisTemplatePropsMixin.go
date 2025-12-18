package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanrooms/previewawscleanroomsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new analysis template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnalysisTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnAnalysisTemplatePropsMixin(&CfnAnalysisTemplateMixinProps{
//   	AnalysisParameters: []interface{}{
//   		&AnalysisParameterProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ErrorMessageConfiguration: &ErrorMessageConfigurationProperty{
//   		Type: jsii.String("type"),
//   	},
//   	Format: jsii.String("format"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	Schema: &AnalysisSchemaProperty{
//   		ReferencedTables: []*string{
//   			jsii.String("referencedTables"),
//   		},
//   	},
//   	Source: &AnalysisSourceProperty{
//   		Artifacts: &AnalysisTemplateArtifactsProperty{
//   			AdditionalArtifacts: []interface{}{
//   				&AnalysisTemplateArtifactProperty{
//   					Location: &S3LocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//   				},
//   			},
//   			EntryPoint: &AnalysisTemplateArtifactProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		Text: jsii.String("text"),
//   	},
//   	SourceMetadata: &AnalysisSourceMetadataProperty{
//   		Artifacts: &AnalysisTemplateArtifactMetadataProperty{
//   			AdditionalArtifactHashes: []interface{}{
//   				&HashProperty{
//   					Sha256: jsii.String("sha256"),
//   				},
//   			},
//   			EntryPointHash: &HashProperty{
//   				Sha256: jsii.String("sha256"),
//   			},
//   		},
//   	},
//   	SyntheticDataParameters: &SyntheticDataParametersProperty{
//   		MlSyntheticDataParameters: &MLSyntheticDataParametersProperty{
//   			ColumnClassification: &ColumnClassificationDetailsProperty{
//   				ColumnMapping: []interface{}{
//   					&SyntheticDataColumnPropertiesProperty{
//   						ColumnName: jsii.String("columnName"),
//   						ColumnType: jsii.String("columnType"),
//   						IsPredictiveValue: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			Epsilon: jsii.Number(123),
//   			MaxMembershipInferenceAttackScore: jsii.Number(123),
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html
//
type CfnAnalysisTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAnalysisTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAnalysisTemplatePropsMixin
type jsiiProxy_CfnAnalysisTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAnalysisTemplatePropsMixin) Props() *CfnAnalysisTemplateMixinProps {
	var returns *CfnAnalysisTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysisTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::AnalysisTemplate`.
func NewCfnAnalysisTemplatePropsMixin(props *CfnAnalysisTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAnalysisTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAnalysisTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnalysisTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::AnalysisTemplate`.
func NewCfnAnalysisTemplatePropsMixin_Override(c CfnAnalysisTemplatePropsMixin, props *CfnAnalysisTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAnalysisTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnalysisTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnalysisTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnalysisTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAnalysisTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

