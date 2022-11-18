package awsrum


// This structure contains much of the configuration data for the app monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMonitorConfigurationProperty := &appMonitorConfigurationProperty{
//   	allowCookies: jsii.Boolean(false),
//   	enableXRay: jsii.Boolean(false),
//   	excludedPages: []*string{
//   		jsii.String("excludedPages"),
//   	},
//   	favoritePages: []*string{
//   		jsii.String("favoritePages"),
//   	},
//   	guestRoleArn: jsii.String("guestRoleArn"),
//   	identityPoolId: jsii.String("identityPoolId"),
//   	includedPages: []*string{
//   		jsii.String("includedPages"),
//   	},
//   	metricDestinations: []interface{}{
//   		&metricDestinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			destinationArn: jsii.String("destinationArn"),
//   			iamRoleArn: jsii.String("iamRoleArn"),
//   			metricDefinitions: []interface{}{
//   				&metricDefinitionProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					dimensionKeys: map[string]*string{
//   						"dimensionKeysKey": jsii.String("dimensionKeys"),
//   					},
//   					eventPattern: jsii.String("eventPattern"),
//   					unitLabel: jsii.String("unitLabel"),
//   					valueKey: jsii.String("valueKey"),
//   				},
//   			},
//   		},
//   	},
//   	sessionSampleRate: jsii.Number(123),
//   	telemetries: []*string{
//   		jsii.String("telemetries"),
//   	},
//   }
//
type CfnAppMonitor_AppMonitorConfigurationProperty struct {
	// If you set this to `true` , the CloudWatch RUM web client sets two cookies, a session cookie and a user cookie.
	//
	// The cookies allow the CloudWatch RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	AllowCookies interface{} `field:"optional" json:"allowCookies" yaml:"allowCookies"`
	// If you set this to `true` , CloudWatch RUM sends client-side traces to X-Ray for each sampled session.
	//
	// You can then see traces and segments from these user sessions in the RUM dashboard and the CloudWatch ServiceLens console. For more information, see [What is AWS X-Ray ?](https://docs.aws.amazon.com/xray/latest/devguide/aws-xray.html)
	EnableXRay interface{} `field:"optional" json:"enableXRay" yaml:"enableXRay"`
	// A list of URLs in your website or application to exclude from RUM data collection.
	//
	// You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.
	ExcludedPages *[]*string `field:"optional" json:"excludedPages" yaml:"excludedPages"`
	// A list of pages in your application that are to be displayed with a "favorite" icon in the CloudWatch RUM console.
	FavoritePages *[]*string `field:"optional" json:"favoritePages" yaml:"favoritePages"`
	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to CloudWatch RUM.
	GuestRoleArn *string `field:"optional" json:"guestRoleArn" yaml:"guestRoleArn"`
	// The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to CloudWatch RUM.
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	//
	// You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.
	IncludedPages *[]*string `field:"optional" json:"includedPages" yaml:"includedPages"`
	// `CfnAppMonitor.AppMonitorConfigurationProperty.MetricDestinations`.
	MetricDestinations interface{} `field:"optional" json:"metricDestinations" yaml:"metricDestinations"`
	// Specifies the portion of user sessions to use for CloudWatch RUM data collection.
	//
	// Choosing a higher portion gives you more data but also incurs more costs.
	//
	// The range for this value is 0 to 1 inclusive. Setting this to 1 means that 100% of user sessions are sampled, and setting it to 0.1 means that 10% of user sessions are sampled.
	//
	// If you omit this parameter, the default of 0.1 is used, and 10% of sessions will be sampled.
	SessionSampleRate *float64 `field:"optional" json:"sessionSampleRate" yaml:"sessionSampleRate"`
	// An array that lists the types of telemetry data that this app monitor is to collect.
	//
	// - `errors` indicates that RUM collects data about unhandled JavaScript errors raised by your application.
	// - `performance` indicates that RUM collects performance data about how your application and its resources are loaded and rendered. This includes Core Web Vitals.
	// - `http` indicates that RUM collects data about HTTP errors thrown by your application.
	Telemetries *[]*string `field:"optional" json:"telemetries" yaml:"telemetries"`
}

