package awsecs


// A Linux capability.
// Experimental.
type Capability string

const (
	// Experimental.
	Capability_ALL Capability = "ALL"
	// Experimental.
	Capability_AUDIT_CONTROL Capability = "AUDIT_CONTROL"
	// Experimental.
	Capability_AUDIT_WRITE Capability = "AUDIT_WRITE"
	// Experimental.
	Capability_BLOCK_SUSPEND Capability = "BLOCK_SUSPEND"
	// Experimental.
	Capability_CHOWN Capability = "CHOWN"
	// Experimental.
	Capability_DAC_OVERRIDE Capability = "DAC_OVERRIDE"
	// Experimental.
	Capability_DAC_READ_SEARCH Capability = "DAC_READ_SEARCH"
	// Experimental.
	Capability_FOWNER Capability = "FOWNER"
	// Experimental.
	Capability_FSETID Capability = "FSETID"
	// Experimental.
	Capability_IPC_LOCK Capability = "IPC_LOCK"
	// Experimental.
	Capability_IPC_OWNER Capability = "IPC_OWNER"
	// Experimental.
	Capability_KILL Capability = "KILL"
	// Experimental.
	Capability_LEASE Capability = "LEASE"
	// Experimental.
	Capability_LINUX_IMMUTABLE Capability = "LINUX_IMMUTABLE"
	// Experimental.
	Capability_MAC_ADMIN Capability = "MAC_ADMIN"
	// Experimental.
	Capability_MAC_OVERRIDE Capability = "MAC_OVERRIDE"
	// Experimental.
	Capability_MKNOD Capability = "MKNOD"
	// Experimental.
	Capability_NET_ADMIN Capability = "NET_ADMIN"
	// Experimental.
	Capability_NET_BIND_SERVICE Capability = "NET_BIND_SERVICE"
	// Experimental.
	Capability_NET_BROADCAST Capability = "NET_BROADCAST"
	// Experimental.
	Capability_NET_RAW Capability = "NET_RAW"
	// Experimental.
	Capability_SETFCAP Capability = "SETFCAP"
	// Experimental.
	Capability_SETGID Capability = "SETGID"
	// Experimental.
	Capability_SETPCAP Capability = "SETPCAP"
	// Experimental.
	Capability_SETUID Capability = "SETUID"
	// Experimental.
	Capability_SYS_ADMIN Capability = "SYS_ADMIN"
	// Experimental.
	Capability_SYS_BOOT Capability = "SYS_BOOT"
	// Experimental.
	Capability_SYS_CHROOT Capability = "SYS_CHROOT"
	// Experimental.
	Capability_SYS_MODULE Capability = "SYS_MODULE"
	// Experimental.
	Capability_SYS_NICE Capability = "SYS_NICE"
	// Experimental.
	Capability_SYS_PACCT Capability = "SYS_PACCT"
	// Experimental.
	Capability_SYS_PTRACE Capability = "SYS_PTRACE"
	// Experimental.
	Capability_SYS_RAWIO Capability = "SYS_RAWIO"
	// Experimental.
	Capability_SYS_RESOURCE Capability = "SYS_RESOURCE"
	// Experimental.
	Capability_SYS_TIME Capability = "SYS_TIME"
	// Experimental.
	Capability_SYS_TTY_CONFIG Capability = "SYS_TTY_CONFIG"
	// Experimental.
	Capability_SYSLOG Capability = "SYSLOG"
	// Experimental.
	Capability_WAKE_ALARM Capability = "WAKE_ALARM"
)

