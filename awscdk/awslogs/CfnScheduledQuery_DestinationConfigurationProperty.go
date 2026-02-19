package awslogs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigurationProperty := &DestinationConfigurationProperty{
//   	S3Configuration: &S3ConfigurationProperty{
//   		DestinationIdentifier: jsii.String("destinationIdentifier"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-scheduledquery-destinationconfiguration.html
//
type CfnScheduledQuery_DestinationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-scheduledquery-destinationconfiguration.html#cfn-logs-scheduledquery-destinationconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

