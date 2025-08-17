package opt

// Custom types for parameters
type WorkingDir string
type EnvVar string
type Shell string

// Boolean flag types with constants
type ShellFlag bool
const (
	UseShell   ShellFlag = true
	NoShell    ShellFlag = false
)

type IgnoreErrorsFlag bool
const (
	IgnoreErrors   IgnoreErrorsFlag = true
	NoIgnoreErrors IgnoreErrorsFlag = false
)

type QuietFlag bool
const (
	Quiet   QuietFlag = true
	NoQuiet QuietFlag = false
)

type InteractiveFlag bool
const (
	Interactive   InteractiveFlag = true
	NoInteractive InteractiveFlag = false
)

type InheritEnvFlag bool
const (
	InheritEnv   InheritEnvFlag = true
	NoInheritEnv InheritEnvFlag = false
)

// Flags represents the configuration options for the exec command
type Flags struct {
	WorkingDir    WorkingDir       // Working directory for command execution
	EnvVars       []EnvVar         // Environment variables (KEY=VALUE format)
	Shell         Shell            // Shell to use (if UseShell is true)
	UseShell      ShellFlag        // Run command through shell
	IgnoreErrors  IgnoreErrorsFlag // Continue on non-zero exit codes
	Quiet         QuietFlag        // Suppress stderr output
	Interactive   InteractiveFlag  // Allow interactive commands
	InheritEnv    InheritEnvFlag   // Inherit parent environment
}

// Configure methods for the opt system
func (w WorkingDir) Configure(flags *Flags)       { flags.WorkingDir = w }
func (e EnvVar) Configure(flags *Flags)           { flags.EnvVars = append(flags.EnvVars, e) }
func (s Shell) Configure(flags *Flags)            { flags.Shell = s }
func (s ShellFlag) Configure(flags *Flags)        { flags.UseShell = s }
func (i IgnoreErrorsFlag) Configure(flags *Flags) { flags.IgnoreErrors = i }
func (q QuietFlag) Configure(flags *Flags)        { flags.Quiet = q }
func (i InteractiveFlag) Configure(flags *Flags)  { flags.Interactive = i }
func (i InheritEnvFlag) Configure(flags *Flags)   { flags.InheritEnv = i }
