package awsemr


// Properties for defining a `CfnStep`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStepProps := &CfnStepProps{
//   	ActionOnFailure: jsii.String("actionOnFailure"),
//   	HadoopJarStep: &HadoopJarStepConfigProperty{
//   		Jar: jsii.String("jar"),
//
//   		// the properties below are optional
//   		Args: []*string{
//   			jsii.String("args"),
//   		},
//   		MainClass: jsii.String("mainClass"),
//   		StepProperties: []interface{}{
//   			&KeyValueProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	JobFlowId: jsii.String("jobFlowId"),
//   	Name: jsii.String("name"),
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

