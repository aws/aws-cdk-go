package previewawsfmsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsfms/previewawsfmsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS Firewall Manager policy.
//
// A Firewall Manager policy is specific to the individual policy type. If you want to enforce multiple policy types across accounts, you can create multiple policies. You can create more than one policy for each type.
//
// If you add a new account to an organization that you created with AWS Organizations , Firewall Manager automatically applies the policy to the resources in that account that are within scope of the policy.
//
// Policies require some setup to use. For more information, see the sections on prerequisites and getting started under [Firewall Manager prerequisites](https://docs.aws.amazon.com/waf/latest/developerguide/fms-prereq.html) .
//
// Firewall Manager provides the following types of policies:
//
// - *AWS WAF policy* - This policy applies AWS WAF web ACL protections to specified accounts and resources.
// - *Shield Advanced policy* - This policy applies Shield Advanced protection to specified accounts and resources.
// - *Security Groups policy* - This type of policy gives you control over security groups that are in use throughout your organization in AWS Organizations and lets you enforce a baseline set of rules across your organization.
// - *Network ACL policy* - This type of policy gives you control over the network ACLs that are in use throughout your organization in AWS Organizations and lets you enforce a baseline set of first and last network ACL rules across your organization.
// - *Network Firewall policy* - This policy applies Network Firewall protection to your organization's VPCs.
// - *DNS Firewall policy* - This policy applies Amazon Route 53 Resolver DNS Firewall protections to your organization's VPCs.
// - *Third-party firewall policy* - This policy applies third-party firewall protections. Third-party firewalls are available by subscription through the AWS Marketplace console at [AWS Marketplace](https://docs.aws.amazon.com/marketplace) .
//
// - *Palo Alto Networks Cloud NGFW policy* - This policy applies Palo Alto Networks Cloud Next Generation Firewall (NGFW) protections and Palo Alto Networks Cloud NGFW rulestacks to your organization's VPCs.
// - *Fortigate CNF policy* - This policy applies Fortigate Cloud Native Firewall (CNF) protections. Fortigate CNF is a cloud-centered solution that blocks Zero-Day threats and secures cloud infrastructures with industry-leading advanced threat prevention, smart web application firewalls (WAF), and API protection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnPolicyPropsMixin(&CfnPolicyMixinProps{
//   	DeleteAllPolicyResources: jsii.Boolean(false),
//   	ExcludeMap: map[string][]*string{
//   		"account": []*string{
//   			jsii.String("account"),
//   		},
//   		"orgunit": []*string{
//   			jsii.String("orgunit"),
//   		},
//   	},
//   	ExcludeResourceTags: jsii.Boolean(false),
//   	IncludeMap: map[string][]*string{
//   		"account": []*string{
//   			jsii.String("account"),
//   		},
//   		"orgunit": []*string{
//   			jsii.String("orgunit"),
//   		},
//   	},
//   	PolicyDescription: jsii.String("policyDescription"),
//   	PolicyName: jsii.String("policyName"),
//   	RemediationEnabled: jsii.Boolean(false),
//   	ResourcesCleanUp: jsii.Boolean(false),
//   	ResourceSetIds: []*string{
//   		jsii.String("resourceSetIds"),
//   	},
//   	ResourceTagLogicalOperator: jsii.String("resourceTagLogicalOperator"),
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	ResourceTypeList: []*string{
//   		jsii.String("resourceTypeList"),
//   	},
//   	SecurityServicePolicyData: &SecurityServicePolicyDataProperty{
//   		ManagedServiceData: jsii.String("managedServiceData"),
//   		PolicyOption: &PolicyOptionProperty{
//   			NetworkAclCommonPolicy: &NetworkAclCommonPolicyProperty{
//   				NetworkAclEntrySet: &NetworkAclEntrySetProperty{
//   					FirstEntries: []interface{}{
//   						&NetworkAclEntryProperty{
//   							CidrBlock: jsii.String("cidrBlock"),
//   							Egress: jsii.Boolean(false),
//   							IcmpTypeCode: &IcmpTypeCodeProperty{
//   								Code: jsii.Number(123),
//   								Type: jsii.Number(123),
//   							},
//   							Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   							PortRange: &PortRangeProperty{
//   								From: jsii.Number(123),
//   								To: jsii.Number(123),
//   							},
//   							Protocol: jsii.String("protocol"),
//   							RuleAction: jsii.String("ruleAction"),
//   						},
//   					},
//   					ForceRemediateForFirstEntries: jsii.Boolean(false),
//   					ForceRemediateForLastEntries: jsii.Boolean(false),
//   					LastEntries: []interface{}{
//   						&NetworkAclEntryProperty{
//   							CidrBlock: jsii.String("cidrBlock"),
//   							Egress: jsii.Boolean(false),
//   							IcmpTypeCode: &IcmpTypeCodeProperty{
//   								Code: jsii.Number(123),
//   								Type: jsii.Number(123),
//   							},
//   							Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   							PortRange: &PortRangeProperty{
//   								From: jsii.Number(123),
//   								To: jsii.Number(123),
//   							},
//   							Protocol: jsii.String("protocol"),
//   							RuleAction: jsii.String("ruleAction"),
//   						},
//   					},
//   				},
//   			},
//   			NetworkFirewallPolicy: &NetworkFirewallPolicyProperty{
//   				FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   			},
//   			ThirdPartyFirewallPolicy: &ThirdPartyFirewallPolicyProperty{
//   				FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Tags: []PolicyTagProperty{
//   		&PolicyTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html
//
type CfnPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyPropsMixin
type jsiiProxy_CfnPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPolicyPropsMixin) Props() *CfnPolicyMixinProps {
	var returns *CfnPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FMS::Policy`.
func NewCfnPolicyPropsMixin(props *CfnPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FMS::Policy`.
func NewCfnPolicyPropsMixin_Override(c CfnPolicyPropsMixin, props *CfnPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

