package task

import (
	"fmt"
	"net/url"
	"scrago/core/util"
	"scrago/custom/setting"
	"sync"
)

type Deduplicate interface {
	De(u url.URL) int
	Exist(u url.URL) int
	Remove(u url.URL) int
}

type DefaultDeduplicate struct {
	b    [setting.BloomFilterSize]bool
	lock sync.Mutex
}

func (d *DefaultDeduplicate) Exist(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := util.HashCode(s)
	position := util.Reserve(hashCode, setting.DeduplicationOffset)

	exists := d.b[position]
	if exists {
		return 1
	} else {
		return 0
	}
}

func NewDeduplicate() Deduplicate {
	return &DefaultDeduplicate{}
}

func (d *DefaultDeduplicate) De(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := util.HashCode(s)
	position := util.Reserve(hashCode, setting.DeduplicationOffset)

	exists := d.b[position]
	if exists {
		return 1
	} else {
		d.lock.Lock()
		defer d.lock.Unlock()
		if d.b[position] == false {
			d.b[position] = true
			return 0
		} else {
			return 1
		}
	}
}

func (d *DefaultDeduplicate) Remove(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := util.HashCode(s)
	position := util.Reserve(hashCode, setting.DeduplicationOffset)

	exists := d.b[position]
	if exists {
		d.b[position] = false
		return 1
	} else {
		return 0
	}
}
