package awsglue


// Offset-based pagination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   offsetConfigurationProperty := &OffsetConfigurationProperty{
//   	LimitParameter: &ExtractedParameterProperty{
//   		DefaultValue: jsii.String("defaultValue"),
//   		Key: jsii.String("key"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		Value: &ResponseExtractionMappingProperty{
//   			ContentPath: jsii.String("contentPath"),
//   			HeaderKey: jsii.String("headerKey"),
//   		},
//   	},
//   	OffsetParameter: &ExtractedParameterProperty{
//   		DefaultValue: jsii.String("defaultValue"),
//   		Key: jsii.String("key"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		Value: &ResponseExtractionMappingProperty{
//   			ContentPath: jsii.String("contentPath"),
//   			HeaderKey: jsii.String("headerKey"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-offsetconfiguration.html
//
type CfnConnectionType_OffsetConfigurationProperty struct {
	// Parameter extraction configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-offsetconfiguration.html#cfn-glue-connectiontype-offsetconfiguration-limitparameter
	//
	LimitParameter interface{} `field:"required" json:"limitParameter" yaml:"limitParameter"`
	// Parameter extraction configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-offsetconfiguration.html#cfn-glue-connectiontype-offsetconfiguration-offsetparameter
	//
	OffsetParameter interface{} `field:"required" json:"offsetParameter" yaml:"offsetParameter"`
}

