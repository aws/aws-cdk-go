package awsresourceexplorer2


// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &cfnIndexProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnIndexProps struct {
	// `AWS::ResourceExplorer2::Index.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `AWS::ResourceExplorer2::Index.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

