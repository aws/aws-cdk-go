package mixinsawssagemaker


// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringOutputProperty := &MonitoringOutputProperty{
//   	S3Output: &S3OutputProperty{
//   		LocalPath: jsii.String("localPath"),
//   		S3UploadMode: jsii.String("s3UploadMode"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-monitoringoutput.html
//
type CfnDataQualityJobDefinitionPropsMixin_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-monitoringoutput.html#cfn-sagemaker-dataqualityjobdefinition-monitoringoutput-s3output
	//
	S3Output interface{} `field:"optional" json:"s3Output" yaml:"s3Output"`
}

