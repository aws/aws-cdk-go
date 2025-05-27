package awsaps


// Query logging configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryLoggingConfigurationProperty := &QueryLoggingConfigurationProperty{
//   	Destinations: []interface{}{
//   		&LoggingDestinationProperty{
//   			CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   			Filters: &LoggingFilterProperty{
//   				QspThreshold: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-queryloggingconfiguration.html
//
type CfnWorkspace_QueryLoggingConfigurationProperty struct {
	// The destinations configuration for query logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-queryloggingconfiguration.html#cfn-aps-workspace-queryloggingconfiguration-destinations
	//
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
}

