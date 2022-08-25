package awsappmesh


// An object that represents the methods by which a subject alternative name on a peer Transport Layer Security (TLS) certificate can be matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subjectAlternativeNameMatchersProperty := &subjectAlternativeNameMatchersProperty{
//   	exact: []*string{
//   		jsii.String("exact"),
//   	},
//   }
//
type CfnVirtualGateway_SubjectAlternativeNameMatchersProperty struct {
	// The values sent must match the specified values exactly.
	Exact *[]*string `field:"optional" json:"exact" yaml:"exact"`
}

