package awsappflow


// Properties for defining a `CfnConnectorProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProfileProps := &cfnConnectorProfileProps{
//   	connectionMode: jsii.String("connectionMode"),
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	connectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	connectorLabel: jsii.String("connectorLabel"),
//   	connectorProfileConfig: &connectorProfileConfigProperty{
//   		connectorProfileCredentials: &connectorProfileCredentialsProperty{
//   			amplitude: &amplitudeConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				secretKey: jsii.String("secretKey"),
//   			},
//   			customConnector: &customConnectorProfileCredentialsProperty{
//   				authenticationType: jsii.String("authenticationType"),
//
//   				// the properties below are optional
//   				apiKey: &apiKeyCredentialsProperty{
//   					apiKey: jsii.String("apiKey"),
//
//   					// the properties below are optional
//   					apiSecretKey: jsii.String("apiSecretKey"),
//   				},
//   				basic: &basicAuthCredentialsProperty{
//   					password: jsii.String("password"),
//   					username: jsii.String("username"),
//   				},
//   				custom: &customAuthCredentialsProperty{
//   					customAuthenticationType: jsii.String("customAuthenticationType"),
//
//   					// the properties below are optional
//   					credentialsMap: map[string]*string{
//   						"credentialsMapKey": jsii.String("credentialsMap"),
//   					},
//   				},
//   				oauth2: &oAuth2CredentialsProperty{
//   					accessToken: jsii.String("accessToken"),
//   					clientId: jsii.String("clientId"),
//   					clientSecret: jsii.String("clientSecret"),
//   					oAuthRequest: &connectorOAuthRequestProperty{
//   						authCode: jsii.String("authCode"),
//   						redirectUri: jsii.String("redirectUri"),
//   					},
//   					refreshToken: jsii.String("refreshToken"),
//   				},
//   			},
//   			datadog: &datadogConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				applicationKey: jsii.String("applicationKey"),
//   			},
//   			dynatrace: &dynatraceConnectorProfileCredentialsProperty{
//   				apiToken: jsii.String("apiToken"),
//   			},
//   			googleAnalytics: &googleAnalyticsConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			inforNexus: &inforNexusConnectorProfileCredentialsProperty{
//   				accessKeyId: jsii.String("accessKeyId"),
//   				datakey: jsii.String("datakey"),
//   				secretAccessKey: jsii.String("secretAccessKey"),
//   				userId: jsii.String("userId"),
//   			},
//   			marketo: &marketoConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			redshift: &redshiftConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			salesforce: &salesforceConnectorProfileCredentialsProperty{
//   				accessToken: jsii.String("accessToken"),
//   				clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			sapoData: &sAPODataConnectorProfileCredentialsProperty{
//   				basicAuthCredentials: &basicAuthCredentialsProperty{
//   					password: jsii.String("password"),
//   					username: jsii.String("username"),
//   				},
//   				oAuthCredentials: &oAuthCredentialsProperty{
//   					accessToken: jsii.String("accessToken"),
//   					clientId: jsii.String("clientId"),
//   					clientSecret: jsii.String("clientSecret"),
//   					connectorOAuthRequest: &connectorOAuthRequestProperty{
//   						authCode: jsii.String("authCode"),
//   						redirectUri: jsii.String("redirectUri"),
//   					},
//   					refreshToken: jsii.String("refreshToken"),
//   				},
//   			},
//   			serviceNow: &serviceNowConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			singular: &singularConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   			},
//   			slack: &slackConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			snowflake: &snowflakeConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			trendmicro: &trendmicroConnectorProfileCredentialsProperty{
//   				apiSecretKey: jsii.String("apiSecretKey"),
//   			},
//   			veeva: &veevaConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			zendesk: &zendeskConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   		},
//   		connectorProfileProperties: &connectorProfilePropertiesProperty{
//   			customConnector: &customConnectorProfilePropertiesProperty{
//   				oAuth2Properties: &oAuth2PropertiesProperty{
//   					oAuth2GrantType: jsii.String("oAuth2GrantType"),
//   					tokenUrl: jsii.String("tokenUrl"),
//   					tokenUrlCustomProperties: map[string]*string{
//   						"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   					},
//   				},
//   				profileProperties: map[string]*string{
//   					"profilePropertiesKey": jsii.String("profileProperties"),
//   				},
//   			},
//   			datadog: &datadogConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			dynatrace: &dynatraceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			inforNexus: &inforNexusConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			marketo: &marketoConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			redshift: &redshiftConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				clusterIdentifier: jsii.String("clusterIdentifier"),
//   				dataApiRoleArn: jsii.String("dataApiRoleArn"),
//   				databaseName: jsii.String("databaseName"),
//   				databaseUrl: jsii.String("databaseUrl"),
//   				isRedshiftServerless: jsii.Boolean(false),
//   				workgroupName: jsii.String("workgroupName"),
//   			},
//   			salesforce: &salesforceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   				isSandboxEnvironment: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataConnectorProfilePropertiesProperty{
//   				applicationHostUrl: jsii.String("applicationHostUrl"),
//   				applicationServicePath: jsii.String("applicationServicePath"),
//   				clientNumber: jsii.String("clientNumber"),
//   				logonLanguage: jsii.String("logonLanguage"),
//   				oAuthProperties: &oAuthPropertiesProperty{
//   					authCodeUrl: jsii.String("authCodeUrl"),
//   					oAuthScopes: []*string{
//   						jsii.String("oAuthScopes"),
//   					},
//   					tokenUrl: jsii.String("tokenUrl"),
//   				},
//   				portNumber: jsii.Number(123),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   			},
//   			serviceNow: &serviceNowConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			slack: &slackConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			snowflake: &snowflakeConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				stage: jsii.String("stage"),
//   				warehouse: jsii.String("warehouse"),
//
//   				// the properties below are optional
//   				accountName: jsii.String("accountName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   				region: jsii.String("region"),
//   			},
//   			veeva: &veevaConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			zendesk: &zendeskConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   		},
//   	},
//   	kmsArn: jsii.String("kmsArn"),
//   }
//
type CfnConnectorProfileProps struct {
	// Indicates the connection mode and if it is public or private.
	ConnectionMode *string `field:"required" json:"connectionMode" yaml:"connectionMode"`
	// The name of the connector profile.
	//
	// The name is unique for each `ConnectorProfile` in the AWS account .
	ConnectorProfileName *string `field:"required" json:"connectorProfileName" yaml:"connectorProfileName"`
	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// `AWS::AppFlow::ConnectorProfile.ConnectorLabel`.
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// Defines the connector-specific configuration and credentials.
	ConnectorProfileConfig interface{} `field:"optional" json:"connectorProfileConfig" yaml:"connectorProfileConfig"`
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}

