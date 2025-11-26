package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a security profile.
//
// For information about security profiles, see [Security Profiles](https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html) in the *Amazon Connect Administrator Guide* . For a mapping of the API name and user interface name of the security profile permissions, see [List of security profile permissions](https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnSecurityProfilePropsMixin(&CfnSecurityProfileMixinProps{
//   	AllowedAccessControlHierarchyGroupId: jsii.String("allowedAccessControlHierarchyGroupId"),
//   	AllowedAccessControlTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Applications: []interface{}{
//   		&ApplicationProperty{
//   			ApplicationPermissions: []*string{
//   				jsii.String("applicationPermissions"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GranularAccessControlConfiguration: &GranularAccessControlConfigurationProperty{
//   		DataTableAccessControlConfiguration: &DataTableAccessControlConfigurationProperty{
//   			PrimaryAttributeAccessControlConfiguration: &PrimaryAttributeAccessControlConfigurationItemProperty{
//   				PrimaryAttributeValues: []interface{}{
//   					&PrimaryAttributeValueProperty{
//   						AccessType: jsii.String("accessType"),
//   						AttributeName: jsii.String("attributeName"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	HierarchyRestrictedResources: []*string{
//   		jsii.String("hierarchyRestrictedResources"),
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	SecurityProfileName: jsii.String("securityProfileName"),
//   	TagRestrictedResources: []*string{
//   		jsii.String("tagRestrictedResources"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html
//
type CfnSecurityProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSecurityProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityProfilePropsMixin
type jsiiProxy_CfnSecurityProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSecurityProfilePropsMixin) Props() *CfnSecurityProfileMixinProps {
	var returns *CfnSecurityProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::SecurityProfile`.
func NewCfnSecurityProfilePropsMixin(props *CfnSecurityProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSecurityProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnSecurityProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::SecurityProfile`.
func NewCfnSecurityProfilePropsMixin_Override(c CfnSecurityProfilePropsMixin, props *CfnSecurityProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnSecurityProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSecurityProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnSecurityProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnSecurityProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSecurityProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

