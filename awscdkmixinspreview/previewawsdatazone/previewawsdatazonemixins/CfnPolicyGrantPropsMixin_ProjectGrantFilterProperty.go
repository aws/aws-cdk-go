package previewawsdatazonemixins


// The project grant filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   projectGrantFilterProperty := &ProjectGrantFilterProperty{
//   	DomainUnitFilter: &DomainUnitFilterForProjectProperty{
//   		DomainUnit: jsii.String("domainUnit"),
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-projectgrantfilter.html
//
type CfnPolicyGrantPropsMixin_ProjectGrantFilterProperty struct {
	// The domain unit filter of the project grant filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-projectgrantfilter.html#cfn-datazone-policygrant-projectgrantfilter-domainunitfilter
	//
	DomainUnitFilter interface{} `field:"optional" json:"domainUnitFilter" yaml:"domainUnitFilter"`
}

