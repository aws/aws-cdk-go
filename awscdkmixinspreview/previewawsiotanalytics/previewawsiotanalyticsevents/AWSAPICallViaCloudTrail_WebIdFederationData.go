package previewawsiotanalyticsevents


// Type definition for WebIdFederationData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webIdFederationData := &WebIdFederationData{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	FederatedProvider: []*string{
//   		jsii.String("federatedProvider"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_WebIdFederationData struct {
	// attributes property.
	//
	// Specify an array of string values to match this event if the actual value of attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// federatedProvider property.
	//
	// Specify an array of string values to match this event if the actual value of federatedProvider is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FederatedProvider *[]*string `field:"optional" json:"federatedProvider" yaml:"federatedProvider"`
}

