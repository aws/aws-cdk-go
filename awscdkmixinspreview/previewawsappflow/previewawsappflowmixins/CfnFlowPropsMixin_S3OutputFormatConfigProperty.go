package previewawsappflowmixins


// The configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3OutputFormatConfigProperty := &S3OutputFormatConfigProperty{
//   	AggregationConfig: &AggregationConfigProperty{
//   		AggregationType: jsii.String("aggregationType"),
//   		TargetFileSize: jsii.Number(123),
//   	},
//   	FileType: jsii.String("fileType"),
//   	PrefixConfig: &PrefixConfigProperty{
//   		PathPrefixHierarchy: []*string{
//   			jsii.String("pathPrefixHierarchy"),
//   		},
//   		PrefixFormat: jsii.String("prefixFormat"),
//   		PrefixType: jsii.String("prefixType"),
//   	},
//   	PreserveSourceDataTyping: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html
//
type CfnFlowPropsMixin_S3OutputFormatConfigProperty struct {
	// The aggregation settings that you can use to customize the output format of your flow data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html#cfn-appflow-flow-s3outputformatconfig-aggregationconfig
	//
	AggregationConfig interface{} `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Indicates the file type that Amazon AppFlow places in the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html#cfn-appflow-flow-s3outputformatconfig-filetype
	//
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
	// Determines the prefix that Amazon AppFlow applies to the folder name in the Amazon S3 bucket.
	//
	// You can name folders according to the flow frequency and date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html#cfn-appflow-flow-s3outputformatconfig-prefixconfig
	//
	PrefixConfig interface{} `field:"optional" json:"prefixConfig" yaml:"prefixConfig"`
	// If your file output format is Parquet, use this parameter to set whether Amazon AppFlow preserves the data types in your source data when it writes the output to Amazon S3.
	//
	// - `true` : Amazon AppFlow preserves the data types when it writes to Amazon S3. For example, an integer or `1` in your source data is still an integer in your output.
	// - `false` : Amazon AppFlow converts all of the source data into strings when it writes to Amazon S3. For example, an integer of `1` in your source data becomes the string `"1"` in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html#cfn-appflow-flow-s3outputformatconfig-preservesourcedatatyping
	//
	PreserveSourceDataTyping interface{} `field:"optional" json:"preserveSourceDataTyping" yaml:"preserveSourceDataTyping"`
}

