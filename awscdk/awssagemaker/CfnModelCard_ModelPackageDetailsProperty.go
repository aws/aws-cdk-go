package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageDetailsProperty := &ModelPackageDetailsProperty{
//   	ApprovalDescription: jsii.String("approvalDescription"),
//   	CreatedBy: &ModelPackageCreatorProperty{
//   		UserProfileName: jsii.String("userProfileName"),
//   	},
//   	Domain: jsii.String("domain"),
//   	InferenceSpecification: &InferenceSpecificationProperty{
//   		Containers: []interface{}{
//   			&ContainerProperty{
//   				Image: jsii.String("image"),
//
//   				// the properties below are optional
//   				ModelDataUrl: jsii.String("modelDataUrl"),
//   				NearestModelName: jsii.String("nearestModelName"),
//   			},
//   		},
//   	},
//   	ModelApprovalStatus: jsii.String("modelApprovalStatus"),
//   	ModelPackageArn: jsii.String("modelPackageArn"),
//   	ModelPackageDescription: jsii.String("modelPackageDescription"),
//   	ModelPackageGroupName: jsii.String("modelPackageGroupName"),
//   	ModelPackageName: jsii.String("modelPackageName"),
//   	ModelPackageStatus: jsii.String("modelPackageStatus"),
//   	ModelPackageVersion: jsii.Number(123),
//   	SourceAlgorithms: []interface{}{
//   		&SourceAlgorithmProperty{
//   			AlgorithmName: jsii.String("algorithmName"),
//
//   			// the properties below are optional
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   		},
//   	},
//   	Task: jsii.String("task"),
//   }
//
type CfnModelCard_ModelPackageDetailsProperty struct {
	// `CfnModelCard.ModelPackageDetailsProperty.ApprovalDescription`.
	ApprovalDescription *string `field:"optional" json:"approvalDescription" yaml:"approvalDescription"`
	// `CfnModelCard.ModelPackageDetailsProperty.CreatedBy`.
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// `CfnModelCard.ModelPackageDetailsProperty.Domain`.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// `CfnModelCard.ModelPackageDetailsProperty.InferenceSpecification`.
	InferenceSpecification interface{} `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// `CfnModelCard.ModelPackageDetailsProperty.ModelApprovalStatus`.
	ModelApprovalStatus *string `field:"optional" json:"modelApprovalStatus" yaml:"modelApprovalStatus"`
	// `CfnModelCard.ModelPackageDetailsProperty.ModelPackageArn`.
	ModelPackageArn *string `field:"optional" json:"modelPackageArn" yaml:"modelPackageArn"`
	// `CfnModelCard.ModelPackageDetailsProperty.ModelPackageDescription`.
	ModelPackageDescription *string `field:"optional" json:"modelPackageDescription" yaml:"modelPackageDescription"`
	// `CfnModelCard.ModelPackageDetailsProperty.ModelPackageGroupName`.
	ModelPackageGroupName *string `field:"optional" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// `CfnModelCard.ModelPackageDetailsProperty.ModelPackageName`.
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// `CfnModelCard.ModelPackageDetailsProperty.ModelPackageStatus`.
	ModelPackageStatus *string `field:"optional" json:"modelPackageStatus" yaml:"modelPackageStatus"`
	// `CfnModelCard.ModelPackageDetailsProperty.ModelPackageVersion`.
	ModelPackageVersion *float64 `field:"optional" json:"modelPackageVersion" yaml:"modelPackageVersion"`
	// `CfnModelCard.ModelPackageDetailsProperty.SourceAlgorithms`.
	SourceAlgorithms interface{} `field:"optional" json:"sourceAlgorithms" yaml:"sourceAlgorithms"`
	// `CfnModelCard.ModelPackageDetailsProperty.Task`.
	Task *string `field:"optional" json:"task" yaml:"task"`
}

