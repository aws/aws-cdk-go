package mixinsawsfrauddetector

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsfrauddetector/mixinsawsfrauddetector/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manages a detector and associated detector versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDetectorPropsMixin := awscdkmixinspreview.Mixins.NewCfnDetectorPropsMixin(&CfnDetectorMixinProps{
//   	AssociatedModels: []interface{}{
//   		&ModelProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DetectorId: jsii.String("detectorId"),
//   	DetectorVersionStatus: jsii.String("detectorVersionStatus"),
//   	EventType: &EventTypeProperty{
//   		Arn: jsii.String("arn"),
//   		CreatedTime: jsii.String("createdTime"),
//   		Description: jsii.String("description"),
//   		EntityTypes: []interface{}{
//   			&EntityTypeProperty{
//   				Arn: jsii.String("arn"),
//   				CreatedTime: jsii.String("createdTime"),
//   				Description: jsii.String("description"),
//   				Inline: jsii.Boolean(false),
//   				LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				Name: jsii.String("name"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		EventVariables: []interface{}{
//   			&EventVariableProperty{
//   				Arn: jsii.String("arn"),
//   				CreatedTime: jsii.String("createdTime"),
//   				DataSource: jsii.String("dataSource"),
//   				DataType: jsii.String("dataType"),
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				Inline: jsii.Boolean(false),
//   				LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				Name: jsii.String("name"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				VariableType: jsii.String("variableType"),
//   			},
//   		},
//   		Inline: jsii.Boolean(false),
//   		Labels: []interface{}{
//   			&LabelProperty{
//   				Arn: jsii.String("arn"),
//   				CreatedTime: jsii.String("createdTime"),
//   				Description: jsii.String("description"),
//   				Inline: jsii.Boolean(false),
//   				LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				Name: jsii.String("name"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   		Name: jsii.String("name"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	RuleExecutionMode: jsii.String("ruleExecutionMode"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			DetectorId: jsii.String("detectorId"),
//   			Expression: jsii.String("expression"),
//   			Language: jsii.String("language"),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Outcomes: []interface{}{
//   				&OutcomeProperty{
//   					Arn: jsii.String("arn"),
//   					CreatedTime: jsii.String("createdTime"),
//   					Description: jsii.String("description"),
//   					Inline: jsii.Boolean(false),
//   					LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   					Name: jsii.String("name"),
//   					Tags: []CfnTag{
//   						&CfnTag{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			RuleId: jsii.String("ruleId"),
//   			RuleVersion: jsii.String("ruleVersion"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html
//
type CfnDetectorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDetectorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDetectorPropsMixin
type jsiiProxy_CfnDetectorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDetectorPropsMixin) Props() *CfnDetectorMixinProps {
	var returns *CfnDetectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FraudDetector::Detector`.
func NewCfnDetectorPropsMixin(props *CfnDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDetectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDetectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDetectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnDetectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FraudDetector::Detector`.
func NewCfnDetectorPropsMixin_Override(c CfnDetectorPropsMixin, props *CfnDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnDetectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDetectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDetectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnDetectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDetectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnDetectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDetectorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDetectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

