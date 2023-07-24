package users

type UserDB struct {
	Users map[string][]Dog
}

type Dog struct {
	Breed   string
	Picture string
}

func NewDB() *UserDB {
	return &UserDB{
		Users: make(map[string][]Dog, 0),
	}
}

func (db UserDB) addUser(user string) {
	db.Users[user] = make([]Dog, 0)
}

func (db UserDB) addDog(user, breed, picture string) {

	_, ok := db.Users[user]
	// If the key does not exists
	if !ok {
		db.addUser(user)
	}

	newDog := Dog{
		Breed:   breed,
		Picture: picture,
	}
	db.Users[user] = append(db.Users[user], newDog)
}

func (db *UserDB) deleteDog(user, picture string) bool {
	dogs, ok := db.Users[user]
	if !ok {
		// User not found, nothing to delete
		return false
	}

	// Find the index of the dog with the given picture
	index := -1
	for i, dog := range dogs {
		if dog.Picture == picture {
			index = i
			break
		}
	}

	if index == -1 {
		// Dog with the given breed not found for this user
		return false
	}

	// Remove the dog from the slice
	db.Users[user] = append(dogs[:index], dogs[index+1:]...)
	return true
}

func (db *UserDB) deleteAllDog(user string) {
	db.Users[user] = nil
}

func (db *UserDB) getDogs(user string) []Dog {
	return db.Users[user]
}
