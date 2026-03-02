package awsrds


// The license model.
type LicenseModel string

const (
	// License included.
	LicenseModel_LICENSE_INCLUDED LicenseModel = "LICENSE_INCLUDED"
	// Bring your own license.
	LicenseModel_BRING_YOUR_OWN_LICENSE LicenseModel = "BRING_YOUR_OWN_LICENSE"
	// General public license.
	LicenseModel_GENERAL_PUBLIC_LICENSE LicenseModel = "GENERAL_PUBLIC_LICENSE"
)

