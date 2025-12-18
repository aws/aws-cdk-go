package awsentityresolution


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerProfilesIntegrationConfigProperty := &CustomerProfilesIntegrationConfigProperty{
//   	DomainArn: jsii.String("domainArn"),
//   	ObjectTypeArn: jsii.String("objectTypeArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-customerprofilesintegrationconfig.html
//
type CfnMatchingWorkflow_CustomerProfilesIntegrationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-customerprofilesintegrationconfig.html#cfn-entityresolution-matchingworkflow-customerprofilesintegrationconfig-domainarn
	//
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-customerprofilesintegrationconfig.html#cfn-entityresolution-matchingworkflow-customerprofilesintegrationconfig-objecttypearn
	//
	ObjectTypeArn *string `field:"required" json:"objectTypeArn" yaml:"objectTypeArn"`
}

