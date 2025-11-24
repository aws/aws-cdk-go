package mixinsawslex


// Specifies the S3 bucket location where audio logs are stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioLogDestinationProperty := &AudioLogDestinationProperty{
//   	S3Bucket: &S3BucketLogDestinationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		LogPrefix: jsii.String("logPrefix"),
//   		S3BucketArn: jsii.String("s3BucketArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-audiologdestination.html
//
type CfnBotAliasPropsMixin_AudioLogDestinationProperty struct {
	// The S3 bucket location where audio logs are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-audiologdestination.html#cfn-lex-botalias-audiologdestination-s3bucket
	//
	S3Bucket interface{} `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

