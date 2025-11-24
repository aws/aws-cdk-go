package mixinsawsdatazone


// Specifies whether to create a project from project profile policy grant details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createProjectFromProjectProfilePolicyGrantDetailProperty := &CreateProjectFromProjectProfilePolicyGrantDetailProperty{
//   	IncludeChildDomainUnits: jsii.Boolean(false),
//   	ProjectProfiles: []*string{
//   		jsii.String("projectProfiles"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createprojectfromprojectprofilepolicygrantdetail.html
//
type CfnPolicyGrantPropsMixin_CreateProjectFromProjectProfilePolicyGrantDetailProperty struct {
	// Specifies whether to include child domain units when creating a project from project profile policy grant details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createprojectfromprojectprofilepolicygrantdetail.html#cfn-datazone-policygrant-createprojectfromprojectprofilepolicygrantdetail-includechilddomainunits
	//
	IncludeChildDomainUnits interface{} `field:"optional" json:"includeChildDomainUnits" yaml:"includeChildDomainUnits"`
	// Specifies project profiles when creating a project from project profile policy grant details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createprojectfromprojectprofilepolicygrantdetail.html#cfn-datazone-policygrant-createprojectfromprojectprofilepolicygrantdetail-projectprofiles
	//
	ProjectProfiles *[]*string `field:"optional" json:"projectProfiles" yaml:"projectProfiles"`
}

