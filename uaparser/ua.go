package uaparser

type Browser struct {
	Name    string
	Version string
	Major   string
}

type Cpu struct {
	Architecture string
}

type Engine struct {
	Name    string
	Version string
}

type OS struct {
	Name    string
	Version string
}

type Device struct {
	Vendor string
	Model  string
	Type   string
}

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

func Parse(ua []byte) (*UserAgent, error) {
	uagent := &UserAgent{
		UA: ua,
	}

	err := uagent.parse()

	return uagent, err
}
