package awsmwaaserverless


// Contains the environment variables that tasks in a workflow use during execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentProperty := &EnvironmentProperty{
//   	Variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-environment.html
//
type CfnWorkflow_EnvironmentProperty struct {
	// A map of environment variable names to values.
	//
	// Up to 50 variables. Names must start with a letter and contain only letters, numbers, and underscores. Names starting with AWS_ or AIRFLOW__ are reserved, except AIRFLOW__WORKERS__SECRETS_BACKEND, AIRFLOW__WORKERS__SECRETS_BACKEND_KWARGS, and names prefixed with AIRFLOW__WORKERS__SECRETS_BACKEND_KWARG__.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-environment.html#cfn-mwaaserverless-workflow-environment-variables
	//
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
}

