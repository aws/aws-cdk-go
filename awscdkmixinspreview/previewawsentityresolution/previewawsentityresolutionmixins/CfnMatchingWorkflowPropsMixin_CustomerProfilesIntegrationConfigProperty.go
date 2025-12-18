package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customerProfilesIntegrationConfigProperty := &CustomerProfilesIntegrationConfigProperty{
//   	DomainArn: jsii.String("domainArn"),
//   	ObjectTypeArn: jsii.String("objectTypeArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-customerprofilesintegrationconfig.html
//
type CfnMatchingWorkflowPropsMixin_CustomerProfilesIntegrationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-customerprofilesintegrationconfig.html#cfn-entityresolution-matchingworkflow-customerprofilesintegrationconfig-domainarn
	//
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-customerprofilesintegrationconfig.html#cfn-entityresolution-matchingworkflow-customerprofilesintegrationconfig-objecttypearn
	//
	ObjectTypeArn *string `field:"optional" json:"objectTypeArn" yaml:"objectTypeArn"`
}

