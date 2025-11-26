package previewawsecsevents


// Type definition for AwsvpcConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   awsvpcConfiguration := &AwsvpcConfiguration{
//   	AssignPublicIp: []*string{
//   		jsii.String("assignPublicIp"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail_AwsvpcConfiguration struct {
	// assignPublicIp property.
	//
	// Specify an array of string values to match this event if the actual value of assignPublicIp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AssignPublicIp *[]*string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// subnets property.
	//
	// Specify an array of string values to match this event if the actual value of subnets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

