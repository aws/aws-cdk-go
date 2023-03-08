package awsemr


// Properties for defining a `CfnStep`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStepProps := &cfnStepProps{
//   	actionOnFailure: jsii.String("actionOnFailure"),
//   	hadoopJarStep: &hadoopJarStepConfigProperty{
//   		jar: jsii.String("jar"),
//
//   		// the properties below are optional
//   		args: []*string{
//   			jsii.String("args"),
//   		},
//   		mainClass: jsii.String("mainClass"),
//   		stepProperties: []interface{}{
//   			&keyValueProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	jobFlowId: jsii.String("jobFlowId"),
//   	name: jsii.String("name"),
//   }
//
type CfnStepProps struct {
	// This specifies what action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	ActionOnFailure *string `field:"required" json:"actionOnFailure" yaml:"actionOnFailure"`
	// The `HadoopJarStepConfig` property type specifies a job flow step consisting of a JAR file whose main function will be executed.
	//
	// The main function submits a job for the cluster to execute as a step on the master node, and then waits for the job to finish or fail before executing subsequent steps.
	HadoopJarStep interface{} `field:"required" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// A string that uniquely identifies the cluster (job flow).
	JobFlowId *string `field:"required" json:"jobFlowId" yaml:"jobFlowId"`
	// The name of the cluster step.
	Name *string `field:"required" json:"name" yaml:"name"`
}

