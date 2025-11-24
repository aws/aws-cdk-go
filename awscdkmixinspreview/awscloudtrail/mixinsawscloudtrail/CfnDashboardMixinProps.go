package mixinsawscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDashboardPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDashboardMixinProps := &CfnDashboardMixinProps{
//   	Name: jsii.String("name"),
//   	RefreshSchedule: &RefreshScheduleProperty{
//   		Frequency: &FrequencyProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		Status: jsii.String("status"),
//   		TimeOfDay: jsii.String("timeOfDay"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminationProtectionEnabled: jsii.Boolean(false),
//   	Widgets: []interface{}{
//   		&WidgetProperty{
//   			QueryParameters: []*string{
//   				jsii.String("queryParameters"),
//   			},
//   			QueryStatement: jsii.String("queryStatement"),
//   			ViewProperties: map[string]*string{
//   				"viewPropertiesKey": jsii.String("viewProperties"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html
//
type CfnDashboardMixinProps struct {
	// The name of the dashboard. The name must be unique to your account.
	//
	// To create the Highlights dashboard, the name must be `AWSCloudTrail-Highlights` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schedule for a dashboard refresh.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-refreshschedule
	//
	RefreshSchedule interface{} `field:"optional" json:"refreshSchedule" yaml:"refreshSchedule"`
	// A list of tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether termination protection is enabled for the dashboard.
	//
	// If termination protection is enabled, you cannot delete the dashboard until termination protection is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-terminationprotectionenabled
	//
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// An array of widgets for a custom dashboard. A custom dashboard can have a maximum of ten widgets.
	//
	// You do not need to specify widgets for the Highlights dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html#cfn-cloudtrail-dashboard-widgets
	//
	Widgets interface{} `field:"optional" json:"widgets" yaml:"widgets"`
}

