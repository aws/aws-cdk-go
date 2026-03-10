package previewawsxrayevents


// Type definition for Event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var names interface{}
//
//   event := &Event{
//   	ClientRequestImpactStatistics: &ClientRequestImpactStatistics{
//   		FaultCount: []*string{
//   			jsii.String("faultCount"),
//   		},
//   		OkCount: []*string{
//   			jsii.String("okCount"),
//   		},
//   		TotalCount: []*string{
//   			jsii.String("totalCount"),
//   		},
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	RootCauseServiceRequestImpactStatistics: &ClientRequestImpactStatistics{
//   		FaultCount: []*string{
//   			jsii.String("faultCount"),
//   		},
//   		OkCount: []*string{
//   			jsii.String("okCount"),
//   		},
//   		TotalCount: []*string{
//   			jsii.String("totalCount"),
//   		},
//   	},
//   	Summary: []*string{
//   		jsii.String("summary"),
//   	},
//   	TopAnomalousServices: []AwsxRayInsightUpdateItem{
//   		&AwsxRayInsightUpdateItem{
//   			ServiceId: &ServiceId{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Names: []interface{}{
//   					names,
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AWSXRayInsightUpdate_Event struct {
	// ClientRequestImpactStatistics property.
	//
	// Specify an array of string values to match this event if the actual value of ClientRequestImpactStatistics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientRequestImpactStatistics *AWSXRayInsightUpdate_ClientRequestImpactStatistics `field:"optional" json:"clientRequestImpactStatistics" yaml:"clientRequestImpactStatistics"`
	// EventTime property.
	//
	// Specify an array of string values to match this event if the actual value of EventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// RootCauseServiceRequestImpactStatistics property.
	//
	// Specify an array of string values to match this event if the actual value of RootCauseServiceRequestImpactStatistics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RootCauseServiceRequestImpactStatistics *AWSXRayInsightUpdate_ClientRequestImpactStatistics `field:"optional" json:"rootCauseServiceRequestImpactStatistics" yaml:"rootCauseServiceRequestImpactStatistics"`
	// Summary property.
	//
	// Specify an array of string values to match this event if the actual value of Summary is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Summary *[]*string `field:"optional" json:"summary" yaml:"summary"`
	// TopAnomalousServices property.
	//
	// Specify an array of string values to match this event if the actual value of TopAnomalousServices is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TopAnomalousServices *[]*AWSXRayInsightUpdate_AwsxRayInsightUpdateItem `field:"optional" json:"topAnomalousServices" yaml:"topAnomalousServices"`
}

