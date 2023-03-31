package awsinspectorv2


// Details on the criteria used to define the filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterCriteriaProperty := &filterCriteriaProperty{
//   	awsAccountId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	componentId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	componentType: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2InstanceImageId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2InstanceSubnetId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2InstanceVpcId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageArchitecture: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageHash: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImagePushedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	ecrImageRegistry: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageRepositoryName: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageTags: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	findingArn: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	findingStatus: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	findingType: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	firstObservedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	inspectorScore: []interface{}{
//   		&numberFilterProperty{
//   			lowerInclusive: jsii.Number(123),
//   			upperInclusive: jsii.Number(123),
//   		},
//   	},
//   	lastObservedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	networkProtocol: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	portRange: []interface{}{
//   		&portRangeFilterProperty{
//   			beginInclusive: jsii.Number(123),
//   			endInclusive: jsii.Number(123),
//   		},
//   	},
//   	relatedVulnerabilities: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	resourceId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	resourceTags: []interface{}{
//   		&mapFilterProperty{
//   			comparison: jsii.String("comparison"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	resourceType: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	severity: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	title: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	updatedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	vendorSeverity: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vulnerabilityId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vulnerabilitySource: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vulnerablePackages: []interface{}{
//   		&packageFilterProperty{
//   			architecture: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			epoch: &numberFilterProperty{
//   				lowerInclusive: jsii.Number(123),
//   				upperInclusive: jsii.Number(123),
//   			},
//   			name: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			release: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			sourceLayerHash: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			version: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnFilter_FilterCriteriaProperty struct {
	// Details of the AWS account IDs used to filter findings.
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Details of the component IDs used to filter findings.
	ComponentId interface{} `field:"optional" json:"componentId" yaml:"componentId"`
	// Details of the component types used to filter findings.
	ComponentType interface{} `field:"optional" json:"componentType" yaml:"componentType"`
	// Details of the Amazon EC2 instance image IDs used to filter findings.
	Ec2InstanceImageId interface{} `field:"optional" json:"ec2InstanceImageId" yaml:"ec2InstanceImageId"`
	// Details of the Amazon EC2 instance subnet IDs used to filter findings.
	Ec2InstanceSubnetId interface{} `field:"optional" json:"ec2InstanceSubnetId" yaml:"ec2InstanceSubnetId"`
	// Details of the Amazon EC2 instance VPC IDs used to filter findings.
	Ec2InstanceVpcId interface{} `field:"optional" json:"ec2InstanceVpcId" yaml:"ec2InstanceVpcId"`
	// Details of the Amazon ECR image architecture types used to filter findings.
	EcrImageArchitecture interface{} `field:"optional" json:"ecrImageArchitecture" yaml:"ecrImageArchitecture"`
	// Details of the Amazon ECR image hashes used to filter findings.
	EcrImageHash interface{} `field:"optional" json:"ecrImageHash" yaml:"ecrImageHash"`
	// Details on the Amazon ECR image push date and time used to filter findings.
	EcrImagePushedAt interface{} `field:"optional" json:"ecrImagePushedAt" yaml:"ecrImagePushedAt"`
	// Details on the Amazon ECR registry used to filter findings.
	EcrImageRegistry interface{} `field:"optional" json:"ecrImageRegistry" yaml:"ecrImageRegistry"`
	// Details on the name of the Amazon ECR repository used to filter findings.
	EcrImageRepositoryName interface{} `field:"optional" json:"ecrImageRepositoryName" yaml:"ecrImageRepositoryName"`
	// The tags attached to the Amazon ECR container image.
	EcrImageTags interface{} `field:"optional" json:"ecrImageTags" yaml:"ecrImageTags"`
	// Details on the finding ARNs used to filter findings.
	FindingArn interface{} `field:"optional" json:"findingArn" yaml:"findingArn"`
	// Details on the finding status types used to filter findings.
	FindingStatus interface{} `field:"optional" json:"findingStatus" yaml:"findingStatus"`
	// Details on the finding types used to filter findings.
	FindingType interface{} `field:"optional" json:"findingType" yaml:"findingType"`
	// Details on the date and time a finding was first seen used to filter findings.
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// The Amazon Inspector score to filter on.
	InspectorScore interface{} `field:"optional" json:"inspectorScore" yaml:"inspectorScore"`
	// Details on the date and time a finding was last seen used to filter findings.
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// Details on the ingress source addresses used to filter findings.
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// Details on the port ranges used to filter findings.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// Details on the related vulnerabilities used to filter findings.
	RelatedVulnerabilities interface{} `field:"optional" json:"relatedVulnerabilities" yaml:"relatedVulnerabilities"`
	// Details on the resource IDs used to filter findings.
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Details on the resource tags used to filter findings.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Details on the resource types used to filter findings.
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Details on the severity used to filter findings.
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// Details on the finding title used to filter findings.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// Details on the date and time a finding was last updated at used to filter findings.
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Details on the vendor severity used to filter findings.
	VendorSeverity interface{} `field:"optional" json:"vendorSeverity" yaml:"vendorSeverity"`
	// Details on the vulnerability ID used to filter findings.
	VulnerabilityId interface{} `field:"optional" json:"vulnerabilityId" yaml:"vulnerabilityId"`
	// Details on the vulnerability score to filter findings by.
	VulnerabilitySource interface{} `field:"optional" json:"vulnerabilitySource" yaml:"vulnerabilitySource"`
	// Details on the vulnerable packages used to filter findings.
	VulnerablePackages interface{} `field:"optional" json:"vulnerablePackages" yaml:"vulnerablePackages"`
}

