package awslightsail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lightsail::Container` resource specifies a container service.
//
// A Lightsail container service is a compute resource to which you can deploy containers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnContainerPropsMixin := awscdkcfnpropertymixins.Aws_lightsail.NewCfnContainerPropsMixin(&CfnContainerMixinProps{
//   	ContainerServiceDeployment: &ContainerServiceDeploymentProperty{
//   		Containers: []interface{}{
//   			&ContainerProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				ContainerName: jsii.String("containerName"),
//   				Environment: []interface{}{
//   					&EnvironmentVariableProperty{
//   						Value: jsii.String("value"),
//   						Variable: jsii.String("variable"),
//   					},
//   				},
//   				Image: jsii.String("image"),
//   				Ports: []interface{}{
//   					&PortInfoProperty{
//   						Port: jsii.String("port"),
//   						Protocol: jsii.String("protocol"),
//   					},
//   				},
//   			},
//   		},
//   		PublicEndpoint: &PublicEndpointProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			HealthCheckConfig: &HealthCheckConfigProperty{
//   				HealthyThreshold: jsii.Number(123),
//   				IntervalSeconds: jsii.Number(123),
//   				Path: jsii.String("path"),
//   				SuccessCodes: jsii.String("successCodes"),
//   				TimeoutSeconds: jsii.Number(123),
//   				UnhealthyThreshold: jsii.Number(123),
//   			},
//   		},
//   	},
//   	IsDisabled: jsii.Boolean(false),
//   	Power: jsii.String("power"),
//   	PrivateRegistryAccess: &PrivateRegistryAccessProperty{
//   		EcrImagePullerRole: &EcrImagePullerRoleProperty{
//   			IsActive: jsii.Boolean(false),
//   			PrincipalArn: jsii.String("principalArn"),
//   		},
//   	},
//   	PublicDomainNames: []interface{}{
//   		&PublicDomainNameProperty{
//   			CertificateName: jsii.String("certificateName"),
//   			DomainNames: []*string{
//   				jsii.String("domainNames"),
//   			},
//   		},
//   	},
//   	Scale: jsii.Number(123),
//   	ServiceName: jsii.String("serviceName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html
//
type CfnContainerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnContainerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContainerPropsMixin
type jsiiProxy_CfnContainerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnContainerPropsMixin) Props() *CfnContainerMixinProps {
	var returns *CfnContainerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lightsail::Container`.
func NewCfnContainerPropsMixin(props *CfnContainerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnContainerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContainerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::Container`.
func NewCfnContainerPropsMixin_Override(c CfnContainerPropsMixin, props *CfnContainerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnContainerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnContainerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

