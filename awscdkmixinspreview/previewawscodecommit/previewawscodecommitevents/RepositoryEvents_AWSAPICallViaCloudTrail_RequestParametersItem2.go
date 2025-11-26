package previewawscodecommitevents


// Type definition for RequestParametersItem_2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem2 := &RequestParametersItem2{
//   	DestinationReference: []*string{
//   		jsii.String("destinationReference"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem2 struct {
	// destinationReference property.
	//
	// Specify an array of string values to match this event if the actual value of destinationReference is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationReference *[]*string `field:"optional" json:"destinationReference" yaml:"destinationReference"`
}

