package awsiotanalytics


// Information needed to run the "containerAction" to produce data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	actionName: jsii.String("actionName"),
//
//   	// the properties below are optional
//   	containerAction: &containerActionProperty{
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		image: jsii.String("image"),
//   		resourceConfiguration: &resourceConfigurationProperty{
//   			computeType: jsii.String("computeType"),
//   			volumeSizeInGb: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		variables: []interface{}{
//   			&variableProperty{
//   				variableName: jsii.String("variableName"),
//
//   				// the properties below are optional
//   				datasetContentVersionValue: &datasetContentVersionValueProperty{
//   					datasetName: jsii.String("datasetName"),
//   				},
//   				doubleValue: jsii.Number(123),
//   				outputFileUriValue: &outputFileUriValueProperty{
//   					fileName: jsii.String("fileName"),
//   				},
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	queryAction: &queryActionProperty{
//   		sqlQuery: jsii.String("sqlQuery"),
//
//   		// the properties below are optional
//   		filters: []interface{}{
//   			&filterProperty{
//   				deltaTime: &deltaTimeProperty{
//   					offsetSeconds: jsii.Number(123),
//   					timeExpression: jsii.String("timeExpression"),
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

