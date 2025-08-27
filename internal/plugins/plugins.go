package plugins

type Plugin interface {
	Name() string
	Run(args []string)
}

var registry = make(map[string]Plugin)

func Register(p Plugin) {
	registry[p.Name()] = p
}

func Get(name string) (Plugin, bool) {
	p, ok := registry[name]
	return p, ok
}

func List() []string {
	keys := []string{}
	for k := range registry {
		keys = append(keys, k)
	}
	return keys
}
