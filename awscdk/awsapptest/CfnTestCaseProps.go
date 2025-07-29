package awsapptest


// Properties for defining a `CfnTestCase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTestCaseProps := &CfnTestCaseProps{
//   	Name: jsii.String("name"),
//   	Steps: []interface{}{
//   		&StepProperty{
//   			Action: &StepActionProperty{
//   				CompareAction: &CompareActionProperty{
//   					Input: &InputProperty{
//   						File: &InputFileProperty{
//   							FileMetadata: &FileMetadataProperty{
//   								DatabaseCdc: &DatabaseCDCProperty{
//   									SourceMetadata: &SourceDatabaseMetadataProperty{
//   										CaptureTool: jsii.String("captureTool"),
//   										Type: jsii.String("type"),
//   									},
//   									TargetMetadata: &TargetDatabaseMetadataProperty{
//   										CaptureTool: jsii.String("captureTool"),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   								DataSets: []interface{}{
//   									&DataSetProperty{
//   										Ccsid: jsii.String("ccsid"),
//   										Format: jsii.String("format"),
//   										Length: jsii.Number(123),
//   										Name: jsii.String("name"),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   							SourceLocation: jsii.String("sourceLocation"),
//   							TargetLocation: jsii.String("targetLocation"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Output: &OutputProperty{
//   						File: &OutputFileProperty{
//   							FileLocation: jsii.String("fileLocation"),
//   						},
//   					},
//   				},
//   				MainframeAction: &MainframeActionProperty{
//   					ActionType: &MainframeActionTypeProperty{
//   						Batch: &BatchProperty{
//   							BatchJobName: jsii.String("batchJobName"),
//
//   							// the properties below are optional
//   							BatchJobParameters: map[string]*string{
//   								"batchJobParametersKey": jsii.String("batchJobParameters"),
//   							},
//   							ExportDataSetNames: []*string{
//   								jsii.String("exportDataSetNames"),
//   							},
//   						},
//   						Tn3270: &TN3270Property{
//   							Script: &ScriptProperty{
//   								ScriptLocation: jsii.String("scriptLocation"),
//   								Type: jsii.String("type"),
//   							},
//
//   							// the properties below are optional
//   							ExportDataSetNames: []*string{
//   								jsii.String("exportDataSetNames"),
//   							},
//   						},
//   					},
//   					Resource: jsii.String("resource"),
//
//   					// the properties below are optional
//   					Properties: &MainframeActionPropertiesProperty{
//   						DmsTaskArn: jsii.String("dmsTaskArn"),
//   					},
//   				},
//   				ResourceAction: &ResourceActionProperty{
//   					CloudFormationAction: &CloudFormationActionProperty{
//   						Resource: jsii.String("resource"),
//
//   						// the properties below are optional
//   						ActionType: jsii.String("actionType"),
//   					},
//   					M2ManagedApplicationAction: &M2ManagedApplicationActionProperty{
//   						ActionType: jsii.String("actionType"),
//   						Resource: jsii.String("resource"),
//
//   						// the properties below are optional
//   						Properties: &M2ManagedActionPropertiesProperty{
//   							ForceStop: jsii.Boolean(false),
//   							ImportDataSetLocation: jsii.String("importDataSetLocation"),
//   						},
//   					},
//   					M2NonManagedApplicationAction: &M2NonManagedApplicationActionProperty{
//   						ActionType: jsii.String("actionType"),
//   						Resource: jsii.String("resource"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html
//
type CfnTestCaseProps struct {
	// The name of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The steps in the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-steps
	//
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// The description of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The specified tags of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

