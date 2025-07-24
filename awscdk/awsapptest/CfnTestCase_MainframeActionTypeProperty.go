package awsapptest


// Specifies the mainframe action type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mainframeActionTypeProperty := &MainframeActionTypeProperty{
//   	Batch: &BatchProperty{
//   		BatchJobName: jsii.String("batchJobName"),
//
//   		// the properties below are optional
//   		BatchJobParameters: map[string]*string{
//   			"batchJobParametersKey": jsii.String("batchJobParameters"),
//   		},
//   		ExportDataSetNames: []*string{
//   			jsii.String("exportDataSetNames"),
//   		},
//   	},
//   	Tn3270: &TN3270Property{
//   		Script: &ScriptProperty{
//   			ScriptLocation: jsii.String("scriptLocation"),
//   			Type: jsii.String("type"),
//   		},
//
//   		// the properties below are optional
//   		ExportDataSetNames: []*string{
//   			jsii.String("exportDataSetNames"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactiontype.html
//
type CfnTestCase_MainframeActionTypeProperty struct {
	// The batch of the mainframe action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactiontype.html#cfn-apptest-testcase-mainframeactiontype-batch
	//
	Batch interface{} `field:"optional" json:"batch" yaml:"batch"`
	// The tn3270 port of the mainframe action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactiontype.html#cfn-apptest-testcase-mainframeactiontype-tn3270
	//
	Tn3270 interface{} `field:"optional" json:"tn3270" yaml:"tn3270"`
}

