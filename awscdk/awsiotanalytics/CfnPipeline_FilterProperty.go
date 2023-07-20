package awsiotanalytics


// An activity that filters a message based on its attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Filter: jsii.String("filter"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-filter.html
//
type CfnPipeline_FilterProperty struct {
	// An expression that looks like an SQL WHERE clause that must return a Boolean value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-filter.html#cfn-iotanalytics-pipeline-filter-filter
	//
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The name of the 'filter' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-filter.html#cfn-iotanalytics-pipeline-filter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-filter.html#cfn-iotanalytics-pipeline-filter-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}

