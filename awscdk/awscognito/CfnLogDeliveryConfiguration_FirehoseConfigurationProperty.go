package awscognito


// Configuration for the Amazon Data Firehose stream destination of user activity log export with advanced security features.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseConfigurationProperty := &FirehoseConfigurationProperty{
//   	StreamArn: jsii.String("streamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-firehoseconfiguration.html
//
type CfnLogDeliveryConfiguration_FirehoseConfigurationProperty struct {
	// The ARN of an Amazon Data Firehose stream that's the destination for advanced security features log export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-firehoseconfiguration.html#cfn-cognito-logdeliveryconfiguration-firehoseconfiguration-streamarn
	//
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

