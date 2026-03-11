package awsimagebuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new container recipe.
//
// Container recipes define how images are configured, tested, and assessed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnContainerRecipePropsMixin := awscdkcfnpropertymixins.Aws_imagebuilder.NewCfnContainerRecipePropsMixin(&CfnContainerRecipeMixinProps{
//   	Components: []interface{}{
//   		&ComponentConfigurationProperty{
//   			ComponentArn: jsii.String("componentArn"),
//   			Parameters: []interface{}{
//   				&ComponentParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ContainerType: jsii.String("containerType"),
//   	Description: jsii.String("description"),
//   	DockerfileTemplateData: jsii.String("dockerfileTemplateData"),
//   	DockerfileTemplateUri: jsii.String("dockerfileTemplateUri"),
//   	ImageOsVersionOverride: jsii.String("imageOsVersionOverride"),
//   	InstanceConfiguration: &InstanceConfigurationProperty{
//   		BlockDeviceMappings: []interface{}{
//   			&InstanceBlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsInstanceBlockDeviceSpecificationProperty{
//   					DeleteOnTermination: jsii.Boolean(false),
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					SnapshotId: jsii.String("snapshotId"),
//   					Throughput: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				NoDevice: jsii.String("noDevice"),
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		Image: jsii.String("image"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	ParentImage: jsii.String("parentImage"),
//   	PlatformOverride: jsii.String("platformOverride"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetRepository: &TargetContainerRepositoryProperty{
//   		RepositoryName: jsii.String("repositoryName"),
//   		Service: jsii.String("service"),
//   	},
//   	Version: jsii.String("version"),
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html
//
type CfnContainerRecipePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnContainerRecipeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContainerRecipePropsMixin
type jsiiProxy_CfnContainerRecipePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnContainerRecipePropsMixin) Props() *CfnContainerRecipeMixinProps {
	var returns *CfnContainerRecipeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::ContainerRecipe`.
func NewCfnContainerRecipePropsMixin(props *CfnContainerRecipeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnContainerRecipePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContainerRecipePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerRecipePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::ContainerRecipe`.
func NewCfnContainerRecipePropsMixin_Override(c CfnContainerRecipePropsMixin, props *CfnContainerRecipeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnContainerRecipePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerRecipePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerRecipePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerRecipePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnContainerRecipePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

