package previewawsbatchevents


// Type definition for Container_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var networkInterfaces interface{}
//
//   container1 := &Container1{
//   	ContainerInstanceArn: []*string{
//   		jsii.String("containerInstanceArn"),
//   	},
//   	ExitCode: []*string{
//   		jsii.String("exitCode"),
//   	},
//   	LogStreamName: []*string{
//   		jsii.String("logStreamName"),
//   	},
//   	NetworkInterfaces: []interface{}{
//   		networkInterfaces,
//   	},
//   	TaskArn: []*string{
//   		jsii.String("taskArn"),
//   	},
//   }
//
// Experimental.
type BatchJobStateChange_Container1 struct {
	// containerInstanceArn property.
	//
	// Specify an array of string values to match this event if the actual value of containerInstanceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerInstanceArn *[]*string `field:"optional" json:"containerInstanceArn" yaml:"containerInstanceArn"`
	// exitCode property.
	//
	// Specify an array of string values to match this event if the actual value of exitCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExitCode *[]*string `field:"optional" json:"exitCode" yaml:"exitCode"`
	// logStreamName property.
	//
	// Specify an array of string values to match this event if the actual value of logStreamName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LogStreamName *[]*string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// networkInterfaces property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaces is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaces *[]interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// taskArn property.
	//
	// Specify an array of string values to match this event if the actual value of taskArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskArn *[]*string `field:"optional" json:"taskArn" yaml:"taskArn"`
}

