package awspcs


// Additional cgroup configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cgroupCustomSettingProperty := &CgroupCustomSettingProperty{
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-cgroupcustomsetting.html
//
type CfnCluster_CgroupCustomSettingProperty struct {
	// The cgroup.conf parameter name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-cgroupcustomsetting.html#cfn-pcs-cluster-cgroupcustomsetting-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The value for the cgroup.conf parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-cgroupcustomsetting.html#cfn-pcs-cluster-cgroupcustomsetting-parametervalue
	//
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

