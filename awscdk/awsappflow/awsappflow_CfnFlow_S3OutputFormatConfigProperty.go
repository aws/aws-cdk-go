package awsappflow


// The configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputFormatConfigProperty := &s3OutputFormatConfigProperty{
//   	aggregationConfig: &aggregationConfigProperty{
//   		aggregationType: jsii.String("aggregationType"),
//   		targetFileSize: jsii.Number(123),
//   	},
//   	fileType: jsii.String("fileType"),
//   	prefixConfig: &prefixConfigProperty{
//   		pathPrefixHierarchy: []*string{
//   			jsii.String("pathPrefixHierarchy"),
//   		},
//   		prefixFormat: jsii.String("prefixFormat"),
//   		prefixType: jsii.String("prefixType"),
//   	},
//   	preserveSourceDataTyping: jsii.Boolean(false),
//   }
//
type CfnFlow_S3OutputFormatConfigProperty struct {
	// The aggregation settings that you can use to customize the output format of your flow data.
	AggregationConfig interface{} `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Indicates the file type that Amazon AppFlow places in the Amazon S3 bucket.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
	// Determines the prefix that Amazon AppFlow applies to the folder name in the Amazon S3 bucket.
	//
	// You can name folders according to the flow frequency and date.
	PrefixConfig interface{} `field:"optional" json:"prefixConfig" yaml:"prefixConfig"`
	// `CfnFlow.S3OutputFormatConfigProperty.PreserveSourceDataTyping`.
	PreserveSourceDataTyping interface{} `field:"optional" json:"preserveSourceDataTyping" yaml:"preserveSourceDataTyping"`
}

