package awscloudtrail


// Contains information about a widget on a CloudTrail Lake dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   widgetProperty := &WidgetProperty{
//   	QueryStatement: jsii.String("queryStatement"),
//
//   	// the properties below are optional
//   	QueryParameters: []*string{
//   		jsii.String("queryParameters"),
//   	},
//   	ViewProperties: map[string]*string{
//   		"viewPropertiesKey": jsii.String("viewProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-widget.html
//
type CfnDashboard_WidgetProperty struct {
	// The query statement for the widget.
	//
	// For custom dashboard widgets, you can query across multiple event data stores as long as all event data stores exist in your account.
	//
	// > When a query uses `?` with `eventTime` , `?` must be surrounded by single quotes as follows: `'?'` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-widget.html#cfn-cloudtrail-dashboard-widget-querystatement
	//
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// The optional query parameters.
	//
	// The following query parameters are valid: `$StartTime$` , `$EndTime$` , and `$Period$` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-widget.html#cfn-cloudtrail-dashboard-widget-queryparameters
	//
	QueryParameters *[]*string `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The view properties for the widget.
	//
	// For more information about view properties, see [View properties for widgets](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-widget-properties.html) in the *AWS CloudTrail User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-widget.html#cfn-cloudtrail-dashboard-widget-viewproperties
	//
	ViewProperties interface{} `field:"optional" json:"viewProperties" yaml:"viewProperties"`
}

