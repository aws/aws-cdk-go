package awsglue


// Configuration that defines supported authentication types and required properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectorAuthenticationConfigurationProperty := &ConnectorAuthenticationConfigurationProperty{
//   	AuthenticationTypes: []*string{
//   		jsii.String("authenticationTypes"),
//   	},
//   	BasicAuthenticationProperties: &BasicAuthenticationPropertiesProperty{
//   		Password: &SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		Username: &SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   	},
//   	CustomAuthenticationProperties: &CustomAuthenticationPropertiesProperty{
//   		AuthenticationParameters: []interface{}{
//   			&SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	OAuth2Properties: &ConnectorOAuth2PropertiesProperty{
//   		AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   			AuthorizationCode: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			ClientId: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			ClientSecret: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			ContentType: jsii.String("contentType"),
//   			Prompt: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			RedirectUri: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			RequestMethod: jsii.String("requestMethod"),
//   			Scope: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			TokenUrl: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			TokenUrlParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   			ClientId: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			ClientSecret: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			ContentType: jsii.String("contentType"),
//   			RequestMethod: jsii.String("requestMethod"),
//   			Scope: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			TokenUrl: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			TokenUrlParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		JwtBearerProperties: &JWTBearerPropertiesProperty{
//   			ContentType: jsii.String("contentType"),
//   			JwtToken: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			RequestMethod: jsii.String("requestMethod"),
//   			TokenUrl: &ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			TokenUrlParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthenticationconfiguration.html
//
type CfnConnectionTypePropsMixin_ConnectorAuthenticationConfigurationProperty struct {
	// A list of authentication types supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthenticationconfiguration.html#cfn-glue-connectiontype-connectorauthenticationconfiguration-authenticationtypes
	//
	AuthenticationTypes *[]*string `field:"optional" json:"authenticationTypes" yaml:"authenticationTypes"`
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

