package awsmgn

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmgn/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::MGN::NetworkMigrationDefinition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNetworkMigrationDefinitionPropsMixin := awscdkcfnpropertymixins.Aws_mgn.NewCfnNetworkMigrationDefinitionPropsMixin(&CfnNetworkMigrationDefinitionMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ScopeTags: map[string]*string{
//   		"scopeTagsKey": jsii.String("scopeTags"),
//   	},
//   	SourceConfigurations: []interface{}{
//   		&SourceConfigurationProperty{
//   			SourceEnvironment: jsii.String("sourceEnvironment"),
//   			SourceS3Configuration: &SourceS3ConfigurationProperty{
//   				S3Bucket: jsii.String("s3Bucket"),
//   				S3BucketOwner: jsii.String("s3BucketOwner"),
//   				S3Key: jsii.String("s3Key"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetDeployment: jsii.String("targetDeployment"),
//   	TargetNetwork: &TargetNetworkProperty{
//   		InboundCidr: jsii.String("inboundCidr"),
//   		InspectionCidr: jsii.String("inspectionCidr"),
//   		OutboundCidr: jsii.String("outboundCidr"),
//   		Topology: jsii.String("topology"),
//   	},
//   	TargetS3Configuration: &TargetS3ConfigurationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3BucketOwner: jsii.String("s3BucketOwner"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html
//
type CfnNetworkMigrationDefinitionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNetworkMigrationDefinitionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkMigrationDefinitionPropsMixin
type jsiiProxy_CfnNetworkMigrationDefinitionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNetworkMigrationDefinitionPropsMixin) Props() *CfnNetworkMigrationDefinitionMixinProps {
	var returns *CfnNetworkMigrationDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkMigrationDefinitionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MGN::NetworkMigrationDefinition`.
func NewCfnNetworkMigrationDefinitionPropsMixin(props *CfnNetworkMigrationDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNetworkMigrationDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkMigrationDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkMigrationDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MGN::NetworkMigrationDefinition`.
func NewCfnNetworkMigrationDefinitionPropsMixin_Override(c CfnNetworkMigrationDefinitionPropsMixin, props *CfnNetworkMigrationDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNetworkMigrationDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkMigrationDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkMigrationDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkMigrationDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNetworkMigrationDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

