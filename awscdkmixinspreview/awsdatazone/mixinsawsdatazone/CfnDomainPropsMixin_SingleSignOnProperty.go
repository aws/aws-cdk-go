package mixinsawsdatazone


// The single sign-on details in Amazon DataZone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   singleSignOnProperty := &SingleSignOnProperty{
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	Type: jsii.String("type"),
//   	UserAssignment: jsii.String("userAssignment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-domain-singlesignon.html
//
type CfnDomainPropsMixin_SingleSignOnProperty struct {
	// The ARN of the IDC instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-domain-singlesignon.html#cfn-datazone-domain-singlesignon-idcinstancearn
	//
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The type of single sign-on in Amazon DataZone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-domain-singlesignon.html#cfn-datazone-domain-singlesignon-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The single sign-on user assignment in Amazon DataZone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-domain-singlesignon.html#cfn-datazone-domain-singlesignon-userassignment
	//
	UserAssignment *string `field:"optional" json:"userAssignment" yaml:"userAssignment"`
}

