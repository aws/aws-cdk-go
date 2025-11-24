package mixinsawsbcmdataexports


// The destinations used for data exports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationConfigurationsProperty := &DestinationConfigurationsProperty{
//   	S3Destination: &S3DestinationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3OutputConfigurations: &S3OutputConfigurationsProperty{
//   			Compression: jsii.String("compression"),
//   			Format: jsii.String("format"),
//   			OutputType: jsii.String("outputType"),
//   			Overwrite: jsii.String("overwrite"),
//   		},
//   		S3Prefix: jsii.String("s3Prefix"),
//   		S3Region: jsii.String("s3Region"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-destinationconfigurations.html
//
type CfnExportPropsMixin_DestinationConfigurationsProperty struct {
	// An object that describes the destination of the data exports file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-destinationconfigurations.html#cfn-bcmdataexports-export-destinationconfigurations-s3destination
	//
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

