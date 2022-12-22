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
//   cfnAppMonitorProps := &cfnAppMonitorProps{
//   	domain: jsii.String("domain"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	appMonitorConfiguration: &appMonitorConfigurationProperty{
//   		allowCookies: jsii.Boolean(false),
//   		enableXRay: jsii.Boolean(false),
//   		excludedPages: []*string{
//   			jsii.String("excludedPages"),
//   		},
//   		favoritePages: []*string{
//   			jsii.String("favoritePages"),
//   		},
//   		guestRoleArn: jsii.String("guestRoleArn"),
//   		identityPoolId: jsii.String("identityPoolId"),
//   		includedPages: []*string{
//   			jsii.String("includedPages"),
//   		},
//   		metricDestinations: []interface{}{
//   			&metricDestinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				destinationArn: jsii.String("destinationArn"),
//   				iamRoleArn: jsii.String("iamRoleArn"),
//   				metricDefinitions: []interface{}{
//   					&metricDefinitionProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						dimensionKeys: map[string]*string{
//   							"dimensionKeysKey": jsii.String("dimensionKeys"),
//   						},
//   						eventPattern: jsii.String("eventPattern"),
//   						unitLabel: jsii.String("unitLabel"),
//   						valueKey: jsii.String("valueKey"),
//   					},
//   				},
//   			},
//   		},
//   		sessionSampleRate: jsii.Number(123),
//   		telemetries: []*string{
//   			jsii.String("telemetries"),
//   		},
//   	},
//   	cwLogEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

