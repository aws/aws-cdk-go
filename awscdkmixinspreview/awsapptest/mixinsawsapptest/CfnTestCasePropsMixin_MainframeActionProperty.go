package mixinsawsapptest


// Specifies the mainframe action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mainframeActionProperty := &MainframeActionProperty{
//   	ActionType: &MainframeActionTypeProperty{
//   		Batch: &BatchProperty{
//   			BatchJobName: jsii.String("batchJobName"),
//   			BatchJobParameters: map[string]*string{
//   				"batchJobParametersKey": jsii.String("batchJobParameters"),
//   			},
//   			ExportDataSetNames: []*string{
//   				jsii.String("exportDataSetNames"),
//   			},
//   		},
//   		Tn3270: &TN3270Property{
//   			ExportDataSetNames: []*string{
//   				jsii.String("exportDataSetNames"),
//   			},
//   			Script: &ScriptProperty{
//   				ScriptLocation: jsii.String("scriptLocation"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Properties: &MainframeActionPropertiesProperty{
//   		DmsTaskArn: jsii.String("dmsTaskArn"),
//   	},
//   	Resource: jsii.String("resource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html
//
type CfnTestCasePropsMixin_MainframeActionProperty struct {
	// The action type of the mainframe action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html#cfn-apptest-testcase-mainframeaction-actiontype
	//
	ActionType interface{} `field:"optional" json:"actionType" yaml:"actionType"`
	// The properties of the mainframe action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html#cfn-apptest-testcase-mainframeaction-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The resource of the mainframe action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html#cfn-apptest-testcase-mainframeaction-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

