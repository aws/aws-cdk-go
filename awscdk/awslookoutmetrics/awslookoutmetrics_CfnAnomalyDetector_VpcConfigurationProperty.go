package awslookoutmetrics


// Contains configuration information about the Amazon Virtual Private Cloud (VPC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &vpcConfigurationProperty{
//   	securityGroupIdList: []*string{
//   		jsii.String("securityGroupIdList"),
//   	},
//   	subnetIdList: []*string{
//   		jsii.String("subnetIdList"),
//   	},
//   }
//
type CfnAnomalyDetector_VpcConfigurationProperty struct {
	// An array of strings containing the list of security groups.
	SecurityGroupIdList *[]*string `field:"required" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// An array of strings containing the Amazon VPC subnet IDs (e.g., `subnet-0bb1c79de3EXAMPLE` .
	SubnetIdList *[]*string `field:"required" json:"subnetIdList" yaml:"subnetIdList"`
}

