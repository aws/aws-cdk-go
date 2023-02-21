package awsiotanalytics


// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
//
// The default file format is JSON. You can specify only one format.
//
// You can't change the file format after you create the data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonConfiguration interface{}
//
//   fileFormatConfigurationProperty := &fileFormatConfigurationProperty{
//   	jsonConfiguration: jsonConfiguration,
//   	parquetConfiguration: &parquetConfigurationProperty{
//   		schemaDefinition: &schemaDefinitionProperty{
//   			columns: []interface{}{
//   				&columnProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDatastore_FileFormatConfigurationProperty struct {
	// Contains the configuration information of the JSON format.
	JsonConfiguration interface{} `field:"optional" json:"jsonConfiguration" yaml:"jsonConfiguration"`
	// Contains the configuration information of the Parquet format.
	ParquetConfiguration interface{} `field:"optional" json:"parquetConfiguration" yaml:"parquetConfiguration"`
}

