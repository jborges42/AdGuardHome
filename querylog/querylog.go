package querylog

import (
	"net"
	"time"

	"github.com/AdguardTeam/AdGuardHome/dnsfilter"
	"github.com/miekg/dns"
)

// QueryLog - main interface
type QueryLog interface {
	Close()

	// Set new configuration at runtime
	// Currently only 'Interval' field is supported.
	Configure(conf Config)

	Add(question *dns.Msg, answer *dns.Msg, result *dnsfilter.Result, elapsed time.Duration, addr net.Addr, upstream string)
	GetData() []map[string]interface{}
}

// Config - configuration object
type Config struct {
	BaseDir  string        // directory where log file is stored
	Interval time.Duration // interval to rotate logs
}

// New - create instance
func New(conf Config) QueryLog {
	return newQueryLog(conf)
}
