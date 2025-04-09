package awsappflow


// Properties for defining a `CfnConnectorProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProfileProps := &CfnConnectorProfileProps{
//   	ConnectionMode: jsii.String("connectionMode"),
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	ConnectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	ConnectorLabel: jsii.String("connectorLabel"),
//   	ConnectorProfileConfig: &ConnectorProfileConfigProperty{
//   		ConnectorProfileCredentials: &ConnectorProfileCredentialsProperty{
//   			Amplitude: &AmplitudeConnectorProfileCredentialsProperty{
//   				ApiKey: jsii.String("apiKey"),
//   				SecretKey: jsii.String("secretKey"),
//   			},
//   			CustomConnector: &CustomConnectorProfileCredentialsProperty{
//   				AuthenticationType: jsii.String("authenticationType"),
//
//   				// the properties below are optional
//   				ApiKey: &ApiKeyCredentialsProperty{
//   					ApiKey: jsii.String("apiKey"),
//
//   					// the properties below are optional
//   					ApiSecretKey: jsii.String("apiSecretKey"),
//   				},
//   				Basic: &BasicAuthCredentialsProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   				Custom: &CustomAuthCredentialsProperty{
//   					CustomAuthenticationType: jsii.String("customAuthenticationType"),
//
//   					// the properties below are optional
//   					CredentialsMap: map[string]*string{
//   						"credentialsMapKey": jsii.String("credentialsMap"),
//   					},
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
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
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
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
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
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
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
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				AccessToken: jsii.String("accessToken"),
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
//
//   				// the properties below are optional
//   				InstanceUrl: jsii.String("instanceUrl"),
//   				IsSandboxEnvironment: jsii.Boolean(false),
//   			},
//   			Redshift: &RedshiftConnectorProfilePropertiesProperty{
//   				BucketName: jsii.String("bucketName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				ClusterIdentifier: jsii.String("clusterIdentifier"),
//   				DataApiRoleArn: jsii.String("dataApiRoleArn"),
//   				DatabaseName: jsii.String("databaseName"),
//   				DatabaseUrl: jsii.String("databaseUrl"),
//   				IsRedshiftServerless: jsii.Boolean(false),
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
//   				BucketName: jsii.String("bucketName"),
//   				Stage: jsii.String("stage"),
//   				Warehouse: jsii.String("warehouse"),
//
//   				// the properties below are optional
//   				AccountName: jsii.String("accountName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   				Region: jsii.String("region"),
//   			},
//   			Veeva: &VeevaConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   			Zendesk: &ZendeskConnectorProfilePropertiesProperty{
//   				InstanceUrl: jsii.String("instanceUrl"),
//   			},
//   		},
//   	},
//   	KmsArn: jsii.String("kmsArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html
//
type CfnConnectorProfileProps struct {
	// Indicates the connection mode and if it is public or private.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectionmode
	//
	ConnectionMode *string `field:"required" json:"connectionMode" yaml:"connectionMode"`
	// The name of the connector profile.
	//
	// The name is unique for each `ConnectorProfile` in the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectorprofilename
	//
	ConnectorProfileName *string `field:"required" json:"connectorProfileName" yaml:"connectorProfileName"`
	// The type of connector, such as Salesforce, Amplitude, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectortype
	//
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// The label for the connector profile being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectorlabel
	//
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// Defines the connector-specific configuration and credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectorprofileconfig
	//
	ConnectorProfileConfig interface{} `field:"optional" json:"connectorProfileConfig" yaml:"connectorProfileConfig"`
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-kmsarn
	//
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}

