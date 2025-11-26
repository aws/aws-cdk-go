package previewawsredshiftserverlessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsredshiftserverless/previewawsredshiftserverlessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The collection of compute resources in Amazon Redshift Serverless.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkgroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkgroupPropsMixin(&CfnWorkgroupMixinProps{
//   	BaseCapacity: jsii.Number(123),
//   	ConfigParameters: []interface{}{
//   		&ConfigParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	EnhancedVpcRouting: jsii.Boolean(false),
//   	MaxCapacity: jsii.Number(123),
//   	NamespaceName: jsii.String("namespaceName"),
//   	Port: jsii.Number(123),
//   	PricePerformanceTarget: &PerformanceTargetProperty{
//   		Level: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	PubliclyAccessible: jsii.Boolean(false),
//   	RecoveryPointId: jsii.String("recoveryPointId"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnapshotArn: jsii.String("snapshotArn"),
//   	SnapshotName: jsii.String("snapshotName"),
//   	SnapshotOwnerAccount: jsii.String("snapshotOwnerAccount"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackName: jsii.String("trackName"),
//   	Workgroup: &WorkgroupProperty{
//   		BaseCapacity: jsii.Number(123),
//   		ConfigParameters: []interface{}{
//   			&ConfigParameterProperty{
//   				ParameterKey: jsii.String("parameterKey"),
//   				ParameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   		CreationDate: jsii.String("creationDate"),
//   		Endpoint: &EndpointProperty{
//   			Address: jsii.String("address"),
//   			Port: jsii.Number(123),
//   			VpcEndpoints: []interface{}{
//   				&VpcEndpointProperty{
//   					NetworkInterfaces: []interface{}{
//   						&NetworkInterfaceProperty{
//   							AvailabilityZone: jsii.String("availabilityZone"),
//   							NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   							PrivateIpAddress: jsii.String("privateIpAddress"),
//   							SubnetId: jsii.String("subnetId"),
//   						},
//   					},
//   					VpcEndpointId: jsii.String("vpcEndpointId"),
//   					VpcId: jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   		EnhancedVpcRouting: jsii.Boolean(false),
//   		MaxCapacity: jsii.Number(123),
//   		NamespaceName: jsii.String("namespaceName"),
//   		PricePerformanceTarget: &PerformanceTargetProperty{
//   			Level: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		PubliclyAccessible: jsii.Boolean(false),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Status: jsii.String("status"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		TrackName: jsii.String("trackName"),
//   		WorkgroupArn: jsii.String("workgroupArn"),
//   		WorkgroupId: jsii.String("workgroupId"),
//   		WorkgroupName: jsii.String("workgroupName"),
//   	},
//   	WorkgroupName: jsii.String("workgroupName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html
//
type CfnWorkgroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkgroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkgroupPropsMixin
type jsiiProxy_CfnWorkgroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkgroupPropsMixin) Props() *CfnWorkgroupMixinProps {
	var returns *CfnWorkgroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RedshiftServerless::Workgroup`.
func NewCfnWorkgroupPropsMixin(props *CfnWorkgroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkgroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkgroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkgroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RedshiftServerless::Workgroup`.
func NewCfnWorkgroupPropsMixin_Override(c CfnWorkgroupPropsMixin, props *CfnWorkgroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkgroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkgroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkgroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkgroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkgroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

