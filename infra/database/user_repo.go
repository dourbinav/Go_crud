package database

import (
    "errors"
    "keycloakwithgo/domain/user"
)

type InMemoryUserRepo struct {
    users map[int]user.User
    nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{
        users:  make(map[int]user.User),
        nextID: 1,
    }
}

func (r *InMemoryUserRepo) Create(u *user.User) (*user.User, error) {
    u.ID = r.nextID
    r.users[u.ID] = *u
    r.nextID++
    return u, nil
}

func (r *InMemoryUserRepo) FindAll() ([]user.User, error) {
    var result []user.User
    for _, v := range r.users {
        result = append(result, v)
    }
    return result, nil
}

func (r *InMemoryUserRepo) FindByID(id int) (*user.User, error) {
    u, ok := r.users[id]
    if !ok {
        return nil, errors.New("user not found")
    }
    return &u, nil
}

func (r *InMemoryUserRepo) Update(id int, u *user.User) (*user.User, error) {
    _, ok := r.users[id]
    if !ok {
        return nil, errors.New("user not found")
    }
    u.ID = id
    r.users[id] = *u
    return u, nil
}

func (r *InMemoryUserRepo) Delete(id int) error {
    _, ok := r.users[id]
    if !ok {
        return errors.New("user not found")
    }
    delete(r.users, id)
    return nil
}
