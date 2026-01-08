package previewawscodebuildevents


// Type definition for Network-interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterface := &NetworkInterface{
//   	EniId: []*string{
//   		jsii.String("eniId"),
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildStateChange_NetworkInterface struct {
	// eni-id property.
	//
	// Specify an array of string values to match this event if the actual value of eni-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EniId *[]*string `field:"optional" json:"eniId" yaml:"eniId"`
	// subnet-id property.
	//
	// Specify an array of string values to match this event if the actual value of subnet-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

