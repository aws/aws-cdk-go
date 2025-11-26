package previewawsimagebuildermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsimagebuilder/previewawsimagebuildermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new infrastructure configuration.
//
// An infrastructure configuration defines the environment in which your image will be built and tested.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInfrastructureConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnInfrastructureConfigurationPropsMixin(&CfnInfrastructureConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceMetadataOptions: &InstanceMetadataOptionsProperty{
//   		HttpPutResponseHopLimit: jsii.Number(123),
//   		HttpTokens: jsii.String("httpTokens"),
//   	},
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   	InstanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	KeyPair: jsii.String("keyPair"),
//   	Logging: &LoggingProperty{
//   		S3Logs: &S3LogsProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Placement: &PlacementProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		HostId: jsii.String("hostId"),
//   		HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   		Tenancy: jsii.String("tenancy"),
//   	},
//   	ResourceTags: map[string]*string{
//   		"resourceTagsKey": jsii.String("resourceTags"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TerminateInstanceOnFailure: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html
//
type CfnInfrastructureConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInfrastructureConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInfrastructureConfigurationPropsMixin
type jsiiProxy_CfnInfrastructureConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInfrastructureConfigurationPropsMixin) Props() *CfnInfrastructureConfigurationMixinProps {
	var returns *CfnInfrastructureConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::InfrastructureConfiguration`.
func NewCfnInfrastructureConfigurationPropsMixin(props *CfnInfrastructureConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInfrastructureConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInfrastructureConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInfrastructureConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::InfrastructureConfiguration`.
func NewCfnInfrastructureConfigurationPropsMixin_Override(c CfnInfrastructureConfigurationPropsMixin, props *CfnInfrastructureConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInfrastructureConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInfrastructureConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInfrastructureConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInfrastructureConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnInfrastructureConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

