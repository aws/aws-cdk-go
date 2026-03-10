package previewawsemrevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.emr@EMRInstanceGroupStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRInstanceGroupStateChangeProps := &EMRInstanceGroupStateChangeProps{
//   	BidPrice: []*string{
//   		jsii.String("bidPrice"),
//   	},
//   	BidPriceAsPercentageOfOnDemandPrice: []*string{
//   		jsii.String("bidPriceAsPercentageOfOnDemandPrice"),
//   	},
//   	ClusterId: []*string{
//   		jsii.String("clusterId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	InstanceGroupId: []*string{
//   		jsii.String("instanceGroupId"),
//   	},
//   	InstanceGroupType: []*string{
//   		jsii.String("instanceGroupType"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	Market: []*string{
//   		jsii.String("market"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	RequestedInstanceCount: []*string{
//   		jsii.String("requestedInstanceCount"),
//   	},
//   	RunningInstanceCount: []*string{
//   		jsii.String("runningInstanceCount"),
//   	},
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   }
//
// Experimental.
type EMRInstanceGroupStateChange_EMRInstanceGroupStateChangeProps struct {
	// bidPrice property.
	//
	// Specify an array of string values to match this event if the actual value of bidPrice is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BidPrice *[]*string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// bidPriceAsPercentageOfOnDemandPrice property.
	//
	// Specify an array of string values to match this event if the actual value of bidPriceAsPercentageOfOnDemandPrice is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BidPriceAsPercentageOfOnDemandPrice *[]*string `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// clusterId property.
	//
	// Specify an array of string values to match this event if the actual value of clusterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClusterId *[]*string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// instanceGroupId property.
	//
	// Specify an array of string values to match this event if the actual value of instanceGroupId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceGroupId *[]*string `field:"optional" json:"instanceGroupId" yaml:"instanceGroupId"`
	// instanceGroupType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceGroupType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceGroupType *[]*string `field:"optional" json:"instanceGroupType" yaml:"instanceGroupType"`
	// instanceType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// market property.
	//
	// Specify an array of string values to match this event if the actual value of market is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Market *[]*string `field:"optional" json:"market" yaml:"market"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// requestedInstanceCount property.
	//
	// Specify an array of string values to match this event if the actual value of requestedInstanceCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestedInstanceCount *[]*string `field:"optional" json:"requestedInstanceCount" yaml:"requestedInstanceCount"`
	// runningInstanceCount property.
	//
	// Specify an array of string values to match this event if the actual value of runningInstanceCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RunningInstanceCount *[]*string `field:"optional" json:"runningInstanceCount" yaml:"runningInstanceCount"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
}

