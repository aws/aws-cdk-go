package awsappmesh


// An object that represents the subject alternative names secured by the certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subjectAlternativeNamesProperty := &subjectAlternativeNamesProperty{
//   	match: &subjectAlternativeNameMatchersProperty{
//   		exact: []*string{
//   			jsii.String("exact"),
//   		},
//   	},
//   }
//
type CfnVirtualNode_SubjectAlternativeNamesProperty struct {
	// An object that represents the criteria for determining a SANs match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
}

