package awsgreengrass


// Settings for an Secrets Manager machine learning resource.
//
// For more information, see [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `SageMakerMachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerMachineLearningModelResourceDataProperty := &SageMakerMachineLearningModelResourceDataProperty{
//   	DestinationPath: jsii.String("destinationPath"),
//   	SageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   	// the properties below are optional
//   	OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   		GroupOwner: jsii.String("groupOwner"),
//   		GroupPermission: jsii.String("groupPermission"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata.html
//
type CfnResourceDefinition_SageMakerMachineLearningModelResourceDataProperty struct {
	// The absolute local path of the resource inside the Lambda environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata.html#cfn-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata-destinationpath
	//
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The Amazon Resource Name (ARN) of the Amazon SageMaker training job that represents the source model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata.html#cfn-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata-sagemakerjobarn
	//
	SageMakerJobArn *string `field:"required" json:"sageMakerJobArn" yaml:"sageMakerJobArn"`
	// The owner setting for the downloaded machine learning resource.
	//
	// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata.html#cfn-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata-ownersetting
	//
	OwnerSetting interface{} `field:"optional" json:"ownerSetting" yaml:"ownerSetting"`
}

