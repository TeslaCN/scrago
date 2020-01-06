package dedup

import (
	"fmt"
	"github.com/TeslaCN/scrago/core/setting"
	"net/url"
	"reflect"
	"sync"
)

func init() {
	setting.AddDeduplicationType("DefaultDeduplicate", reflect.TypeOf(&DefaultDeduplicate{}))
}

// URL去重声明
type Deduplicate interface {

	// De 去重
	// 返回结果：
	// 目标已存在 大于0
	// 否则返回 0
	De(u url.URL) int

	// Exist 是否重复
	// 返回结果：
	// 目标存在 大于0
	// 不存在 0
	Exist(u url.URL) int

	// Remove 撤销去重
	// 返回结果：
	// 目标存在 大于0
	// 原本就不存在 返回0
	Remove(u url.URL) int
}

// 默认去重实现，基于进程内存实现
type DefaultDeduplicate struct {
	b    []bool
	lock sync.Mutex
}

func (d *DefaultDeduplicate) Exist(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := HashCode(s)
	position := Reserve(hashCode, setting.GetDeduplicationOffset())

	exists := d.b[position]
	if exists {
		return 1
	} else {
		return 0
	}
}

func NewDeduplicate() Deduplicate {
	d := &DefaultDeduplicate{}
	d.b = make([]bool, setting.GetBloomFilterSize())
	return d
}

func (d *DefaultDeduplicate) De(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := HashCode(s)
	position := Reserve(hashCode, setting.GetDeduplicationOffset())

	d.lock.Lock()
	defer d.lock.Unlock()
	exists := d.b[position]
	if exists {
		return 1
	} else {
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
	hashCode := HashCode(s)
	position := Reserve(hashCode, setting.GetDeduplicationOffset())

	d.lock.Lock()
	defer d.lock.Unlock()
	exists := d.b[position]
	if exists {
		d.b[position] = false
		return 1
	} else {
		return 0
	}
}
