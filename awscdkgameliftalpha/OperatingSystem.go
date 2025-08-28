package awscdkgameliftalpha


// The operating system that the game server binaries are built to run on.
// Experimental.
type OperatingSystem string

const (
	// Amazon Linux operating system.
	// Experimental.
	OperatingSystem_AMAZON_LINUX OperatingSystem = "AMAZON_LINUX"
	// Amazon Linux 2 operating system.
	// Experimental.
	OperatingSystem_AMAZON_LINUX_2 OperatingSystem = "AMAZON_LINUX_2"
	// Amazon Linux 2023 operating system.
	// Experimental.
	OperatingSystem_AMAZON_LINUX_2023 OperatingSystem = "AMAZON_LINUX_2023"
	// Windows Server 2012 operating system.
	// Deprecated: If you have active fleets using the Windows Server 2012 operating system,
	// you can continue to create new builds using this OS until October 10, 2023, when Microsoft ends its support.
	// All others must use Windows Server 2016 when creating new Windows-based builds.
	OperatingSystem_WINDOWS_2012 OperatingSystem = "WINDOWS_2012"
	// Windows Server 2016 operating system.
	// Experimental.
	OperatingSystem_WINDOWS_2016 OperatingSystem = "WINDOWS_2016"
)

