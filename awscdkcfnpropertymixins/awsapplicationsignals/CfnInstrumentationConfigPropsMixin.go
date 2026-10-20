package awsapplicationsignals

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a dynamic instrumentation configuration for a specific code or endpoint location within a service and environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnInstrumentationConfigPropsMixin := awscdkcfnpropertymixins.Aws_applicationsignals.NewCfnInstrumentationConfigPropsMixin(&CfnInstrumentationConfigMixinProps{
//   	AttributeFilters: []interface{}{
//   		map[string]*string{
//   			"attributeFiltersKey": jsii.String("attributeFilters"),
//   		},
//   	},
//   	CaptureConfiguration: &CaptureConfigurationProperty{
//   		CodeCapture: &CodeCaptureConfigurationProperty{
//   			CaptureArguments: []*string{
//   				jsii.String("captureArguments"),
//   			},
//   			CaptureLimits: &CaptureLimitsConfigProperty{
//   				MaxCollectionDepth: jsii.Number(123),
//   				MaxCollectionWidth: jsii.Number(123),
//   				MaxFieldsPerObject: jsii.Number(123),
//   				MaxHits: jsii.Number(123),
//   				MaxObjectDepth: jsii.Number(123),
//   				MaxStackFrames: jsii.Number(123),
//   				MaxStackTraceSize: jsii.Number(123),
//   				MaxStringLength: jsii.Number(123),
//   			},
//   			CaptureLocals: []*string{
//   				jsii.String("captureLocals"),
//   			},
//   			CaptureReturn: jsii.Boolean(false),
//   			CaptureStackTrace: jsii.Boolean(false),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Environment: jsii.String("environment"),
//   	ExpiresAt: jsii.String("expiresAt"),
//   	InstrumentationType: jsii.String("instrumentationType"),
//   	Location: &LocationProperty{
//   		CodeLocation: &CodeLocationProperty{
//   			ClassName: jsii.String("className"),
//   			CodeUnit: jsii.String("codeUnit"),
//   			FilePath: jsii.String("filePath"),
//   			Language: jsii.String("language"),
//   			LineNumber: jsii.Number(123),
//   			MethodName: jsii.String("methodName"),
//   		},
//   	},
//   	Service: jsii.String("service"),
//   	SignalType: jsii.String("signalType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html
//
type CfnInstrumentationConfigPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnInstrumentationConfigMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstrumentationConfigPropsMixin
type jsiiProxy_CfnInstrumentationConfigPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnInstrumentationConfigPropsMixin) Props() *CfnInstrumentationConfigMixinProps {
	var returns *CfnInstrumentationConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstrumentationConfigPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApplicationSignals::InstrumentationConfig`.
func NewCfnInstrumentationConfigPropsMixin(props *CfnInstrumentationConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnInstrumentationConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstrumentationConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstrumentationConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_applicationsignals.CfnInstrumentationConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApplicationSignals::InstrumentationConfig`.
func NewCfnInstrumentationConfigPropsMixin_Override(c CfnInstrumentationConfigPropsMixin, props *CfnInstrumentationConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_applicationsignals.CfnInstrumentationConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnInstrumentationConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstrumentationConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_applicationsignals.CfnInstrumentationConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstrumentationConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_applicationsignals.CfnInstrumentationConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstrumentationConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInstrumentationConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

