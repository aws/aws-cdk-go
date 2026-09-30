package awsglue


// OAuth2 configuration container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorOAuth2PropertiesProperty := &ConnectorOAuth2PropertiesProperty{
//   	OAuth2GrantType: jsii.String("oAuth2GrantType"),
//
//   	// the properties below are optional
//   	AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   		AuthorizationCode: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		ClientId: &SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		ClientSecret: &SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		ContentType: jsii.String("contentType"),
//   		Prompt: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		RedirectUri: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		RequestMethod: jsii.String("requestMethod"),
//   		Scope: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		TokenUrl: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		TokenUrlParameters: []interface{}{
//   			&ConnectorPropertyProperty{
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
//   		},
//   	},
//   	ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   		ClientId: &SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		ClientSecret: &SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		ContentType: jsii.String("contentType"),
//   		RequestMethod: jsii.String("requestMethod"),
//   		Scope: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		TokenUrl: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		TokenUrlParameters: []interface{}{
//   			&ConnectorPropertyProperty{
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
//   		},
//   	},
//   	JwtBearerProperties: &JWTBearerPropertiesProperty{
//   		ContentType: jsii.String("contentType"),
//   		JwtToken: &SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		RequestMethod: jsii.String("requestMethod"),
//   		TokenUrl: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   		TokenUrlParameters: []interface{}{
//   			&ConnectorPropertyProperty{
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectoroauth2properties.html
//
type CfnConnectionType_ConnectorOAuth2PropertiesProperty struct {
	// The OAuth2 grant type to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectoroauth2properties.html#cfn-glue-connectiontype-connectoroauth2properties-oauth2granttype
	//
	OAuth2GrantType *string `field:"required" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
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
}

