package previewawsecsevents


// Type definition for ResponseElementsItemItem_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var networkInterfaces interface{}
//
//   responseElementsItemItem1 := &ResponseElementsItemItem1{
//   	ContainerArn: []*string{
//   		jsii.String("containerArn"),
//   	},
//   	Cpu: []*string{
//   		jsii.String("cpu"),
//   	},
//   	Image: []*string{
//   		jsii.String("image"),
//   	},
//   	LastStatus: []*string{
//   		jsii.String("lastStatus"),
//   	},
//   	Memory: []*string{
//   		jsii.String("memory"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
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
type ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItemItem1 struct {
	// containerArn property.
	//
	// Specify an array of string values to match this event if the actual value of containerArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerArn *[]*string `field:"optional" json:"containerArn" yaml:"containerArn"`
	// cpu property.
	//
	// Specify an array of string values to match this event if the actual value of cpu is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cpu *[]*string `field:"optional" json:"cpu" yaml:"cpu"`
	// image property.
	//
	// Specify an array of string values to match this event if the actual value of image is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Image *[]*string `field:"optional" json:"image" yaml:"image"`
	// lastStatus property.
	//
	// Specify an array of string values to match this event if the actual value of lastStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastStatus *[]*string `field:"optional" json:"lastStatus" yaml:"lastStatus"`
	// memory property.
	//
	// Specify an array of string values to match this event if the actual value of memory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Memory *[]*string `field:"optional" json:"memory" yaml:"memory"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
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

