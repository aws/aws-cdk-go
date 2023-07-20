package awssagemaker


// The ground truth labels for the dataset used for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringGroundTruthS3InputProperty := &MonitoringGroundTruthS3InputProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-monitoringgroundtruths3input.html
//
type CfnModelBiasJobDefinition_MonitoringGroundTruthS3InputProperty struct {
	// The address of the Amazon S3 location of the ground truth labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-monitoringgroundtruths3input.html#cfn-sagemaker-modelbiasjobdefinition-monitoringgroundtruths3input-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

