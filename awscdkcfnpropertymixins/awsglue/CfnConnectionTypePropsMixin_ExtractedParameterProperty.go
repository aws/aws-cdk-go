package awsglue


// Parameter extraction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   extractedParameterProperty := &ExtractedParameterProperty{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Key: jsii.String("key"),
//   	PropertyLocation: jsii.String("propertyLocation"),
//   	Value: &ResponseExtractionMappingProperty{
//   		ContentPath: jsii.String("contentPath"),
//   		HeaderKey: jsii.String("headerKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-extractedparameter.html
//
type CfnConnectionTypePropsMixin_ExtractedParameterProperty struct {
	// The default value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-extractedparameter.html#cfn-glue-connectiontype-extractedparameter-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The parameter key name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-extractedparameter.html#cfn-glue-connectiontype-extractedparameter-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Specifies where to place the parameter in requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-extractedparameter.html#cfn-glue-connectiontype-extractedparameter-propertylocation
	//
	PropertyLocation *string `field:"optional" json:"propertyLocation" yaml:"propertyLocation"`
	// Defines how to extract values from HTTP responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-extractedparameter.html#cfn-glue-connectiontype-extractedparameter-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

