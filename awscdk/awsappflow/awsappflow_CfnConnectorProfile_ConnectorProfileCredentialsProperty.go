package awsappflow


// The connector-specific credentials required by a connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorProfileCredentialsProperty := &connectorProfileCredentialsProperty{
//   	amplitude: &amplitudeConnectorProfileCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//   		secretKey: jsii.String("secretKey"),
//   	},
//   	customConnector: &customConnectorProfileCredentialsProperty{
//   		authenticationType: jsii.String("authenticationType"),
//
//   		// the properties below are optional
//   		apiKey: &apiKeyCredentialsProperty{
//   			apiKey: jsii.String("apiKey"),
//
//   			// the properties below are optional
//   			apiSecretKey: jsii.String("apiSecretKey"),
//   		},
//   		basic: &basicAuthCredentialsProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		custom: &customAuthCredentialsProperty{
//   			customAuthenticationType: jsii.String("customAuthenticationType"),
//
//   			// the properties below are optional
//   			credentialsMap: map[string]*string{
//   				"credentialsMapKey": jsii.String("credentialsMap"),
//   			},
//   		},
//   		oauth2: &oAuth2CredentialsProperty{
//   			accessToken: jsii.String("accessToken"),
//   			clientId: jsii.String("clientId"),
//   			clientSecret: jsii.String("clientSecret"),
//   			oAuthRequest: &connectorOAuthRequestProperty{
//   				authCode: jsii.String("authCode"),
//   				redirectUri: jsii.String("redirectUri"),
//   			},
//   			refreshToken: jsii.String("refreshToken"),
//   		},
//   	},
//   	datadog: &datadogConnectorProfileCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//   		applicationKey: jsii.String("applicationKey"),
//   	},
//   	dynatrace: &dynatraceConnectorProfileCredentialsProperty{
//   		apiToken: jsii.String("apiToken"),
//   	},
//   	googleAnalytics: &googleAnalyticsConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	inforNexus: &inforNexusConnectorProfileCredentialsProperty{
//   		accessKeyId: jsii.String("accessKeyId"),
//   		datakey: jsii.String("datakey"),
//   		secretAccessKey: jsii.String("secretAccessKey"),
//   		userId: jsii.String("userId"),
//   	},
//   	marketo: &marketoConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   	redshift: &redshiftConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	salesforce: &salesforceConnectorProfileCredentialsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	sapoData: &sAPODataConnectorProfileCredentialsProperty{
//   		basicAuthCredentials: &basicAuthCredentialsProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		oAuthCredentials: &oAuthCredentialsProperty{
//   			accessToken: jsii.String("accessToken"),
//   			clientId: jsii.String("clientId"),
//   			clientSecret: jsii.String("clientSecret"),
//   			connectorOAuthRequest: &connectorOAuthRequestProperty{
//   				authCode: jsii.String("authCode"),
//   				redirectUri: jsii.String("redirectUri"),
//   			},
//   			refreshToken: jsii.String("refreshToken"),
//   		},
//   	},
//   	serviceNow: &serviceNowConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	singular: &singularConnectorProfileCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//   	},
//   	slack: &slackConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   	snowflake: &snowflakeConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	trendmicro: &trendmicroConnectorProfileCredentialsProperty{
//   		apiSecretKey: jsii.String("apiSecretKey"),
//   	},
//   	veeva: &veevaConnectorProfileCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	zendesk: &zendeskConnectorProfileCredentialsProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		accessToken: jsii.String("accessToken"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   	},
//   }
//
type CfnConnectorProfile_ConnectorProfileCredentialsProperty struct {
	// The connector-specific credentials required when using Amplitude.
	Amplitude interface{} `field:"optional" json:"amplitude" yaml:"amplitude"`
	// `CfnConnectorProfile.ConnectorProfileCredentialsProperty.CustomConnector`.
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

