package previewawsec2events


// Type definition for SecurityGroupId.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityGroupId := &SecurityGroupId{
//   	Content: []*string{
//   		jsii.String("content"),
//   	},
//   	Tag: []*string{
//   		jsii.String("tag"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_SecurityGroupId struct {
	// content property.
	//
	// Specify an array of string values to match this event if the actual value of content is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Content *[]*string `field:"optional" json:"content" yaml:"content"`
	// tag property.
	//
	// Specify an array of string values to match this event if the actual value of tag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
}

