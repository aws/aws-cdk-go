package mixinsawsaps


// The query logging configuration in an Amazon Managed Service for Prometheus workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnWorkspacePropsMixin_QueryLoggingConfigurationProperty struct {
	// Defines a destination and its associated filtering criteria for query logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-queryloggingconfiguration.html#cfn-aps-workspace-queryloggingconfiguration-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

