package awsglue


// JWT bearer token configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jWTBearerPropertiesProperty := &JWTBearerPropertiesProperty{
//   	ContentType: jsii.String("contentType"),
//   	JwtToken: &SecretConnectorPropertyProperty{
//   		Name: jsii.String("name"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		KeyOverride: jsii.String("keyOverride"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   	},
//   	RequestMethod: jsii.String("requestMethod"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-jwtbearerproperties.html
//
type CfnConnectionType_JWTBearerPropertiesProperty struct {
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

