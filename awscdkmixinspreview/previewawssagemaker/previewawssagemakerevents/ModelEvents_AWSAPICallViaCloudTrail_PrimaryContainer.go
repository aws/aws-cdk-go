package previewawssagemakerevents


// Type definition for PrimaryContainer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   primaryContainer := &PrimaryContainer{
//   	ContainerHostname: []*string{
//   		jsii.String("containerHostname"),
//   	},
//   	Image: []*string{
//   		jsii.String("image"),
//   	},
//   	ModelDataUrl: []*string{
//   		jsii.String("modelDataUrl"),
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_PrimaryContainer struct {
	// containerHostname property.
	//
	// Specify an array of string values to match this event if the actual value of containerHostname is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerHostname *[]*string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// image property.
	//
	// Specify an array of string values to match this event if the actual value of image is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Image *[]*string `field:"optional" json:"image" yaml:"image"`
	// modelDataUrl property.
	//
	// Specify an array of string values to match this event if the actual value of modelDataUrl is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelDataUrl *[]*string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
}

