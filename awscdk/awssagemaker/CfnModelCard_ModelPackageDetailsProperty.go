package awssagemaker


// Metadata information related to model package version.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html
//
type CfnModelCard_ModelPackageDetailsProperty struct {
	// A description provided for the model approval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-approvaldescription
	//
	ApprovalDescription *string `field:"optional" json:"approvalDescription" yaml:"approvalDescription"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-createdby
	//
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// The machine learning domain of the model package you specified.
	//
	// Common machine learning domains include computer vision and natural language processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-inferencespecification
	//
	InferenceSpecification interface{} `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// Current approval status of model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-modelapprovalstatus
	//
	ModelApprovalStatus *string `field:"optional" json:"modelApprovalStatus" yaml:"modelApprovalStatus"`
	// The Amazon Resource Name (ARN) of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-modelpackagearn
	//
	ModelPackageArn *string `field:"optional" json:"modelPackageArn" yaml:"modelPackageArn"`
	// A brief summary of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-modelpackagedescription
	//
	ModelPackageDescription *string `field:"optional" json:"modelPackageDescription" yaml:"modelPackageDescription"`
	// If the model is a versioned model, the name of the model group that the versioned model belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-modelpackagegroupname
	//
	ModelPackageGroupName *string `field:"optional" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// Name of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-modelpackagename
	//
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// Current status of model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-modelpackagestatus
	//
	ModelPackageStatus *string `field:"optional" json:"modelPackageStatus" yaml:"modelPackageStatus"`
	// Version of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-modelpackageversion
	//
	ModelPackageVersion *float64 `field:"optional" json:"modelPackageVersion" yaml:"modelPackageVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-sourcealgorithms
	//
	SourceAlgorithms interface{} `field:"optional" json:"sourceAlgorithms" yaml:"sourceAlgorithms"`
	// The machine learning task you specified that your model package accomplishes.
	//
	// Common machine learning tasks include object detection and image classification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagedetails.html#cfn-sagemaker-modelcard-modelpackagedetails-task
	//
	Task *string `field:"optional" json:"task" yaml:"task"`
}

