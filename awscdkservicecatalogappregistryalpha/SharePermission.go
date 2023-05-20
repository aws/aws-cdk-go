package awscdkservicecatalogappregistryalpha


// Supported permissions for sharing applications or attribute groups with principals using AWS RAM.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var application application
//
//   application.shareApplication(jsii.String("MyShareId"), &ShareOptions{
//   	Name: jsii.String("MyShare"),
//   	Accounts: []*string{
//   		jsii.String("123456789012"),
//   		jsii.String("234567890123"),
//   	},
//   	SharePermission: appreg.SharePermission_ALLOW_ACCESS,
//   })
//
// Experimental.
type SharePermission string

const (
	// Allows principals in the share to only view the application or attribute group.
	// Experimental.
	SharePermission_READ_ONLY SharePermission = "READ_ONLY"
	// Allows principals in the share to associate resources and attribute groups with applications.
	// Experimental.
	SharePermission_ALLOW_ACCESS SharePermission = "ALLOW_ACCESS"
)

