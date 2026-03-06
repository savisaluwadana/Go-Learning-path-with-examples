package contacts

import (
	"errors"
	"sort"
	"strings"
)

var (
	ErrNotFound     = errors.New("contact not found")
	ErrInvalidInput = errors.New("invalid contact input")
)

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

type Manager struct {
	nextID   int
	contacts map[int]Contact
}

func NewManager() *Manager {
	return &Manager{
		nextID:   1,
		contacts: make(map[int]Contact),
	}
}

func (m *Manager) Add(name, email, phone string) (Contact, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	phone = strings.TrimSpace(phone)

	if name == "" || email == "" || phone == "" {
		return Contact{}, ErrInvalidInput
	}

	contact := Contact{
		ID:    m.nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	m.contacts[contact.ID] = contact
	m.nextID++

	return contact, nil
}

func (m *Manager) Get(id int) (Contact, bool) {
	contact, ok := m.contacts[id]
	return contact, ok
}

func (m *Manager) List() []Contact {
	list := make([]Contact, 0, len(m.contacts))
	for _, c := range m.contacts {
		list = append(list, c)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}

func (m *Manager) Update(id int, name, email, phone string) (Contact, error) {
	if _, ok := m.contacts[id]; !ok {
		return Contact{}, ErrNotFound
	}

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	phone = strings.TrimSpace(phone)
	if name == "" || email == "" || phone == "" {
		return Contact{}, ErrInvalidInput
	}

	updated := Contact{
		ID:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	m.contacts[id] = updated
	return updated, nil
}

func (m *Manager) Delete(id int) error {
	if _, ok := m.contacts[id]; !ok {
		return ErrNotFound
	}

	delete(m.contacts, id)
	return nil
}
