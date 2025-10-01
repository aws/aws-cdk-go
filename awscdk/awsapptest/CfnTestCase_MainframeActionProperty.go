package awsapptest


// Specifies the mainframe action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mainframeActionProperty := &MainframeActionProperty{
//   	ActionType: &MainframeActionTypeProperty{
//   		Batch: &BatchProperty{
//   			BatchJobName: jsii.String("batchJobName"),
//
//   			// the properties below are optional
//   			BatchJobParameters: map[string]*string{
//   				"batchJobParametersKey": jsii.String("batchJobParameters"),
//   			},
//   			ExportDataSetNames: []*string{
//   				jsii.String("exportDataSetNames"),
//   			},
//   		},
//   		Tn3270: &TN3270Property{
//   			Script: &ScriptProperty{
//   				ScriptLocation: jsii.String("scriptLocation"),
//   				Type: jsii.String("type"),
//   			},
//
//   			// the properties below are optional
//   			ExportDataSetNames: []*string{
//   				jsii.String("exportDataSetNames"),
//   			},
//   		},
//   	},
//   	Resource: jsii.String("resource"),
//
//   	// the properties below are optional
//   	Properties: &MainframeActionPropertiesProperty{
//   		DmsTaskArn: jsii.String("dmsTaskArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html
//
type CfnTestCase_MainframeActionProperty struct {
	// The action type of the mainframe action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html#cfn-apptest-testcase-mainframeaction-actiontype
	//
	ActionType interface{} `field:"required" json:"actionType" yaml:"actionType"`
	// The resource of the mainframe action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html#cfn-apptest-testcase-mainframeaction-resource
	//
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The properties of the mainframe action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-mainframeaction.html#cfn-apptest-testcase-mainframeaction-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

