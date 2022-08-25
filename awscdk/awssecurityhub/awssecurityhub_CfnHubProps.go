package awssecurityhub


// Properties for defining a `CfnHub`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnHubProps := &cfnHubProps{
//   	tags: tags,
//   }
//
type CfnHubProps struct {
	// The tags to add to the hub resource.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

