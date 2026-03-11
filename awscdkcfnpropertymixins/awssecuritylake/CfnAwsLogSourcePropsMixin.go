package awssecuritylake

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds a natively supported AWS service as an AWS source.
//
// Enables source types for member accounts in required AWS Regions, based on the parameters you specify. You can choose any source type in any Region for either accounts that are part of a trusted organization or standalone accounts. Once you add an AWS service as a source, Security Lake starts collecting logs and events from it.
//
// > If you want to create multiple sources using `AWS::SecurityLake::AwsLogSource` , you must use the `DependsOn` attribute to create the sources sequentially. With the `DependsOn` attribute you can specify that the creation of a specific `AWSLogSource` follows another. When you add a `DependsOn` attribute to a resource, that resource is created only after the creation of the resource specified in the `DependsOn` attribute. For an example, see [Add AWS log sources](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#aws-resource-securitylake-awslogsource--examples) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAwsLogSourcePropsMixin := awscdkcfnpropertymixins.Aws_securitylake.NewCfnAwsLogSourcePropsMixin(&CfnAwsLogSourceMixinProps{
//   	Accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	DataLakeArn: jsii.String("dataLakeArn"),
//   	SourceName: jsii.String("sourceName"),
//   	SourceVersion: jsii.String("sourceVersion"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html
//
type CfnAwsLogSourcePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAwsLogSourceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAwsLogSourcePropsMixin
type jsiiProxy_CfnAwsLogSourcePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAwsLogSourcePropsMixin) Props() *CfnAwsLogSourceMixinProps {
	var returns *CfnAwsLogSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAwsLogSourcePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityLake::AwsLogSource`.
func NewCfnAwsLogSourcePropsMixin(props *CfnAwsLogSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAwsLogSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAwsLogSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAwsLogSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnAwsLogSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityLake::AwsLogSource`.
func NewCfnAwsLogSourcePropsMixin_Override(c CfnAwsLogSourcePropsMixin, props *CfnAwsLogSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnAwsLogSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAwsLogSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAwsLogSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnAwsLogSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAwsLogSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnAwsLogSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAwsLogSourcePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAwsLogSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

