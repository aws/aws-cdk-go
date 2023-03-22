package awsemr


// `StepConfig` is a property of the `AWS::EMR::Cluster` resource.
//
// The `StepConfig` property type specifies a cluster (job flow) step, which runs only on the master node. Steps are used to submit data processing jobs to the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepConfigProperty := &stepConfigProperty{
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
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	actionOnFailure: jsii.String("actionOnFailure"),
//   }
//
type CfnCluster_StepConfigProperty struct {
	// The JAR file used for the step.
	HadoopJarStep interface{} `field:"required" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// The name of the step.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	ActionOnFailure *string `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
}

