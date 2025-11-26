package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Connection` resource specifies an AWS Glue connection to a data source.
//
// For more information, see [Adding a Connection to Your Data Store](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html) and [Connection Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-connections.html#aws-glue-api-catalog-connections-Connection) in the *AWS Glue Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var athenaProperties interface{}
//   var connectionProperties interface{}
//   var customAuthenticationCredentials interface{}
//   var pythonProperties interface{}
//   var sparkProperties interface{}
//   var tokenUrlParametersMap interface{}
//
//   cfnConnectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnConnectionPropsMixin(&CfnConnectionMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	ConnectionInput: &ConnectionInputProperty{
//   		AthenaProperties: athenaProperties,
//   		AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//   			BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			CustomAuthenticationCredentials: customAuthenticationCredentials,
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			OAuth2Properties: &OAuth2PropertiesInputProperty{
//   				AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   					AuthorizationCode: jsii.String("authorizationCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   					AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   					UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   				},
//   				OAuth2Credentials: &OAuth2CredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					JwtToken: jsii.String("jwtToken"),
//   					RefreshToken: jsii.String("refreshToken"),
//   					UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   				},
//   				OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   				TokenUrl: jsii.String("tokenUrl"),
//   				TokenUrlParametersMap: tokenUrlParametersMap,
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		ConnectionProperties: connectionProperties,
//   		ConnectionType: jsii.String("connectionType"),
//   		Description: jsii.String("description"),
//   		MatchCriteria: []*string{
//   			jsii.String("matchCriteria"),
//   		},
//   		Name: jsii.String("name"),
//   		PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			SecurityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   		PythonProperties: pythonProperties,
//   		SparkProperties: sparkProperties,
//   		ValidateCredentials: jsii.Boolean(false),
//   		ValidateForComputeEnvironments: []*string{
//   			jsii.String("validateForComputeEnvironments"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html
//
type CfnConnectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConnectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectionPropsMixin
type jsiiProxy_CfnConnectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConnectionPropsMixin) Props() *CfnConnectionMixinProps {
	var returns *CfnConnectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Connection`.
func NewCfnConnectionPropsMixin(props *CfnConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConnectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Connection`.
func NewCfnConnectionPropsMixin_Override(c CfnConnectionPropsMixin, props *CfnConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConnectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnConnectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConnectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

