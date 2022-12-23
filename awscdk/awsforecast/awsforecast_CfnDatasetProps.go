package awsforecast


// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryptionConfig interface{}
//   var schema interface{}
//
//   cfnDatasetProps := &cfnDatasetProps{
//   	datasetName: jsii.String("datasetName"),
//   	datasetType: jsii.String("datasetType"),
//   	domain: jsii.String("domain"),
//   	schema: schema,
//
//   	// the properties below are optional
//   	dataFrequency: jsii.String("dataFrequency"),
//   	encryptionConfig: encryptionConfig,
//   	tags: []tagsItemsProperty{
//   		&tagsItemsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatasetProps struct {
	// The name of the dataset.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// The dataset type.
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The domain associated with the dataset.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The schema for the dataset.
	//
	// The schema attributes and their order must match the fields in your data. The dataset `Domain` and `DatasetType` that you choose determine the minimum required fields in your training data. For information about the required fields for a specific dataset domain and type, see [Dataset Domains and Dataset Types](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-domains-ds-types.html) .
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// The frequency of data collection. This parameter is required for RELATED_TIME_SERIES datasets.
	//
	// Valid intervals are Y (Year), M (Month), W (Week), D (Day), H (Hour), 30min (30 minutes), 15min (15 minutes), 10min (10 minutes), 5min (5 minutes), and 1min (1 minute). For example, "D" indicates every day and "15min" indicates every 15 minutes.
	DataFrequency *string `field:"optional" json:"dataFrequency" yaml:"dataFrequency"`
	// A Key Management Service (KMS) key and the Identity and Access Management (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*CfnDataset_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

