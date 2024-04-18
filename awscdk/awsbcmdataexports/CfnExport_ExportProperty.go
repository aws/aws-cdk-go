package awsbcmdataexports


// The details that are available for an export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exportProperty := &ExportProperty{
//   	DataQuery: &DataQueryProperty{
//   		QueryStatement: jsii.String("queryStatement"),
//
//   		// the properties below are optional
//   		TableConfigurations: map[string]interface{}{
//   			"tableConfigurationsKey": map[string]*string{
//   				"tableConfigurationsKey": jsii.String("tableConfigurations"),
//   			},
//   		},
//   	},
//   	DestinationConfigurations: &DestinationConfigurationsProperty{
//   		S3Destination: &S3DestinationProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3OutputConfigurations: &S3OutputConfigurationsProperty{
//   				Compression: jsii.String("compression"),
//   				Format: jsii.String("format"),
//   				OutputType: jsii.String("outputType"),
//   				Overwrite: jsii.String("overwrite"),
//   			},
//   			S3Prefix: jsii.String("s3Prefix"),
//   			S3Region: jsii.String("s3Region"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RefreshCadence: &RefreshCadenceProperty{
//   		Frequency: jsii.String("frequency"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExportArn: jsii.String("exportArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-export.html
//
type CfnExport_ExportProperty struct {
	// The data query for this specific data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-export.html#cfn-bcmdataexports-export-export-dataquery
	//
	DataQuery interface{} `field:"required" json:"dataQuery" yaml:"dataQuery"`
	// The destination configuration for this specific data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-export.html#cfn-bcmdataexports-export-export-destinationconfigurations
	//
	DestinationConfigurations interface{} `field:"required" json:"destinationConfigurations" yaml:"destinationConfigurations"`
	// The name of this specific data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-export.html#cfn-bcmdataexports-export-export-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The cadence for AWS to update the export in your S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-export.html#cfn-bcmdataexports-export-export-refreshcadence
	//
	RefreshCadence interface{} `field:"required" json:"refreshCadence" yaml:"refreshCadence"`
	// The description for this specific data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-export.html#cfn-bcmdataexports-export-export-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) for this export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-export.html#cfn-bcmdataexports-export-export-exportarn
	//
	ExportArn *string `field:"optional" json:"exportArn" yaml:"exportArn"`
}

