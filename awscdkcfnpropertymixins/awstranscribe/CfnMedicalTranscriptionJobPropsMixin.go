package awstranscribe

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awstranscribe/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for an Amazon Transcribe Medical Transcription Job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMedicalTranscriptionJobPropsMixin := awscdkcfnpropertymixins.Aws_transcribe.NewCfnMedicalTranscriptionJobPropsMixin(&CfnMedicalTranscriptionJobMixinProps{
//   	LanguageCode: jsii.String("languageCode"),
//   	Media: &MediaProperty{
//   		MediaFileUri: jsii.String("mediaFileUri"),
//   	},
//   	MediaFormat: jsii.String("mediaFormat"),
//   	MediaSampleRateHertz: jsii.Number(123),
//   	MedicalTranscriptionJobName: jsii.String("medicalTranscriptionJobName"),
//   	Settings: &MedicalTranscriptionSettingProperty{
//   		ChannelIdentification: jsii.Boolean(false),
//   		ShowAlternatives: jsii.Boolean(false),
//   	},
//   	Specialty: jsii.String("specialty"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicaltranscriptionjob.html
//
type CfnMedicalTranscriptionJobPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMedicalTranscriptionJobMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMedicalTranscriptionJobPropsMixin
type jsiiProxy_CfnMedicalTranscriptionJobPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMedicalTranscriptionJobPropsMixin) Props() *CfnMedicalTranscriptionJobMixinProps {
	var returns *CfnMedicalTranscriptionJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMedicalTranscriptionJobPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transcribe::MedicalTranscriptionJob`.
func NewCfnMedicalTranscriptionJobPropsMixin(props *CfnMedicalTranscriptionJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMedicalTranscriptionJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMedicalTranscriptionJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMedicalTranscriptionJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transcribe::MedicalTranscriptionJob`.
func NewCfnMedicalTranscriptionJobPropsMixin_Override(c CfnMedicalTranscriptionJobPropsMixin, props *CfnMedicalTranscriptionJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMedicalTranscriptionJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMedicalTranscriptionJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMedicalTranscriptionJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMedicalTranscriptionJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMedicalTranscriptionJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

