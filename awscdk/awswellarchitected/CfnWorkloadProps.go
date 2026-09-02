package awswellarchitected


// Properties for defining a `CfnWorkload`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkloadProps := &CfnWorkloadProps{
//   	Description: jsii.String("description"),
//   	Environment: jsii.String("environment"),
//   	Lenses: []*string{
//   		jsii.String("lenses"),
//   	},
//   	WorkloadName: jsii.String("workloadName"),
//
//   	// the properties below are optional
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//   	ArchitecturalDesign: jsii.String("architecturalDesign"),
//   	AwsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   	DiscoveryConfig: &DiscoveryConfigProperty{
//   		TrustedAdvisorIntegrationStatus: jsii.String("trustedAdvisorIntegrationStatus"),
//   		WorkloadResourceDefinition: []*string{
//   			jsii.String("workloadResourceDefinition"),
//   		},
//   	},
//   	Industry: jsii.String("industry"),
//   	IndustryType: jsii.String("industryType"),
//   	NonAwsRegions: []*string{
//   		jsii.String("nonAwsRegions"),
//   	},
//   	Notes: jsii.String("notes"),
//   	ReviewOwner: jsii.String("reviewOwner"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html
//
type CfnWorkloadProps struct {
	// The description for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The environment for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-environment
	//
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The list of lenses associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-lenses
	//
	Lenses *[]*string `field:"required" json:"lenses" yaml:"lenses"`
	// The name of the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-workloadname
	//
	WorkloadName *string `field:"required" json:"workloadName" yaml:"workloadName"`
	// The list of Amazon Web Services account IDs associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-accountids
	//
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// The URL of the architectural design for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-architecturaldesign
	//
	ArchitecturalDesign *string `field:"optional" json:"architecturalDesign" yaml:"architecturalDesign"`
	// The list of Amazon Web Services Regions associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-awsregions
	//
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// Discovery configuration associated to the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-discoveryconfig
	//
	DiscoveryConfig interface{} `field:"optional" json:"discoveryConfig" yaml:"discoveryConfig"`
	// The industry for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-industry
	//
	Industry *string `field:"optional" json:"industry" yaml:"industry"`
	// The industry type for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-industrytype
	//
	IndustryType *string `field:"optional" json:"industryType" yaml:"industryType"`
	// The list of non-Amazon Web Services Regions associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-nonawsregions
	//
	NonAwsRegions *[]*string `field:"optional" json:"nonAwsRegions" yaml:"nonAwsRegions"`
	// The notes associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The review owner of the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-reviewowner
	//
	ReviewOwner *string `field:"optional" json:"reviewOwner" yaml:"reviewOwner"`
	// The tags associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-tags
	//
	Tags *[]*CfnWorkload_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

