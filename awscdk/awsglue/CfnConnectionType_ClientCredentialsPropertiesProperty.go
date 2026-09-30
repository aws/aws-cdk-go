package awsglue


// OAuth2 client credentials configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientCredentialsPropertiesProperty := &ClientCredentialsPropertiesProperty{
//   	ClientId: &SecretConnectorPropertyProperty{
//   		Name: jsii.String("name"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		KeyOverride: jsii.String("keyOverride"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   	},
//   	ClientSecret: &SecretConnectorPropertyProperty{
//   		Name: jsii.String("name"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		KeyOverride: jsii.String("keyOverride"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   	},
//   	ContentType: jsii.String("contentType"),
//   	RequestMethod: jsii.String("requestMethod"),
//   	Scope: &ConnectorPropertyProperty{
//   		Name: jsii.String("name"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   	},
//   	TokenUrl: &ConnectorPropertyProperty{
//   		Name: jsii.String("name"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   	},
//   	TokenUrlParameters: []interface{}{
//   		&ConnectorPropertyProperty{
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-clientcredentialsproperties.html
//
type CfnConnectionType_ClientCredentialsPropertiesProperty struct {
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

