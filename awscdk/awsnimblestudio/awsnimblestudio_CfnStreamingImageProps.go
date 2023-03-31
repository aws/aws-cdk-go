package awsnimblestudio


// Properties for defining a `CfnStreamingImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamingImageProps := &cfnStreamingImageProps{
//   	ec2ImageId: jsii.String("ec2ImageId"),
//   	name: jsii.String("name"),
//   	studioId: jsii.String("studioId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnStreamingImageProps struct {
	// The ID of an EC2 machine image with which to create the streaming image.
	Ec2ImageId *string `field:"required" json:"ec2ImageId" yaml:"ec2ImageId"`
	// A friendly name for a streaming image resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// A human-readable description of the streaming image.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

