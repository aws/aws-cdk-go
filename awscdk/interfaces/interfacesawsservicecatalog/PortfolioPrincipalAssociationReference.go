package interfacesawsservicecatalog


// A reference to a PortfolioPrincipalAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioPrincipalAssociationReference := &PortfolioPrincipalAssociationReference{
//   	PortfolioId: jsii.String("portfolioId"),
//   	PrincipalArn: jsii.String("principalArn"),
//   }
//
type PortfolioPrincipalAssociationReference struct {
	// The PortfolioId of the PortfolioPrincipalAssociation resource.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The PrincipalARN of the PortfolioPrincipalAssociation resource.
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
}

