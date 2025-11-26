package previewawsrobomakermixins


// Information about a source configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceConfigProperty := &SourceConfigProperty{
//   	Architecture: jsii.String("architecture"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html
//
type CfnSimulationApplicationPropsMixin_SourceConfigProperty struct {
	// The target processor architecture for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html#cfn-robomaker-simulationapplication-sourceconfig-architecture
	//
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// The Amazon S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html#cfn-robomaker-simulationapplication-sourceconfig-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The s3 object key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html#cfn-robomaker-simulationapplication-sourceconfig-s3key
	//
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
}

