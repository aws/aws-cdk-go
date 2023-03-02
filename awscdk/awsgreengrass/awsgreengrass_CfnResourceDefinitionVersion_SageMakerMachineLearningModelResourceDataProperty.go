package awsgreengrass


// Settings for an Secrets Manager machine learning resource.
//
// For more information, see [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `SageMakerMachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerMachineLearningModelResourceDataProperty := &sageMakerMachineLearningModelResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
//   	sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   	// the properties below are optional
//   	ownerSetting: &resourceDownloadOwnerSettingProperty{
//   		groupOwner: jsii.String("groupOwner"),
//   		groupPermission: jsii.String("groupPermission"),
//   	},
//   }
//
type CfnResourceDefinitionVersion_SageMakerMachineLearningModelResourceDataProperty struct {
	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The Amazon Resource Name (ARN) of the Amazon SageMaker training job that represents the source model.
	SageMakerJobArn *string `field:"required" json:"sageMakerJobArn" yaml:"sageMakerJobArn"`
	// The owner setting for the downloaded machine learning resource.
	//
	// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
	OwnerSetting interface{} `field:"optional" json:"ownerSetting" yaml:"ownerSetting"`
}

