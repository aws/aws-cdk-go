package previewawsappmeshmixins


// An object that represents the subject alternative names secured by the certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subjectAlternativeNamesProperty := &SubjectAlternativeNamesProperty{
//   	Match: &SubjectAlternativeNameMatchersProperty{
//   		Exact: []*string{
//   			jsii.String("exact"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-subjectalternativenames.html
//
type CfnVirtualGatewayPropsMixin_SubjectAlternativeNamesProperty struct {
	// An object that represents the criteria for determining a SANs match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-subjectalternativenames.html#cfn-appmesh-virtualgateway-subjectalternativenames-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

