package awsconnectcampaignsv2


// Communication time configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   communicationTimeConfigProperty := &CommunicationTimeConfigProperty{
//   	LocalTimeZoneConfig: &LocalTimeZoneConfigProperty{
//   		DefaultTimeZone: jsii.String("defaultTimeZone"),
//   		LocalTimeZoneDetection: []*string{
//   			jsii.String("localTimeZoneDetection"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Email: &TimeWindowProperty{
//   		OpenHours: &OpenHoursProperty{
//   			DailyHours: []interface{}{
//   				&DailyHourProperty{
//   					Key: jsii.String("key"),
//   					Value: []interface{}{
//   						&TimeRangeProperty{
//   							EndTime: jsii.String("endTime"),
//   							StartTime: jsii.String("startTime"),
//   						},
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		RestrictedPeriods: &RestrictedPeriodsProperty{
//   			RestrictedPeriodList: []interface{}{
//   				&RestrictedPeriodProperty{
//   					EndDate: jsii.String("endDate"),
//   					StartDate: jsii.String("startDate"),
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	Sms: &TimeWindowProperty{
//   		OpenHours: &OpenHoursProperty{
//   			DailyHours: []interface{}{
//   				&DailyHourProperty{
//   					Key: jsii.String("key"),
//   					Value: []interface{}{
//   						&TimeRangeProperty{
//   							EndTime: jsii.String("endTime"),
//   							StartTime: jsii.String("startTime"),
//   						},
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		RestrictedPeriods: &RestrictedPeriodsProperty{
//   			RestrictedPeriodList: []interface{}{
//   				&RestrictedPeriodProperty{
//   					EndDate: jsii.String("endDate"),
//   					StartDate: jsii.String("startDate"),
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	Telephony: &TimeWindowProperty{
//   		OpenHours: &OpenHoursProperty{
//   			DailyHours: []interface{}{
//   				&DailyHourProperty{
//   					Key: jsii.String("key"),
//   					Value: []interface{}{
//   						&TimeRangeProperty{
//   							EndTime: jsii.String("endTime"),
//   							StartTime: jsii.String("startTime"),
//   						},
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		RestrictedPeriods: &RestrictedPeriodsProperty{
//   			RestrictedPeriodList: []interface{}{
//   				&RestrictedPeriodProperty{
//   					EndDate: jsii.String("endDate"),
//   					StartDate: jsii.String("startDate"),
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html
//
type CfnCampaign_CommunicationTimeConfigProperty struct {
	// The local timezone configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-localtimezoneconfig
	//
	LocalTimeZoneConfig interface{} `field:"required" json:"localTimeZoneConfig" yaml:"localTimeZoneConfig"`
	// The communication time configuration for the email channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-email
	//
	Email interface{} `field:"optional" json:"email" yaml:"email"`
	// The communication time configuration for the SMS channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-sms
	//
	Sms interface{} `field:"optional" json:"sms" yaml:"sms"`
	// The communication time configuration for the telephony channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-telephony
	//
	Telephony interface{} `field:"optional" json:"telephony" yaml:"telephony"`
}

