package awsimagebuilderalpha


// The rendered component data value, for use in CloudFormation.
//
// - For inline components, data is the component text
// - For S3-backed components, uri is the S3 URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   componentDataConfig := &ComponentDataConfig{
//   	Data: jsii.String("data"),
//   	Uri: jsii.String("uri"),
//   }
//
// Experimental.
type ComponentDataConfig struct {
	// The rendered component data, for use in CloudFormation.
	// Default: - none if uri is set.
	//
	// Experimental.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The rendered component data URI, for use in CloudFormation.
	// Default: - none if data is set.
	//
	// Experimental.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

