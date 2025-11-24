package mixinsawsconnectcampaignsv2


// Communication time configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   communicationTimeConfigProperty := &CommunicationTimeConfigProperty{
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
//   		RestrictedPeriods: &RestrictedPeriodsProperty{
//   			RestrictedPeriodList: []interface{}{
//   				&RestrictedPeriodProperty{
//   					EndDate: jsii.String("endDate"),
//   					Name: jsii.String("name"),
//   					StartDate: jsii.String("startDate"),
//   				},
//   			},
//   		},
//   	},
//   	LocalTimeZoneConfig: &LocalTimeZoneConfigProperty{
//   		DefaultTimeZone: jsii.String("defaultTimeZone"),
//   		LocalTimeZoneDetection: []*string{
//   			jsii.String("localTimeZoneDetection"),
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
//   		RestrictedPeriods: &RestrictedPeriodsProperty{
//   			RestrictedPeriodList: []interface{}{
//   				&RestrictedPeriodProperty{
//   					EndDate: jsii.String("endDate"),
//   					Name: jsii.String("name"),
//   					StartDate: jsii.String("startDate"),
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
//   		RestrictedPeriods: &RestrictedPeriodsProperty{
//   			RestrictedPeriodList: []interface{}{
//   				&RestrictedPeriodProperty{
//   					EndDate: jsii.String("endDate"),
//   					Name: jsii.String("name"),
//   					StartDate: jsii.String("startDate"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html
//
type CfnCampaignPropsMixin_CommunicationTimeConfigProperty struct {
	// The communication time configuration for the email channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-email
	//
	Email interface{} `field:"optional" json:"email" yaml:"email"`
	// The local timezone configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-localtimezoneconfig
	//
	LocalTimeZoneConfig interface{} `field:"optional" json:"localTimeZoneConfig" yaml:"localTimeZoneConfig"`
	// The communication time configuration for the SMS channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-sms
	//
	Sms interface{} `field:"optional" json:"sms" yaml:"sms"`
	// The communication time configuration for the telephony channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationtimeconfig.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig-telephony
	//
	Telephony interface{} `field:"optional" json:"telephony" yaml:"telephony"`
}

