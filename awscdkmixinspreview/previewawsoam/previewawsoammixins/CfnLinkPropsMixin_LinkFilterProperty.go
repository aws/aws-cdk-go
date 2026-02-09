package previewawsoammixins


// When used in `MetricConfiguration` this field specifies which metric namespaces are to be shared with the monitoring account.
//
// When used in `LogGroupConfiguration` this field specifies which log groups are to share their log events with the monitoring account. Use the term `LogGroupName` and one or more of the following operands.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   linkFilterProperty := &LinkFilterProperty{
//   	Filter: jsii.String("filter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-oam-link-linkfilter.html
//
type CfnLinkPropsMixin_LinkFilterProperty struct {
	// When used in `MetricConfiguration` this field specifies which metric namespaces are to be shared with the monitoring account.
	//
	// When used in `LogGroupConfiguration` this field specifies which log groups are to share their log events with the monitoring account. Use the term `LogGroupName` and one or more of the following operands.
	//
	// Use single quotation marks (') around log group names and metric namespaces.
	//
	// The matching of log group names and metric namespaces is case sensitive. Each filter has a limit of five conditional operands. Conditional operands are `AND` and `OR` .
	//
	// - `=` and `!=`
	// - `AND`
	// - `OR`
	// - `LIKE` and `NOT LIKE` . These can be used only as prefix searches. Include a `%` at the end of the string that you want to search for and include.
	// - `IN` and `NOT IN` , using parentheses `( )`
	//
	// Examples:
	//
	// - `Namespace NOT LIKE 'AWS/%'` includes only namespaces that don't start with `AWS/` , such as custom namespaces.
	// - `Namespace IN ('AWS/EC2', 'AWS/ELB', 'AWS/S3')` includes only the metrics in the EC2, Elastic Load Balancing , and Amazon S3 namespaces.
	// - `Namespace = 'AWS/EC2' OR Namespace NOT LIKE 'AWS/%'` includes only the EC2 namespace and your custom namespaces.
	// - `LogGroupName IN ('This-Log-Group', 'Other-Log-Group')` includes only the log groups with names `This-Log-Group` and `Other-Log-Group` .
	// - `LogGroupName NOT IN ('Private-Log-Group', 'Private-Log-Group-2')` includes all log groups except the log groups with names `Private-Log-Group` and `Private-Log-Group-2` .
	// - `LogGroupName LIKE 'aws/lambda/%' OR LogGroupName LIKE 'AWSLogs%'` includes all log groups that have names that start with `aws/lambda/` or `AWSLogs` .
	//
	// > If you are updating a link that uses filters, you can specify `*` as the only value for the `filter` parameter to delete the filter and share all log groups with the monitoring account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-oam-link-linkfilter.html#cfn-oam-link-linkfilter-filter
	//
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

