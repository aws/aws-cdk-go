package awsappflow


// The configuration that determines how Amazon AppFlow formats the flow output data when Upsolver is used as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   upsolverS3OutputFormatConfigProperty := &upsolverS3OutputFormatConfigProperty{
//   	prefixConfig: &prefixConfigProperty{
//   		prefixFormat: jsii.String("prefixFormat"),
//   		prefixType: jsii.String("prefixType"),
//   	},
//
//   	// the properties below are optional
//   	aggregationConfig: &aggregationConfigProperty{
//   		aggregationType: jsii.String("aggregationType"),
//   	},
//   	fileType: jsii.String("fileType"),
//   }
//
type CfnFlow_UpsolverS3OutputFormatConfigProperty struct {
	// Specifies elements that Amazon AppFlow includes in the file and folder names in the flow destination.
	PrefixConfig interface{} `field:"required" json:"prefixConfig" yaml:"prefixConfig"`
	// The aggregation settings that you can use to customize the output format of your flow data.
	AggregationConfig interface{} `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Indicates the file type that Amazon AppFlow places in the Upsolver Amazon S3 bucket.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

