package previewawsdatazonemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatazone/previewawsdatazonemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds a policy grant (an authorization policy) to a specified entity, including domain units, environment blueprint configurations, or environment profiles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var allDomainUnitsGrantFilter interface{}
//   var allUsersGrantFilter interface{}
//   var createEnvironment interface{}
//   var createEnvironmentFromBlueprint interface{}
//   var delegateCreateEnvironmentProfile interface{}
//
//   cfnPolicyGrantPropsMixin := awscdkmixinspreview.Mixins.NewCfnPolicyGrantPropsMixin(&CfnPolicyGrantMixinProps{
//   	Detail: &PolicyGrantDetailProperty{
//   		AddToProjectMemberPool: &AddToProjectMemberPoolPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateAssetType: &CreateAssetTypePolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateDomainUnit: &CreateDomainUnitPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateEnvironment: createEnvironment,
//   		CreateEnvironmentFromBlueprint: createEnvironmentFromBlueprint,
//   		CreateEnvironmentProfile: &CreateEnvironmentProfilePolicyGrantDetailProperty{
//   			DomainUnitId: jsii.String("domainUnitId"),
//   		},
//   		CreateFormType: &CreateFormTypePolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateGlossary: &CreateGlossaryPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateProject: &CreateProjectPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateProjectFromProjectProfile: &CreateProjectFromProjectProfilePolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   			ProjectProfiles: []*string{
//   				jsii.String("projectProfiles"),
//   			},
//   		},
//   		DelegateCreateEnvironmentProfile: delegateCreateEnvironmentProfile,
//   		OverrideDomainUnitOwners: &OverrideDomainUnitOwnersPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		OverrideProjectOwners: &OverrideProjectOwnersPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   	},
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EntityIdentifier: jsii.String("entityIdentifier"),
//   	EntityType: jsii.String("entityType"),
//   	PolicyType: jsii.String("policyType"),
//   	Principal: &PolicyGrantPrincipalProperty{
//   		DomainUnit: &DomainUnitPolicyGrantPrincipalProperty{
//   			DomainUnitDesignation: jsii.String("domainUnitDesignation"),
//   			DomainUnitGrantFilter: &DomainUnitGrantFilterProperty{
//   				AllDomainUnitsGrantFilter: allDomainUnitsGrantFilter,
//   			},
//   			DomainUnitIdentifier: jsii.String("domainUnitIdentifier"),
//   		},
//   		Group: &GroupPolicyGrantPrincipalProperty{
//   			GroupIdentifier: jsii.String("groupIdentifier"),
//   		},
//   		Project: &ProjectPolicyGrantPrincipalProperty{
//   			ProjectDesignation: jsii.String("projectDesignation"),
//   			ProjectGrantFilter: &ProjectGrantFilterProperty{
//   				DomainUnitFilter: &DomainUnitFilterForProjectProperty{
//   					DomainUnit: jsii.String("domainUnit"),
//   					IncludeChildDomainUnits: jsii.Boolean(false),
//   				},
//   			},
//   			ProjectIdentifier: jsii.String("projectIdentifier"),
//   		},
//   		User: &UserPolicyGrantPrincipalProperty{
//   			AllUsersGrantFilter: allUsersGrantFilter,
//   			UserIdentifier: jsii.String("userIdentifier"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html
//
type CfnPolicyGrantPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPolicyGrantMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyGrantPropsMixin
type jsiiProxy_CfnPolicyGrantPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPolicyGrantPropsMixin) Props() *CfnPolicyGrantMixinProps {
	var returns *CfnPolicyGrantMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyGrantPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataZone::PolicyGrant`.
func NewCfnPolicyGrantPropsMixin(props *CfnPolicyGrantMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPolicyGrantPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyGrantPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyGrantPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataZone::PolicyGrant`.
func NewCfnPolicyGrantPropsMixin_Override(c CfnPolicyGrantPropsMixin, props *CfnPolicyGrantMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPolicyGrantPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyGrantPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyGrantPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnPolicyGrantPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyGrantPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPolicyGrantPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

