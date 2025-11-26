package previewawsdatazonemixins


// The project policy grant principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   projectPolicyGrantPrincipalProperty := &ProjectPolicyGrantPrincipalProperty{
//   	ProjectDesignation: jsii.String("projectDesignation"),
//   	ProjectGrantFilter: &ProjectGrantFilterProperty{
//   		DomainUnitFilter: &DomainUnitFilterForProjectProperty{
//   			DomainUnit: jsii.String("domainUnit"),
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   	},
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-projectpolicygrantprincipal.html
//
type CfnPolicyGrantPropsMixin_ProjectPolicyGrantPrincipalProperty struct {
	// The project designation of the project policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-projectpolicygrantprincipal.html#cfn-datazone-policygrant-projectpolicygrantprincipal-projectdesignation
	//
	ProjectDesignation *string `field:"optional" json:"projectDesignation" yaml:"projectDesignation"`
	// The project grant filter of the project policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-projectpolicygrantprincipal.html#cfn-datazone-policygrant-projectpolicygrantprincipal-projectgrantfilter
	//
	ProjectGrantFilter interface{} `field:"optional" json:"projectGrantFilter" yaml:"projectGrantFilter"`
	// The project ID of the project policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-projectpolicygrantprincipal.html#cfn-datazone-policygrant-projectpolicygrantprincipal-projectidentifier
	//
	ProjectIdentifier *string `field:"optional" json:"projectIdentifier" yaml:"projectIdentifier"`
}

