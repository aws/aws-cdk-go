package previewawsiotanalyticsmixins


// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
//
// The default file format is JSON. You can specify only one format.
//
// You can't change the file format after you create the data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var jsonConfiguration interface{}
//
//   fileFormatConfigurationProperty := &FileFormatConfigurationProperty{
//   	JsonConfiguration: jsonConfiguration,
//   	ParquetConfiguration: &ParquetConfigurationProperty{
//   		SchemaDefinition: &SchemaDefinitionProperty{
//   			Columns: []interface{}{
//   				&ColumnProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-fileformatconfiguration.html
//
type CfnDatastorePropsMixin_FileFormatConfigurationProperty struct {
	// Contains the configuration information of the JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-fileformatconfiguration.html#cfn-iotanalytics-datastore-fileformatconfiguration-jsonconfiguration
	//
	JsonConfiguration interface{} `field:"optional" json:"jsonConfiguration" yaml:"jsonConfiguration"`
	// Contains the configuration information of the Parquet format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-fileformatconfiguration.html#cfn-iotanalytics-datastore-fileformatconfiguration-parquetconfiguration
	//
	ParquetConfiguration interface{} `field:"optional" json:"parquetConfiguration" yaml:"parquetConfiguration"`
}

