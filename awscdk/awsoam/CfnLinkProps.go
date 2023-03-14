package awsoam


// Properties for defining a `CfnLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLinkProps := &CfnLinkProps{
//   	LabelTemplate: jsii.String("labelTemplate"),
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	SinkIdentifier: jsii.String("sinkIdentifier"),
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnLinkProps struct {
	// Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
	//
	// You can include the following variables in your template:
	//
	// - `$AccountName` is the name of the account
	// - `$AccountEmail` is a globally-unique email address, which includes the email domain, such as `mariagarcia@example.com`
	// - `$AccountEmailNoDomain` is an email address without the domain name, such as `mariagarcia`.
	LabelTemplate *string `field:"required" json:"labelTemplate" yaml:"labelTemplate"`
	// An array of strings that define which types of data that the source account shares with the monitoring account.
	//
	// Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace` .
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// The ARN of the sink in the monitoring account that you want to link to.
	//
	// You can use [ListSinks](https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html) to find the ARNs of sinks.
	SinkIdentifier *string `field:"required" json:"sinkIdentifier" yaml:"sinkIdentifier"`
	// An array of key-value pairs to apply to the link.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

