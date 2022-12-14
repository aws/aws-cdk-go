package awsoam


// Properties for defining a `CfnLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLinkProps := &cfnLinkProps{
//   	labelTemplate: jsii.String("labelTemplate"),
//   	resourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	sinkIdentifier: jsii.String("sinkIdentifier"),
//
//   	// the properties below are optional
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnLinkProps struct {
	// `AWS::Oam::Link.LabelTemplate`.
	LabelTemplate *string `field:"required" json:"labelTemplate" yaml:"labelTemplate"`
	// `AWS::Oam::Link.ResourceTypes`.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// `AWS::Oam::Link.SinkIdentifier`.
	SinkIdentifier *string `field:"required" json:"sinkIdentifier" yaml:"sinkIdentifier"`
	// `AWS::Oam::Link.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

