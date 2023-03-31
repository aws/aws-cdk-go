package awsroute53resolver


// Properties for defining a `CfnResolverQueryLoggingConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverQueryLoggingConfigProps := &cfnResolverQueryLoggingConfigProps{
//   	destinationArn: jsii.String("destinationArn"),
//   	name: jsii.String("name"),
//   }
//
type CfnResolverQueryLoggingConfigProps struct {
	// The ARN of the resource that you want Resolver to send query logs: an Amazon S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The name of the query logging configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

