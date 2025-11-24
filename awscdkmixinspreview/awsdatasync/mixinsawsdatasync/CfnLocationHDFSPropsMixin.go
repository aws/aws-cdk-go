package mixinsawsdatasync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsdatasync/mixinsawsdatasync/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataSync::LocationHDFS` resource specifies an endpoint for a Hadoop Distributed File System (HDFS).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationHDFSPropsMixin := awscdkmixinspreview.Mixins.NewCfnLocationHDFSPropsMixin(&CfnLocationHDFSMixinProps{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	AuthenticationType: jsii.String("authenticationType"),
//   	BlockSize: jsii.Number(123),
//   	KerberosKeytab: jsii.String("kerberosKeytab"),
//   	KerberosKrb5Conf: jsii.String("kerberosKrb5Conf"),
//   	KerberosPrincipal: jsii.String("kerberosPrincipal"),
//   	KmsKeyProviderUri: jsii.String("kmsKeyProviderUri"),
//   	NameNodes: []interface{}{
//   		&NameNodeProperty{
//   			Hostname: jsii.String("hostname"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   	QopConfiguration: &QopConfigurationProperty{
//   		DataTransferProtection: jsii.String("dataTransferProtection"),
//   		RpcProtection: jsii.String("rpcProtection"),
//   	},
//   	ReplicationFactor: jsii.Number(123),
//   	SimpleUser: jsii.String("simpleUser"),
//   	Subdirectory: jsii.String("subdirectory"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html
//
type CfnLocationHDFSPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLocationHDFSMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLocationHDFSPropsMixin
type jsiiProxy_CfnLocationHDFSPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLocationHDFSPropsMixin) Props() *CfnLocationHDFSMixinProps {
	var returns *CfnLocationHDFSMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFSPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::LocationHDFS`.
func NewCfnLocationHDFSPropsMixin(props *CfnLocationHDFSMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLocationHDFSPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLocationHDFSPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLocationHDFSPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::LocationHDFS`.
func NewCfnLocationHDFSPropsMixin_Override(c CfnLocationHDFSPropsMixin, props *CfnLocationHDFSMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLocationHDFSPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLocationHDFSPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationHDFSPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLocationHDFSPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLocationHDFSPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

