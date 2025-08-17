package command

type WorkingDir string
type EnvVar string
type Shell string

type ShellFlag bool

const (
	UseShell ShellFlag = true
	NoShell  ShellFlag = false
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

type flags struct {
	WorkingDir   WorkingDir
	EnvVars      []EnvVar
	Shell        Shell
	UseShell     ShellFlag
	IgnoreErrors IgnoreErrorsFlag
	Quiet        QuietFlag
	Interactive  InteractiveFlag
	InheritEnv   InheritEnvFlag
}

func (w WorkingDir) Configure(flags *flags)       { flags.WorkingDir = w }
func (e EnvVar) Configure(flags *flags)           { flags.EnvVars = append(flags.EnvVars, e) }
func (s Shell) Configure(flags *flags)            { flags.Shell = s }
func (s ShellFlag) Configure(flags *flags)        { flags.UseShell = s }
func (i IgnoreErrorsFlag) Configure(flags *flags) { flags.IgnoreErrors = i }
func (q QuietFlag) Configure(flags *flags)        { flags.Quiet = q }
func (i InteractiveFlag) Configure(flags *flags)  { flags.Interactive = i }
func (i InheritEnvFlag) Configure(flags *flags)   { flags.InheritEnv = i }
