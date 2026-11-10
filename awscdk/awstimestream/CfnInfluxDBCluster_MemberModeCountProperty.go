package awstimestream


// The number of DB cluster members reporting one instance mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberModeCountProperty := &MemberModeCountProperty{
//   	Count: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-membermodecount.html
//
type CfnInfluxDBCluster_MemberModeCountProperty struct {
	// The number of DB cluster members reporting this instance mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-membermodecount.html#cfn-timestream-influxdbcluster-membermodecount-count
	//
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The instance mode reported by DB cluster members.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-membermodecount.html#cfn-timestream-influxdbcluster-membermodecount-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

