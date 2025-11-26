package previewawsemrmixins


// `StepConfig` is a property of the `AWS::EMR::Cluster` resource.
//
// The `StepConfig` property type specifies a cluster (job flow) step, which runs only on the master node. Steps are used to submit data processing jobs to the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stepConfigProperty := &StepConfigProperty{
//   	ActionOnFailure: jsii.String("actionOnFailure"),
//   	HadoopJarStep: &HadoopJarStepConfigProperty{
//   		Args: []*string{
//   			jsii.String("args"),
//   		},
//   		Jar: jsii.String("jar"),
//   		MainClass: jsii.String("mainClass"),
//   		StepProperties: []interface{}{
//   			&KeyValueProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html
//
type CfnClusterPropsMixin_StepConfigProperty struct {
	// The action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html#cfn-emr-cluster-stepconfig-actiononfailure
	//
	ActionOnFailure *string `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
	// The JAR file used for the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html#cfn-emr-cluster-stepconfig-hadoopjarstep
	//
	HadoopJarStep interface{} `field:"optional" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// The name of the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-stepconfig.html#cfn-emr-cluster-stepconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

