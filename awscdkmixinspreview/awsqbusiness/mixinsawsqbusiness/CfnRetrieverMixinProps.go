package mixinsawsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRetrieverPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRetrieverMixinProps := &CfnRetrieverMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Configuration: &RetrieverConfigurationProperty{
//   		KendraIndexConfiguration: &KendraIndexConfigurationProperty{
//   			IndexId: jsii.String("indexId"),
//   		},
//   		NativeIndexConfiguration: &NativeIndexConfigurationProperty{
//   			IndexId: jsii.String("indexId"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-retriever.html
//
type CfnRetrieverMixinProps struct {
	// The identifier of the Amazon Q Business application using the retriever.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-retriever.html#cfn-qbusiness-retriever-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Provides information on how the retriever used for your Amazon Q Business application is configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-retriever.html#cfn-qbusiness-retriever-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The name of your retriever.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-retriever.html#cfn-qbusiness-retriever-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The ARN of an IAM role used by Amazon Q Business to access the basic authentication credentials stored in a Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-retriever.html#cfn-qbusiness-retriever-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A list of key-value pairs that identify or categorize the retriever.
	//
	// You can also use tags to help control access to the retriever. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-retriever.html#cfn-qbusiness-retriever-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of your retriever.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-retriever.html#cfn-qbusiness-retriever-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

