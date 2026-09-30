package awsglue


// Configuration for handling paginated responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paginationConfigurationProperty := &PaginationConfigurationProperty{
//   	CursorConfiguration: &CursorConfigurationProperty{
//   		NextPage: &ExtractedParameterProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Key: jsii.String("key"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			Value: &ResponseExtractionMappingProperty{
//   				ContentPath: jsii.String("contentPath"),
//   				HeaderKey: jsii.String("headerKey"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		LimitParameter: &ExtractedParameterProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Key: jsii.String("key"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			Value: &ResponseExtractionMappingProperty{
//   				ContentPath: jsii.String("contentPath"),
//   				HeaderKey: jsii.String("headerKey"),
//   			},
//   		},
//   	},
//   	OffsetConfiguration: &OffsetConfigurationProperty{
//   		LimitParameter: &ExtractedParameterProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Key: jsii.String("key"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			Value: &ResponseExtractionMappingProperty{
//   				ContentPath: jsii.String("contentPath"),
//   				HeaderKey: jsii.String("headerKey"),
//   			},
//   		},
//   		OffsetParameter: &ExtractedParameterProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Key: jsii.String("key"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			Value: &ResponseExtractionMappingProperty{
//   				ContentPath: jsii.String("contentPath"),
//   				HeaderKey: jsii.String("headerKey"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-paginationconfiguration.html
//
type CfnConnectionType_PaginationConfigurationProperty struct {
	// Cursor-based pagination configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-paginationconfiguration.html#cfn-glue-connectiontype-paginationconfiguration-cursorconfiguration
	//
	CursorConfiguration interface{} `field:"optional" json:"cursorConfiguration" yaml:"cursorConfiguration"`
	// Offset-based pagination configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-paginationconfiguration.html#cfn-glue-connectiontype-paginationconfiguration-offsetconfiguration
	//
	OffsetConfiguration interface{} `field:"optional" json:"offsetConfiguration" yaml:"offsetConfiguration"`
}

