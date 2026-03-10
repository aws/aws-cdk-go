package previewawsxrayevents


// Type definition for AWSXRayInsightUpdateItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var names interface{}
//
//   awsxRayInsightUpdateItem := &AwsxRayInsightUpdateItem{
//   	ServiceId: &ServiceId{
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		Names: []interface{}{
//   			names,
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSXRayInsightUpdate_AwsxRayInsightUpdateItem struct {
	// ServiceId property.
	//
	// Specify an array of string values to match this event if the actual value of ServiceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServiceId *AWSXRayInsightUpdate_ServiceId `field:"optional" json:"serviceId" yaml:"serviceId"`
}

