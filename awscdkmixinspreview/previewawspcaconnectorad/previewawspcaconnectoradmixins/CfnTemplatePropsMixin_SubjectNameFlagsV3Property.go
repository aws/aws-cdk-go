package previewawspcaconnectoradmixins


// Information to include in the subject name and alternate subject name of the certificate.
//
// The subject name can be common name, directory path, DNS as common name, or left blank. You can optionally include email to the subject name for user templates. If you leave the subject name blank then you must set a subject alternate name. The subject alternate name (SAN) can include globally unique identifier (GUID), DNS, domain DNS, email, service principal name (SPN), and user principal name (UPN). You can leave the SAN blank. If you leave the SAN blank, then you must set a subject name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subjectNameFlagsV3Property := &SubjectNameFlagsV3Property{
//   	RequireCommonName: jsii.Boolean(false),
//   	RequireDirectoryPath: jsii.Boolean(false),
//   	RequireDnsAsCn: jsii.Boolean(false),
//   	RequireEmail: jsii.Boolean(false),
//   	SanRequireDirectoryGuid: jsii.Boolean(false),
//   	SanRequireDns: jsii.Boolean(false),
//   	SanRequireDomainDns: jsii.Boolean(false),
//   	SanRequireEmail: jsii.Boolean(false),
//   	SanRequireSpn: jsii.Boolean(false),
//   	SanRequireUpn: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html
//
type CfnTemplatePropsMixin_SubjectNameFlagsV3Property struct {
	// Include the common name in the subject name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-requirecommonname
	//
	RequireCommonName interface{} `field:"optional" json:"requireCommonName" yaml:"requireCommonName"`
	// Include the directory path in the subject name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-requiredirectorypath
	//
	RequireDirectoryPath interface{} `field:"optional" json:"requireDirectoryPath" yaml:"requireDirectoryPath"`
	// Include the DNS as common name in the subject name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-requirednsascn
	//
	RequireDnsAsCn interface{} `field:"optional" json:"requireDnsAsCn" yaml:"requireDnsAsCn"`
	// Include the subject's email in the subject name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-requireemail
	//
	RequireEmail interface{} `field:"optional" json:"requireEmail" yaml:"requireEmail"`
	// Include the globally unique identifier (GUID) in the subject alternate name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-sanrequiredirectoryguid
	//
	SanRequireDirectoryGuid interface{} `field:"optional" json:"sanRequireDirectoryGuid" yaml:"sanRequireDirectoryGuid"`
	// Include the DNS in the subject alternate name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-sanrequiredns
	//
	SanRequireDns interface{} `field:"optional" json:"sanRequireDns" yaml:"sanRequireDns"`
	// Include the domain DNS in the subject alternate name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-sanrequiredomaindns
	//
	SanRequireDomainDns interface{} `field:"optional" json:"sanRequireDomainDns" yaml:"sanRequireDomainDns"`
	// Include the subject's email in the subject alternate name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-sanrequireemail
	//
	SanRequireEmail interface{} `field:"optional" json:"sanRequireEmail" yaml:"sanRequireEmail"`
	// Include the service principal name (SPN) in the subject alternate name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-sanrequirespn
	//
	SanRequireSpn interface{} `field:"optional" json:"sanRequireSpn" yaml:"sanRequireSpn"`
	// Include the user principal name (UPN) in the subject alternate name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-subjectnameflagsv3.html#cfn-pcaconnectorad-template-subjectnameflagsv3-sanrequireupn
	//
	SanRequireUpn interface{} `field:"optional" json:"sanRequireUpn" yaml:"sanRequireUpn"`
}

