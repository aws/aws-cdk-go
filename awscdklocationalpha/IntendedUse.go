package awscdklocationalpha


// Intend use for the results of an operation.
// Experimental.
type IntendedUse string

const (
	// The results won't be stored.
	// Experimental.
	IntendedUse_SINGLE_USE IntendedUse = "SINGLE_USE"
	// The result can be cached or stored in a database.
	// Experimental.
	IntendedUse_STORAGE IntendedUse = "STORAGE"
)

