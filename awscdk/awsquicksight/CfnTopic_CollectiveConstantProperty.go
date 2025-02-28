package awsquicksight


// A structure that represents a collective constant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectiveConstantProperty := &CollectiveConstantProperty{
//   	ValueList: []*string{
//   		jsii.String("valueList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-collectiveconstant.html
//
type CfnTopic_CollectiveConstantProperty struct {
	// A list of values for the collective constant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-collectiveconstant.html#cfn-quicksight-topic-collectiveconstant-valuelist
	//
	ValueList *[]*string `field:"optional" json:"valueList" yaml:"valueList"`
}

