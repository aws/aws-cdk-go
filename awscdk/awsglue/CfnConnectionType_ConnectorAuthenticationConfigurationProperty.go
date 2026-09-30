package awsglue


// Configuration that defines supported authentication types and required properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorAuthenticationConfigurationProperty := &ConnectorAuthenticationConfigurationProperty{
//   	AuthenticationTypes: []*string{
//   		jsii.String("authenticationTypes"),
//   	},
//
//   	// the properties below are optional
//   	BasicAuthenticationProperties: &BasicAuthenticationPropertiesProperty{
//   		Password: &SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		Username: &SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   	},
//   	CustomAuthenticationProperties: &CustomAuthenticationPropertiesProperty{
//   		AuthenticationParameters: []interface{}{
//   			&SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   		},
//   	},
//   	OAuth2Properties: &ConnectorOAuth2PropertiesProperty{
//   		OAuth2GrantType: jsii.String("oAuth2GrantType"),
//
//   		// the properties below are optional
//   		AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   			AuthorizationCode: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			ClientId: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			ClientSecret: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			ContentType: jsii.String("contentType"),
//   			Prompt: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			RedirectUri: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			RequestMethod: jsii.String("requestMethod"),
//   			Scope: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			TokenUrl: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			TokenUrlParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   			},
//   		},
//   		ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   			ClientId: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			ClientSecret: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			ContentType: jsii.String("contentType"),
//   			RequestMethod: jsii.String("requestMethod"),
//   			Scope: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			TokenUrl: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			TokenUrlParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   			},
//   		},
//   		JwtBearerProperties: &JWTBearerPropertiesProperty{
//   			ContentType: jsii.String("contentType"),
//   			JwtToken: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			RequestMethod: jsii.String("requestMethod"),
//   			TokenUrl: &ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			TokenUrlParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthenticationconfiguration.html
//
type CfnConnectionType_ConnectorAuthenticationConfigurationProperty struct {
	// A list of authentication types supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthenticationconfiguration.html#cfn-glue-connectiontype-connectorauthenticationconfiguration-authenticationtypes
	//
	AuthenticationTypes *[]*string `field:"required" json:"authenticationTypes" yaml:"authenticationTypes"`
	// Basic authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthenticationconfiguration.html#cfn-glue-connectiontype-connectorauthenticationconfiguration-basicauthenticationproperties
	//
	BasicAuthenticationProperties interface{} `field:"optional" json:"basicAuthenticationProperties" yaml:"basicAuthenticationProperties"`
	// Custom authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthenticationconfiguration.html#cfn-glue-connectiontype-connectorauthenticationconfiguration-customauthenticationproperties
	//
	CustomAuthenticationProperties interface{} `field:"optional" json:"customAuthenticationProperties" yaml:"customAuthenticationProperties"`
	// OAuth2 configuration container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthenticationconfiguration.html#cfn-glue-connectiontype-connectorauthenticationconfiguration-oauth2properties
	//
	OAuth2Properties interface{} `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
}

