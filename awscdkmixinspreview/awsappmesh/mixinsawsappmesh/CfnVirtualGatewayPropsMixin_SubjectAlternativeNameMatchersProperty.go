package mixinsawsappmesh


// An object that represents the methods by which a subject alternative name on a peer Transport Layer Security (TLS) certificate can be matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subjectAlternativeNameMatchersProperty := &SubjectAlternativeNameMatchersProperty{
//   	Exact: []*string{
//   		jsii.String("exact"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-subjectalternativenamematchers.html
//
type CfnVirtualGatewayPropsMixin_SubjectAlternativeNameMatchersProperty struct {
	// The values sent must match the specified values exactly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-subjectalternativenamematchers.html#cfn-appmesh-virtualgateway-subjectalternativenamematchers-exact
	//
	Exact *[]*string `field:"optional" json:"exact" yaml:"exact"`
}

