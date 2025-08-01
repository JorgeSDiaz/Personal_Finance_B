package repository

import (
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
)

type inMemoryRepository struct {
	data   []*model.User
	mu     *sync.Mutex
	nextID int
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		mu:     &sync.Mutex{},
		data:   make([]*model.User, 0),
		nextID: 1,
	}
}

func (r *inMemoryRepository) getId() int {
	id := r.nextID
	r.nextID++

	return id
}

func (r *inMemoryRepository) FindAll() []*model.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.data
}

func (r *inMemoryRepository) Save(userRequest *model.UserRegistryRequest) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if existUserWithEmailInData(userRequest.Email, r.data) {
		return nil, fmt.Errorf("email already exist")
	}

	userToSave := translateRegistryToSave(*userRequest)

	userToSave.ID = r.getId()
	userToSave.CreatedAt = time.Now()
	userToSave.UpdatedAt = time.Now()

	r.data = append(r.data, &userToSave)
	return &userToSave, nil
}

func existUserWithEmailInData(email string, data []*model.User) bool {
	for _, user := range data {
		if strings.EqualFold(user.Email, email) {
			return true
		}
	}

	return false
}

func translateRegistryToSave(userRequest model.UserRegistryRequest) model.User {
	return model.User{
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Password: base64.StdEncoding.EncodeToString([]byte(userRequest.Password)),
	}
}

func (r *inMemoryRepository) GetUserByEmail(userRequest *model.UserLogInRequest) *model.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.data {
		if strings.EqualFold(user.Email, userRequest.Email) {
			return user
		}
	}

	return nil
}
