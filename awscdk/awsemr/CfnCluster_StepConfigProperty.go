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
//   stepConfigProperty := &StepConfigProperty{
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
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ActionOnFailure: jsii.String("actionOnFailure"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html
//
type CfnCluster_StepConfigProperty struct {
	// The JAR file used for the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html#cfn-emr-cluster-stepconfig-hadoopjarstep
	//
	HadoopJarStep interface{} `field:"required" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// The name of the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html#cfn-emr-cluster-stepconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html#cfn-emr-cluster-stepconfig-actiononfailure
	//
	ActionOnFailure *string `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
}

