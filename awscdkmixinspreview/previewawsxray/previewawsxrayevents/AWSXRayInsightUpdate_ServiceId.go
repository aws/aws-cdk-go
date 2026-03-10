package previewawsxrayevents


// Type definition for ServiceId.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var names interface{}
//
//   serviceId := &ServiceId{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Names: []interface{}{
//   		names,
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type AWSXRayInsightUpdate_ServiceId struct {
	// AccountId property.
	//
	// Specify an array of string values to match this event if the actual value of AccountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// Name property.
	//
	// Specify an array of string values to match this event if the actual value of Name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// Names property.
	//
	// Specify an array of string values to match this event if the actual value of Names is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Names *[]interface{} `field:"optional" json:"names" yaml:"names"`
	// Type property.
	//
	// Specify an array of string values to match this event if the actual value of Type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

