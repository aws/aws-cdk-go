package previewawss3objectlambdamixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyStatusProperty := &PolicyStatusProperty{
//   	IsPublic: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-policystatus.html
//
type CfnAccessPointPropsMixin_PolicyStatusProperty struct {
	// Specifies whether the Object lambda Access Point Policy is Public or not.
	//
	// Object lambda Access Points are private by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-policystatus.html#cfn-s3objectlambda-accesspoint-policystatus-ispublic
	//
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
}

