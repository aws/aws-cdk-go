package awsappflow


// The configuration that determines how Amazon AppFlow formats the flow output data when Upsolver is used as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   upsolverS3OutputFormatConfigProperty := &UpsolverS3OutputFormatConfigProperty{
//   	PrefixConfig: &PrefixConfigProperty{
//   		PathPrefixHierarchy: []*string{
//   			jsii.String("pathPrefixHierarchy"),
//   		},
//   		PrefixFormat: jsii.String("prefixFormat"),
//   		PrefixType: jsii.String("prefixType"),
//   	},
//
//   	// the properties below are optional
//   	AggregationConfig: &AggregationConfigProperty{
//   		AggregationType: jsii.String("aggregationType"),
//   		TargetFileSize: jsii.Number(123),
//   	},
//   	FileType: jsii.String("fileType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolvers3outputformatconfig.html
//
type CfnFlow_UpsolverS3OutputFormatConfigProperty struct {
	// Specifies elements that Amazon AppFlow includes in the file and folder names in the flow destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolvers3outputformatconfig.html#cfn-appflow-flow-upsolvers3outputformatconfig-prefixconfig
	//
	PrefixConfig interface{} `field:"required" json:"prefixConfig" yaml:"prefixConfig"`
	// The aggregation settings that you can use to customize the output format of your flow data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolvers3outputformatconfig.html#cfn-appflow-flow-upsolvers3outputformatconfig-aggregationconfig
	//
	AggregationConfig interface{} `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Indicates the file type that Amazon AppFlow places in the Upsolver Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolvers3outputformatconfig.html#cfn-appflow-flow-upsolvers3outputformatconfig-filetype
	//
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

