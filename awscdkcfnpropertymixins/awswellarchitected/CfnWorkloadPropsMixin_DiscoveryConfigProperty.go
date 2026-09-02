package awswellarchitected


// Discovery configuration associated to the workload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   discoveryConfigProperty := &DiscoveryConfigProperty{
//   	TrustedAdvisorIntegrationStatus: jsii.String("trustedAdvisorIntegrationStatus"),
//   	WorkloadResourceDefinition: []*string{
//   		jsii.String("workloadResourceDefinition"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-workload-discoveryconfig.html
//
type CfnWorkloadPropsMixin_DiscoveryConfigProperty struct {
	// Discovery integration status in respect to Trusted Advisor for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-workload-discoveryconfig.html#cfn-wellarchitected-workload-discoveryconfig-trustedadvisorintegrationstatus
	//
	TrustedAdvisorIntegrationStatus *string `field:"optional" json:"trustedAdvisorIntegrationStatus" yaml:"trustedAdvisorIntegrationStatus"`
	// The mode to use for identifying resources associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-workload-discoveryconfig.html#cfn-wellarchitected-workload-discoveryconfig-workloadresourcedefinition
	//
	WorkloadResourceDefinition *[]*string `field:"optional" json:"workloadResourceDefinition" yaml:"workloadResourceDefinition"`
}

