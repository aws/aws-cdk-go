package awssagemaker


// The inputs for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringInputProperty := &monitoringInputProperty{
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringInputProperty struct {
	// The endpoint for a monitoring job.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
}

