package previewawsgreengrassmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgreengrass/previewawsgreengrassmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Greengrass::ResourceDefinitionVersion` resource represents a resource definition version for AWS IoT Greengrass .
//
// A resource definition version contains a list of resources. (In CloudFormation , resources are named *resource instances* .)
//
// > To create a resource definition version, you must specify the ID of the resource definition that you want to associate with the version. For information about creating a resource definition, see [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) .
// >
// > After you create a resource definition version that contains the resources you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceDefinitionVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnResourceDefinitionVersionPropsMixin(&CfnResourceDefinitionVersionMixinProps{
//   	ResourceDefinitionId: jsii.String("resourceDefinitionId"),
//   	Resources: []interface{}{
//   		&ResourceInstanceProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			ResourceDataContainer: &ResourceDataContainerProperty{
//   				LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   					GroupOwnerSetting: &GroupOwnerSettingProperty{
//   						AutoAddGroupOwner: jsii.Boolean(false),
//   						GroupOwner: jsii.String("groupOwner"),
//   					},
//   					SourcePath: jsii.String("sourcePath"),
//   				},
//   				LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					GroupOwnerSetting: &GroupOwnerSettingProperty{
//   						AutoAddGroupOwner: jsii.Boolean(false),
//   						GroupOwner: jsii.String("groupOwner"),
//   					},
//   					SourcePath: jsii.String("sourcePath"),
//   				},
//   				S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   						GroupOwner: jsii.String("groupOwner"),
//   						GroupPermission: jsii.String("groupPermission"),
//   					},
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   				SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   						GroupOwner: jsii.String("groupOwner"),
//   						GroupPermission: jsii.String("groupPermission"),
//   					},
//   					SageMakerJobArn: jsii.String("sageMakerJobArn"),
//   				},
//   				SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   					AdditionalStagingLabelsToDownload: []*string{
//   						jsii.String("additionalStagingLabelsToDownload"),
//   					},
//   					Arn: jsii.String("arn"),
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html
//
type CfnResourceDefinitionVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResourceDefinitionVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourceDefinitionVersionPropsMixin
type jsiiProxy_CfnResourceDefinitionVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResourceDefinitionVersionPropsMixin) Props() *CfnResourceDefinitionVersionMixinProps {
	var returns *CfnResourceDefinitionVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefinitionVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Greengrass::ResourceDefinitionVersion`.
func NewCfnResourceDefinitionVersionPropsMixin(props *CfnResourceDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResourceDefinitionVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourceDefinitionVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourceDefinitionVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnResourceDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Greengrass::ResourceDefinitionVersion`.
func NewCfnResourceDefinitionVersionPropsMixin_Override(c CfnResourceDefinitionVersionPropsMixin, props *CfnResourceDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnResourceDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResourceDefinitionVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourceDefinitionVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnResourceDefinitionVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceDefinitionVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnResourceDefinitionVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceDefinitionVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnResourceDefinitionVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

