package awsinternetmonitor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitorProps := &CfnMonitorProps{
//   	MonitorName: jsii.String("monitorName"),
//
//   	// the properties below are optional
//   	HealthEventsConfig: &HealthEventsConfigProperty{
//   		AvailabilityLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   			HealthScoreThreshold: jsii.Number(123),
//   			MinTrafficImpact: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		AvailabilityScoreThreshold: jsii.Number(123),
//   		PerformanceLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   			HealthScoreThreshold: jsii.Number(123),
//   			MinTrafficImpact: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		PerformanceScoreThreshold: jsii.Number(123),
//   	},
//   	IncludeLinkedAccounts: jsii.Boolean(false),
//   	InternetMeasurementsLogDelivery: &InternetMeasurementsLogDeliveryProperty{
//   		S3Config: &S3ConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			LogDeliveryStatus: jsii.String("logDeliveryStatus"),
//   		},
//   	},
//   	LinkedAccountId: jsii.String("linkedAccountId"),
//   	MaxCityNetworksToMonitor: jsii.Number(123),
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	ResourcesToAdd: []*string{
//   		jsii.String("resourcesToAdd"),
//   	},
//   	ResourcesToRemove: []*string{
//   		jsii.String("resourcesToRemove"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficPercentageToMonitor: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html
//
type CfnMonitorProps struct {
	// The name of the monitor.
	//
	// A monitor name can contain only alphanumeric characters, dashes (-), periods (.), and underscores (_).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-monitorname
	//
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// A complex type with the configuration information that determines the threshold and other conditions for when Internet Monitor creates a health event for an overall performance or availability issue, across an application's geographies.
	//
	// Defines the percentages, for overall performance scores and availability scores for an application, that are the thresholds for when Internet Monitor creates a health event. You can override the defaults to set a custom threshold for overall performance or availability scores, or both.
	//
	// You can also set thresholds for local health scores,, where Internet Monitor creates a health event when scores cross a threshold for one or more city-networks, in addition to creating an event when an overall score crosses a threshold.
	//
	// If you don't set a health event threshold, the default value is 95%.
	//
	// For local thresholds, you also set a minimum percentage of overall traffic that is impacted by an issue before Internet Monitor creates an event. In addition, you can disable local thresholds, for performance scores, availability scores, or both.
	//
	// For more information, see [Change health event thresholds](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-overview.html#IMUpdateThresholdFromOverview) in the Internet Monitor section of the *CloudWatch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-healtheventsconfig
	//
	HealthEventsConfig interface{} `field:"optional" json:"healthEventsConfig" yaml:"healthEventsConfig"`
	// A boolean option that you can set to `TRUE` to include monitors for linked accounts in a list of monitors, when you've set up cross-account sharing in Internet Monitor.
	//
	// You configure cross-account sharing by using Amazon CloudWatch Observability Access Manager. For more information, see [Internet Monitor cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cwim-cross-account.html) in the Amazon CloudWatch User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-includelinkedaccounts
	//
	IncludeLinkedAccounts interface{} `field:"optional" json:"includeLinkedAccounts" yaml:"includeLinkedAccounts"`
	// Publish internet measurements for a monitor for all city-networks (up to the 500,000 service limit) to another location, such as an Amazon S3 bucket.
	//
	// Measurements are also published to Amazon CloudWatch Logs for the first 500 (by traffic volume) city-networks (client locations and ASNs, typically internet service providers or ISPs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-internetmeasurementslogdelivery
	//
	InternetMeasurementsLogDelivery interface{} `field:"optional" json:"internetMeasurementsLogDelivery" yaml:"internetMeasurementsLogDelivery"`
	// The account ID for an account that you've set up cross-account sharing for in Internet Monitor.
	//
	// You configure cross-account sharing by using Amazon CloudWatch Observability Access Manager. For more information, see [Internet Monitor cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cwim-cross-account.html) in the Amazon CloudWatch User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-linkedaccountid
	//
	LinkedAccountId *string `field:"optional" json:"linkedAccountId" yaml:"linkedAccountId"`
	// The maximum number of city-networks to monitor for your resources.
	//
	// A city-network is the location (city) where clients access your application resources from and the network, such as an internet service provider, that clients access the resources through.
	//
	// For more information, see [Choosing a city-network maximum value](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html) in *Using Amazon CloudWatch Internet Monitor* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-maxcitynetworkstomonitor
	//
	MaxCityNetworksToMonitor *float64 `field:"optional" json:"maxCityNetworksToMonitor" yaml:"maxCityNetworksToMonitor"`
	// The resources that have been added for the monitor, listed by their Amazon Resource Names (ARNs).
	//
	// Use this option to add or remove resources when making an update.
	//
	// > Be aware that if you include content in the `Resources` field when you update a monitor, the `ResourcesToAdd` and `ResourcesToRemove` fields must be empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The resources to include in a monitor, which you provide as a set of Amazon Resource Names (ARNs).
	//
	// Resources can be Amazon Virtual Private Cloud VPCs, Network Load Balancers (NLBs), Amazon CloudFront distributions, or Amazon WorkSpaces directories.
	//
	// You can add a combination of VPCs and CloudFront distributions, or you can add WorkSpaces directories, or you can add NLBs. You can't add NLBs or WorkSpaces directories together with any other resources.
	//
	// If you add only VPC resources, at least one VPC must have an Internet Gateway attached to it, to make sure that it has internet connectivity.
	//
	// > You can specify this field for a monitor update only if the `Resources` field is empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-resourcestoadd
	//
	ResourcesToAdd *[]*string `field:"optional" json:"resourcesToAdd" yaml:"resourcesToAdd"`
	// The resources to remove from a monitor, which you provide as a set of Amazon Resource Names (ARNs).
	//
	// > You can specify this field for a monitor update only if the `Resources` field is empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-resourcestoremove
	//
	ResourcesToRemove *[]*string `field:"optional" json:"resourcesToRemove" yaml:"resourcesToRemove"`
	// The status of a monitor.
	//
	// The accepted values that you can specify for `Status` are `ACTIVE` and `INACTIVE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags for a monitor, listed as a set of *key:value* pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The percentage of the internet-facing traffic for your application that you want to monitor.
	//
	// You can also, optionally, set a limit for the number of city-networks (client locations and ASNs, typically internet service providers) that Internet Monitor will monitor traffic for. The city-networks maximum limit caps the number of city-networks that Internet Monitor monitors for your application, regardless of the percentage of traffic that you choose to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html#cfn-internetmonitor-monitor-trafficpercentagetomonitor
	//
	TrafficPercentageToMonitor *float64 `field:"optional" json:"trafficPercentageToMonitor" yaml:"trafficPercentageToMonitor"`
}

