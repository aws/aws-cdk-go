package previewawsautoscalingevents


// Type definition for RequestParametersItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem := &RequestParametersItem{
//   	Key: []*string{
//   		jsii.String("key"),
//   	},
//   	PropagateAtLaunch: []*string{
//   		jsii.String("propagateAtLaunch"),
//   	},
//   	ResourceId: []*string{
//   		jsii.String("resourceId"),
//   	},
//   	ResourceType: []*string{
//   		jsii.String("resourceType"),
//   	},
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParametersItem struct {
	// key property.
	//
	// Specify an array of string values to match this event if the actual value of key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
	// propagateAtLaunch property.
	//
	// Specify an array of string values to match this event if the actual value of propagateAtLaunch is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PropagateAtLaunch *[]*string `field:"optional" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// resourceId property.
	//
	// Specify an array of string values to match this event if the actual value of resourceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceId *[]*string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// resourceType property.
	//
	// Specify an array of string values to match this event if the actual value of resourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceType *[]*string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// value property.
	//
	// Specify an array of string values to match this event if the actual value of value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

