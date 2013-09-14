package datastore

import (
	"net/http"
	"appengine"
	"appengine/datastore"
)

type Progress struct {
	BookIndex  int    `json:"bookIndex"`
	Progress   int    `json:"progress"`
	Percentage int    `json:"percentage"`
	StoredOn   string `json:"storedOn"`
	DeviceName string `json:"deviceName"`
}

type Datastore interface {
	Get(key string) *Progress
	Put(key string, p *Progress)
	SetContext(r *http.Request)
}

type Mock struct {
	data map[string]*Progress
}

func (m *Mock) SetContext(*http.Request) {}

func (m *Mock) Get(key string) *Progress {
	if d, ok := m.data[key]; ok {
		return d
	}

	return nil
}

func (m *Mock) Put(key string, p *Progress) {
	m.data[key] = p
}

type Appengine struct {
	request *http.Request
}

func (m *Appengine) SetContext(r *http.Request) {
	m.request = r
}

func (m *Appengine) Get(key string) *Progress {
	c := appengine.NewContext(m.request)
	var p Progress
	k := datastore.NewKey(c, "Progress", key, 0, nil)
	err := datastore.Get(c, k, &p)
	if err != nil {
		c.Errorf("Error in Get: %s", err.Error())
		return nil
	}
	return &p
}

func (m *Appengine) Put(key string, p *Progress) {
	c := appengine.NewContext(m.request)
	k := datastore.NewKey(c, "Progress", key, 0, nil)
	_, err := datastore.Put(c, k, p)
	if err != nil {
		c.Errorf("Error in Put: %s", err.Error())
	}
}
