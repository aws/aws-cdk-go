package mixinsawsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnProjectPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProjectMixinProps := &CfnProjectMixinProps{
//   	DatasetName: jsii.String("datasetName"),
//   	Name: jsii.String("name"),
//   	RecipeName: jsii.String("recipeName"),
//   	RoleArn: jsii.String("roleArn"),
//   	Sample: &SampleProperty{
//   		Size: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html
//
type CfnProjectMixinProps struct {
	// The dataset that the project is to act upon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-datasetname
	//
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// The unique name of a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of a recipe that will be developed during a project session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-recipename
	//
	RecipeName *string `field:"optional" json:"recipeName" yaml:"recipeName"`
	// The Amazon Resource Name (ARN) of the role that will be assumed for this project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The sample size and sampling type to apply to the data.
	//
	// If this parameter isn't specified, then the sample consists of the first 500 rows from the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-sample
	//
	Sample interface{} `field:"optional" json:"sample" yaml:"sample"`
	// Metadata tags that have been applied to the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

