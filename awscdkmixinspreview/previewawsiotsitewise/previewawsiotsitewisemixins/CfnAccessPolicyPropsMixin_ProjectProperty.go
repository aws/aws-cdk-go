package previewawsiotsitewisemixins


// Identifies a specific AWS IoT SiteWise Monitor project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   projectProperty := &ProjectProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-project.html
//
type CfnAccessPolicyPropsMixin_ProjectProperty struct {
	// The ID of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-project.html#cfn-iotsitewise-accesspolicy-project-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

