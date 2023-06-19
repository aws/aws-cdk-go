package awsrum

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAppMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppMonitorProps := &CfnAppMonitorProps{
//   	Domain: jsii.String("domain"),
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppMonitorProps struct {
	// The top-level internet domain name for which your application has administrative authority.
	//
	// This parameter is required.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// A name for the app monitor.
	//
	// This parameter is required.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A structure that contains much of the configuration data for the app monitor.
	//
	// If you are using Amazon Cognito for authorization, you must include this structure in your request, and it must include the ID of the Amazon Cognito identity pool to use for authorization. If you don't include `AppMonitorConfiguration` , you must set up your own authorization method. For more information, see [Authorize your application to send data to AWS](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-get-started-authorization.html) .
	//
	// If you omit this argument, the sample rate used for CloudWatch RUM is set to 10% of the user sessions.
	AppMonitorConfiguration interface{} `field:"optional" json:"appMonitorConfiguration" yaml:"appMonitorConfiguration"`
	// Specifies whether this app monitor allows the web client to define and send custom events.
	//
	// If you omit this parameter, custom events are `DISABLED` .
	CustomEvents interface{} `field:"optional" json:"customEvents" yaml:"customEvents"`
	// Data collected by CloudWatch RUM is kept by RUM for 30 days and then deleted.
	//
	// This parameter specifies whether CloudWatch RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges.
	//
	// If you omit this parameter, the default is `false` .
	CwLogEnabled interface{} `field:"optional" json:"cwLogEnabled" yaml:"cwLogEnabled"`
	// Assigns one or more tags (key-value pairs) to the app monitor.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with an app monitor.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

