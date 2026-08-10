package awsbedrockagentcore


// Allow-list of upstream workloads permitted to reach this resource via the workload identity chain.
//
// When set, the data plane enforces that the introspected workload chain's caller matches one of the configured hosting environments or workload identities; absent means no chain enforcement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowedWorkloadConfigurationProperty := &AllowedWorkloadConfigurationProperty{
//   	HostingEnvironments: []interface{}{
//   		&HostingEnvironmentProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	WorkloadIdentities: []*string{
//   		jsii.String("workloadIdentities"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-allowedworkloadconfiguration.html
//
type CfnRuntime_AllowedWorkloadConfigurationProperty struct {
	// List of allow-listed hosting environments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-allowedworkloadconfiguration.html#cfn-bedrockagentcore-runtime-allowedworkloadconfiguration-hostingenvironments
	//
	HostingEnvironments interface{} `field:"optional" json:"hostingEnvironments" yaml:"hostingEnvironments"`
	// List of allow-listed workload identity names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-allowedworkloadconfiguration.html#cfn-bedrockagentcore-runtime-allowedworkloadconfiguration-workloadidentities
	//
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

