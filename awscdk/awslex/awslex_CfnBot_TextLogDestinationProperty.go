package awslex


// Specifies the Amazon CloudWatch Logs destination log group for conversation text logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogDestinationProperty := &textLogDestinationProperty{
//   	cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		logPrefix: jsii.String("logPrefix"),
//   	},
//   }
//
type CfnBot_TextLogDestinationProperty struct {
	// Specifies the Amazon CloudWatch Logs log group where text and metadata logs are delivered.
	CloudWatch interface{} `field:"required" json:"cloudWatch" yaml:"cloudWatch"`
}

