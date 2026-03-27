package previewawss3events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.s3@AWSAPICallViaCloudTrail event.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn Function
//
//
//   // Works with L2 Rule
//   // Works with L2 Rule
//   events.NewRule(scope, jsii.String("Rule"), &RuleProps{
//   	EventPattern: awscdkmixinspreview.AWSAPICallViaCloudTrail_AwsAPICallViaCloudTrailPattern(&AWSAPICallViaCloudTrailProps{
//   		TlsDetails: &TlsDetails{
//   			TlsVersion: []*string{
//   				jsii.String("TLSv1.3"),
//   			},
//   		},
//   		EventMetadata: &AWSEventMetadataProps{
//   			Region: []*string{
//   				jsii.String("us-east-1"),
//   			},
//   		},
//   	}),
//   	Targets: []IRuleTarget{
//   		targets.NewLambdaFunction(fn),
//   	},
//   })
//
//   // Also works with L1 CfnRule
//   // Also works with L1 CfnRule
//   events.NewCfnRule(scope, jsii.String("CfnRule"), &CfnRuleProps{
//   	State: jsii.String("ENABLED"),
//   	EventPattern: awscdkmixinspreview.AWSAPICallViaCloudTrail_*AwsAPICallViaCloudTrailPattern(&AWSAPICallViaCloudTrailProps{
//   		TlsDetails: &TlsDetails{
//   			TlsVersion: []*string{
//   				jsii.String("TLSv1.3"),
//   			},
//   		},
//   		EventMetadata: &AWSEventMetadataProps{
//   			Region: []*string{
//   				jsii.String("us-east-1"),
//   			},
//   		},
//   	}),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Arn: fn.functionArn,
//   			Id: jsii.String("L1"),
//   		},
//   	},
//   })
//
// Experimental.
type AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// additionalEventData property.
	//
	// Specify an array of string values to match this event if the actual value of additionalEventData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdditionalEventData *AWSAPICallViaCloudTrail_AdditionalEventData `field:"optional" json:"additionalEventData" yaml:"additionalEventData"`
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventCategory property.
	//
	// Specify an array of string values to match this event if the actual value of eventCategory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventCategory *[]*string `field:"optional" json:"eventCategory" yaml:"eventCategory"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// managementEvent property.
	//
	// Specify an array of string values to match this event if the actual value of managementEvent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ManagementEvent *[]*string `field:"optional" json:"managementEvent" yaml:"managementEvent"`
	// readOnly property.
	//
	// Specify an array of string values to match this event if the actual value of readOnly is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReadOnly *[]*string `field:"optional" json:"readOnly" yaml:"readOnly"`
	// recipientAccountId property.
	//
	// Specify an array of string values to match this event if the actual value of recipientAccountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecipientAccountId *[]*string `field:"optional" json:"recipientAccountId" yaml:"recipientAccountId"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// resources property.
	//
	// Specify an array of string values to match this event if the actual value of resources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Resources *[]*AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem `field:"optional" json:"resources" yaml:"resources"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *[]*string `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// tlsDetails property.
	//
	// Specify an array of string values to match this event if the actual value of tlsDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TlsDetails *AWSAPICallViaCloudTrail_TlsDetails `field:"optional" json:"tlsDetails" yaml:"tlsDetails"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
	// vpcEndpointId property.
	//
	// Specify an array of string values to match this event if the actual value of vpcEndpointId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcEndpointId *[]*string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

