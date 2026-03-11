package awsdatasync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a transfer *location* for a Microsoft Azure Blob Storage container.
//
// AWS DataSync can use this location as a transfer source or destination. You can make transfers with or without a [DataSync agent](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-creating-agent) that connects to your container.
//
// Before you begin, make sure you know [how DataSync accesses Azure Blob Storage](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access) and works with [access tiers](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers) and [blob types](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#blob-types) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLocationAzureBlobPropsMixin := awscdkcfnpropertymixins.Aws_datasync.NewCfnLocationAzureBlobPropsMixin(&CfnLocationAzureBlobMixinProps{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	AzureAccessTier: jsii.String("azureAccessTier"),
//   	AzureBlobAuthenticationType: jsii.String("azureBlobAuthenticationType"),
//   	AzureBlobContainerUrl: jsii.String("azureBlobContainerUrl"),
//   	AzureBlobSasConfiguration: &AzureBlobSasConfigurationProperty{
//   		AzureBlobSasToken: jsii.String("azureBlobSasToken"),
//   	},
//   	AzureBlobType: jsii.String("azureBlobType"),
//   	CmkSecretConfig: &CmkSecretConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	CustomSecretConfig: &CustomSecretConfigProperty{
//   		SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	Subdirectory: jsii.String("subdirectory"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html
//
type CfnLocationAzureBlobPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLocationAzureBlobMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLocationAzureBlobPropsMixin
type jsiiProxy_CfnLocationAzureBlobPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLocationAzureBlobPropsMixin) Props() *CfnLocationAzureBlobMixinProps {
	var returns *CfnLocationAzureBlobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationAzureBlobPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::LocationAzureBlob`.
func NewCfnLocationAzureBlobPropsMixin(props *CfnLocationAzureBlobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLocationAzureBlobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLocationAzureBlobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLocationAzureBlobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::LocationAzureBlob`.
func NewCfnLocationAzureBlobPropsMixin_Override(c CfnLocationAzureBlobPropsMixin, props *CfnLocationAzureBlobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLocationAzureBlobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLocationAzureBlobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationAzureBlobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLocationAzureBlobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLocationAzureBlobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

