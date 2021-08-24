package fastdialer

// DefaultResolvers trusted
var DefaultResolvers = []string{
	"1.1.1.1:53",
	"1.0.0.1:53",
	"8.8.8.8:53",
	"8.8.4.4:53",
}

type CacheType uint8

const (
	Memory CacheType = iota
	Disk
)

type DiskDBType uint8

const (
	LevelDB DiskDBType = iota
	Pogreb
)

type Options struct {
	BaseResolvers      []string
	MaxRetries         int
	HostsFile          bool
	ResolversFile      bool
	EnableFallback     bool
	Allow              []string
	Deny               []string
	CacheType          CacheType
	CacheMemoryMaxSize int // used by Memory cache type
	DiskDbType         DiskDBType
	WithDialerHistory  bool
	WithCleanup        bool
}

// DefaultOptions of the cache
var DefaultOptions = Options{
	BaseResolvers: DefaultResolvers,
	MaxRetries:    5,
	HostsFile:     true,
	ResolversFile: true,
	CacheType:     Disk,
}
