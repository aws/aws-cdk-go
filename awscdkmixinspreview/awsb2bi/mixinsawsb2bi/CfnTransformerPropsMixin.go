package mixinsawsb2bi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsb2bi/mixinsawsb2bi/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a transformer. AWS B2B Data Interchange currently supports two scenarios:.
//
// - *Inbound EDI* : the AWS customer receives an EDI file from their trading partner. AWS B2B Data Interchange converts this EDI file into a JSON or XML file with a service-defined structure. A mapping template provided by the customer, in JSONata or XSLT format, is optionally applied to this file to produce a JSON or XML file with the structure the customer requires.
// - *Outbound EDI* : the AWS customer has a JSON or XML file containing data that they wish to use in an EDI file. A mapping template, provided by the customer (in either JSONata or XSLT format) is applied to this file to generate a JSON or XML file in the service-defined structure. This file is then converted to an EDI file.
//
// > The following fields are provided for backwards compatibility only: `fileFormat` , `mappingTemplate` , `ediType` , and `sampleDocument` .
// >
// > - Use the `mapping` data type in place of `mappingTemplate` and `fileFormat`
// > - Use the `sampleDocuments` data type in place of `sampleDocument`
// > - Use either the `inputConversion` or `outputConversion` in place of `ediType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerPropsMixin := awscdkmixinspreview.Mixins.NewCfnTransformerPropsMixin(&CfnTransformerMixinProps{
//   	EdiType: &EdiTypeProperty{
//   		X12Details: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	FileFormat: jsii.String("fileFormat"),
//   	InputConversion: &InputConversionProperty{
//   		AdvancedOptions: &AdvancedOptionsProperty{
//   			X12: &X12AdvancedOptionsProperty{
//   				SplitOptions: &X12SplitOptionsProperty{
//   					SplitBy: jsii.String("splitBy"),
//   				},
//   				ValidationOptions: &X12ValidationOptionsProperty{
//   					ValidationRules: []interface{}{
//   						&X12ValidationRuleProperty{
//   							CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   								CodesToAdd: []*string{
//   									jsii.String("codesToAdd"),
//   								},
//   								CodesToRemove: []*string{
//   									jsii.String("codesToRemove"),
//   								},
//   								ElementId: jsii.String("elementId"),
//   							},
//   							ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   								ElementId: jsii.String("elementId"),
//   								MaxLength: jsii.Number(123),
//   								MinLength: jsii.Number(123),
//   							},
//   							ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   								ElementPosition: jsii.String("elementPosition"),
//   								Requirement: jsii.String("requirement"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		FormatOptions: &FormatOptionsProperty{
//   			X12: &X12DetailsProperty{
//   				TransactionSet: jsii.String("transactionSet"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   		FromFormat: jsii.String("fromFormat"),
//   	},
//   	Mapping: &MappingProperty{
//   		Template: jsii.String("template"),
//   		TemplateLanguage: jsii.String("templateLanguage"),
//   	},
//   	MappingTemplate: jsii.String("mappingTemplate"),
//   	Name: jsii.String("name"),
//   	OutputConversion: &OutputConversionProperty{
//   		AdvancedOptions: &AdvancedOptionsProperty{
//   			X12: &X12AdvancedOptionsProperty{
//   				SplitOptions: &X12SplitOptionsProperty{
//   					SplitBy: jsii.String("splitBy"),
//   				},
//   				ValidationOptions: &X12ValidationOptionsProperty{
//   					ValidationRules: []interface{}{
//   						&X12ValidationRuleProperty{
//   							CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   								CodesToAdd: []*string{
//   									jsii.String("codesToAdd"),
//   								},
//   								CodesToRemove: []*string{
//   									jsii.String("codesToRemove"),
//   								},
//   								ElementId: jsii.String("elementId"),
//   							},
//   							ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   								ElementId: jsii.String("elementId"),
//   								MaxLength: jsii.Number(123),
//   								MinLength: jsii.Number(123),
//   							},
//   							ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   								ElementPosition: jsii.String("elementPosition"),
//   								Requirement: jsii.String("requirement"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		FormatOptions: &FormatOptionsProperty{
//   			X12: &X12DetailsProperty{
//   				TransactionSet: jsii.String("transactionSet"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   		ToFormat: jsii.String("toFormat"),
//   	},
//   	SampleDocument: jsii.String("sampleDocument"),
//   	SampleDocuments: &SampleDocumentsProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Keys: []interface{}{
//   			&SampleDocumentKeysProperty{
//   				Input: jsii.String("input"),
//   				Output: jsii.String("output"),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html
//
type CfnTransformerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTransformerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransformerPropsMixin
type jsiiProxy_CfnTransformerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTransformerPropsMixin) Props() *CfnTransformerMixinProps {
	var returns *CfnTransformerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::B2BI::Transformer`.
func NewCfnTransformerPropsMixin(props *CfnTransformerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTransformerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransformerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransformerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::B2BI::Transformer`.
func NewCfnTransformerPropsMixin_Override(c CfnTransformerPropsMixin, props *CfnTransformerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTransformerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransformerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransformerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTransformerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

