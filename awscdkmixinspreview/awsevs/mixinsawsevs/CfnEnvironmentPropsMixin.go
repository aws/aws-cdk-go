package mixinsawsevs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsevs/mixinsawsevs/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon EVS environment that runs VCF software, such as SDDC Manager, NSX Manager, and vCenter Server.
//
// During environment creation, Amazon EVS performs validations on DNS settings, provisions VLAN subnets and hosts, and deploys the supplied version of VCF.
//
// It can take several hours to create an environment. After the deployment completes, you can configure VCF in the vSphere user interface according to your needs.
//
// > You cannot use the `dedicatedHostId` and `placementGroupId` parameters together in the same `CreateEnvironment` action. This results in a `ValidationException` response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentPropsMixin(&CfnEnvironmentMixinProps{
//   	ConnectivityInfo: &ConnectivityInfoProperty{
//   		PrivateRouteServerPeerings: []*string{
//   			jsii.String("privateRouteServerPeerings"),
//   		},
//   	},
//   	EnvironmentName: jsii.String("environmentName"),
//   	Hosts: []interface{}{
//   		&HostInfoForCreateProperty{
//   			DedicatedHostId: jsii.String("dedicatedHostId"),
//   			HostName: jsii.String("hostName"),
//   			InstanceType: jsii.String("instanceType"),
//   			KeyName: jsii.String("keyName"),
//   			PlacementGroupId: jsii.String("placementGroupId"),
//   		},
//   	},
//   	InitialVlans: &InitialVlansProperty{
//   		EdgeVTep: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		ExpansionVlan1: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		ExpansionVlan2: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		Hcx: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		HcxNetworkAclId: jsii.String("hcxNetworkAclId"),
//   		IsHcxPublic: jsii.Boolean(false),
//   		NsxUpLink: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VmkManagement: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VmManagement: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VMotion: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VSan: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VTep: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LicenseInfo: &LicenseInfoProperty{
//   		SolutionKey: jsii.String("solutionKey"),
//   		VsanKey: jsii.String("vsanKey"),
//   	},
//   	ServiceAccessSecurityGroups: &ServiceAccessSecurityGroupsProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   	ServiceAccessSubnetId: jsii.String("serviceAccessSubnetId"),
//   	SiteId: jsii.String("siteId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TermsAccepted: jsii.Boolean(false),
//   	VcfHostnames: &VcfHostnamesProperty{
//   		CloudBuilder: jsii.String("cloudBuilder"),
//   		Nsx: jsii.String("nsx"),
//   		NsxEdge1: jsii.String("nsxEdge1"),
//   		NsxEdge2: jsii.String("nsxEdge2"),
//   		NsxManager1: jsii.String("nsxManager1"),
//   		NsxManager2: jsii.String("nsxManager2"),
//   		NsxManager3: jsii.String("nsxManager3"),
//   		SddcManager: jsii.String("sddcManager"),
//   		VCenter: jsii.String("vCenter"),
//   	},
//   	VcfVersion: jsii.String("vcfVersion"),
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html
//
type CfnEnvironmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentPropsMixin
type jsiiProxy_CfnEnvironmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Props() *CfnEnvironmentMixinProps {
	var returns *CfnEnvironmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EVS::Environment`.
func NewCfnEnvironmentPropsMixin(props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EVS::Environment`.
func NewCfnEnvironmentPropsMixin_Override(c CfnEnvironmentPropsMixin, props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEnvironmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

