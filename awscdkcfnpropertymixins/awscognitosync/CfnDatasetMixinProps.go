package awscognitosync


// Properties for CfnDatasetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDatasetMixinProps := &CfnDatasetMixinProps{
//   	DatasetName: jsii.String("datasetName"),
//   	IdentityId: jsii.String("identityId"),
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognitosync-dataset.html
//
type CfnDatasetMixinProps struct {
	// A string of up to 128 characters.
	//
	// Allowed characters are a-z, A-Z, 0-9, underscore, dash, and dot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognitosync-dataset.html#cfn-cognitosync-dataset-datasetname
	//
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognitosync-dataset.html#cfn-cognitosync-dataset-identityid
	//
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognitosync-dataset.html#cfn-cognitosync-dataset-identitypoolid
	//
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
}

