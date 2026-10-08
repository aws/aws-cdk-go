package awsnetworksecuritymanager


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associatedScopeProperty := &AssociatedScopeProperty{
//   	ScopeArn: jsii.String("scopeArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-deployment-associatedscope.html
//
type CfnDeployment_AssociatedScopeProperty struct {
	// ARN of the associated scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-deployment-associatedscope.html#cfn-networksecuritymanager-deployment-associatedscope-scopearn
	//
	ScopeArn *string `field:"required" json:"scopeArn" yaml:"scopeArn"`
}

