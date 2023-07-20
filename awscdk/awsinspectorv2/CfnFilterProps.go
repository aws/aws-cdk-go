package awsinspectorv2


// Properties for defining a `CfnFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFilterProps := &CfnFilterProps{
//   	FilterAction: jsii.String("filterAction"),
//   	FilterCriteria: &FilterCriteriaProperty{
//   		AwsAccountId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComponentId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComponentType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Ec2InstanceImageId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Ec2InstanceSubnetId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Ec2InstanceVpcId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EcrImageArchitecture: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EcrImageHash: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EcrImagePushedAt: []interface{}{
//   			&DateFilterProperty{
//   				EndInclusive: jsii.Number(123),
//   				StartInclusive: jsii.Number(123),
//   			},
//   		},
//   		EcrImageRegistry: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EcrImageRepositoryName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EcrImageTags: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingStatus: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FirstObservedAt: []interface{}{
//   			&DateFilterProperty{
//   				EndInclusive: jsii.Number(123),
//   				StartInclusive: jsii.Number(123),
//   			},
//   		},
//   		InspectorScore: []interface{}{
//   			&NumberFilterProperty{
//   				LowerInclusive: jsii.Number(123),
//   				UpperInclusive: jsii.Number(123),
//   			},
//   		},
//   		LastObservedAt: []interface{}{
//   			&DateFilterProperty{
//   				EndInclusive: jsii.Number(123),
//   				StartInclusive: jsii.Number(123),
//   			},
//   		},
//   		NetworkProtocol: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		PortRange: []interface{}{
//   			&PortRangeFilterProperty{
//   				BeginInclusive: jsii.Number(123),
//   				EndInclusive: jsii.Number(123),
//   			},
//   		},
//   		RelatedVulnerabilities: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceTags: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//
//   				// the properties below are optional
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Severity: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Title: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		UpdatedAt: []interface{}{
//   			&DateFilterProperty{
//   				EndInclusive: jsii.Number(123),
//   				StartInclusive: jsii.Number(123),
//   			},
//   		},
//   		VendorSeverity: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		VulnerabilityId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		VulnerabilitySource: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		VulnerablePackages: []interface{}{
//   			&PackageFilterProperty{
//   				Architecture: &StringFilterProperty{
//   					Comparison: jsii.String("comparison"),
//   					Value: jsii.String("value"),
//   				},
//   				Epoch: &NumberFilterProperty{
//   					LowerInclusive: jsii.Number(123),
//   					UpperInclusive: jsii.Number(123),
//   				},
//   				Name: &StringFilterProperty{
//   					Comparison: jsii.String("comparison"),
//   					Value: jsii.String("value"),
//   				},
//   				Release: &StringFilterProperty{
//   					Comparison: jsii.String("comparison"),
//   					Value: jsii.String("value"),
//   				},
//   				SourceLayerHash: &StringFilterProperty{
//   					Comparison: jsii.String("comparison"),
//   					Value: jsii.String("value"),
//   				},
//   				Version: &StringFilterProperty{
//   					Comparison: jsii.String("comparison"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-filter.html
//
type CfnFilterProps struct {
	// The action that is to be applied to the findings that match the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-filter.html#cfn-inspectorv2-filter-filteraction
	//
	FilterAction *string `field:"required" json:"filterAction" yaml:"filterAction"`
	// Details on the filter criteria associated with this filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-filter.html#cfn-inspectorv2-filter-filtercriteria
	//
	FilterCriteria interface{} `field:"required" json:"filterCriteria" yaml:"filterCriteria"`
	// The name of the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-filter.html#cfn-inspectorv2-filter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-filter.html#cfn-inspectorv2-filter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

