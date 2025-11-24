package mixinsawsapptest


// Specifies the mainframe action type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mainframeActionTypeProperty := &MainframeActionTypeProperty{
//   	Batch: &BatchProperty{
//   		BatchJobName: jsii.String("batchJobName"),
//   		BatchJobParameters: map[string]*string{
//   			"batchJobParametersKey": jsii.String("batchJobParameters"),
//   		},
//   		ExportDataSetNames: []*string{
//   			jsii.String("exportDataSetNames"),
//   		},
//   	},
//   	Tn3270: &TN3270Property{
//   		ExportDataSetNames: []*string{
//   			jsii.String("exportDataSetNames"),
//   		},
//   		Script: &ScriptProperty{
//   			ScriptLocation: jsii.String("scriptLocation"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactiontype.html
//
type CfnTestCasePropsMixin_MainframeActionTypeProperty struct {
	// The batch of the mainframe action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactiontype.html#cfn-apptest-testcase-mainframeactiontype-batch
	//
	Batch interface{} `field:"optional" json:"batch" yaml:"batch"`
	// The tn3270 port of the mainframe action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeactiontype.html#cfn-apptest-testcase-mainframeactiontype-tn3270
	//
	Tn3270 interface{} `field:"optional" json:"tn3270" yaml:"tn3270"`
}

