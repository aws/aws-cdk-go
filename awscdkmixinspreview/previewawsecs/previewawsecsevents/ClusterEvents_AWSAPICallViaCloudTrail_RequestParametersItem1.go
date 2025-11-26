package previewawsecsevents


// Type definition for RequestParametersItem_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var networkBindings interface{}
//
//   requestParametersItem1 := &RequestParametersItem1{
//   	ContainerName: []*string{
//   		jsii.String("containerName"),
//   	},
//   	ExitCode: []*string{
//   		jsii.String("exitCode"),
//   	},
//   	NetworkBindings: []interface{}{
//   		networkBindings,
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail_RequestParametersItem1 struct {
	// containerName property.
	//
	// Specify an array of string values to match this event if the actual value of containerName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerName *[]*string `field:"optional" json:"containerName" yaml:"containerName"`
	// exitCode property.
	//
	// Specify an array of string values to match this event if the actual value of exitCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExitCode *[]*string `field:"optional" json:"exitCode" yaml:"exitCode"`
	// networkBindings property.
	//
	// Specify an array of string values to match this event if the actual value of networkBindings is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkBindings *[]interface{} `field:"optional" json:"networkBindings" yaml:"networkBindings"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

