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
type CfnTopic_CollectiveConstantProperty struct {
	// A list of values for the collective constant.
	ValueList *[]*string `field:"optional" json:"valueList" yaml:"valueList"`
}

