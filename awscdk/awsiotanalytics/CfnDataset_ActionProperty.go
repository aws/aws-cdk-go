package awsiotanalytics


// Information needed to run the "containerAction" to produce data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	ActionName: jsii.String("actionName"),
//
//   	// the properties below are optional
//   	ContainerAction: &ContainerActionProperty{
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		Image: jsii.String("image"),
//   		ResourceConfiguration: &ResourceConfigurationProperty{
//   			ComputeType: jsii.String("computeType"),
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		Variables: []interface{}{
//   			&VariableProperty{
//   				VariableName: jsii.String("variableName"),
//
//   				// the properties below are optional
//   				DatasetContentVersionValue: &DatasetContentVersionValueProperty{
//   					DatasetName: jsii.String("datasetName"),
//   				},
//   				DoubleValue: jsii.Number(123),
//   				OutputFileUriValue: &OutputFileUriValueProperty{
//   					FileName: jsii.String("fileName"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	QueryAction: &QueryActionProperty{
//   		SqlQuery: jsii.String("sqlQuery"),
//
//   		// the properties below are optional
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				DeltaTime: &DeltaTimeProperty{
//   					OffsetSeconds: jsii.Number(123),
//   					TimeExpression: jsii.String("timeExpression"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataset_ActionProperty struct {
	// The name of the data set action by which data set contents are automatically created.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Information which allows the system to run a containerized application in order to create the data set contents.
	//
	// The application must be in a Docker container along with any needed support libraries.
	ContainerAction interface{} `field:"optional" json:"containerAction" yaml:"containerAction"`
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	QueryAction interface{} `field:"optional" json:"queryAction" yaml:"queryAction"`
}

