package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes a Network Access Scope.
//
// A Network Access Scope defines outbound (egress) and inbound (ingress) traffic patterns, including sources, destinations, paths, and traffic types.
//
// Network Access Analyzer identifies unintended network access to your resources on AWS . When you start an analysis on a Network Access Scope, Network Access Analyzer produces findings. For more information, see the [Network Access Analyzer User Guide](https://docs.aws.amazon.com/vpc/latest/network-access-analyzer/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkInsightsAccessScopePropsMixin := awscdkmixinspreview.Mixins.NewCfnNetworkInsightsAccessScopePropsMixin(&CfnNetworkInsightsAccessScopeMixinProps{
//   	ExcludePaths: []interface{}{
//   		&AccessScopePathRequestProperty{
//   			Destination: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			Source: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			ThroughResources: []interface{}{
//   				&ThroughResourcesStatementRequestProperty{
//   					ResourceStatement: &ResourceStatementRequestProperty{
//   						Resources: []*string{
//   							jsii.String("resources"),
//   						},
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	MatchPaths: []interface{}{
//   		&AccessScopePathRequestProperty{
//   			Destination: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			Source: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			ThroughResources: []interface{}{
//   				&ThroughResourcesStatementRequestProperty{
//   					ResourceStatement: &ResourceStatementRequestProperty{
//   						Resources: []*string{
//   							jsii.String("resources"),
//   						},
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscope.html
//
type CfnNetworkInsightsAccessScopePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNetworkInsightsAccessScopeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkInsightsAccessScopePropsMixin
type jsiiProxy_CfnNetworkInsightsAccessScopePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNetworkInsightsAccessScopePropsMixin) Props() *CfnNetworkInsightsAccessScopeMixinProps {
	var returns *CfnNetworkInsightsAccessScopeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkInsightsAccessScopePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NetworkInsightsAccessScope`.
func NewCfnNetworkInsightsAccessScopePropsMixin(props *CfnNetworkInsightsAccessScopeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNetworkInsightsAccessScopePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkInsightsAccessScopePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkInsightsAccessScopePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NetworkInsightsAccessScope`.
func NewCfnNetworkInsightsAccessScopePropsMixin_Override(c CfnNetworkInsightsAccessScopePropsMixin, props *CfnNetworkInsightsAccessScopeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNetworkInsightsAccessScopePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkInsightsAccessScopePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkInsightsAccessScopePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkInsightsAccessScopePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNetworkInsightsAccessScopePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

