package awsapptest


// Specifies the TN3270 protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tN3270Property := &TN3270Property{
//   	Script: &ScriptProperty{
//   		ScriptLocation: jsii.String("scriptLocation"),
//   		Type: jsii.String("type"),
//   	},
//
//   	// the properties below are optional
//   	ExportDataSetNames: []*string{
//   		jsii.String("exportDataSetNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-tn3270.html
//
type CfnTestCase_TN3270Property struct {
	// The script of the TN3270 protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-tn3270.html#cfn-apptest-testcase-tn3270-script
	//
	Script interface{} `field:"required" json:"script" yaml:"script"`
	// The data set names of the TN3270 protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-tn3270.html#cfn-apptest-testcase-tn3270-exportdatasetnames
	//
	ExportDataSetNames *[]*string `field:"optional" json:"exportDataSetNames" yaml:"exportDataSetNames"`
}

