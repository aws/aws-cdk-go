package awsrum


// This structure contains much of the configuration data for the app monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMonitorConfigurationProperty := &AppMonitorConfigurationProperty{
//   	AllowCookies: jsii.Boolean(false),
//   	EnableXRay: jsii.Boolean(false),
//   	ExcludedPages: []*string{
//   		jsii.String("excludedPages"),
//   	},
//   	FavoritePages: []*string{
//   		jsii.String("favoritePages"),
//   	},
//   	GuestRoleArn: jsii.String("guestRoleArn"),
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   	IncludedPages: []*string{
//   		jsii.String("includedPages"),
//   	},
//   	MetricDestinations: []interface{}{
//   		&MetricDestinationProperty{
//   			Destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			DestinationArn: jsii.String("destinationArn"),
//   			IamRoleArn: jsii.String("iamRoleArn"),
//   			MetricDefinitions: []interface{}{
//   				&MetricDefinitionProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					DimensionKeys: map[string]*string{
//   						"dimensionKeysKey": jsii.String("dimensionKeys"),
//   					},
//   					EventPattern: jsii.String("eventPattern"),
//   					Namespace: jsii.String("namespace"),
//   					UnitLabel: jsii.String("unitLabel"),
//   					ValueKey: jsii.String("valueKey"),
//   				},
//   			},
//   		},
//   	},
//   	SessionSampleRate: jsii.Number(123),
//   	Telemetries: []*string{
//   		jsii.String("telemetries"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html
//
type CfnAppMonitor_AppMonitorConfigurationProperty struct {
	// If you set this to `true` , the CloudWatch RUM web client sets two cookies, a session cookie and a user cookie.
	//
	// The cookies allow the CloudWatch RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-allowcookies
	//
	AllowCookies interface{} `field:"optional" json:"allowCookies" yaml:"allowCookies"`
	// If you set this to `true` , CloudWatch RUM sends client-side traces to X-Ray for each sampled session.
	//
	// You can then see traces and segments from these user sessions in the RUM dashboard and the CloudWatch ServiceLens console. For more information, see [What is AWS X-Ray ?](https://docs.aws.amazon.com/xray/latest/devguide/aws-xray.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-enablexray
	//
	EnableXRay interface{} `field:"optional" json:"enableXRay" yaml:"enableXRay"`
	// A list of URLs in your website or application to exclude from RUM data collection.
	//
	// You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-excludedpages
	//
	ExcludedPages *[]*string `field:"optional" json:"excludedPages" yaml:"excludedPages"`
	// A list of pages in your application that are to be displayed with a "favorite" icon in the CloudWatch RUM console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-favoritepages
	//
	FavoritePages *[]*string `field:"optional" json:"favoritePages" yaml:"favoritePages"`
	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to CloudWatch RUM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-guestrolearn
	//
	GuestRoleArn *string `field:"optional" json:"guestRoleArn" yaml:"guestRoleArn"`
	// The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to CloudWatch RUM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-identitypoolid
	//
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	//
	// You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-includedpages
	//
	IncludedPages *[]*string `field:"optional" json:"includedPages" yaml:"includedPages"`
	// An array of structures that each define a destination that this app monitor will send extended metrics to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-metricdestinations
	//
	MetricDestinations interface{} `field:"optional" json:"metricDestinations" yaml:"metricDestinations"`
	// Specifies the portion of user sessions to use for CloudWatch RUM data collection.
	//
	// Choosing a higher portion gives you more data but also incurs more costs.
	//
	// The range for this value is 0 to 1 inclusive. Setting this to 1 means that 100% of user sessions are sampled, and setting it to 0.1 means that 10% of user sessions are sampled.
	//
	// If you omit this parameter, the default of 0.1 is used, and 10% of sessions will be sampled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-sessionsamplerate
	//
	SessionSampleRate *float64 `field:"optional" json:"sessionSampleRate" yaml:"sessionSampleRate"`
	// An array that lists the types of telemetry data that this app monitor is to collect.
	//
	// - `errors` indicates that RUM collects data about unhandled JavaScript errors raised by your application.
	// - `performance` indicates that RUM collects performance data about how your application and its resources are loaded and rendered. This includes Core Web Vitals.
	// - `http` indicates that RUM collects data about HTTP errors thrown by your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-appmonitorconfiguration.html#cfn-rum-appmonitor-appmonitorconfiguration-telemetries
	//
	Telemetries *[]*string `field:"optional" json:"telemetries" yaml:"telemetries"`
}

