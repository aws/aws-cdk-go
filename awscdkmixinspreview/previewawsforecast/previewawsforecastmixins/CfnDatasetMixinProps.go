package previewawsforecastmixins


// Properties for CfnDatasetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var encryptionConfig interface{}
//   var schema interface{}
//
//   cfnDatasetMixinProps := &CfnDatasetMixinProps{
//   	DataFrequency: jsii.String("dataFrequency"),
//   	DatasetName: jsii.String("datasetName"),
//   	DatasetType: jsii.String("datasetType"),
//   	Domain: jsii.String("domain"),
//   	EncryptionConfig: encryptionConfig,
//   	Schema: schema,
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html
//
type CfnDatasetMixinProps struct {
	// The frequency of data collection. This parameter is required for RELATED_TIME_SERIES datasets.
	//
	// Valid intervals are an integer followed by Y (Year), M (Month), W (Week), D (Day), H (Hour), and min (Minute). For example, "1D" indicates every day and "15min" indicates every 15 minutes. You cannot specify a value that would overlap with the next larger frequency. That means, for example, you cannot specify a frequency of 60 minutes, because that is equivalent to 1 hour. The valid values for each frequency are the following:
	//
	// - Minute - 1-59
	// - Hour - 1-23
	// - Day - 1-6
	// - Week - 1-4
	// - Month - 1-11
	// - Year - 1
	//
	// Thus, if you want every other week forecasts, specify "2W". Or, if you want quarterly forecasts, you specify "3M".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html#cfn-forecast-dataset-datafrequency
	//
	DataFrequency *string `field:"optional" json:"dataFrequency" yaml:"dataFrequency"`
	// The name of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html#cfn-forecast-dataset-datasetname
	//
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// The dataset type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html#cfn-forecast-dataset-datasettype
	//
	DatasetType *string `field:"optional" json:"datasetType" yaml:"datasetType"`
	// The domain associated with the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html#cfn-forecast-dataset-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// A Key Management Service (KMS) key and the Identity and Access Management (IAM) role that Amazon Forecast can assume to access the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html#cfn-forecast-dataset-encryptionconfig
	//
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The schema for the dataset.
	//
	// The schema attributes and their order must match the fields in your data. The dataset `Domain` and `DatasetType` that you choose determine the minimum required fields in your training data. For information about the required fields for a specific dataset domain and type, see [Dataset Domains and Dataset Types](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-domains-ds-types.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html#cfn-forecast-dataset-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-forecast-dataset.html#cfn-forecast-dataset-tags
	//
	Tags *[]*CfnDatasetPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

