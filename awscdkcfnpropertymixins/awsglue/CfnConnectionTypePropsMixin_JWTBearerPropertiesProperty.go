package awsglue


// JWT bearer token configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   jWTBearerPropertiesProperty := &JWTBearerPropertiesProperty{
//   	ContentType: jsii.String("contentType"),
//   	JwtToken: &SecretConnectorPropertyProperty{
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	RequestMethod: jsii.String("requestMethod"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-jwtbearerproperties.html
//
type CfnConnectionTypePropsMixin_JWTBearerPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-jwtbearerproperties.html#cfn-glue-connectiontype-jwtbearerproperties-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Defines a secret property configuration.
	//
	// SECRET-type properties cannot have DefaultValue or AllowedValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-jwtbearerproperties.html#cfn-glue-connectiontype-jwtbearerproperties-jwttoken
	//
	JwtToken interface{} `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-jwtbearerproperties.html#cfn-glue-connectiontype-jwtbearerproperties-requestmethod
	//
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-jwtbearerproperties.html#cfn-glue-connectiontype-jwtbearerproperties-tokenurl
	//
	TokenUrl interface{} `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-jwtbearerproperties.html#cfn-glue-connectiontype-jwtbearerproperties-tokenurlparameters
	//
	TokenUrlParameters interface{} `field:"optional" json:"tokenUrlParameters" yaml:"tokenUrlParameters"`
}

