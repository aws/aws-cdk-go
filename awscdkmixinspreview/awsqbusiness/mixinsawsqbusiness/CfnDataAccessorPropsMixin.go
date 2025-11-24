package mixinsawsqbusiness

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsqbusiness/mixinsawsqbusiness/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new data accessor for an ISV to access data from a Amazon Q Business application.
//
// The data accessor is an entity that represents the ISV's access to the Amazon Q Business application's data. It includes the IAM role ARN for the ISV, a friendly name, and a set of action configurations that define the specific actions the ISV is allowed to perform and any associated data filters. When the data accessor is created, an IAM Identity Center application is also created to manage the ISV's identity and authentication for accessing the Amazon Q Business application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var attributeFilterProperty_ AttributeFilterProperty
//
//   cfnDataAccessorPropsMixin := awscdkmixinspreview.Mixins.NewCfnDataAccessorPropsMixin(&CfnDataAccessorMixinProps{
//   	ActionConfigurations: []interface{}{
//   		&ActionConfigurationProperty{
//   			Action: jsii.String("action"),
//   			FilterConfiguration: &ActionFilterConfigurationProperty{
//   				DocumentAttributeFilter: &AttributeFilterProperty{
//   					AndAllFilters: []interface{}{
//   						attributeFilterProperty_,
//   					},
//   					ContainsAll: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					ContainsAny: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					EqualsTo: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					GreaterThan: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					GreaterThanOrEquals: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					LessThan: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					LessThanOrEquals: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					NotFilter: attributeFilterProperty_,
//   					OrAllFilters: []interface{}{
//   						attributeFilterProperty_,
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ApplicationId: jsii.String("applicationId"),
//   	AuthenticationDetail: &DataAccessorAuthenticationDetailProperty{
//   		AuthenticationConfiguration: &DataAccessorAuthenticationConfigurationProperty{
//   			IdcTrustedTokenIssuerConfiguration: &DataAccessorIdcTrustedTokenIssuerConfigurationProperty{
//   				IdcTrustedTokenIssuerArn: jsii.String("idcTrustedTokenIssuerArn"),
//   			},
//   		},
//   		AuthenticationType: jsii.String("authenticationType"),
//   		ExternalIds: []*string{
//   			jsii.String("externalIds"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Principal: jsii.String("principal"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html
//
type CfnDataAccessorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataAccessorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataAccessorPropsMixin
type jsiiProxy_CfnDataAccessorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataAccessorPropsMixin) Props() *CfnDataAccessorMixinProps {
	var returns *CfnDataAccessorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataAccessorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QBusiness::DataAccessor`.
func NewCfnDataAccessorPropsMixin(props *CfnDataAccessorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataAccessorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataAccessorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataAccessorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QBusiness::DataAccessor`.
func NewCfnDataAccessorPropsMixin_Override(c CfnDataAccessorPropsMixin, props *CfnDataAccessorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataAccessorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataAccessorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataAccessorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataAccessorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDataAccessorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

