package previewawsec2events


// Type definition for InstancesSet_1Item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instancesSet1Item := &InstancesSet1Item{
//   	ImageId: []*string{
//   		jsii.String("imageId"),
//   	},
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   	MaxCount: []*string{
//   		jsii.String("maxCount"),
//   	},
//   	MinCount: []*string{
//   		jsii.String("minCount"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_InstancesSet1Item struct {
	// imageId property.
	//
	// Specify an array of string values to match this event if the actual value of imageId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageId *[]*string `field:"optional" json:"imageId" yaml:"imageId"`
	// instanceId property.
	//
	// Specify an array of string values to match this event if the actual value of instanceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// maxCount property.
	//
	// Specify an array of string values to match this event if the actual value of maxCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxCount *[]*string `field:"optional" json:"maxCount" yaml:"maxCount"`
	// minCount property.
	//
	// Specify an array of string values to match this event if the actual value of minCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MinCount *[]*string `field:"optional" json:"minCount" yaml:"minCount"`
}

