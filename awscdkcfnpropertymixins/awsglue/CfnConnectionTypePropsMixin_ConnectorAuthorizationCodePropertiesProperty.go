package awsglue


// OAuth2 authorization code configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectorAuthorizationCodePropertiesProperty := &ConnectorAuthorizationCodePropertiesProperty{
//   	AuthorizationCode: &ConnectorPropertyProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	ClientId: &SecretConnectorPropertyProperty{
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	ClientSecret: &SecretConnectorPropertyProperty{
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	ContentType: jsii.String("contentType"),
//   	Prompt: &ConnectorPropertyProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	RedirectUri: &ConnectorPropertyProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	RequestMethod: jsii.String("requestMethod"),
//   	Scope: &ConnectorPropertyProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	TokenUrl: &ConnectorPropertyProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	TokenUrlParameters: []interface{}{
//   		&ConnectorPropertyProperty{
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html
//
type CfnConnectionTypePropsMixin_ConnectorAuthorizationCodePropertiesProperty struct {
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-authorizationcode
	//
	AuthorizationCode interface{} `field:"optional" json:"authorizationCode" yaml:"authorizationCode"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-authorizationcodeurl
	//
	AuthorizationCodeUrl interface{} `field:"optional" json:"authorizationCodeUrl" yaml:"authorizationCodeUrl"`
	// Defines a secret property configuration.
	//
	// SECRET-type properties cannot have DefaultValue or AllowedValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-clientid
	//
	ClientId interface{} `field:"optional" json:"clientId" yaml:"clientId"`
	// Defines a secret property configuration.
	//
	// SECRET-type properties cannot have DefaultValue or AllowedValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-clientsecret
	//
	ClientSecret interface{} `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-prompt
	//
	Prompt interface{} `field:"optional" json:"prompt" yaml:"prompt"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-redirecturi
	//
	RedirectUri interface{} `field:"optional" json:"redirectUri" yaml:"redirectUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-requestmethod
	//
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-scope
	//
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-tokenurl
	//
	TokenUrl interface{} `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorauthorizationcodeproperties.html#cfn-glue-connectiontype-connectorauthorizationcodeproperties-tokenurlparameters
	//
	TokenUrlParameters interface{} `field:"optional" json:"tokenUrlParameters" yaml:"tokenUrlParameters"`
}

