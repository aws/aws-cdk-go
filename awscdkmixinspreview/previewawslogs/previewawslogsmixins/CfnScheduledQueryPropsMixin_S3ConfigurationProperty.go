package previewawslogsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3ConfigurationProperty := &S3ConfigurationProperty{
//   	DestinationIdentifier: jsii.String("destinationIdentifier"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-scheduledquery-s3configuration.html
//
type CfnScheduledQueryPropsMixin_S3ConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-scheduledquery-s3configuration.html#cfn-logs-scheduledquery-s3configuration-destinationidentifier
	//
	DestinationIdentifier *string `field:"optional" json:"destinationIdentifier" yaml:"destinationIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-scheduledquery-s3configuration.html#cfn-logs-scheduledquery-s3configuration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

