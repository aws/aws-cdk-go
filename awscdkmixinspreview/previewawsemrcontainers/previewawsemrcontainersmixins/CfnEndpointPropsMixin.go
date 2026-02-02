package previewawsemrcontainersmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsemrcontainers/previewawsemrcontainersmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Schema of AWS::EMRContainers::Endpoint Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var eMREKSConfigurationProperty_ EMREKSConfigurationProperty
//
//   cfnEndpointPropsMixin := awscdkmixinspreview.Mixins.NewCfnEndpointPropsMixin(&CfnEndpointMixinProps{
//   	ConfigurationOverrides: &ConfigurationOverridesProperty{
//   		ApplicationConfiguration: []interface{}{
//   			&EMREKSConfigurationProperty{
//   				Classification: jsii.String("classification"),
//   				Configurations: []interface{}{
//   					eMREKSConfigurationProperty_,
//   				},
//   				Properties: map[string]*string{
//   					"propertiesKey": jsii.String("properties"),
//   				},
//   			},
//   		},
//   		MonitoringConfiguration: &MonitoringConfigurationProperty{
//   			CloudWatchMonitoringConfiguration: &CloudWatchMonitoringConfigurationProperty{
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   			},
//   			ContainerLogRotationConfiguration: &ContainerLogRotationConfigurationProperty{
//   				MaxFilesToKeep: jsii.Number(123),
//   				RotationSize: jsii.String("rotationSize"),
//   			},
//   			PersistentAppUi: jsii.String("persistentAppUi"),
//   			S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   				LogUri: jsii.String("logUri"),
//   			},
//   		},
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	VirtualClusterId: jsii.String("virtualClusterId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html
//
type CfnEndpointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEndpointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEndpointPropsMixin
type jsiiProxy_CfnEndpointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEndpointPropsMixin) Props() *CfnEndpointMixinProps {
	var returns *CfnEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMRContainers::Endpoint`.
func NewCfnEndpointPropsMixin(props *CfnEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMRContainers::Endpoint`.
func NewCfnEndpointPropsMixin_Override(c CfnEndpointPropsMixin, props *CfnEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

