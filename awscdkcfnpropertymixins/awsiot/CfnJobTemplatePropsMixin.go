package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a job template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var abortConfig interface{}
//   var jobExecutionsRolloutConfig interface{}
//   var mergeStrategy IMergeStrategy
//   var presignedUrlConfig interface{}
//   var timeoutConfig interface{}
//
//   cfnJobTemplatePropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnJobTemplatePropsMixin(&CfnJobTemplateMixinProps{
//   	AbortConfig: abortConfig,
//   	Description: jsii.String("description"),
//   	DestinationPackageVersions: []*string{
//   		jsii.String("destinationPackageVersions"),
//   	},
//   	Document: jsii.String("document"),
//   	DocumentSource: jsii.String("documentSource"),
//   	JobArn: jsii.String("jobArn"),
//   	JobExecutionsRetryConfig: &JobExecutionsRetryConfigProperty{
//   		RetryCriteriaList: []interface{}{
//   			&RetryCriteriaProperty{
//   				FailureType: jsii.String("failureType"),
//   				NumberOfRetries: jsii.Number(123),
//   			},
//   		},
//   	},
//   	JobExecutionsRolloutConfig: jobExecutionsRolloutConfig,
//   	JobTemplateId: jsii.String("jobTemplateId"),
//   	MaintenanceWindows: []interface{}{
//   		&MaintenanceWindowProperty{
//   			DurationInMinutes: jsii.Number(123),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	PresignedUrlConfig: presignedUrlConfig,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutConfig: timeoutConfig,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html
//
type CfnJobTemplatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnJobTemplateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnJobTemplatePropsMixin
type jsiiProxy_CfnJobTemplatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnJobTemplatePropsMixin) Props() *CfnJobTemplateMixinProps {
	var returns *CfnJobTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::JobTemplate`.
func NewCfnJobTemplatePropsMixin(props *CfnJobTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnJobTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnJobTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::JobTemplate`.
func NewCfnJobTemplatePropsMixin_Override(c CfnJobTemplatePropsMixin, props *CfnJobTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnJobTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnJobTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

