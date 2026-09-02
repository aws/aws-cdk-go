package awswellarchitected


// Properties for CfnWorkloadPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWorkloadMixinProps := &CfnWorkloadMixinProps{
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//   	ArchitecturalDesign: jsii.String("architecturalDesign"),
//   	AwsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   	Description: jsii.String("description"),
//   	DiscoveryConfig: &DiscoveryConfigProperty{
//   		TrustedAdvisorIntegrationStatus: jsii.String("trustedAdvisorIntegrationStatus"),
//   		WorkloadResourceDefinition: []*string{
//   			jsii.String("workloadResourceDefinition"),
//   		},
//   	},
//   	Environment: jsii.String("environment"),
//   	Industry: jsii.String("industry"),
//   	IndustryType: jsii.String("industryType"),
//   	Lenses: []*string{
//   		jsii.String("lenses"),
//   	},
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
//   	WorkloadName: jsii.String("workloadName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html
//
type CfnWorkloadMixinProps struct {
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
	// The description for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Discovery configuration associated to the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-discoveryconfig
	//
	DiscoveryConfig interface{} `field:"optional" json:"discoveryConfig" yaml:"discoveryConfig"`
	// The environment for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-environment
	//
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The industry for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-industry
	//
	Industry *string `field:"optional" json:"industry" yaml:"industry"`
	// The industry type for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-industrytype
	//
	IndustryType *string `field:"optional" json:"industryType" yaml:"industryType"`
	// The list of lenses associated with the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-lenses
	//
	Lenses *[]*string `field:"optional" json:"lenses" yaml:"lenses"`
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
	Tags *[]*CfnWorkloadPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html#cfn-wellarchitected-workload-workloadname
	//
	WorkloadName *string `field:"optional" json:"workloadName" yaml:"workloadName"`
}

