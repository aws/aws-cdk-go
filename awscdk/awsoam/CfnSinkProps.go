package awsoam


// Properties for defining a `CfnSink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnSinkProps := &CfnSinkProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Policy: policy,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSinkProps struct {
	// A name for the sink.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IAM policy that grants permissions to source accounts to link to this sink.
	//
	// The policy can grant permission in the following ways:
	//
	// - Include organization IDs or organization paths to permit all accounts in an organization
	// - Include account IDs to permit the specified accounts.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// An array of key-value pairs to apply to the sink.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

