package awscloudtrail


// The dashboard widget.
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
	// The SQL query statement on one or more event data stores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-widget.html#cfn-cloudtrail-dashboard-widget-querystatement
	//
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// The placeholder keys in the QueryStatement.
	//
	// For example: $StartTime$, $EndTime$, $Period$.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-widget.html#cfn-cloudtrail-dashboard-widget-queryparameters
	//
	QueryParameters *[]*string `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The view properties of the widget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-widget.html#cfn-cloudtrail-dashboard-widget-viewproperties
	//
	ViewProperties interface{} `field:"optional" json:"viewProperties" yaml:"viewProperties"`
}

