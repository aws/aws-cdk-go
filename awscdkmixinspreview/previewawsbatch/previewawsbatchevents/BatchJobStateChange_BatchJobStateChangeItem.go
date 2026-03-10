package previewawsbatchevents


// Type definition for BatchJobStateChangeItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var networkInterfaces interface{}
//
//   batchJobStateChangeItem := &BatchJobStateChangeItem{
//   	Container: &Container1{
//   		ContainerInstanceArn: []*string{
//   			jsii.String("containerInstanceArn"),
//   		},
//   		ExitCode: []*string{
//   			jsii.String("exitCode"),
//   		},
//   		LogStreamName: []*string{
//   			jsii.String("logStreamName"),
//   		},
//   		NetworkInterfaces: []interface{}{
//   			networkInterfaces,
//   		},
//   		TaskArn: []*string{
//   			jsii.String("taskArn"),
//   		},
//   	},
//   	StartedAt: []*string{
//   		jsii.String("startedAt"),
//   	},
//   	StatusReason: []*string{
//   		jsii.String("statusReason"),
//   	},
//   	StoppedAt: []*string{
//   		jsii.String("stoppedAt"),
//   	},
//   }
//
// Experimental.
type BatchJobStateChange_BatchJobStateChangeItem struct {
	// container property.
	//
	// Specify an array of string values to match this event if the actual value of container is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Container *BatchJobStateChange_Container1 `field:"optional" json:"container" yaml:"container"`
	// startedAt property.
	//
	// Specify an array of string values to match this event if the actual value of startedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedAt *[]*string `field:"optional" json:"startedAt" yaml:"startedAt"`
	// statusReason property.
	//
	// Specify an array of string values to match this event if the actual value of statusReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusReason *[]*string `field:"optional" json:"statusReason" yaml:"statusReason"`
	// stoppedAt property.
	//
	// Specify an array of string values to match this event if the actual value of stoppedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppedAt *[]*string `field:"optional" json:"stoppedAt" yaml:"stoppedAt"`
}

