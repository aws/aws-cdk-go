package awshealthlake

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awshealthlake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Data Transformation Profile in AWS HealthLake that converts healthcare data from a source format (such as C-CDA or CSV) into FHIR R4.
//
// A profile is immutable once created; to change its template content, replace the resource. Only its tags can be updated in place.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDataTransformationProfilePropsMixin := awscdkcfnpropertymixins.Aws_healthlake.NewCfnDataTransformationProfilePropsMixin(&CfnDataTransformationProfileMixinProps{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ProfileDescription: jsii.String("profileDescription"),
//   	ProfileName: jsii.String("profileName"),
//   	Source: &SourceProperty{
//   		ExistingVersionedProfileId: &ExistingVersionedProfileSourceProperty{
//   			ProfileId: jsii.String("profileId"),
//   			Version: jsii.Number(123),
//   		},
//   		ProfileMapping: &ProfileMappingSourceProperty{
//   			ProfileMapping: map[string]*string{
//   				"profileMappingKey": jsii.String("profileMapping"),
//   			},
//   		},
//   		StarterProfile: &StarterProfileSourceProperty{
//   			StarterProfileName: jsii.String("starterProfileName"),
//   		},
//   	},
//   	SourceFormat: jsii.String("sourceFormat"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html
//
type CfnDataTransformationProfilePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDataTransformationProfileMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataTransformationProfilePropsMixin
type jsiiProxy_CfnDataTransformationProfilePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDataTransformationProfilePropsMixin) Props() *CfnDataTransformationProfileMixinProps {
	var returns *CfnDataTransformationProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataTransformationProfilePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::HealthLake::DataTransformationProfile`.
func NewCfnDataTransformationProfilePropsMixin(props *CfnDataTransformationProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDataTransformationProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataTransformationProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataTransformationProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::HealthLake::DataTransformationProfile`.
func NewCfnDataTransformationProfilePropsMixin_Override(c CfnDataTransformationProfilePropsMixin, props *CfnDataTransformationProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDataTransformationProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataTransformationProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataTransformationProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataTransformationProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataTransformationProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

