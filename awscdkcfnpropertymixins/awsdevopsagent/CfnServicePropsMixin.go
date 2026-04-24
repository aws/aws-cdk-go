package awsdevopsagent

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdevopsagent/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::DevOpsAgent::Service resource registers external services (like Dynatrace, MCP servers, GitLab) for integration with DevOpsAgent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnServicePropsMixin := awscdkcfnpropertymixins.Aws_devopsagent.NewCfnServicePropsMixin(&CfnServiceMixinProps{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	ServiceDetails: &ServiceDetailsProperty{
//   		AzureIdentity: &AzureIdentityServiceDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			TenantId: jsii.String("tenantId"),
//   			WebIdentityRoleArn: jsii.String("webIdentityRoleArn"),
//   			WebIdentityTokenAudiences: []*string{
//   				jsii.String("webIdentityTokenAudiences"),
//   			},
//   		},
//   		Dynatrace: &DynatraceServiceDetailsProperty{
//   			AccountUrn: jsii.String("accountUrn"),
//   			AuthorizationConfig: &DynatraceAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   		},
//   		GitLab: &GitLabDetailsProperty{
//   			GroupId: jsii.String("groupId"),
//   			TargetUrl: jsii.String("targetUrl"),
//   			TokenType: jsii.String("tokenType"),
//   			TokenValue: jsii.String("tokenValue"),
//   		},
//   		McpServer: &MCPServerDetailsProperty{
//   			AuthorizationConfig: &MCPServerAuthorizationConfigProperty{
//   				ApiKey: &ApiKeyDetailsProperty{
//   					ApiKeyHeader: jsii.String("apiKeyHeader"),
//   					ApiKeyName: jsii.String("apiKeyName"),
//   					ApiKeyValue: jsii.String("apiKeyValue"),
//   				},
//   				OAuthClientCredentials: &MCPServerOAuthClientCredentialsConfigProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   					ExchangeUrl: jsii.String("exchangeUrl"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		McpServerNewRelic: &NewRelicServiceDetailsProperty{
//   			AuthorizationConfig: &NewRelicAuthorizationConfigProperty{
//   				ApiKey: &NewRelicApiKeyConfigProperty{
//   					AccountId: jsii.String("accountId"),
//   					AlertPolicyIds: []*string{
//   						jsii.String("alertPolicyIds"),
//   					},
//   					ApiKey: jsii.String("apiKey"),
//   					ApplicationIds: []*string{
//   						jsii.String("applicationIds"),
//   					},
//   					EntityGuids: []*string{
//   						jsii.String("entityGuids"),
//   					},
//   					Region: jsii.String("region"),
//   				},
//   			},
//   		},
//   		McpServerSigV4: &MCPServerSigV4DetailsProperty{
//   			AuthorizationConfig: &MCPServerSigV4AuthorizationConfigProperty{
//   				CustomHeaders: map[string]*string{
//   					"customHeadersKey": jsii.String("customHeaders"),
//   				},
//   				Region: jsii.String("region"),
//   				RoleArn: jsii.String("roleArn"),
//   				Service: jsii.String("service"),
//   			},
//   			Description: jsii.String("description"),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		McpServerSplunk: &MCPServerSplunkDetailsProperty{
//   			AuthorizationConfig: &MCPServerSplunkAuthorizationConfigProperty{
//   				BearerToken: &BearerTokenDetailsProperty{
//   					AuthorizationHeader: jsii.String("authorizationHeader"),
//   					TokenName: jsii.String("tokenName"),
//   					TokenValue: jsii.String("tokenValue"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		PagerDuty: &PagerDutyDetailsProperty{
//   			AuthorizationConfig: &PagerDutyAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//   		},
//   		ServiceNow: &ServiceNowServiceDetailsProperty{
//   			AuthorizationConfig: &ServiceNowAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   			InstanceUrl: jsii.String("instanceUrl"),
//   		},
//   	},
//   	ServiceType: jsii.String("serviceType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html
//
type CfnServicePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnServiceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServicePropsMixin
type jsiiProxy_CfnServicePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnServicePropsMixin) Props() *CfnServiceMixinProps {
	var returns *CfnServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServicePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DevOpsAgent::Service`.
func NewCfnServicePropsMixin(props *CfnServiceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DevOpsAgent::Service`.
func NewCfnServicePropsMixin_Override(c CfnServicePropsMixin, props *CfnServiceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServicePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

