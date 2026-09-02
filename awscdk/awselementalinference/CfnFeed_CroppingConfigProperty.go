package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   croppingConfigProperty := &CroppingConfigProperty{
//   	TemplateGroups: []interface{}{
//   		&TemplateGroupProperty{
//   			Name: jsii.String("name"),
//   			TemplateUris: []*string{
//   				jsii.String("templateUris"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-croppingconfig.html
//
type CfnFeed_CroppingConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-croppingconfig.html#cfn-elementalinference-feed-croppingconfig-templategroups
	//
	TemplateGroups interface{} `field:"optional" json:"templateGroups" yaml:"templateGroups"`
}

