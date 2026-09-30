package awsbcm


// The health of the scheduled report as of its last refresh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   healthStatusProperty := &HealthStatusProperty{
//   	LastRefreshedAt: jsii.String("lastRefreshedAt"),
//   	StatusCode: jsii.String("statusCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-healthstatus.html
//
type CfnScheduledReportPropsMixin_HealthStatusProperty struct {
	// The time at which the health status was last refreshed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-healthstatus.html#cfn-bcm-scheduledreport-healthstatus-lastrefreshedat
	//
	LastRefreshedAt *string `field:"optional" json:"lastRefreshedAt" yaml:"lastRefreshedAt"`
	// The health status of a scheduled report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-healthstatus.html#cfn-bcm-scheduledreport-healthstatus-statuscode
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

