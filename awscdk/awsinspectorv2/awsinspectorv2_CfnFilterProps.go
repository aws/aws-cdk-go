package awsinspectorv2


// Properties for defining a `CfnFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFilterProps := &cfnFilterProps{
//   	filterAction: jsii.String("filterAction"),
//   	filterCriteria: &filterCriteriaProperty{
//   		awsAccountId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceImageId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceSubnetId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceVpcId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageArchitecture: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageHash: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImagePushedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		ecrImageRegistry: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageRepositoryName: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageTags: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingArn: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingStatus: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		firstObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		inspectorScore: []interface{}{
//   			&numberFilterProperty{
//   				lowerInclusive: jsii.Number(123),
//   				upperInclusive: jsii.Number(123),
//   			},
//   		},
//   		lastObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		networkProtocol: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		portRange: []interface{}{
//   			&portRangeFilterProperty{
//   				beginInclusive: jsii.Number(123),
//   				endInclusive: jsii.Number(123),
//   			},
//   		},
//   		relatedVulnerabilities: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceTags: []interface{}{
//   			&mapFilterProperty{
//   				comparison: jsii.String("comparison"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		severity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		title: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		updatedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		vendorSeverity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilityId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilitySource: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerablePackages: []interface{}{
//   			&packageFilterProperty{
//   				architecture: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				epoch: &numberFilterProperty{
//   					lowerInclusive: jsii.Number(123),
//   					upperInclusive: jsii.Number(123),
//   				},
//   				name: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				release: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				sourceLayerHash: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				version: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnFilterProps struct {
	// The action that is to be applied to the findings that match the filter.
	FilterAction *string `field:"required" json:"filterAction" yaml:"filterAction"`
	// Details on the filter criteria associated with this filter.
	FilterCriteria interface{} `field:"required" json:"filterCriteria" yaml:"filterCriteria"`
	// The name of the filter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the filter.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

