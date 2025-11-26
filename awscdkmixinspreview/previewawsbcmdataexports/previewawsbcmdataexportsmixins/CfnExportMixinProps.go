package previewawsbcmdataexportsmixins


// Properties for CfnExportPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnExportMixinProps := &CfnExportMixinProps{
//   	Export: &ExportProperty{
//   		DataQuery: &DataQueryProperty{
//   			QueryStatement: jsii.String("queryStatement"),
//   			TableConfigurations: map[string]interface{}{
//   				"tableConfigurationsKey": map[string]*string{
//   					"tableConfigurationsKey": jsii.String("tableConfigurations"),
//   				},
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		DestinationConfigurations: &DestinationConfigurationsProperty{
//   			S3Destination: &S3DestinationProperty{
//   				S3Bucket: jsii.String("s3Bucket"),
//   				S3OutputConfigurations: &S3OutputConfigurationsProperty{
//   					Compression: jsii.String("compression"),
//   					Format: jsii.String("format"),
//   					OutputType: jsii.String("outputType"),
//   					Overwrite: jsii.String("overwrite"),
//   				},
//   				S3Prefix: jsii.String("s3Prefix"),
//   				S3Region: jsii.String("s3Region"),
//   			},
//   		},
//   		ExportArn: jsii.String("exportArn"),
//   		Name: jsii.String("name"),
//   		RefreshCadence: &RefreshCadenceProperty{
//   			Frequency: jsii.String("frequency"),
//   		},
//   	},
//   	Tags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-export.html
//
type CfnExportMixinProps struct {
	// The details that are available for an export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-export.html#cfn-bcmdataexports-export-export
	//
	Export interface{} `field:"optional" json:"export" yaml:"export"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-export.html#cfn-bcmdataexports-export-tags
	//
	Tags *[]*CfnExportPropsMixin_ResourceTagProperty `field:"optional" json:"tags" yaml:"tags"`
}

