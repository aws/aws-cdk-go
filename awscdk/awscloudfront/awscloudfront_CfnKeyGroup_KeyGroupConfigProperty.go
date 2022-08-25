package awscloudfront


// A key group configuration.
//
// A key group contains a list of public keys that you can use with [CloudFront signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyGroupConfigProperty := &keyGroupConfigProperty{
//   	items: []*string{
//   		jsii.String("items"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   }
//
type CfnKeyGroup_KeyGroupConfigProperty struct {
	// A list of the identifiers of the public keys in the key group.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
	// A name to identify the key group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the key group.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

