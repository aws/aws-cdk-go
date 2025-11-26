package previewawsapptestmixins


// Defines a step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stepProperty := &StepProperty{
//   	Action: &StepActionProperty{
//   		CompareAction: &CompareActionProperty{
//   			Input: &InputProperty{
//   				File: &InputFileProperty{
//   					FileMetadata: &FileMetadataProperty{
//   						DatabaseCdc: &DatabaseCDCProperty{
//   							SourceMetadata: &SourceDatabaseMetadataProperty{
//   								CaptureTool: jsii.String("captureTool"),
//   								Type: jsii.String("type"),
//   							},
//   							TargetMetadata: &TargetDatabaseMetadataProperty{
//   								CaptureTool: jsii.String("captureTool"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						DataSets: []interface{}{
//   							&DataSetProperty{
//   								Ccsid: jsii.String("ccsid"),
//   								Format: jsii.String("format"),
//   								Length: jsii.Number(123),
//   								Name: jsii.String("name"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					SourceLocation: jsii.String("sourceLocation"),
//   					TargetLocation: jsii.String("targetLocation"),
//   				},
//   			},
//   			Output: &OutputProperty{
//   				File: &OutputFileProperty{
//   					FileLocation: jsii.String("fileLocation"),
//   				},
//   			},
//   		},
//   		MainframeAction: &MainframeActionProperty{
//   			ActionType: &MainframeActionTypeProperty{
//   				Batch: &BatchProperty{
//   					BatchJobName: jsii.String("batchJobName"),
//   					BatchJobParameters: map[string]*string{
//   						"batchJobParametersKey": jsii.String("batchJobParameters"),
//   					},
//   					ExportDataSetNames: []*string{
//   						jsii.String("exportDataSetNames"),
//   					},
//   				},
//   				Tn3270: &TN3270Property{
//   					ExportDataSetNames: []*string{
//   						jsii.String("exportDataSetNames"),
//   					},
//   					Script: &ScriptProperty{
//   						ScriptLocation: jsii.String("scriptLocation"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			Properties: &MainframeActionPropertiesProperty{
//   				DmsTaskArn: jsii.String("dmsTaskArn"),
//   			},
//   			Resource: jsii.String("resource"),
//   		},
//   		ResourceAction: &ResourceActionProperty{
//   			CloudFormationAction: &CloudFormationActionProperty{
//   				ActionType: jsii.String("actionType"),
//   				Resource: jsii.String("resource"),
//   			},
//   			M2ManagedApplicationAction: &M2ManagedApplicationActionProperty{
//   				ActionType: jsii.String("actionType"),
//   				Properties: &M2ManagedActionPropertiesProperty{
//   					ForceStop: jsii.Boolean(false),
//   					ImportDataSetLocation: jsii.String("importDataSetLocation"),
//   				},
//   				Resource: jsii.String("resource"),
//   			},
//   			M2NonManagedApplicationAction: &M2NonManagedApplicationActionProperty{
//   				ActionType: jsii.String("actionType"),
//   				Resource: jsii.String("resource"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-step.html
//
type CfnTestCasePropsMixin_StepProperty struct {
	// The action of the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-step.html#cfn-apptest-testcase-step-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// The description of the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-step.html#cfn-apptest-testcase-step-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-step.html#cfn-apptest-testcase-step-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

