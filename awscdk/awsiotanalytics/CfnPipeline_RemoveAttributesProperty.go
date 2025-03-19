package awsiotanalytics


// An activity that removes attributes from a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   removeAttributesProperty := &RemoveAttributesProperty{
//   	Attributes: []*string{
//   		jsii.String("attributes"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-removeattributes.html
//
type CfnPipeline_RemoveAttributesProperty struct {
	// A list of 1-50 attributes to remove from the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-removeattributes.html#cfn-iotanalytics-pipeline-removeattributes-attributes
	//
	Attributes *[]*string `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the 'removeAttributes' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-removeattributes.html#cfn-iotanalytics-pipeline-removeattributes-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-removeattributes.html#cfn-iotanalytics-pipeline-removeattributes-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}

