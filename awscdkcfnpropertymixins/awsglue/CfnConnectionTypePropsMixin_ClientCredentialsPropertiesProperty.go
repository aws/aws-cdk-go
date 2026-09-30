package awsglue


// OAuth2 client credentials configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   clientCredentialsPropertiesProperty := &ClientCredentialsPropertiesProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html
//
type CfnConnectionTypePropsMixin_ClientCredentialsPropertiesProperty struct {
	// Defines a secret property configuration.
	//
	// SECRET-type properties cannot have DefaultValue or AllowedValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html#cfn-glue-connectiontype-clientcredentialsproperties-clientid
	//
	ClientId interface{} `field:"optional" json:"clientId" yaml:"clientId"`
	// Defines a secret property configuration.
	//
	// SECRET-type properties cannot have DefaultValue or AllowedValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html#cfn-glue-connectiontype-clientcredentialsproperties-clientsecret
	//
	ClientSecret interface{} `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html#cfn-glue-connectiontype-clientcredentialsproperties-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html#cfn-glue-connectiontype-clientcredentialsproperties-requestmethod
	//
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html#cfn-glue-connectiontype-clientcredentialsproperties-scope
	//
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html#cfn-glue-connectiontype-clientcredentialsproperties-tokenurl
	//
	TokenUrl interface{} `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html#cfn-glue-connectiontype-clientcredentialsproperties-tokenurlparameters
	//
	TokenUrlParameters interface{} `field:"optional" json:"tokenUrlParameters" yaml:"tokenUrlParameters"`
}

