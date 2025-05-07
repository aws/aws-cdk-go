package awsoam


// Properties for defining a `CfnLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLinkProps := &CfnLinkProps{
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	SinkIdentifier: jsii.String("sinkIdentifier"),
//
//   	// the properties below are optional
//   	LabelTemplate: jsii.String("labelTemplate"),
//   	LinkConfiguration: &LinkConfigurationProperty{
//   		LogGroupConfiguration: &LinkFilterProperty{
//   			Filter: jsii.String("filter"),
//   		},
//   		MetricConfiguration: &LinkFilterProperty{
//   			Filter: jsii.String("filter"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-link.html
//
type CfnLinkProps struct {
	// An array of strings that define which types of data that the source account shares with the monitoring account.
	//
	// Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace | AWS::ApplicationInsights::Application | AWS::InternetMonitor::Monitor` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-link.html#cfn-oam-link-resourcetypes
	//
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// The ARN of the sink in the monitoring account that you want to link to.
	//
	// You can use [ListSinks](https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html) to find the ARNs of sinks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-link.html#cfn-oam-link-sinkidentifier
	//
	SinkIdentifier *string `field:"required" json:"sinkIdentifier" yaml:"sinkIdentifier"`
	// Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
	//
	// You can include the following variables in your template:
	//
	// - `$AccountName` is the name of the account
	// - `$AccountEmail` is a globally-unique email address, which includes the email domain, such as `mariagarcia@example.com`
	// - `$AccountEmailNoDomain` is an email address without the domain name, such as `mariagarcia`
	//
	// > In the  and  Regions, the only supported option is to use custom labels, and the `$AccountName` , `$AccountEmail` , and `$AccountEmailNoDomain` variables all resolve as *account-id* instead of the specified variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-link.html#cfn-oam-link-labeltemplate
	//
	LabelTemplate *string `field:"optional" json:"labelTemplate" yaml:"labelTemplate"`
	// Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-link.html#cfn-oam-link-linkconfiguration
	//
	LinkConfiguration interface{} `field:"optional" json:"linkConfiguration" yaml:"linkConfiguration"`
	// An array of key-value pairs to apply to the link.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-link.html#cfn-oam-link-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

