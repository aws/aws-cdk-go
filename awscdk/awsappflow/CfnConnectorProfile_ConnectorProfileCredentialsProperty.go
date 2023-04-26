package awsappflow


// The connector-specific credentials required by a connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorProfileCredentialsProperty := &ConnectorProfileCredentialsProperty{
//   	Amplitude: &AmplitudeConnectorProfileCredentialsProperty{
//   		ApiKey: jsii.String("apiKey"),
//   		SecretKey: jsii.String("secretKey"),
//   	},
//   	CustomConnector: &CustomConnectorProfileCredentialsProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//
//   		// the properties below are optional
//   		ApiKey: &ApiKeyCredentialsProperty{
//   			ApiKey: jsii.String("apiKey"),
//
//   			// the properties below are optional
//   			ApiSecretKey: jsii.String("apiSecretKey"),
//   		},
//   		Basic: &BasicAuthCredentialsProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		Custom: &CustomAuthCredentialsProperty{
//   			CustomAuthenticationType: jsii.String("customAuthenticationType"),
//
//   			// the properties below are optional
//   			CredentialsMap: map[string]*string{
//   				"credentialsMapKey": jsii.String("credentialsMap"),
//   			},
//   		},
//   		Oauth2: &OAuth2CredentialsProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			OAuthRequest: &ConnectorOAuthRequestProperty{
//   				AuthCode: jsii.String("authCode"),
//   				RedirectUri: jsii.String("redirectUri"),
//   			},
//   			RefreshToken: jsii.String("refreshToken"),
//   		},
//   	},
//   	Datadog: &DatadogConnectorProfileCredentialsProperty{
//   		ApiKey: jsii.String("apiKey"),
//   		ApplicationKey: jsii.String("applicationKey"),
//   	},
//   	Dynatrace: &DynatraceConnectorProfileCredentialsProperty{
//   		ApiToken: jsii.String("apiToken"),
//   	},
//   	GoogleAnalytics: &GoogleAnalyticsConnectorProfileCredentialsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		AccessToken: jsii.String("accessToken"),
//   		ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	InforNexus: &InforNexusConnectorProfileCredentialsProperty{
//   		AccessKeyId: jsii.String("accessKeyId"),
//   		Datakey: jsii.String("datakey"),
//   		SecretAccessKey: jsii.String("secretAccessKey"),
//   		UserId: jsii.String("userId"),
//   	},
//   	Marketo: &MarketoConnectorProfileCredentialsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		AccessToken: jsii.String("accessToken"),
//   		ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   	Pardot: &PardotConnectorProfileCredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   		ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	Redshift: &RedshiftConnectorProfileCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	Salesforce: &SalesforceConnectorProfileCredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   		ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	SapoData: &SAPODataConnectorProfileCredentialsProperty{
//   		BasicAuthCredentials: &BasicAuthCredentialsProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		OAuthCredentials: &OAuthCredentialsProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   				AuthCode: jsii.String("authCode"),
//   				RedirectUri: jsii.String("redirectUri"),
//   			},
//   			RefreshToken: jsii.String("refreshToken"),
//   		},
//   	},
//   	ServiceNow: &ServiceNowConnectorProfileCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	Singular: &SingularConnectorProfileCredentialsProperty{
//   		ApiKey: jsii.String("apiKey"),
//   	},
//   	Slack: &SlackConnectorProfileCredentialsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		AccessToken: jsii.String("accessToken"),
//   		ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   	Snowflake: &SnowflakeConnectorProfileCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	Trendmicro: &TrendmicroConnectorProfileCredentialsProperty{
//   		ApiSecretKey: jsii.String("apiSecretKey"),
//   	},
//   	Veeva: &VeevaConnectorProfileCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	Zendesk: &ZendeskConnectorProfileCredentialsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		AccessToken: jsii.String("accessToken"),
//   		ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   }
//
type CfnConnectorProfile_ConnectorProfileCredentialsProperty struct {
	// The connector-specific credentials required when using Amplitude.
	Amplitude interface{} `field:"optional" json:"amplitude" yaml:"amplitude"`
	// The connector-specific profile credentials that are required when using the custom connector.
	CustomConnector interface{} `field:"optional" json:"customConnector" yaml:"customConnector"`
	// The connector-specific credentials required when using Datadog.
	Datadog interface{} `field:"optional" json:"datadog" yaml:"datadog"`
	// The connector-specific credentials required when using Dynatrace.
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// The connector-specific credentials required when using Google Analytics.
	GoogleAnalytics interface{} `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// The connector-specific credentials required when using Infor Nexus.
	InforNexus interface{} `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// The connector-specific credentials required when using Marketo.
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.Pardot`.
	Pardot interface{} `field:"optional" json:"pardot" yaml:"pardot"`
	// The connector-specific credentials required when using Amazon Redshift.
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// The connector-specific credentials required when using Salesforce.
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The connector-specific profile credentials required when using SAPOData.
	SapoData interface{} `field:"optional" json:"sapoData" yaml:"sapoData"`
	// The connector-specific credentials required when using ServiceNow.
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The connector-specific credentials required when using Singular.
	Singular interface{} `field:"optional" json:"singular" yaml:"singular"`
	// The connector-specific credentials required when using Slack.
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
	// The connector-specific credentials required when using Snowflake.
	Snowflake interface{} `field:"optional" json:"snowflake" yaml:"snowflake"`
	// The connector-specific credentials required when using Trend Micro.
	Trendmicro interface{} `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// The connector-specific credentials required when using Veeva.
	Veeva interface{} `field:"optional" json:"veeva" yaml:"veeva"`
	// The connector-specific credentials required when using Zendesk.
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

