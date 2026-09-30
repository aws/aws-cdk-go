package awsglue


// Cursor-based pagination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cursorConfigurationProperty := &CursorConfigurationProperty{
//   	LimitParameter: &ExtractedParameterProperty{
//   		DefaultValue: jsii.String("defaultValue"),
//   		Key: jsii.String("key"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		Value: &ResponseExtractionMappingProperty{
//   			ContentPath: jsii.String("contentPath"),
//   			HeaderKey: jsii.String("headerKey"),
//   		},
//   	},
//   	NextPage: &ExtractedParameterProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-cursorconfiguration.html
//
type CfnConnectionTypePropsMixin_CursorConfigurationProperty struct {
	// Parameter extraction configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-cursorconfiguration.html#cfn-glue-connectiontype-cursorconfiguration-limitparameter
	//
	LimitParameter interface{} `field:"optional" json:"limitParameter" yaml:"limitParameter"`
	// Parameter extraction configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-cursorconfiguration.html#cfn-glue-connectiontype-cursorconfiguration-nextpage
	//
	NextPage interface{} `field:"optional" json:"nextPage" yaml:"nextPage"`
}

