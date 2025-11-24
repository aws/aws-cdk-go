package mixinsawsb2bi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsb2bi/mixinsawsb2bi/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a partnership between a customer and a trading partner, based on the supplied parameters.
//
// A partnership represents the connection between you and your trading partner. It ties together a profile and one or more trading capabilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPartnershipPropsMixin := awscdkmixinspreview.Mixins.NewCfnPartnershipPropsMixin(&CfnPartnershipMixinProps{
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	CapabilityOptions: &CapabilityOptionsProperty{
//   		InboundEdi: &InboundEdiOptionsProperty{
//   			X12: &X12InboundEdiOptionsProperty{
//   				AcknowledgmentOptions: &X12AcknowledgmentOptionsProperty{
//   					FunctionalAcknowledgment: jsii.String("functionalAcknowledgment"),
//   					TechnicalAcknowledgment: jsii.String("technicalAcknowledgment"),
//   				},
//   			},
//   		},
//   		OutboundEdi: &OutboundEdiOptionsProperty{
//   			X12: &X12EnvelopeProperty{
//   				Common: &X12OutboundEdiHeadersProperty{
//   					ControlNumbers: &X12ControlNumbersProperty{
//   						StartingFunctionalGroupControlNumber: jsii.Number(123),
//   						StartingInterchangeControlNumber: jsii.Number(123),
//   						StartingTransactionSetControlNumber: jsii.Number(123),
//   					},
//   					Delimiters: &X12DelimitersProperty{
//   						ComponentSeparator: jsii.String("componentSeparator"),
//   						DataElementSeparator: jsii.String("dataElementSeparator"),
//   						SegmentTerminator: jsii.String("segmentTerminator"),
//   					},
//   					FunctionalGroupHeaders: &X12FunctionalGroupHeadersProperty{
//   						ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   						ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   						ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   					},
//   					Gs05TimeFormat: jsii.String("gs05TimeFormat"),
//   					InterchangeControlHeaders: &X12InterchangeControlHeadersProperty{
//   						AcknowledgmentRequestedCode: jsii.String("acknowledgmentRequestedCode"),
//   						ReceiverId: jsii.String("receiverId"),
//   						ReceiverIdQualifier: jsii.String("receiverIdQualifier"),
//   						RepetitionSeparator: jsii.String("repetitionSeparator"),
//   						SenderId: jsii.String("senderId"),
//   						SenderIdQualifier: jsii.String("senderIdQualifier"),
//   						UsageIndicatorCode: jsii.String("usageIndicatorCode"),
//   					},
//   					ValidateEdi: jsii.Boolean(false),
//   				},
//   				WrapOptions: &WrapOptionsProperty{
//   					LineLength: jsii.Number(123),
//   					LineTerminator: jsii.String("lineTerminator"),
//   					WrapBy: jsii.String("wrapBy"),
//   				},
//   			},
//   		},
//   	},
//   	Email: jsii.String("email"),
//   	Name: jsii.String("name"),
//   	Phone: jsii.String("phone"),
//   	ProfileId: jsii.String("profileId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html
//
type CfnPartnershipPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPartnershipMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPartnershipPropsMixin
type jsiiProxy_CfnPartnershipPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPartnershipPropsMixin) Props() *CfnPartnershipMixinProps {
	var returns *CfnPartnershipMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartnershipPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::B2BI::Partnership`.
func NewCfnPartnershipPropsMixin(props *CfnPartnershipMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPartnershipPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPartnershipPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPartnershipPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::B2BI::Partnership`.
func NewCfnPartnershipPropsMixin_Override(c CfnPartnershipPropsMixin, props *CfnPartnershipMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPartnershipPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPartnershipPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPartnershipPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnPartnershipPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPartnershipPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPartnershipPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

