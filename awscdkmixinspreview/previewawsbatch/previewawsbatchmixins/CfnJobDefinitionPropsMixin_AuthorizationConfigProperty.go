package previewawsbatchmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authorizationConfigProperty := &AuthorizationConfigProperty{
//   	AccessPointId: jsii.String("accessPointId"),
//   	Iam: jsii.String("iam"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-authorizationconfig.html
//
type CfnJobDefinitionPropsMixin_AuthorizationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-authorizationconfig.html#cfn-batch-jobdefinition-authorizationconfig-accesspointid
	//
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-authorizationconfig.html#cfn-batch-jobdefinition-authorizationconfig-iam
	//
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

