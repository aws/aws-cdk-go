package previewawsiotanalyticsmixins


// Information needed to run the "containerAction" to produce data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	ActionName: jsii.String("actionName"),
//   	ContainerAction: &ContainerActionProperty{
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		Image: jsii.String("image"),
//   		ResourceConfiguration: &ResourceConfigurationProperty{
//   			ComputeType: jsii.String("computeType"),
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   		Variables: []interface{}{
//   			&VariableProperty{
//   				DatasetContentVersionValue: &DatasetContentVersionValueProperty{
//   					DatasetName: jsii.String("datasetName"),
//   				},
//   				DoubleValue: jsii.Number(123),
//   				OutputFileUriValue: &OutputFileUriValueProperty{
//   					FileName: jsii.String("fileName"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   				VariableName: jsii.String("variableName"),
//   			},
//   		},
//   	},
//   	QueryAction: &QueryActionProperty{
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				DeltaTime: &DeltaTimeProperty{
//   					OffsetSeconds: jsii.Number(123),
//   					TimeExpression: jsii.String("timeExpression"),
//   				},
//   			},
//   		},
//   		SqlQuery: jsii.String("sqlQuery"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html
//
type CfnDatasetPropsMixin_ActionProperty struct {
	// The name of the data set action by which data set contents are automatically created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html#cfn-iotanalytics-dataset-action-actionname
	//
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// Information which allows the system to run a containerized application in order to create the data set contents.
	//
	// The application must be in a Docker container along with any needed support libraries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html#cfn-iotanalytics-dataset-action-containeraction
	//
	ContainerAction interface{} `field:"optional" json:"containerAction" yaml:"containerAction"`
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html#cfn-iotanalytics-dataset-action-queryaction
	//
	QueryAction interface{} `field:"optional" json:"queryAction" yaml:"queryAction"`
}

