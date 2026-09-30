package awsglue


// OAuth2 configuration container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectorOAuth2PropertiesProperty := &ConnectorOAuth2PropertiesProperty{
//   	AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   		AuthorizationCode: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		ClientId: &SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		ClientSecret: &SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		ContentType: jsii.String("contentType"),
//   		Prompt: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		RedirectUri: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		RequestMethod: jsii.String("requestMethod"),
//   		Scope: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		TokenUrl: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		TokenUrlParameters: []interface{}{
//   			&ConnectorPropertyProperty{
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
//   		},
//   	},
//   	ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   		ClientId: &SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		ClientSecret: &SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		ContentType: jsii.String("contentType"),
//   		RequestMethod: jsii.String("requestMethod"),
//   		Scope: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		TokenUrl: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		TokenUrlParameters: []interface{}{
//   			&ConnectorPropertyProperty{
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
//   		},
//   	},
//   	JwtBearerProperties: &JWTBearerPropertiesProperty{
//   		ContentType: jsii.String("contentType"),
//   		JwtToken: &SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		RequestMethod: jsii.String("requestMethod"),
//   		TokenUrl: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   		TokenUrlParameters: []interface{}{
//   			&ConnectorPropertyProperty{
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
//   		},
//   	},
//   	OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectoroauth2properties.html
//
type CfnConnectionTypePropsMixin_ConnectorOAuth2PropertiesProperty struct {
	// OAuth2 authorization code configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectoroauth2properties.html#cfn-glue-connectiontype-connectoroauth2properties-authorizationcodeproperties
	//
	AuthorizationCodeProperties interface{} `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// OAuth2 client credentials configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectoroauth2properties.html#cfn-glue-connectiontype-connectoroauth2properties-clientcredentialsproperties
	//
	ClientCredentialsProperties interface{} `field:"optional" json:"clientCredentialsProperties" yaml:"clientCredentialsProperties"`
	// JWT bearer token configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectoroauth2properties.html#cfn-glue-connectiontype-connectoroauth2properties-jwtbearerproperties
	//
	JwtBearerProperties interface{} `field:"optional" json:"jwtBearerProperties" yaml:"jwtBearerProperties"`
	// The OAuth2 grant type to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectoroauth2properties.html#cfn-glue-connectiontype-connectoroauth2properties-oauth2granttype
	//
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
}

