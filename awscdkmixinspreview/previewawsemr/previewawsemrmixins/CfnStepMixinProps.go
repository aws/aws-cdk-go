package previewawsemrmixins


// Properties for CfnStepPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStepMixinProps := &CfnStepMixinProps{
//   	ActionOnFailure: jsii.String("actionOnFailure"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
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
//   	JobFlowId: jsii.String("jobFlowId"),
//   	LogUri: jsii.String("logUri"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html
//
type CfnStepMixinProps struct {
	// This specifies what action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-emr-step-actiononfailure
	//
	ActionOnFailure *string `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
	// The KMS key ARN to encrypt the logs published to the given Amazon S3 destination.
	//
	// When omitted, EMR falls back to cluster-level logging behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-emr-step-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The `HadoopJarStepConfig` property type specifies a job flow step consisting of a JAR file whose main function will be executed.
	//
	// The main function submits a job for the cluster to execute as a step on the master node, and then waits for the job to finish or fail before executing subsequent steps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-emr-step-hadoopjarstep
	//
	HadoopJarStep interface{} `field:"optional" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// A string that uniquely identifies the cluster (job flow).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-emr-step-jobflowid
	//
	JobFlowId *string `field:"optional" json:"jobFlowId" yaml:"jobFlowId"`
	// The Amazon S3 destination URI for log publishing.
	//
	// When omitted, EMR falls back to cluster-level logging behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-emr-step-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// The name of the cluster step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-emr-step-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

