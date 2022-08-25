package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsDestinationProperty := &snsDestinationProperty{
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnConfigurationSetEventDestination_SnsDestinationProperty struct {
	// `CfnConfigurationSetEventDestination.SnsDestinationProperty.TopicARN`.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

