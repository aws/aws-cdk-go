package awslookoutmetrics


// Contains configuration information about the Amazon Virtual Private Cloud (VPC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &VpcConfigurationProperty{
//   	SecurityGroupIdList: []*string{
//   		jsii.String("securityGroupIdList"),
//   	},
//   	SubnetIdList: []*string{
//   		jsii.String("subnetIdList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-vpcconfiguration.html
//
type CfnAnomalyDetector_VpcConfigurationProperty struct {
	// An array of strings containing the list of security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-vpcconfiguration.html#cfn-lookoutmetrics-anomalydetector-vpcconfiguration-securitygroupidlist
	//
	SecurityGroupIdList *[]*string `field:"required" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// An array of strings containing the Amazon VPC subnet IDs (e.g., `subnet-0bb1c79de3EXAMPLE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-vpcconfiguration.html#cfn-lookoutmetrics-anomalydetector-vpcconfiguration-subnetidlist
	//
	SubnetIdList *[]*string `field:"required" json:"subnetIdList" yaml:"subnetIdList"`
}

