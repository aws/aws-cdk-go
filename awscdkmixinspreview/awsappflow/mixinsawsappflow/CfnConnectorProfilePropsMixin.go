package mixinsawsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsappflow/mixinsawsappflow/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppFlow::ConnectorProfile` resource is an Amazon AppFlow resource type that specifies the configuration profile for an instance of a connector.
//
// This includes the provided name, credentials ARN, connection-mode, and so on. The fields that are common to all types of connector profiles are explicitly specified under the `Properties` field. The rest of the connector-specific properties are specified under `Properties/ConnectorProfileConfig` .
//
// > If you want to use CloudFormation to create a connector profile for connectors that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth, or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow console to create the connector profile, and then use that connector profile in the flow creation CloudFormation template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnConnectorProfilePropsMixin(&CfnConnectorProfileMixinProps{
//   	ConnectionMode: jsii.String("connectionMode"),
//   	ConnectorLabel: jsii.String("connectorLabel"),
//   	ConnectorProfileConfig: &ConnectorProfileConfigProperty{
//   		ConnectorProfileCredentials: &ConnectorProfileCredentialsProperty{
//   			Amplitude: &AmplitudeConnectorProfileCredentialsProperty{
//   				ApiKey: jsii.String("apiKey"),
//   				SecretKey: jsii.String("secretKey"),
//   			},
//   			CustomConnector: &CustomConnectorProfileCredentialsProperty{
//   				ApiKey: &ApiKeyCredentialsProperty{
//   					ApiKey: jsii.String("apiKey"),
//   					ApiSecretKey: jsii.String("apiSecretKey"),
//   				},
//   				AuthenticationType: jsii.String("authenticationType"),
//   				Basic: &BasicAuthCredentialsProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   				Custom: &CustomAuthCredentialsProperty{
//   					CredentialsMap: map[string]*string{
//   						"credentialsMapKey": jsii.String("credentialsMap"),
//   					},
//   					CustomAuthenticationType: jsii.String("customAuthenticationType"),
//   				},
//   				Oauth2: &OAuth2CredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					OAuthRequest: &ConnectorOAuthRequestProperty{
//   						AuthCode: jsii.String("authCode"),
//   						RedirectUri: jsii.String("redirectUri"),
//   					},
//   					RefreshToken: jsii.String("refreshToken"),
//   				},
//   			},
//   			Datadog: &DatadogConnectorProfileCredentialsProperty{
//   				ApiKey: jsii.String("apiKey"),
//   				ApplicationKey: jsii.String("applicationKey"),
//   			},
//   			Dynatrace: &DynatraceConnectorProfileCredentialsProperty{
//   				ApiToken: jsii.String("apiToken"),
//   			},
//   			GoogleAnalytics: &GoogleAnalyticsConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				RefreshToken: jsii.String("refreshToken"),
//   			},
//   			InforNexus: &InforNexusConnectorProfileCredentialsProperty{
//   				AccessKeyId: jsii.String("accessKeyId"),
//   				Datakey: jsii.String("datakey"),
//   				SecretAccessKey: jsii.String("secretAccessKey"),
//   				UserId: jsii.String("userId"),
//   			},
//   			Marketo: &MarketoConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			Pardot: &PardotConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				RefreshToken: jsii.String("refreshToken"),
//   			},
//   			Redshift: &RedshiftConnectorProfileCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Salesforce: &SalesforceConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				JwtToken: jsii.String("jwtToken"),
//   				OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   				RefreshToken: jsii.String("refreshToken"),
//   			},
//   			SapoData: &SAPODataConnectorProfileCredentialsProperty{
//   				BasicAuthCredentials: &BasicAuthCredentialsProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   				OAuthCredentials: &OAuthCredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   						AuthCode: jsii.String("authCode"),
//   						RedirectUri: jsii.String("redirectUri"),
//   					},
//   					RefreshToken: jsii.String("refreshToken"),
//   				},
//   			},
//   			ServiceNow: &ServiceNowConnectorProfileCredentialsProperty{
//   				OAuth2Credentials: &OAuth2CredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					OAuthRequest: &ConnectorOAuthRequestProperty{
//   						AuthCode: jsii.String("authCode"),
//   						RedirectUri: jsii.String("redirectUri"),
//   					},
//   					RefreshToken: jsii.String("refreshToken"),
//   				},
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Singular: &SingularConnectorProfileCredentialsProperty{
//   				ApiKey: jsii.String("apiKey"),
//   			},
//   			Slack: &SlackConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			Snowflake: &SnowflakeConnectorProfileCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Trendmicro: &TrendmicroConnectorProfileCredentialsProperty{
//   				ApiSecretKey: jsii.String("apiSecretKey"),
//   			},
//   			Veeva: &VeevaConnectorProfileCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			Zendesk: &ZendeskConnectorProfileCredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   					AuthCode: jsii.String("authCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   		},
//   		ConnectorProfileProperties: &ConnectorProfilePropertiesProperty{
//   			CustomConnector: &CustomConnectorProfilePropertiesProperty{
//   				OAuth2Properties: &OAuth2PropertiesProperty{
//   					OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   					TokenUrl: jsii.String("tokenUrl"),
//   					TokenUrlCustomProperties: map[string]*string{
//   						"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   					},
//   				},
//   				ProfileProperties: map[string]*string{
//   					"profilePropertiesKey": jsii.String("profileProperties"),
//   				},
//   			},
//   			Datadog: &DatadogConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Dynatrace: &DynatraceConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			InforNexus: &InforNexusConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Marketo: &MarketoConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Pardot: &PardotConnectorProfilePropertiesProperty{
//   				BusinessUnitId: jsii.String("businessUnitId"),
//   				InstanceUrl: jsii.String("instanceUrl"),
//   				IsSandboxEnvironment: jsii.Boolean(false),
//   			},
//   			Redshift: &RedshiftConnectorProfilePropertiesProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				ClusterIdentifier: jsii.String("clusterIdentifier"),
//   				DataApiRoleArn: jsii.String("dataApiRoleArn"),
//   				DatabaseName: jsii.String("databaseName"),
//   				DatabaseUrl: jsii.String("databaseUrl"),
//   				IsRedshiftServerless: jsii.Boolean(false),
//   				RoleArn: jsii.String("roleArn"),
//   				WorkgroupName: jsii.String("workgroupName"),
//   			},
//   			Salesforce: &SalesforceConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   				IsSandboxEnvironment: jsii.Boolean(false),
//   				UsePrivateLinkForMetadataAndAuthorization: jsii.Boolean(false),
//   			},
//   			SapoData: &SAPODataConnectorProfilePropertiesProperty{
//   				ApplicationHostUrl: jsii.String("applicationHostUrl"),
//   				ApplicationServicePath: jsii.String("applicationServicePath"),
//   				ClientNumber: jsii.String("clientNumber"),
//   				DisableSso: jsii.Boolean(false),
//   				LogonLanguage: jsii.String("logonLanguage"),
//   				OAuthProperties: &OAuthPropertiesProperty{
//   					AuthCodeUrl: jsii.String("authCodeUrl"),
//   					OAuthScopes: []*string{
//   						jsii.String("oAuthScopes"),
//   					},
//   					TokenUrl: jsii.String("tokenUrl"),
//   				},
//   				PortNumber: jsii.Number(123),
//   				PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   			},
//   			ServiceNow: &ServiceNowConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Slack: &SlackConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Snowflake: &SnowflakeConnectorProfilePropertiesProperty{
//   				AccountName: jsii.String("accountName"),
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   				Region: jsii.String("region"),
//   				Stage: jsii.String("stage"),
//   				Warehouse: jsii.String("warehouse"),
//   			},
//   			Veeva: &VeevaConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Zendesk: &ZendeskConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   		},
//   	},
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	ConnectorType: jsii.String("connectorType"),
//   	KmsArn: jsii.String("kmsArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html
//
type CfnConnectorProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConnectorProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectorProfilePropsMixin
type jsiiProxy_CfnConnectorProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConnectorProfilePropsMixin) Props() *CfnConnectorProfileMixinProps {
	var returns *CfnConnectorProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppFlow::ConnectorProfile`.
func NewCfnConnectorProfilePropsMixin(props *CfnConnectorProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConnectorProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectorProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectorProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appflow.mixins.CfnConnectorProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppFlow::ConnectorProfile`.
func NewCfnConnectorProfilePropsMixin_Override(c CfnConnectorProfilePropsMixin, props *CfnConnectorProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appflow.mixins.CfnConnectorProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConnectorProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appflow.mixins.CfnConnectorProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appflow.mixins.CfnConnectorProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConnectorProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

