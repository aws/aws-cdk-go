package awsresourceexplorer2


// Properties for defining a `CfnDefaultViewAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDefaultViewAssociationProps := &cfnDefaultViewAssociationProps{
//   	viewArn: jsii.String("viewArn"),
//   }
//
type CfnDefaultViewAssociationProps struct {
	// `AWS::ResourceExplorer2::DefaultViewAssociation.ViewArn`.
	ViewArn *string `field:"required" json:"viewArn" yaml:"viewArn"`
}

