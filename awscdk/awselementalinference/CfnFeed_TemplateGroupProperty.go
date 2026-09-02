package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateGroupProperty := &TemplateGroupProperty{
//   	Name: jsii.String("name"),
//   	TemplateUris: []*string{
//   		jsii.String("templateUris"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-templategroup.html
//
type CfnFeed_TemplateGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-templategroup.html#cfn-elementalinference-feed-templategroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-templategroup.html#cfn-elementalinference-feed-templategroup-templateuris
	//
	TemplateUris *[]*string `field:"required" json:"templateUris" yaml:"templateUris"`
}

