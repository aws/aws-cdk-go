package awsrum

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppMonitorProps := &CfnAppMonitorProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AppMonitorConfiguration: &AppMonitorConfigurationProperty{
//   		AllowCookies: jsii.Boolean(false),
//   		EnableXRay: jsii.Boolean(false),
//   		ExcludedPages: []*string{
//   			jsii.String("excludedPages"),
//   		},
//   		FavoritePages: []*string{
//   			jsii.String("favoritePages"),
//   		},
//   		GuestRoleArn: jsii.String("guestRoleArn"),
//   		IdentityPoolId: jsii.String("identityPoolId"),
//   		IncludedPages: []*string{
//   			jsii.String("includedPages"),
//   		},
//   		MetricDestinations: []interface{}{
//   			&MetricDestinationProperty{
//   				Destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				DestinationArn: jsii.String("destinationArn"),
//   				IamRoleArn: jsii.String("iamRoleArn"),
//   				MetricDefinitions: []interface{}{
//   					&MetricDefinitionProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						DimensionKeys: map[string]*string{
//   							"dimensionKeysKey": jsii.String("dimensionKeys"),
//   						},
//   						EventPattern: jsii.String("eventPattern"),
//   						Namespace: jsii.String("namespace"),
//   						UnitLabel: jsii.String("unitLabel"),
//   						ValueKey: jsii.String("valueKey"),
//   					},
//   				},
//   			},
//   		},
//   		SessionSampleRate: jsii.Number(123),
//   		Telemetries: []*string{
//   			jsii.String("telemetries"),
//   		},
//   	},
//   	CustomEvents: &CustomEventsProperty{
//   		Status: jsii.String("status"),
//   	},
//   	CwLogEnabled: jsii.Boolean(false),
//   	DeobfuscationConfiguration: &DeobfuscationConfigurationProperty{
//   		JavaScriptSourceMaps: &JavaScriptSourceMapsProperty{
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	Domain: jsii.String("domain"),
//   	DomainList: []*string{
//   		jsii.String("domainList"),
//   	},
//   	ResourcePolicy: &ResourcePolicyProperty{
//   		PolicyDocument: jsii.String("policyDocument"),
//
//   		// the properties below are optional
//   		PolicyRevisionId: jsii.String("policyRevisionId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html
//
type CfnAppMonitorProps struct {
	// A name for the app monitor.
	//
	// This parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A structure that contains much of the configuration data for the app monitor.
	//
	// If you are using Amazon Cognito for authorization, you must include this structure in your request, and it must include the ID of the Amazon Cognito identity pool to use for authorization. If you don't include `AppMonitorConfiguration` , you must set up your own authorization method. For more information, see [Authorize your application to send data to AWS](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-get-started-authorization.html) .
	//
	// If you omit this argument, the sample rate used for CloudWatch RUM is set to 10% of the user sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-appmonitorconfiguration
	//
	AppMonitorConfiguration interface{} `field:"optional" json:"appMonitorConfiguration" yaml:"appMonitorConfiguration"`
	// Specifies whether this app monitor allows the web client to define and send custom events.
	//
	// If you omit this parameter, custom events are `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-customevents
	//
	CustomEvents interface{} `field:"optional" json:"customEvents" yaml:"customEvents"`
	// Data collected by CloudWatch RUM is kept by RUM for 30 days and then deleted.
	//
	// This parameter specifies whether CloudWatch RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges.
	//
	// If you omit this parameter, the default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-cwlogenabled
	//
	CwLogEnabled interface{} `field:"optional" json:"cwLogEnabled" yaml:"cwLogEnabled"`
	// A structure that contains the configuration for how an app monitor can deobfuscate stack traces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-deobfuscationconfiguration
	//
	DeobfuscationConfiguration interface{} `field:"optional" json:"deobfuscationConfiguration" yaml:"deobfuscationConfiguration"`
	// The top-level internet domain name for which your application has administrative authority.
	//
	// This parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The top-level internet domain names for which your application has administrative authority.
	//
	// The CreateAppMonitor requires either the domain or the domain list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-domainlist
	//
	DomainList *[]*string `field:"optional" json:"domainList" yaml:"domainList"`
	// Use this structure to assign a resource-based policy to a CloudWatch RUM app monitor to control access to it.
	//
	// Each app monitor can have one resource-based policy. The maximum size of the policy is 4 KB. To learn more about using resource policies with RUM, see [Using resource-based policies with CloudWatch RUM](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-resource-policies.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-resourcepolicy
	//
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Assigns one or more tags (key-value pairs) to the app monitor.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with an app monitor.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html#cfn-rum-appmonitor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

