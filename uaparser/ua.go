package uaparser

// Borwser
type Browser struct {
	Name    string
	Version string
	Major   string
}

// Cpu
type Cpu struct {
	Architecture string
}

// Engine
type Engine struct {
	Name    string
	Version string
}

// OS
type OS struct {
	Name    string
	Version string
}

// Device
type Device struct {
	Vendor string
	Model  string
	Type   string
}

// UserAgent
type UserAgent struct {
	UA []byte
	*Browser
	Cpu
	Engine
	Device
	OS
}

func (u *UserAgent) parse() error {
	for _, rgx := range rgxMaps["browser"] {
		if rgx.Match(u.UA) {
			u.Browser = rgx.Parse(u.UA)
			break
		}
	}

	return nil
}

// Parse
func Parse(ua []byte) (*UserAgent, error) {
	uagent := &UserAgent{
		UA: ua,
	}

	err := uagent.parse()

	return uagent, err
}
