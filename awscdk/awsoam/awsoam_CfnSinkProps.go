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
//   cfnSinkProps := &cfnSinkProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	policy: policy,
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSinkProps struct {
	// `AWS::Oam::Sink.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Oam::Sink.Policy`.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// `AWS::Oam::Sink.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

