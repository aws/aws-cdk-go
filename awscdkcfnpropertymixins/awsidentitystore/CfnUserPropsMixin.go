package awsidentitystore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsidentitystore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a user within the specified identity store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUserPropsMixin := awscdkcfnpropertymixins.Aws_identitystore.NewCfnUserPropsMixin(&CfnUserMixinProps{
//   	Addresses: []interface{}{
//   		&AddressesItemsProperty{
//   			Country: jsii.String("country"),
//   			Formatted: jsii.String("formatted"),
//   			Locality: jsii.String("locality"),
//   			PostalCode: jsii.String("postalCode"),
//   			Primary: jsii.Boolean(false),
//   			Region: jsii.String("region"),
//   			StreetAddress: jsii.String("streetAddress"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Birthdate: jsii.String("birthdate"),
//   	DisplayName: jsii.String("displayName"),
//   	Emails: []interface{}{
//   		&EmailsItemsProperty{
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   	Locale: jsii.String("locale"),
//   	Name: &NameProperty{
//   		FamilyName: jsii.String("familyName"),
//   		Formatted: jsii.String("formatted"),
//   		GivenName: jsii.String("givenName"),
//   		HonorificPrefix: jsii.String("honorificPrefix"),
//   		HonorificSuffix: jsii.String("honorificSuffix"),
//   		MiddleName: jsii.String("middleName"),
//   	},
//   	NickName: jsii.String("nickName"),
//   	PhoneNumbers: []interface{}{
//   		&PhoneNumbersItemsProperty{
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Photos: []interface{}{
//   		&PhotosItemsProperty{
//   			Display: jsii.String("display"),
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	PreferredLanguage: jsii.String("preferredLanguage"),
//   	ProfileUrl: jsii.String("profileUrl"),
//   	Roles: []interface{}{
//   		&RolesItemsProperty{
//   			Primary: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timezone: jsii.String("timezone"),
//   	Title: jsii.String("title"),
//   	UserName: jsii.String("userName"),
//   	UserType: jsii.String("userType"),
//   	Website: jsii.String("website"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-user.html
//
type CfnUserPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPropsMixin
type jsiiProxy_CfnUserPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUserPropsMixin) Props() *CfnUserMixinProps {
	var returns *CfnUserMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IdentityStore::User`.
func NewCfnUserPropsMixin(props *CfnUserMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnUserPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IdentityStore::User`.
func NewCfnUserPropsMixin_Override(c CfnUserPropsMixin, props *CfnUserMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnUserPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnUserPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnUserPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

